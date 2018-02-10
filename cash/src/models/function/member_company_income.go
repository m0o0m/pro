package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type MemberCompanyIncomeBean struct{}

//func ConfirmCompanyIncome 确认一条公司入款
func (*MemberCompanyIncomeBean) ConfirmCompanyIncome(newId *input.ConfirmCompany) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//首先更改入款状态
	err := sess.Begin()
	newCompanyIncome := new(schema.MemberCompanyIncome)
	newCompanyIncome.Status = 1
	count, err := sess.Where("id=?", newId.Id).Cols("status").Update(newCompanyIncome)
	if err != nil || count == 0 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
		}
		return count, err
	}
	//更改会员余额
	income := new(MemberCompanyIncomeBean)
	info, flag, err := income.GetInfoById(newId.Id)
	if !flag {
		return count, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	memberBean := new(MemberBean)
	balance := info.DepositMoney + info.DepositDiscount + info.OtherDiscount
	fmt.Println("balance:", balance)
	count, err = memberBean.ChangeBalance(info.MemberId, balance, 1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//获取该条入款详情
	entryBean := new(OnlineEntryRecordBean)
	infoMation, flags, err := entryBean.GetInfo(newId.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if !flags {
		return 0, err
	}
	//根据会员Id取出会员详情
	members := new(MemberBean)
	member, flags, err := members.GetInfoById(infoMation.MemberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if !flags {
		return 0, err
	}
	//更改member_cash_count 现金统计记录
	newCashCountBean := new(MemberCashCountBean)
	count, err = newCashCountBean.ChangeData(member.Id, infoMation.AmountDeposit)
	if err != nil || count == 0 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
		}
		return count, err
	}
	//更改现金系统记录
	newCashRecord := new(schema.MemberCashRecord)
	newCashRecord.CreateTime = time.Now().Unix()
	newCashRecord.SourceType = 2
	newCashRecord.TradeNo = infoMation.ThirdOrderNumber
	newCashRecord.SiteId = member.SiteId
	newCashRecord.SiteIndexId = member.SiteIndexId
	newCashRecord.MemberId = infoMation.MemberId
	newCashRecord.UserName = infoMation.MemberAccount
	newCashRecord.Balance = infoMation.AmountDeposit
	newCashRecord.AgencyId = infoMation.AgencyId
	newCashRecord.Type = 1
	newCashRecord.ClientType = int64(infoMation.SourceDeposit)
	newCashRecord.AgencyAccount = infoMation.AgencyAccount
	newCashRecord.Remark = infoMation.Remark
	newCashRecord.AfterBalance = info.DepositMoney + info.DepositDiscount + info.OtherDiscount
	memberCashrecord := new(MemberCashRecordBean)
	count, err = memberCashrecord.AddNewRecord(newCashRecord)
	if err != nil || count == 0 {
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
		}
		return 0, err
	}
	err = sess.Commit()
	return count, err
}

