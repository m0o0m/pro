package function

import (
	"fmt"
	"github.com/golyu/sql-build"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type ManualAccessBean struct{}

//添加一条人工存款
func (*ManualAccessBean) Add(this *input.ManualAccessAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//根据会员账号获取会员id,站点id,站点前台id,所属代理id
	member, _, err := GetMemberInfo(this.Account, this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//根据代理id获取代理账号
	agency, err := GetAccountById(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	manualAccess := new(schema.ManualAccess)
	manualAccess.AccessType = 1
	manualAccess.Balance = member.Balance
	manualAccess.CodeCount = this.CodeCount
	manualAccess.DepositDiscount = this.DepositDiscount
	manualAccess.DepositType = this.DepositType
	manualAccess.IsCodeCount = this.IsCodeCount
	manualAccess.IsDepositDiscount = this.IsDepositDiscount
	manualAccess.IsRemitDiscount = this.IsRemitDiscount
	manualAccess.IsRoutineCheck = this.IsRoutineCheck
	manualAccess.IsWriteRebate = this.IsWriteRebate
	manualAccess.Money = this.Money
	manualAccess.Remark = this.Remark
	manualAccess.RemitDiscount = this.RemitDiscount
	manualAccess.Account = this.Account
	manualAccess.DoAgencyId = this.DoAgencyId
	manualAccess.DoAgencyAccount = this.DoAgencyAccount
	manualAccess.MemberId = member.Id
	manualAccess.SiteId = member.SiteId
	manualAccess.SiteIndexId = member.SiteIndexId
	manualAccess.AgencyId = member.ThirdAgencyId
	manualAccess.AgencyAccount = agency.Account
	sess.Begin()
	//添加人工存取款数据
	count, err := sess.Table(manualAccess.TableName()).InsertOne(manualAccess)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	members := new(schema.Member)
	var remitDiscountMoney float64
	var depositDiscountMoney float64
	//会员余额 = 存款钱的余额 + 存款金额 + 各种优惠
	if manualAccess.IsDepositDiscount == 1 {
		depositDiscountMoney = manualAccess.DepositDiscount
	}
	if manualAccess.IsRemitDiscount == 1 {
		remitDiscountMoney = manualAccess.RemitDiscount
	}
	members.Balance = member.Balance + this.Money + remitDiscountMoney + depositDiscountMoney
	//会员表中修改余额
	count, err = sess.Table(members.TableName()).Where("id = ?", member.Id).Cols("balance").Update(members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//给现金流水表加数据
	mcr := new(schema.MemberCashRecord)
	mcr.SiteId = member.SiteId
	mcr.SiteIndexId = member.SiteIndexId
	mcr.MemberId = member.Id
	mcr.UserName = this.Account
	mcr.AgencyId = member.ThirdAgencyId
	mcr.AgencyAccount = agency.Account
	mcr.DisBalance = manualAccess.DepositDiscount + manualAccess.RemitDiscount
	mcr.SourceType = 11
	mcr.Balance = this.Money
	mcr.AfterBalance = members.Balance
	mcr.Remark = this.Remark
	mcr.CreateTime = time.Now().Unix()
	count, err = sess.Table(mcr.TableName()).InsertOne(mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//根据会员账号获取会员id,站点id,站点前台id,所属代理id
func GetMemberInfo(memberAccount, siteId, siteIndexId string) (*schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	if siteIndexId != "" {
		sess.Where("site_index_id=?", siteIndexId)
	}
	has, err := sess.Table(member.TableName()).Where("account = ?", memberAccount).
		Where("site_id=?", siteId).
		Select("id,site_id,site_index_id,third_agency_id,balance").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return member, has, err
	}
	return member, has, err
}

//根据会员账号获取会员id,站点id,站点前台id,所属代理id
func (*ManualAccessBean) GetMemberInfo(memberAccount, siteId, siteIndexId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	has, err := sess.Table(member.TableName()).
		Where("site_id = ?", siteId).
		Where("account = ?", memberAccount).
		Where("status = ?", 1).
		Where("delete_time = ?", 0).
		Select("id").Get(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//根据代理id获取代理账号
func GetAccountById(agencyId int64) (agency *schema.Agency, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency = new(schema.Agency)
	_, err = sess.Table(agency.TableName()).Where("id = ?", agencyId).Select("account").Get(agency)
	return
}

//根据代理id获取代理名称
func (*ManualAccessBean) GetUserNameById(agencyId int64) (*schema.Agency, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	_, err := sess.Table(agency.TableName()).
		Where("id = ?", agencyId).
		Select("username").Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agency, err
	}
	return agency, err
}

//获取存取款记录列表
func (*ManualAccessBean) GetList(this *input.ManualAccessList, listParams *global.ListParams, times *global.Times) (back.MemberCashRecordBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var list back.MemberCashRecordBack
	memberCashRecord := new(schema.MemberCashRecord)
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("user_name = ?", this.Account)
	}
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	mcr := make([]back.MemberCashRecord, 0)
	err := sess.Table(memberCashRecord.TableName()).Find(&mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	for k := range mcr {
		list.TotalMoney += mcr[k].Balance
	}
	//获得分页记录
	listParams.Make(sess)
	data := make([]back.MemberCashRecord, 0)
	//重新传入表名和where条件查询记录
	err = sess.Table(memberCashRecord.TableName()).Where(conds).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	for k := range data {
		list.SubTotalMoney += data[k].Balance
	}
	list.MemberCashRecord = data
	list.TotalCount = len(mcr)
	return list, err
}

//查看会员账号是否存在
func IsExistMemberAccount(memberAccount []string, siteId, siteIndexId string) ([]schema.Member, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	var m []schema.Member
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	err := sess.Table(member.TableName()).
		Where("site_id = ?", siteId).In("account", memberAccount).
		Where("delete_time = ?", 0).Find(&m)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return m, err
	}
	return m, err
}

//查看会员层级是否存在
func IsExistMemberLevel(memberLevel []string, siteId, siteIndexId string) ([]schema.Member, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var m []schema.Member
	member := new(schema.Member)
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	err := sess.Table(member.TableName()).Where("site_id = ?", siteId).In("level_id", memberLevel).Where("delete_time = ?", 0).
		GroupBy("level_id").GroupBy("level_id").Find(&m)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return m, err
	}
	return m, err
}

//根据会员层级查看会员id,站点id,站点前台id,所属代理id
func GetMemberByLevel(memberLevel, siteId, siteIndexId string) (members []schema.Member, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	err = sess.Table(member.TableName()).Where("site_id = ?", siteId).Where("site_index_id = ?", siteIndexId).Where("level_id = ?", memberLevel).Select("id,third_agency_id,account,balance").Find(&members)
	return
}

//添加多条人工存款(批量方式：账号)
func (*ManualAccessBean) AddBatchAccount(this *input.AddManualAccess) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	manualAccess := new(schema.ManualAccess)
	members := new(schema.Member)
	var mas []schema.ManualAccess
	var memberIds []int64
	var siteId string
	var siteIndexId string
	mcr := new(schema.MemberCashRecord)
	var mcrd schema.MemberCashRecord
	var mcrds []schema.MemberCashRecord
	sess.Begin()
	//根据会员账号数组获取会员id,站点id,站点前台id,所属代理id,代理账号
	data, err := GetInfoByMemberAccount(this.Account, this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	for k := range data {
		var ma schema.ManualAccess
		var remitDiscountMoney float64
		var depositDiscountMoney float64
		ma.AccessType = 1
		ma.CodeCount = this.CodeCount
		ma.DepositDiscount = this.DepositDiscount
		ma.DepositType = this.DepositType
		ma.IsCodeCount = this.IsCodeCount
		ma.IsDepositDiscount = this.IsDepositDiscount
		ma.IsRemitDiscount = this.IsRemitDiscount
		ma.IsRoutineCheck = this.IsRoutineCheck
		ma.IsWriteRebate = this.IsWriteRebate
		ma.Money = this.Money
		ma.Remark = this.Remark
		ma.RemitDiscount = this.RemitDiscount
		ma.DoAgencyId = this.DoAgencyId
		ma.DoAgencyAccount = this.DoAgencyAccount
		ma.Account = data[k].MemberAccount
		ma.MemberId = data[k].MemberId
		ma.SiteId = data[k].SiteId
		ma.SiteIndexId = data[k].SiteIndexId
		ma.Balance = data[k].Balance
		ma.AgencyId = data[k].AgencyId
		ma.AgencyAccount = data[k].AgencyAccount
		//会员余额 = 存款钱的余额 + 存款金额 + 各种优惠
		if manualAccess.IsDepositDiscount == 1 {
			depositDiscountMoney = this.DepositDiscount
		}
		if manualAccess.IsRemitDiscount == 1 {
			remitDiscountMoney = this.RemitDiscount
		}
		siteId = ma.SiteId
		siteIndexId = ma.SiteIndexId
		members.Balance = this.Money + remitDiscountMoney + depositDiscountMoney
		memberIds = append(memberIds, ma.MemberId)
		mas = append(mas, ma)
		//为现金流水赋值
		mcrd.SiteId = ma.SiteId
		mcrd.SiteIndexId = ma.SiteIndexId
		mcrd.MemberId = ma.MemberId
		mcrd.UserName = ma.Account
		mcrd.AgencyId = ma.AgencyId
		mcrd.AgencyAccount = ma.AgencyAccount
		mcrd.SourceType = 11
		mcrd.Balance = this.Money
		mcrd.DisBalance = this.DepositDiscount + this.RemitDiscount
		mcrd.AfterBalance = members.Balance
		mcrd.Remark = this.Remark
		mcrd.CreateTime = time.Now().Unix()
		mcrds = append(mcrds, mcrd)
	}
	//添加人工存取款数据
	count, err := sess.Table(manualAccess.TableName()).Insert(mas)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	var strs []string
	for k := range memberIds {
		str := fmt.Sprintf("%d", memberIds[k])
		if str != "" {
			strs = append(strs, str)
		}
	}
	ids := strings.Join(strs, ",")
	//会员表中修改余额
	sql := fmt.Sprintf("UPDATE sales_member SET balance = balance + %f WHERE site_id = '%s' AND site_index_id = '%s' AND id IN (%s) ", members.Balance, siteId, siteIndexId, ids)
	_, err = sess.Query(sql)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	//给现金流水表加数据
	count, err = sess.Table(mcr.TableName()).Insert(mcrds)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//根据会员账号数组获取获取会员id,站点id,站点前台id,所属代理id，代理账号
func GetInfoByMemberAccount(account []string, siteId, siteIndexId string) (data []back.Members, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	member := new(schema.Member)
	if siteIndexId != "" {
		sess.Where("sales_member.site_index_id=?", siteIndexId)
	}
	sql := fmt.Sprintf("%s.id = %s.third_agency_id", agency.TableName(), member.TableName())
	//重新传入表名和where条件查询记录
	err = sess.Table(member.TableName()).Select("sales_member.id as member_id,sales_member.site_id,sales_member.site_index_id, sales_member.balance,sales_member.third_agency_id as agency_id,sales_agency.account as agency_account,sales_member.account as member_account").Join("LEFT", agency.TableName(), sql).Where("sales_member.site_id=?", siteId).In(" sales_member.account", account).Find(&data)
	return
}

//添加多条人工存款(批量方式：层级)
func (*ManualAccessBean) AddBatchLevel(this *input.AddManualAccess) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	manualAccess := new(schema.ManualAccess)
	members := new(schema.Member)
	var mas []schema.ManualAccess
	var memberIds []int64
	var siteId string
	var siteIndexId string
	mcr := new(schema.MemberCashRecord)
	var mcrd schema.MemberCashRecord
	var mcrds []schema.MemberCashRecord
	sess.Begin()
	//根据会员层级获取会员id,站点id,站点前台id,所属代理id,代理账号
	data, err := GetInfoByMemberLevel(this.LevelId, this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if len(data) == 0 {
		return 0, err
	}
	for k := range data {
		var ma schema.ManualAccess
		var remitDiscountMoney float64
		var depositDiscountMoney float64
		ma.AccessType = 1
		ma.CodeCount = this.CodeCount
		ma.DepositDiscount = this.DepositDiscount
		ma.DepositType = this.DepositType
		ma.IsCodeCount = this.IsCodeCount
		ma.IsDepositDiscount = this.IsDepositDiscount
		ma.IsRemitDiscount = this.IsRemitDiscount
		ma.IsRoutineCheck = this.IsRoutineCheck
		ma.IsWriteRebate = this.IsWriteRebate
		ma.Money = this.Money
		ma.Remark = this.Remark
		ma.RemitDiscount = this.RemitDiscount
		ma.DoAgencyId = this.DoAgencyId
		ma.DoAgencyAccount = this.DoAgencyAccount
		ma.Account = data[k].MemberAccount
		ma.MemberId = data[k].MemberId
		ma.SiteId = data[k].SiteId
		ma.SiteIndexId = data[k].SiteIndexId
		ma.Balance = data[k].Balance
		ma.AgencyId = data[k].AgencyId
		ma.AgencyAccount = data[k].AgencyAccount
		//会员余额 = 存款钱的余额 + 存款金额 + 各种优惠
		if manualAccess.IsDepositDiscount == 1 {
			depositDiscountMoney = this.DepositDiscount
		}
		if manualAccess.IsRemitDiscount == 1 {
			remitDiscountMoney = this.RemitDiscount
		}
		siteId = ma.SiteId
		siteIndexId = ma.SiteIndexId
		members.Balance = this.Money + remitDiscountMoney + depositDiscountMoney
		memberIds = append(memberIds, ma.MemberId)
		mas = append(mas, ma)
		//为现金流水赋值
		mcrd.SiteId = ma.SiteId
		mcrd.SiteIndexId = ma.SiteIndexId
		mcrd.MemberId = ma.MemberId
		mcrd.UserName = ma.Account
		mcrd.AgencyId = ma.AgencyId
		mcrd.AgencyAccount = ma.AgencyAccount
		mcrd.SourceType = 11
		mcrd.Balance = this.Money
		mcrd.DisBalance = this.DepositDiscount + this.RemitDiscount
		mcrd.AfterBalance = members.Balance
		mcrd.Remark = this.Remark
		mcrd.CreateTime = time.Now().Unix()
		mcrds = append(mcrds, mcrd)
	}
	//添加人工存取款数据
	count, err := sess.Table(manualAccess.TableName()).Insert(mas)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	var strs []string
	for k := range memberIds {
		str := fmt.Sprintf("%d", memberIds[k])
		if str != "" {
			strs = append(strs, str)
		}
	}
	ids := strings.Join(strs, ",")
	//会员表中修改余额
	sql := fmt.Sprintf("UPDATE sales_member SET balance = balance + %f WHERE site_id = '%s' AND site_index_id = '%s' AND id IN (%s) ", members.Balance, siteId, siteIndexId, ids)
	_, err = sess.Query(sql)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	//给现金流水表加数据
	count, err = sess.Table(mcr.TableName()).Insert(mcrds)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//根据会员层级数组获取获取会员id,站点id,站点前台id,所属代理id，代理账号
func GetInfoByMemberLevel(levelId []string, siteId, siteIndexId string) (data []back.Members, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	member := new(schema.Member)
	if siteIndexId != "" {
		sess.Where("sales_member.site_index_id=?", siteIndexId)
	}
	sql := fmt.Sprintf("%s.id = %s.third_agency_id", agency.TableName(), member.TableName())
	//重新传入表名和where条件查询记录
	err = sess.Table(member.TableName()).
		Select("sales_member.id as member_id,sales_member.site_id,sales_member.site_index_id, sales_member.balance,sales_member.third_agency_id as agency_id,sales_member.account as member_account,sales_agency.account as agency_account").
		Join("LEFT", agency.TableName(), sql).
		Where("sales_member.site_id=?", siteId).
		In(" sales_member.level_id", levelId).Find(&data)
	return
}

//人工取款
func (*ManualAccessBean) Withdrawals(this *input.ManualWithdrawalAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//根据会员账号获取会员id,站点id,站点前台id,所属代理id
	member, _, err := GetMemberInfo(this.Account, this.SiteId, this.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	//根据代理id获取代理账号
	agency, err := GetAccountById(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	manualAccess := new(schema.ManualAccess)
	manualAccess.AccessType = 2
	manualAccess.Balance = member.Balance
	manualAccess.DepositType = this.DepositType
	manualAccess.Money = this.Money
	manualAccess.Remark = this.Remark
	manualAccess.Account = this.Account
	manualAccess.DoAgencyId = this.DoAgencyId
	manualAccess.DoAgencyAccount = this.DoAgencyAccount
	manualAccess.MemberId = member.Id
	manualAccess.SiteId = member.SiteId
	manualAccess.SiteIndexId = member.SiteIndexId
	manualAccess.AgencyId = member.ThirdAgencyId
	manualAccess.AgencyAccount = agency.Account
	sess.Begin()
	//在人工存取款表添加取款记录
	count, err := sess.Table(manualAccess.TableName()).InsertOne(manualAccess)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	members := new(schema.Member)
	//新余额 = 旧余额 - 取款金额
	members.Balance = member.Balance - this.Money
	//修改会员表的余额
	count, err = sess.Table(members.TableName()).Where("id = ?", member.Id).Cols("balance").Update(members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//给现金流水表加数据
	mcr := new(schema.MemberCashRecord)
	mcr.SiteId = member.SiteId
	mcr.SiteIndexId = member.SiteIndexId
	mcr.MemberId = member.Id
	mcr.UserName = this.Account
	mcr.AgencyId = member.ThirdAgencyId
	mcr.AgencyAccount = agency.Account
	mcr.SourceType = 3
	mcr.Balance = this.Money
	mcr.AfterBalance = members.Balance
	mcr.Remark = this.Remark
	mcr.CreateTime = time.Now().Unix()
	count, err = sess.Table(mcr.TableName()).InsertOne(mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err

}

//账目汇总-公司入款
func (*ManualAccessBean) CompanyIntoMoney(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	//根据时间段查询
	times.Make("create_time", sess)
	memberCompanyIncome := new(schema.MemberCompanyIncome)
	var mci []schema.MemberCompanyIncome
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("user_name = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(memberCompanyIncome.TableName()).Select("deposit_money").Find(&mci)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(memberCompanyIncome.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(memberCompanyIncome.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range mci {
		money := mci[k].DepositMoney
		mal.Money += money
	}
	return mal, err
}

//账目汇总-线上支付
func (*ManualAccessBean) OnlinePayment(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	onlieEntryRecord := new(schema.OnlineEntryRecord)
	var oer []schema.OnlineEntryRecord
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("member_account = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(onlieEntryRecord.TableName()).Select("amount_deposit").Find(&oer)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(onlieEntryRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(onlieEntryRecord.TableName()).Where(conds).GroupBy("member_id").Count()
	for k := range oer {
		money := oer[k].AmountDeposit
		mal.Money += money
	}
	return mal, err
}

//账目汇总-人工存入
func (*ManualAccessBean) DepositManually(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	manual_access := new(schema.ManualAccess)
	var manual_accesss []schema.ManualAccess
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("account = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	sess.Where("access_type = ?", 1)
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(manual_access.TableName()).Select("money").Find(&manual_accesss)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(manual_access.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(manual_access.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range manual_accesss {
		money := manual_accesss[k].Money
		mal.Money += money
	}
	return mal, err
}

//账目汇总-会员出款被扣除(会员出款被拒绝的)
func (*ManualAccessBean) MemberPaymentDeducted(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	//根据时间段查询
	times.Make("create_time", sess)
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("user_name = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	makeMoney := new(schema.MakeMoney)
	var mm []schema.MakeMoney
	sess.Where("out_status = ?", 4) //拒绝出款状态
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(makeMoney.TableName()).Select("outward_num").Find(&mm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(makeMoney.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(makeMoney.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range mm {
		money := mm[k].OutwardNum
		mal.Money += money
	}
	return mal, err
}

//账目汇总-会员出款
func (*ManualAccessBean) MemberPayment(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	//根据时间段查询
	times.Make("create_time", sess)
	makeMoney := new(schema.MakeMoney)
	var mm []schema.MakeMoney
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("user_name = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	sess.Where("out_status = ?", 1) //已出款状态
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(makeMoney.TableName()).Select("outward_money").Find(&mm)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(makeMoney.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(makeMoney.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range mm {
		money := mm[k].OutwardMoney
		mal.Money += money
	}
	return mal, err
}

//账目汇总-给予优惠
func (*ManualAccessBean) GiveDiscount(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	//根据时间段查询
	times.Make("create_time", sess)
	memberCompanyIncome := new(schema.MemberCompanyIncome)
	var mci []schema.MemberCompanyIncome
	onlieEntryRecord := new(schema.OnlineEntryRecord)
	var oer []schema.OnlineEntryRecord
	manual_access := new(schema.ManualAccess)
	var manual_accesss []schema.ManualAccess
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	var mciCount int64
	var mciPeopleNum int64
	var oerCount int64
	var oerPeopleNum int64
	var masCount int64
	var masPeopleNum int64
	if this.Account != "" {
		//公司入款给予的优惠
		err := sess.Table(memberCompanyIncome.TableName()).Where("user_name = ?", this.Account).Select("deposit_discount,other_discount").Find(&mci)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		//线上入款给予的优惠
		err = sess.Table(onlieEntryRecord.TableName()).Where("member_account = ?", this.Account).Select("deposit_discount,other_deposit_discount").Where(conds).Find(&oer)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		//人工存入给予的优惠
		err = sess.Table(manual_access.TableName()).Where("account = ?", this.Account).Select("deposit_discount,remit_discount").Where(conds).Find(&manual_accesss)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		mciCount, err = sess.Table(memberCompanyIncome.TableName()).Where("user_name = ?", this.Account).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		mciPeopleNum, err = sess.Table(memberCompanyIncome.TableName()).Where("user_name = ?", this.Account).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}

		oerCount, err = sess.Table(onlieEntryRecord.TableName()).Where("member_account = ?", this.Account).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		oerPeopleNum, err = sess.Table(onlieEntryRecord.TableName()).Where("member_account = ?", this.Account).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		masCount, err = sess.Table(manual_access.TableName()).Where("account = ?", this.Account).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		masPeopleNum, err = sess.Table(manual_access.TableName()).Where("account = ?", this.Account).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
	} else {
		err := sess.Table(memberCompanyIncome.TableName()).Select("deposit_discount,other_discount").Where(conds).Find(&mci)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		err = sess.Table(onlieEntryRecord.TableName()).Select("deposit_discount,other_deposit_discount").Where(conds).Find(&oer)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		err = sess.Table(manual_access.TableName()).Select("deposit_discount,remit_discount").Where(conds).Find(&manual_accesss)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		mciCount, err = sess.Table(memberCompanyIncome.TableName()).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		mciPeopleNum, err = sess.Table(memberCompanyIncome.TableName()).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}

		oerCount, err = sess.Table(onlieEntryRecord.TableName()).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		oerPeopleNum, err = sess.Table(onlieEntryRecord.TableName()).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		masCount, err = sess.Table(manual_access.TableName()).Where(conds).Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
		masPeopleNum, err = sess.Table(manual_access.TableName()).Where(conds).GroupBy("member_id").Count()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return mal, err
		}
	}
	var mciMoney float64 //公司入款总优惠金额
	var oerMoney float64 //线上入款总优惠金额
	var masMoney float64 //人工存款总优惠金额
	for k := range mci {
		mciMoney += mci[k].OtherDiscount + mci[k].DepositDiscount
	}
	for k := range oer {
		mciMoney += oer[k].OtherDepositDiscount + oer[k].DepositDiscount
	}
	for k := range manual_accesss {
		mciMoney += manual_accesss[k].RemitDiscount + manual_accesss[k].DepositDiscount
	}
	mal.Money = mciMoney + oerMoney + masMoney
	mal.Count = mciCount + oerCount + masCount
	mal.PeopleNum = mciPeopleNum + oerPeopleNum + masPeopleNum
	return mal, nil
}

//账目汇总-人工提现
func (*ManualAccessBean) ManualWithdrawal(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	manual_access := new(schema.ManualAccess)
	var manual_accesss []schema.ManualAccess
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("account = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	sess.Where("access_type = ?", 2)
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(manual_access.TableName()).Select("money").Find(&manual_accesss)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(manual_access.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(manual_access.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range manual_accesss {
		money := manual_accesss[k].Money
		mal.Money += money
	}
	return mal, err
}

//账目汇总-给予返水
func (*ManualAccessBean) GiveBackWater(this *input.ManualAccessLists, times *global.Times) (back.CompanyIntoMoney, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var mal back.CompanyIntoMoney
	//根据时间段查询
	times.Make("create_time", sess)
	memberCashRecord := new(schema.MemberCashRecord)
	var mcr []schema.MemberCashRecord
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("user_name = ?", this.Account)
	}
	if this.AgencyId > 0 {
		sess.Where("agency_id = ?", this.AgencyId)
	}
	sess.Where("(source_type = ? OR source_type = ?)", 9, 10)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	err := sess.Table(memberCashRecord.TableName()).Select("balance").Find(&mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.Count, err = sess.Table(memberCashRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	mal.PeopleNum, err = sess.Table(memberCashRecord.TableName()).Where(conds).GroupBy("member_id").Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mal, err
	}
	for k := range mcr {
		money := mcr[k].Balance
		mal.Money += money
	}
	return mal, err
}

//查询出款账目汇总的统计信息--人工提出 (总额度,总笔数,总人数)
func (m *ManualAccessBean) GetOutCount(siteId string, sTime, eTime int64) (*back.CashCollectDetails, error) {
	cashCollect := back.CashCollectDetails{}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//SELECT
	//count(*)     AS people_num,
	//	sum(t.money) AS money,
	//	sum(num)     AS num
	//FROM (SELECT
	//sum(money)       money,
	//	count(member_id) num
	//FROM sales_manual_access
	//WHERE access_type = 1
	//GROUP BY member_id) t
	subSql, err := sqlBuild.Select(new(schema.ManualAccess).TableName()).
		Column("sum(money) money,count(member_id) num").
		Where(siteId, "site_id").
		Where(sTime, "create_time>=").
		Where(eTime, "create_time<=").
		GroupBy("member_id").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashCollect, err
	}
	sql, err := sqlBuild.Select("(" + subSql + ") t").
		Column("count(*) people_num,sum(t.num) num,sum(t.money) money").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashCollect, err
	}
	_, err = sess.SQL(sql).Get(&cashCollect)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashCollect, err
	}
	return &cashCollect, err
}

//根据会员id获取会员账号,站点id,站点前台id,所属代理id
func WapGetMemberInfo(memberId int64) (member *schema.Member, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member = new(schema.Member)
	has, err = sess.Table(member.TableName()).Where("id = ?", memberId).
		Select("account,site_id,site_index_id,third_agency_id,balance").Get(member)
	return
}

//根据会员账号获取会员id,站点id,站点前台id,所属代理id
func (*ManualAccessBean) GetWapInfo(memberAccount, siteId, siteIndexId string) (*schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	has, err := sess.Table(member.TableName()).
		Where("site_id = ?", siteId).
		Where("site_index_id = ?", siteIndexId).
		Where("account = ?", memberAccount).
		Where("status = ?", 1).
		Where("delete_time = ?", 0).
		Get(member)
	return member, has, err
}

//根据代理id获取代理账号
func (*ManualAccessBean) GetAgencyAccount(agencyId int64) (agency *schema.Agency, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency = new(schema.Agency)
	_, err = sess.Table(agency.TableName()).Where("id = ?", agencyId).Select("account").Get(agency)
	return
}

//查询会员余额
func (*ManualAccessBean) GetBalance(memberId int64) (float64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	member := new(schema.Member)
	memberinfo := new(back.Balance)
	_, err := sess.Table(member.TableName()).Where("id=?", memberId).Get(memberinfo)
	return memberinfo.Balance, err
}
