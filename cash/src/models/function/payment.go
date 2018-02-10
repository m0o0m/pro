package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type PaymentBean struct{}

//入款银行设定列表
func (*PaymentBean) FindBankInList(this *input.BankInList, listparam *global.ListParams) ([]back.BankInBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank_in := new(schema.SiteBankIncomeSet)
	bank := new(schema.Bank)
	var data []back.BankInBack
	bank_level := new(schema.SiteBankIncomeMemberLevel)
	if this.BankId != 0 {
		sess.Where(bank_in.TableName()+".bank_id=?", this.BankId)
	}
	if this.SiteIndexId != "" {
		sess.Where(bank_in.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where(bank_in.TableName()+".site_id=?", this.SiteId)
	}
	if this.Account != "" {
		sess.Where(bank_in.TableName()+".account=?", this.Account)
	}
	if this.Status != 0 {
		sess.Where(bank_in.TableName()+".status=?", this.Status)
	}
	sess.Where(bank_in.TableName()+".delete_time=?", 0)
	sess.Select("sales_site_bank_income_set.status,sales_site_bank_income_set.remark,sales_site_bank_income_set.id,sales_site_bank_income_set.site_id,sales_site_bank_income_set.site_index_id,sales_site_bank_income_set.stop_balance,sales_site_bank_income_set.site_index_id,sales_site_bank_income_set.bank_id,sales_site_bank_income_set.account,sales_site_bank_income_set.open_bank,sales_site_bank_income_set.payee,sales_bank.title,count(sales_site_bank_income_member_level.set_id)")
	sess.GroupBy("sales_site_bank_income_set.id")
	conds := sess.Conds()
	listparam.Make(sess)
	where1 := fmt.Sprintf("%s.bank_id = %s.id", bank_in.TableName(), bank.TableName())
	where2 := fmt.Sprintf("%s.id = %s.set_id", bank_in.TableName(), bank_level.TableName())
	err := sess.Table(bank_in.TableName()).
		Join("LEFT", bank.TableName(), where1).
		Join("LEFT", bank_level.TableName(), where2).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(bank_in.TableName()).
		Join("LEFT", bank.TableName(), where1).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//添加一条入款银行设定
func (*PaymentBean) Add(this *input.BankInAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	bank_in := new(schema.SiteBankIncomeSet)
	bank_in.SiteIndexId = this.SiteIndexId
	bank_in.SiteId = this.SiteId
	bank_in.BankId = this.BankId
	bank_in.Account = this.Account
	bank_in.StopBalance = this.StopBalance
	bank_in.Status = int8(this.Status)
	bank_in.Payee = this.Payee
	bank_in.Remark = this.Remark
	bank_in.OpenBank = this.OpenBank
	if this.QrCode != "" {
		bank_in.QrCode = this.QrCode
	}
	bank_in.PayTypeId = this.PayTypeId
	count, err := sess.Table(bank_in.TableName()).InsertOne(bank_in)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	level := this.Level
	var bank_level schema.SiteBankIncomeMemberLevel
	var bank_levels []schema.SiteBankIncomeMemberLevel
	for _, v := range level {
		bank_level.SetId = bank_in.Id
		bank_level.LevelId = v
		bank_levels = append(bank_levels, bank_level)
	}
	ses := global.GetXorm().NewSession()
	count, err = ses.Table(bank_level.TableName()).Insert(bank_levels)
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

//查询一条入款银行设定
func (*PaymentBean) GetOnePay(this *input.BankInAdd) (data *schema.SiteBankIncomeSet, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.PayTypeId != 0 {
		sess.Where("pay_type_id=?", this.PayTypeId)
	}
	if this.BankId != 0 {
		sess.Where("bank_id=?", this.BankId)
	}
	sess.Where("delete_time=?", 0)
	data = new(schema.SiteBankIncomeSet)
	has, err = sess.Table(data.TableName()).Get(data)
	return
}

//查询一条入款银行设定的详情
func (*PaymentBean) GetOneBankPaySet(this *input.OneBankPaySet) (*back.OneBankPaySet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	data := new(back.OneBankPaySet)
	bank_in := new(schema.SiteBankIncomeSet)
	has, err := sess.Table(bank_in.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询层级id
func (*PaymentBean) GetLelvelId(id int64) (lelvelId []back.LelvelId, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mbiml := new(schema.SiteBankIncomeMemberLevel)
	sess.Table(mbiml.TableName()).Where("set_id=?", id).Find(&lelvelId)
	return
}

//修改一条入款银行支付设定
func (*PaymentBean) UpdataBankPaySet(this *input.BankInUpdata) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	bank_in := new(schema.SiteBankIncomeSet)
	bank_in.BankId = this.BankId
	bank_in.Account = this.Account
	bank_in.StopBalance = this.StopBalance
	bank_in.Payee = this.Payee
	bank_in.Remark = this.Remark
	bank_in.OpenBank = this.OpenBank
	if this.QrCode != "" {
		bank_in.QrCode = this.QrCode
	} else {
		bank_in.QrCode = ""

	}
	bank_in.PayTypeId = this.PayTypeId
	count, err := sess.Where("id=?", this.Id).Cols("pay_type_id,bank_id,account,open_bank,payee,stop_balance,qr_code,remark").Update(bank_in)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	ses := global.GetXorm().NewSession()
	var bank_level schema.SiteBankIncomeMemberLevel
	count, err = ses.Table(bank_level.TableName()).Where("set_id=?", this.Id).Delete(bank_level)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	level := this.Level
	var bank_levels []schema.SiteBankIncomeMemberLevel
	for _, v := range level {
		bank_level.SetId = this.Id
		bank_level.LevelId = v
		bank_levels = append(bank_levels, bank_level)
	}
	sec := global.GetXorm().NewSession()
	count, err = sec.Table(bank_level.TableName()).Insert(bank_levels)
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

//修改状态
func (*PaymentBean) ChangeStatus(this *input.UpdataStatus) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank_in := new(schema.SiteBankIncomeSet)
	if this.Status == 1 {
		bank_in.Status = 2
	} else if this.Status == 2 {
		bank_in.Status = 1
	}
	count, err := sess.Table(bank_in.TableName()).
		Where("id=?", this.Id).Cols("status").Update(bank_in)
	if err != nil {
		return count, err
	}
	return count, err
}

//删除一条入款银行设定
func (*PaymentBean) DeteleOnePaySet(this *input.DeletePaySet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	bank_in := new(schema.SiteBankIncomeSet)
	bank_in.DeleteTime = time.Now().Unix()
	count, err := sess.Table(bank_in.TableName()).Where("id=?", this.Id).Update(bank_in)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	bank_level := new(schema.SiteBankIncomeMemberLevel)
	count, err = sess.Table(bank_level.TableName()).
		Where("set_id=?", this.Id).Delete(bank_level)
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

//查询存款记录
func (*PaymentBean) CheckingDepositRecords(this *input.DepositRecord, listparam *global.ListParams, times *global.Times) (
	[]back.DepositRecordBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.DepositRecordBack
	if this.SetId != 0 {
		sess.Where("set_id=?", this.SetId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.OrderNum != "" {
		sess.Where("order_num=?", this.OrderNum)
	}
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	//sess.Select("id,deposit_money,order_num,account,remark,create_time")
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	mem_company := new(schema.MemberCompanyIncome)
	m_company := make([]schema.MemberCompanyIncome, 0)
	//根据分页查询数据
	err := sess.Table(mem_company.TableName()).Find(&m_company)
	fmt.Println("err:", err, m_company)
	if len(m_company) <= 0 {
		return nil, 0, err
	}
	//统计总数
	count, err := sess.Table(mem_company.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	//总计
	ses := global.GetXorm().NewSession()
	ses.Select("SUM(deposit_money) as a")
	var alltotal back.AllTotal
	_, err = ses.Table(mem_company.TableName()).Where("set_id=?", this.SetId).Get(&alltotal)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	var i float64
	i = 0
	for _, v := range m_company {
		i = i + v.DepositMoney
	}
	alltotal.Subtotal = i
	for _, v := range m_company {
		var da back.DepositRecordBack
		da.Id = v.Id
		da.DepositMoney = v.DepositMoney
		da.Remark = v.Remark
		da.CreateTime = v.CreateTime
		da.OrderNum = v.OrderNum
		da.UserName = v.Account
		da.Atotal = append(da.Atotal, alltotal)
		data = append(data, da)
	}
	return data, count, err
}

//适用层级
func (*PaymentBean) TopClass(this *input.ApplicationLevel) ([]back.ApplicationLevelBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	ml := new(schema.SiteBankIncomeMemberLevel)
	mml := new(schema.MemberLevel)
	var data []back.ApplicationLevelBack
	if this.SiteId != "" {
		sess.Where(mml.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(mml.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	err := sess.Table(ml.TableName()).
		Join("LEFT", mml.TableName(), ml.TableName()+".level_id="+mml.TableName()+".level_id").
		Where("set_id=?", this.SetId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