//func CancleCompanyIncomen 取消一条公司入款  不再提醒一条公司入款
func (*MemberCompanyIncomeBean) CancleCompanyIncome(cancle *input.CancleIncome, user *global.RedisStruct) (
	int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newCompanyIncome := new(schema.MemberCompanyIncome)
	if cancle.Status == 2 {
		newCompanyIncome.Status = 2
	} else if cancle.Status == 3 {
		newCompanyIncome.Status = 3
	}
	count, err := sess.
		Where("id=?", cancle.Id).
		Where("site_id=?", user.SiteId).
		Where("site_index_id=?", user.SiteIndexId).
		Cols("status").Update(newCompanyIncome)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//根据id获取该数据详情
func (*MemberCompanyIncomeBean) GetInfoById(id int64) (info schema.MemberCompanyIncome, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Table(info.TableName()).Where("id=?", id).Get(&info)
	return
}

//公司入款列表
func (*MemberCompanyIncomeBean) GetInfoList(this *input.CompanyIncomeList, times *global.Times, listParams *global.ListParams) (
	back.CompenyIncomeBackLists, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberCompanyIncome := new(schema.MemberCompanyIncome)
	var cibls back.CompenyIncomeBackLists
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("sales_member_company_income.site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("sales_member_company_income.site_index_id = ?", this.SiteIndexId)
	}
	if this.AgencyAccount != "" {
		sess.Where("sales_member_company_income.agency_account = ?", this.AgencyAccount)
	}
	if this.Level != "" {
		level := strings.Split(this.Level, ",")
		sess.In("sales_member_company_income.level_id", level)
	}
	if this.Status != 0 {
		sess.Where("sales_member_company_income.status = ?", this.Status)
	}
	if this.ClientType != 0 {
		sess.Where("sales_member_company_income.client_type = ?", this.ClientType)
	}
	if this.PaymentAccount != "" {
		//分割成数组
		ids := strings.Split(this.PaymentAccount, ",")
		sess.In("sales_site_bank_income_set.id", ids)
	}
	if this.LowerLimit >= 0 && this.UpperLimit != 0 {
		sess.Where("deposit_money < ?", this.UpperLimit).And("deposit_money > ?", this.LowerLimit)
	}
	if this.SelectBy == 1 {
		sess.Where("sales_member_company_income.account = ?", this.Conditions)
	} else if this.SelectBy == 2 {
		sess.Where("sales_member_company_income.order_num=?", this.Conditions)
	}
	if this.IsDiscount == 1 {
		sess.Where("deposit_discount != 0 OR other_discount != 0")
	} else if this.IsDiscount == 2 {
		sess.Where("deposit_discount = 0").And("other_discount = 0")
	}
	//根据时间段查询
	times.Make("sales_member_company_income.create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	mbclt := make([]back.CompenyIncomeList, 0)
	sbis := new(schema.SiteBankIncomeSet)
	sql := fmt.Sprintf("%s.set_id = %s.id", memberCompanyIncome.TableName(), sbis.TableName())
	err := sess.Table(memberCompanyIncome.TableName()).
		Join("LEFT", sbis.TableName(), sql).
		Find(&mbclt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cibls, err
	}
	//公司入款列表返回的数据
	var data []back.CompenyIncomeBackList
	var totalMoney float64               //总计存入金额
	var totalDepositDiscount float64     //总计存款优惠
	var totalOtherDiscount float64       //总计其他优惠
	var totalDeposit float64             //总计存入总金额
	var pageTotalMoney float64           //小计存入金额
	var pageTotalDepositDiscount float64 //小计存款优惠
	var pageTotalOtherDiscount float64   //小计其他优惠
	var pageTotalDeposit float64         //小计存入总金额
	for k := range mbclt {
		//总计
		totalMoney += mbclt[k].DepositMoney
		totalDepositDiscount += mbclt[k].DepositDiscount
		totalOtherDiscount += mbclt[k].OtherDiscount
		totalDeposit += mbclt[k].DepositCount
	}
	//获得分页记录
	listParams.Make(sess)
	mbcl := make([]back.CompenyIncomeList, 0)
	//重新传入表名和where条件查询记录
	err = sess.Table(memberCompanyIncome.TableName()).
		Join("LEFT", sbis.TableName(), sql).
		Where(conds).
		OrderBy("sales_member_company_income.create_time DESC").
		Find(&mbcl)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cibls, err
	}
	var agencyId []int64
	for k := range mbcl {
		agencyId = append(agencyId, mbcl[k].DoAgencyId)
	}
	//根据操作人id获取操作人名称
	agency, err := GetAgencyNameById(agencyId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cibls, err
	}
	var cibl back.CompenyIncomeBackList
	for k := range mbcl {
		for i := range agency {
			if mbcl[k].DoAgencyId == agency[i].Id {
				cibl.OperateName = agency[i].Username
			}
		}
		cibl.UpdateTime = mbcl[k].UpdateTime
		cibl.DepositTime = mbcl[k].DepositTime
		cibl.Id = mbcl[k].Id
		cibl.Status = mbcl[k].Status
		cibl.LevelId = mbcl[k].LevelId
		cibl.Account = mbcl[k].Account
		cibl.AgencyAccount = mbcl[k].AgencyAccount
		cibl.CreateTime = mbcl[k].CreateTime
		cibl.Remark = mbcl[k].Remark
		cibl.DepositCount = mbcl[k].DepositCount
		cibl.DepositMoney = mbcl[k].DepositMoney
		cibl.Payee = mbcl[k].Payee
		cibl.BankAccount = mbcl[k].BankAccount
		cibl.OrderNum = mbcl[k].OrderNum
		cibl.OtherDiscount = mbcl[k].OtherDiscount
		//小计
		pageTotalMoney += mbcl[k].DepositMoney
		pageTotalDepositDiscount += mbcl[k].DepositDiscount
		pageTotalOtherDiscount += mbcl[k].OtherDiscount
		pageTotalDeposit += mbcl[k].DepositCount
		if mbcl[k].IsFirstDeposit == 1 {
			cibl.IsFirstDeposit = "否"
		} else if mbcl[k].IsFirstDeposit == 2 {
			cibl.IsFirstDeposit = "是"
		}
		cibl.DepositDiscount = mbcl[k].DepositDiscount
		if mbcl[k].ClientType == 1 {
			cibl.ClientType = "pc"
		} else if mbcl[k].ClientType == 2 {
			cibl.ClientType = "wap"
		} else if mbcl[k].ClientType == 3 {
			cibl.ClientType = "android"
		} else if mbcl[k].ClientType == 4 {
			cibl.ClientType = "ios"
		}
		data = append(data, cibl)
	}

	//获得符合条件的记录数
	count, err := sess.Table(memberCompanyIncome.TableName()).
		Join("LEFT", sbis.TableName(), sql).
		Where(conds).
		Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cibls, err
	}
	cibls.AllData = data
	cibls.TotalMoney = totalMoney
	cibls.TotalDeposit = totalDeposit
	cibls.TotalDepositDiscount = totalDepositDiscount
	cibls.TotalOtherDiscount = totalOtherDiscount
	cibls.PageTotalMoney = pageTotalMoney
	cibls.PageTotalDeposit = pageTotalDeposit
	cibls.PageTotalDepositDiscount = pageTotalDepositDiscount
	cibls.PageTotalOtherDiscount = pageTotalOtherDiscount
	cibls.TotalCount = count
	cibls.PageCount = len(mbcl)
	return cibls, err
}

