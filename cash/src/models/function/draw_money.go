package function

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/schema"
	"time"
)

type DrawMoneyBean struct{}

//获取会员账号，余额，出款银行卡信息
func (*DrawMoneyBean) GetMemberInfo(memberId int64) (*back.WapMemberInfo, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	wapBank := make([]back.WapBank, 0)
	memberInfo := new(back.WapMemberInfo)
	member := new(schema.Member)
	memberBank := new(schema.MemberBank)
	bank := new(schema.Bank)
	//获取会员账号，余额
	_, err := sess.Where("id=?", memberId).Select("account,balance").Where("delete_time=0").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberInfo, err
	}
	//获取银行卡信息
	sql := fmt.Sprintf("%s.bank_id=%s.id", memberBank.TableName(), bank.TableName())
	sess.Select(bank.TableName() + ".id," + bank.TableName() + ".title," + memberBank.TableName() + ".card")
	err = sess.Table(memberBank.TableName()).Join("LEFT", bank.TableName(), sql).
		Where("member_id=?", memberId).Where(memberBank.TableName() + ".delete_time=0").Find(&wapBank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberInfo, err
	}
	//给返回的结构体赋值
	memberInfo.Account = member.Account
	memberInfo.Balance = member.Balance
	memberInfo.WapBank = wapBank
	return memberInfo, err
}

//获取会员余额和取款密码
func (*DrawMoneyBean) GetMemberBalanceAndPassword(memberId int64) (float64, string, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	var balance float64
	var password string
	has, err := sess.Where("id=?", memberId).Select("balance,draw_password").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return balance, password, has, err
	}
	balance = member.Balance
	password = member.DrawPassword
	return balance, password, has, err
}

