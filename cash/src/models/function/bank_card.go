package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type BankCardBean struct{}

//入款银行(get列表)
/*
先查询bank表
再查询剔除表
遍历改状态
*/
func (*BankCardBean) GetAllBank(this *input.InComeList, listparam *global.ListParams) ([]back.BankTwoBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.BankTwoBack
	if this.IsIncome != 0 {
		sess.Where("is_income=?", 1)
	}
	if this.BankName != "" {
		sess.Where("title=?", this.BankName)
	}
	listparam.Make(sess)
	conds := sess.Conds()
	bank := new(schema.Bank)
	//查询所有银行
	err := sess.Table(bank.TableName()).Where("status=?", 1).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, 0, err
	}
	if data == nil {
		return nil, 0, err
	}
	//统计银行数量
	count, err := sess.Table(bank.TableName()).Where("status=?", 1).
		Where("delete_time=?", 0).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	ses := global.GetXorm().NewSession()
	if this.SiteId != "" {
		ses.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		ses.Where("site_index_id=?", this.SiteIndexId)
	}
	bank_income_del := new(schema.BankIncomeDel)
	var income_del []schema.BankIncomeDel
	//查询入款银行提出表中的数据
	err = ses.Table(bank_income_del.TableName()).Find(&income_del)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	if income_del == nil {
		return data, count, err
	}
	for k, v := range data {
		for _, vs := range income_del {
			if v.Id == vs.BankId {
				data[k].Status = 2
			}
		}
	}
	return data, count, err
}