//根据代理id获取名称
func GetAgencyNameById(ids []int64) (agency []schema.Agency, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	a := new(schema.Agency)
	err = sess.Table(a.TableName()).
		In("id", ids).
		Select("id,username").
		Find(&agency)
	return
}

//获取某个站点下面没有确认得公司入款
func (*MemberCompanyIncomeBean) GetNotConfirm(site, siteindex string) (infolist []schema.MemberCompanyIncome, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newincom := new(schema.MemberCompanyIncome)
	err = sess.Table(newincom.TableName()).Where("site_id=?", site).Where("status=?", 0).Where("site_index_id=?", siteindex).Find(&infolist)
	count, err = sess.Table(newincom.TableName()).Where("site_id=?", site).Where("status=?", 0).Where("site_Index_id=?", siteindex).Count()
	return
}

//获取站点下所有的代理
func (*MemberCompanyIncomeBean) GetAgency(this *input.SiteId) ([]back.GetAgency, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	a := new(schema.Agency)
	var agency []back.GetAgency
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	err := sess.Table(a.TableName()).Where("site_id=?", this.SiteId).
		Where("role_id = ?", 4).Find(&agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return agency, err
	}
	return agency, err
}

//获取站点下所有的收款账号
func (*MemberCompanyIncomeBean) GetSetAccount(this *input.SiteId) (
	[]back.GetSetAgency, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sbis := new(schema.SiteBankIncomeSet)
	var setAccount []back.GetSetAgency
	bank := new(schema.Bank)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sql := fmt.Sprintf("%s.bank_id=%s.id", sbis.TableName(), bank.TableName())
	err := sess.Table(sbis.TableName()).Join("LEFT", bank.TableName(), sql).Where("site_id=?",
		this.SiteId).Find(&setAccount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return setAccount, err
	}
	return setAccount, err
}

//获取站点下所有的收款账号信息
func (*MemberCompanyIncomeBean) GetSetAccountInfo(this *input.SiteId) (setAccount []back.GetPayeeInfo, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sbis := new(schema.SiteBankIncomeSet)
	sbiss := make([]*schema.SiteBankIncomeSet, 0)
	bank := new(schema.Bank)
	banks := make([]schema.Bank, 0)
	payType := new(schema.PaidType)
	payTypes := make([]schema.PaidType, 0)
	sess.Where("type_status=?", 1)
	err = sess.Table(payType.TableName()).Cols("id,paid_type_name,type_status").Find(&payTypes)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(payTypes) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到可用的支付类型")
		return
	}
	ids := make([]int, 0)
	payTypeMap := make(map[int]*schema.PaidType)
	for k, v := range payTypes {
		ids = append(ids, v.Id)
		payTypeMap[v.Id] = &payTypes[k]
	}
	sess.Where("status=?", 1)
	sess.Where("delete_time=?", 0)
	sess.Where("is_out=?", 1)
	sess.In("pay_type_id", ids)
	err = sess.Table(bank.TableName()).Cols("id,title").Find(&banks)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(banks) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到可用的出款银行")
		return
	}
	ids2 := make([]int64, 0)
	bankMap := make(map[int64]string)
	for _, v := range banks {
		ids2 = append(ids2, v.Id)
		bankMap[v.Id] = v.Title
	}
	sess.In("pay_type_id", ids)
	sess.In("bank_id", ids2)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	err = sess.Table(sbis.TableName()).Cols("id,account,open_bank,qr_code,payee,stop_balance,bank_id,pay_type_id,sort").OrderBy("pay_type_id,sort desc").Find(&sbiss)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if len(sbiss) == 0 {
		global.GlobalLogger.Error("error:%s", "未找到本站点可用的出款银行")
		return
	}
	for _, v := range sbiss {
		setAccount = append(setAccount, back.GetPayeeInfo{
			v.Id,
			payTypeMap[int(v.PayTypeId)].PaidTypeName,
			bankMap[v.BankId],
			v.Account,
			v.OpenBank,
			v.Payee,
			v.StopBalance,
			v.BankId,
			v.QrCode,
			int(v.PayTypeId),
			v.Sort})
	}
	return
}

//添加一条公司入款
func (*MemberCompanyIncomeBean) AddCompanyIncome(this *schema.MemberCompanyIncome) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.Table(this.TableName()).InsertOne(this)
	return
}

//查询会员是否首次入款
func (*MemberCompanyIncomeBean) GetInfoByAccount(siteId, siteIndexId, account string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	income := new(schema.MemberCompanyIncome)
	has, err = sess.Table(income.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("account=?", account).
		Where("status=?", 1).Get(income)
	return
}