//查看出款管理表是否有待审核的数据
func (*DrawMoneyBean) IsExistPendingReview(memberId int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makeMoney := new(schema.MakeMoney)
	has, err := sess.Where("member_id=?", memberId).
		Where("out_status=5").
		Get(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//会员银行是否存在
func (*DrawMoneyBean) IsExistMemberBank(memberId, bankId int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberBank := new(schema.MemberBank)
	has, err := sess.Where("member_id=?", memberId).Where("bank_id=?", bankId).
		Where("delete_time=0").Get(memberBank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//取款
func (*DrawMoneyBean) Withdrawal(this *global.MemberRedisToken, memberBalance, money float64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//会员表
	member := new(schema.Member)
	//代理表
	agency := new(schema.Agency)
	//打码统计表
	betReportAccount := new(schema.BetReportAccount)
	//现金流水表
	memberCashRecord := new(schema.MemberCashRecord)
	//出款管理表
	makeMoney := new(schema.MakeMoney)
	//出款稽核表
	member_audit := new(schema.MemberAudit)
	memberAudit := make([]back.WapMemberAudit, 0)
	//查看会员未出款稽核数据
	err := sess.Table(member_audit.TableName()).Where("site_id=?", this.Site).Where("site_index_id=?",
		this.SiteIndex).Where("member_id=?", this.Id).Find(&memberAudit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	var beginTime int64            //开始时间
	var totalBetValid float64      //有效总投注
	var totalMultipleMoney float64 //综合稽核总额
	var totalNormalMoney float64   //常态稽核总额
	var adminMoney float64         //行政费用
	var depositMoney float64       //优惠金额
	var relaxMoney int64           //放宽额度
	var totalMoney float64         //存款金额
	for i := range memberAudit {
		beginTime = memberAudit[i].BeginTime
		totalNormalMoney += memberAudit[i].NormalMoney
		totalMultipleMoney += memberAudit[i].MultipleMoney
		adminMoney = memberAudit[i].AdminMoney
		depositMoney = memberAudit[i].DepositMoney
		relaxMoney = memberAudit[i].RelaxMoney
		totalMoney += memberAudit[i].Money
	}
	//查看有效投注
	betValid := make([]back.WapBetValid, 0)
	err = sess.Table(betReportAccount.TableName()).Where("day_time >= ? AND day_time <=?", beginTime,
		time.Now().Unix()).Find(&betValid)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	for i := range betValid {
		totalBetValid += betValid[i].BetValid
	}
	//计算得出是否需要扣除优惠
	if totalBetValid >= totalMultipleMoney-float64(relaxMoney) {
		//如果有效总投注大就不用扣除优惠
		depositMoney = 0
	}
	//计算得出是否需要扣除行政费
	if totalBetValid >= totalNormalMoney-float64(relaxMoney) {
		//如果有效总投注大就不用扣除行政费
		adminMoney = 0
	}
	//查看会员层级和代理id
	has, err := sess.Where("id=?", this.Id).Select("level_id,third_agency_id").Get(member)
	if err != nil || !has {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//根据代理id获取代理账号
	has, err = sess.Where("id=?", member.ThirdAgencyId).Select("account").Get(agency)
	if err != nil || !has {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	sess.Begin()
	//给现金流水表赋值
	memberCashRecord.AgencyId = member.ThirdAgencyId
	memberCashRecord.AgencyAccount = agency.Account
	memberCashRecord.UserName = this.Account
	memberCashRecord.MemberId = this.Id
	memberCashRecord.SiteIndexId = this.SiteIndex
	memberCashRecord.SiteId = this.Site
	memberCashRecord.Balance = money
	memberCashRecord.AfterBalance = memberBalance
	memberCashRecord.SourceType = 5
	memberCashRecord.Type = 2
	memberCashRecord.ClientType = 2
	//给现金流水表增加数据
	count, err := sess.InsertOne(memberCashRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//给出款稽核表赋值
	member_audit.MemberId = this.Id
	member_audit.Account = this.Account
	member_audit.SiteIndexId = this.SiteIndex
	member_audit.SiteId = this.Site
	if beginTime == 0 {
		member_audit.BeginTime = time.Now().Unix()
	} else {
		member_audit.BeginTime = beginTime
	}
	member_audit.EndTime = time.Now().Unix()
	member_audit.NormalMoney = totalNormalMoney
	member_audit.MultipleMoney = totalMultipleMoney
	member_audit.AdminMoney = adminMoney
	member_audit.Money = totalMoney
	member_audit.DepositMoney = depositMoney
	member_audit.RelaxMoney = relaxMoney
	member_audit.Status = 1
	//给出款稽核表增加数据
	count, err = sess.InsertOne(member_audit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	//给出款管理表赋值
	makeMoney.AgencyAccount = agency.Account
	makeMoney.AgencyId = member.ThirdAgencyId
	makeMoney.SiteId = this.Site
	makeMoney.SiteIndexId = this.SiteIndex
	makeMoney.UserName = this.Account
	makeMoney.MemberId = this.Id
	makeMoney.LevelId = member.LevelId
	makeMoney.OutwardNum = money
	makeMoney.Balance = memberBalance
	makeMoney.FavourableMoney = depositMoney
	makeMoney.ExpeneseMoney = adminMoney
	makeMoney.OutwardMoney = money - depositMoney - adminMoney //实际出款金额=提出-行政费-优惠金额
	makeMoney.ClientType = 1
	makeMoney.OutStatus = 5 //待审核状态
	//给出款管理表增加数据
	count, err = sess.InsertOne(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	//操作后会员余额
	member.Balance = memberBalance
	//修改会员表余额
	count, err = sess.Where("id=?", this.Id).Cols("balance").Update(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//取款进度
func (*DrawMoneyBean) DrawalProgress(memberId int64) (*back.WapDrawalProgress, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	wapDrawalProgress := new(back.WapDrawalProgress)
	//出款管理
	makeMoney := new(schema.MakeMoney)
	//出款稽核
	memberAudit := new(schema.MemberAudit)
	//现金流水
	memberCashRecord := new(schema.MemberCashRecord)
	sql1 := fmt.Sprintf("%s.member_id = %s.member_id  AND %s.create_time = %s.end_time",
		makeMoney.TableName(), memberAudit.TableName(), makeMoney.TableName(), memberAudit.TableName())
	sql2 := fmt.Sprintf("%s.member_id = %s.member_id AND %s.create_time = %s.create_time",
		makeMoney.TableName(), memberCashRecord.TableName(), makeMoney.TableName(), memberCashRecord.TableName())
	//联和三张表，查询会员id为当前操作会员id，并用desc查出最新一条数据
	_, err := sess.Table(makeMoney.TableName()).Join("LEFT", memberAudit.TableName(), sql1).
		Join("LEFT", memberCashRecord.TableName(), sql2).Where(makeMoney.TableName()+".member_id=?",
		memberId).Desc(makeMoney.TableName() + ".id").Get(wapDrawalProgress)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return wapDrawalProgress, err
	}
	return wapDrawalProgress, err
}

//取款数据整理
func (*DrawMoneyBean) GetDrawData(this *global.MemberRedisToken) ([]back.WapMemberAudit, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//出款稽核表
	member_audit := new(schema.MemberAudit)
	memberAudit := make([]back.WapMemberAudit, 0)
	err := sess.Table(member_audit.TableName()).
		Select("sum(normal_money) as normal_money,sum(multiple_money) as multiple_money,sum(money) as money,sum(admin_money) as admin_money,sum(deposit_money) as deposit_money,relax_money,begin_time").
		Where("account=?", this.Account).
		Where("site_id=?", this.Site).
		Where("status=?", 1).
		OrderBy("id asc").
		Find(&memberAudit)
	if err != nil {
		return nil, err
	}
	if memberAudit[0].Money == 0 {
		return nil, err
	}
	return memberAudit, err
}

//查询统计表会员稽核期间的打码情况
func (*DrawMoneyBean) GetReportValid(account, siteId string, sTime, eTime int64) (data []back.AuditBet, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if account != "" {
		sess.Where("account=?", account)
	}
	sess.Where("site_id=?", siteId)
	sess.Where("day_time>=?", sTime)
	sess.Where("day_time<?", eTime) //不要带等于
	sess.Select("id,site_id,site_index_id,account,sum(bet_valid) as bet_valid,game_type")
	bra := new(schema.BetReportAccount)
	err = sess.Table(bra.TableName()).
		GroupBy("game_type").Find(&data)
	return data, err
}

//查询注单表会员当日稽核期间的打码情况
func (*DrawMoneyBean) GetBetValid(account, siteId string, sTime, eTime int64) (data []back.AuditBetRecord, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if account != "" {
		sess.Where("username=?", account)
	}
	sess.Where("site_id=?", siteId)
	sess.Where("bet_timeline>=?", sTime)
	sess.Where("bet_timeline<=?", eTime)
	sess.Where("status=?", 1)
	sess.Select("site_id,index_id,username,sum(bet_yx) as bet_valid,game_type")
	bra := new(schema.BetRecordInfo)
	err = sess.Table(bra.TableName()).
		GroupBy("game_type").Find(&data)
	return data, err
}

//获取当前出款手续费
func (*DrawMoneyBean) GetOutCharge(siteId, siteIndexId string, money float64) (*back.OutCharge, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	poundage := new(schema.SitePoundage)
	has, err := sess.Table(poundage.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Get(poundage)
	outCharge := new(back.OutCharge)
	outCharge.OutMoney = money * poundage.OutPoundageRatio / 100
	return outCharge, has, err
}

//查询是否第一次出款
func (*DrawMoneyBean) IsFirst(memberId int64) (*schema.MakeMoney, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makeMoney := new(schema.MakeMoney)
	has, err := sess.Table(makeMoney.TableName()).Where("member_id=?", memberId).Get(makeMoney)
	return makeMoney, has, err
}

//查询站点出款设置 sales_site_pay_set
func (*DrawMoneyBean) GetSiteSet(siteId, siteIndexId string) (*back.SitePaySet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitepayset := new(schema.SitePaySet)
	data := new(back.SitePaySet)
	has, err := sess.Table(sitepayset.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).Get(data)
	return data, has, err
}

//扣除会员出款金额
func (*DrawMoneyBean) Deduction(memberId int64, deductionMoney float64, sessArgs ...*xorm.Session) (int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	member := new(schema.Member)
	member.Id = memberId
	member.Balance = deductionMoney
	sess.Cols("balance")
	data, err := sess.Table(member.TableName()).Where("id=?", memberId).Update(member)
	return data, err
}

//写入出款记录
func (*DrawMoneyBean) WriteChargeRecord(this *back.SaveMakeMoney, sessArgs ...*xorm.Session) (data int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	cashrecord := new(schema.MakeMoney)
	data, err = sess.Table(cashrecord.TableName()).Insert(this)
	return
}

//写入现金表记录
func (*DrawMoneyBean) WriteCashRecord(this *back.WapCashRecord, sessArgs ...*xorm.Session) (data int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	cashrecord := new(schema.MemberCashRecord)
	data, err = sess.Table(cashrecord.TableName()).Insert(this)
	return
}

//查询站点出款手续费及上限
func (*DrawMoneyBean) GetPoundageSet(siteId, siteIndexId string) (*back.Poundage, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	poundage := new(schema.SitePoundage)
	data := new(back.Poundage)
	has, err := sess.Table(poundage.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).Get(data)
	return data, has, err
}

//会员银行是否存在
func (*DrawMoneyBean) IsExistMemberBankById(memberId, Id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberBank := new(schema.MemberBank)
	has, err := sess.Where("member_id=?", memberId).Where("id=?", Id).
		Where("delete_time=0").Get(memberBank)
	return has, err
}

//查询会员出款是否有未出款
func (*DrawMoneyBean) IsDrawData(siteId, siteIndexId string, memberId int64) (*schema.MakeMoney, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makemoney := new(schema.MakeMoney)
	data := new(schema.MakeMoney)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("member_id=?", memberId)
	sess.Where("out_status=?", 5)
	has, err := sess.Table(makemoney.TableName()).Get(data)
	return data, has, err
}

//查询会员当天的出款次数
func (*DrawMoneyBean) NumOutMoney(account, siteId string, sTime, eTime int64) (num int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if account != "" {
		sess.Where("account=?", account)
	}
	sess.Where("site_id=?", siteId)
	sess.Where("create_time>=?", sTime)
	sess.Where("create_time<=?", eTime)
	num, err = sess.Table(new(schema.MakeMoney).TableName()).Count()
	return
}