//剔除(入款)
func (*BankCardBean) RejectBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	income_bank := new(schema.BankIncomeDel)
	income_bank.SiteId = this.SiteId
	income_bank.SiteIndexId = this.SiteIndexId
	income_bank.BankId = this.Id
	count, err := sess.Table(income_bank.TableName()).InsertOne(income_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//开启（入款）
func (*BankCardBean) OpenBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	income_bank := new(schema.BankIncomeDel)
	count, err := sess.Table(income_bank.TableName()).
		Where("bank_id=?", this.Id).
		Delete(income_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//出款银行
func (*BankCardBean) GetAllBankOut(this *input.IsOutList, listparam *global.ListParams) ([]back.BankTwoBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.BankTwoBack
	if this.Isout != 0 {
		sess.Where("is_out=?", 1)
	}
	if this.BankName != "" {
		sess.Where("title=?", this.BankName)
	}
	listparam.Make(sess)
	conds := sess.Conds()
	bank := new(schema.Bank)
	//查询所有银行
	err := sess.Table(bank.TableName()).Where("status=?", 1).Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	if data == nil {
		return nil, 0, err
	}
	//统计银行数量
	count, err := sess.Table(bank.TableName()).Where("status=?", 1).Where("delete_time=?", 0).
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	bank_out_del := new(schema.BankOutDel)
	var out_del []schema.BankOutDel
	//查询入款银行提出表中的数据
	err = sess.Table(bank_out_del.TableName()).Find(&out_del)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	if out_del == nil {
		return data, count, err
	}
	for k, v := range data {
		for _, vs := range out_del {
			if v.Id == vs.BankId {
				data[k].Status = 2
			}
		}
	}
	return data, count, err
}

//前台获取出款银行列表
func (*BankCardBean) IndexAllBankOut(siteId, siteIndexId string) (data []back.BankTwoBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	outBankDel := new(schema.BankOutDel)
	var outBankList []schema.BankOutDel
	var notInStr []int64
	err = sess.Table(outBankDel.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Find(&outBankList)
	for _, v := range outBankList {
		notInStr = append(notInStr, v.BankId)
	}
	bank := new(schema.Bank)
	//sess.NotIn("id", "select bank_id from sales_bank_out_del WHERE site_id = "+siteId+" AND sitr_index_id = "+siteIndexId)
	//查询所有银行
	err = sess.Table(bank.TableName()).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		NotIn("id", notInStr).
		Find(&data)
	return
}

//剔除(出款)
func (*BankCardBean) RejectOutBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	out_bank := new(schema.BankOutDel)
	out_bank.SiteId = this.SiteId
	out_bank.SiteIndexId = this.SiteIndexId
	out_bank.BankId = this.Id
	count, err := sess.Table(out_bank.TableName()).InsertOne(out_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//开启(出款)
func (*BankCardBean) OpenOutBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	out_bank := new(schema.BankOutDel)
	count, err := sess.Table(out_bank.TableName()).Where("bank_id=?", this.Id).Delete(out_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//三方银行
func (*BankCardBean) GetAllBankThird(this *input.IsThirdList) ([]schema.BankOutDel, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var third_del []schema.BankOutDel
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	bank_third_del := new(schema.BankThirdDel)
	//查询入款银行提出表中的数据
	err := sess.Table(bank_third_del.TableName()).Find(&third_del)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return third_del, err
	}
	return third_del, err
}

//剔除(出款)
func (*BankCardBean) RejectThirdBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	third_bank := new(schema.BankThirdDel)
	third_bank.SiteId = this.SiteId
	third_bank.SiteIndexId = this.SiteIndexId
	third_bank.BankId = this.Id
	count, err := sess.Table(third_bank.TableName()).InsertOne(third_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//开启(出款)
func (*BankCardBean) OpenThirdBank(this *input.OpenAndRejectBank) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	third_bank := new(schema.BankThirdDel)
	count, err := sess.Table(third_bank.TableName()).Where("bank_id=?", this.Id).
		Delete(third_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询银行列表(admin)
func (*BankCardBean) BankCardList(this *input.BankCardList, listparams *global.ListParams) (data []back.BankCardList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.BankName != "" {
		sess.Where("title=?", this.BankName)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	listparams.Make(sess)
	bank := new(schema.Bank)
	err = sess.Table(bank.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(bank.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//根据银行名称查询该银行在bank表中是否已经存在
func (*BankCardBean) BeOneBankCardByName(this *input.BankCardAdd) (data *schema.Bank, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Title != "" {
		sess.Where("title=?", this.Title)
	}
	sess.Where("delete_time=?", 0)
	data = new(schema.Bank)
	has, err = sess.Table(data.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//添加银行卡(admin)
func (*BankCardBean) BankCardAdd(this *input.BankCardAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	bank.Title = this.Title
	bank.Status = this.Status
	bank.BankWebsiteUrl = this.BankWebsiteUrl
	//bank.Icon = this.Icon
	bank.IsIncome = this.IsIncome
	bank.IsOut = this.IsOut
	bank.PayTypeId = 1
	count, err := sess.InsertOne(bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改银行
func (*BankCardBean) BankCardChange(this *input.BankCardChange) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	bank.Title = this.Title
	bank.Status = this.Status
	bank.BankWebsiteUrl = this.BankWebsiteUrl
	//bank.Icon = this.Icon
	bank.IsIncome = this.IsIncome
	bank.IsOut = this.IsOut
	count, err = sess.Table(bank.TableName()).Where("id=?", this.Id).
		Cols("title,icon,is_income,is_out,status,bank_website_url").
		Update(bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改银行状态
func (*BankCardBean) BankCardStatus(this *input.BankCardStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	bank.Status = this.Status
	sess.Table(bank.TableName())
	sess.Where("delete_time=?", 0)
	count, err = sess.Where("id=?", this.Id).
		Cols("status").Update(bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//银行卡删除
func (*BankCardBean) BankCardDelete(this *input.BankCardDelete) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	bank.DeleteTime = time.Now().Unix()
	count, err = sess.Table(bank.TableName()).Where("id=?", this.Id).
		Cols("delete_time").Update(bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//判断银行是否存在
func (*BankCardBean) BeBankCard(id int64) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	has, err = sess.Table(bank.TableName()).Where("delete_time=?", 0).Where("id=?", id).
		Get(bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查询银行列表(admin)不分页
func (*BankCardBean) BankCardRedis() (data []schema.Bank, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	sess.Where("delete_time=?", 0)
	err = sess.Table(bank.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//根据Id判断是否为入款银行
func (*BankCardBean) DepositById(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	bank := new(schema.Bank)
	sess.Table(bank.TableName())
	sess.Where("delete_time=?", 0).
		Where("is_income=?", 1)
	have, err := sess.Get(bank)
	return have, err
}

//wap入款银行列表下拉框
func (*BankCardBean) GetAllIncomeBank(siteId, siteIndexId string) (data []back.SiteBank, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	var incomeDel []int64
	err = sess.Table(new(schema.BankIncomeDel).TableName()).Select("bank_id").Find(&incomeDel)
	if err != nil {
		return
	}
	sess.Where("is_income=?", 1)
	bank := new(schema.Bank)
	//查询所有银行
	err = sess.Table(bank.TableName()).
		Select("id,title").
		NotIn("id", incomeDel).
		Where("status=?", 1).
		Where("delete_time=?", 0).
		Find(&data)
	return
}

//银行下拉框
func (*BankCardBean) BankCardListDrop() ([]back.BankCardListDrop, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("delete_time=?", 0)
	sess.Where("status=?", 1)
	sess.Where("is_out=?", 1)
	bank := new(schema.Bank)
	var data []back.BankCardListDrop
	err := sess.Table(bank.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//银行列表/剔除的银行
func (*BankCardBean) BankCardListDropDel(this *input.AgencyBankOutByDrop) ([]back.BankCardListDropOutDel, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	bd := new(schema.BankOutDel)
	var data []back.BankCardListDropOutDel
	err := sess.Table(bd.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//登陆人信息
func (*BankCardBean) SiteIndexIdByAgencyId(id int64, siteId string) (*schema.Agency, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	sess.Where("site_id=?", siteId)
	ag := new(schema.Agency)
	has, err := sess.Get(ag)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ag, has, err
	}
	return ag, has, err
}

//会员信息
func (*BankCardBean) SiteIndexIdByMemberId(id int64, siteId string) (*schema.Member, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", id)
	sess.Where("delete_time=?", 0)
	sess.Where("site_id=?", siteId)
	mB := new(schema.Member)
	has, err := sess.Get(mB)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return mB, has, err
	}
	return mB, has, err
}
