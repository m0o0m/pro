package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type PaymentSetBean struct{}

//公共币种列表
func (*PaymentSetBean) PublicPaySetList(this *schema.PublicPaySet, listparam *global.ListParams) (
	[]back.PublicPaySetBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	public_pay := new(schema.PublicPaySet)
	var data []back.PublicPaySetBack
	//获得分页记录
	listparam.Make(sess)
	err := sess.Table(public_pay.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(public_pay.TableName()).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询一条公共币种
func (*PaymentSetBean) OnePublicPaySet(this *input.OnePublicPaySet) (
	*back.PublicPaySet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id != 0 {
		sess.Where("id=?", this.Id)
	}
	data := new(back.PublicPaySet)
	pay_set := new(schema.PublicPaySet)
	has, err := sess.Table(pay_set.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//查询一条公司自设定的支付参数
func (*PaymentSetBean) OnePaymentSet(this *input.PaymentSetAdd) (
	*back.PaymentSet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Title != "" {
		sess.Where("title=?", this.Title)
	}
	sess.Where("delete_time=?", 0)
	pay_set := new(schema.SitePaySet)
	data := new(back.PaymentSet)
	has, err := sess.Table(pay_set.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//添加一条公司自设定支付参数
func (*PaymentSetBean) AddPaySet(this *schema.SitePaySet) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err := sess.Table(this.TableName()).InsertOne(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//公司支付设定列表
func (*PaymentSetBean) PaymentSetList(this *input.PaymentSetList, listparam *global.ListParams) (
	[]back.PaymentSetListBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.PaymentSetListBack
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	site_pay_set := new(schema.SitePaySet)
	err := sess.Table(site_pay_set.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(site_pay_set.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//公司支付设定配置
func (*PaymentSetBean) PaymentSetUp(this *input.PaymentSetUp) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site_pay_set := new(schema.SitePaySet)
	site_pay_set.IsFree = this.IsFree
	site_pay_set.FreeNum = this.FreeNum
	site_pay_set.OutCharge = this.OutCharge
	site_pay_set.OnceQuotaChangeLimmit = this.OnceQuotaChangeLimmit
	site_pay_set.OnlineIsDepositDiscount = this.OnlineIsDepositDiscount
	site_pay_set.OnlineIsDeposit = this.OnlineIsDeposit
	site_pay_set.OnlineDiscountStandard = this.OnlineDiscountStandard
	site_pay_set.OnlineDiscountPercent = this.OnlineDiscountPercent
	site_pay_set.OnlineDepositMax = this.OnlineDepositMax
	site_pay_set.OnlineDepositMin = this.OnlineDepositMin
	site_pay_set.OnlineDiscountUp = this.OnlineDiscountUp
	site_pay_set.OnlineOtherDiscountStandard = this.OnlineOtherDiscountStandard
	site_pay_set.OnlineOtherDiscountPercent = this.OnlineOtherDiscountPercent
	site_pay_set.OnlineOtherDiscountUp = this.OnlineOtherDiscountUp
	site_pay_set.OnlineOtherDiscountUpDay = this.OnlineOtherDiscountUpDay
	site_pay_set.OnlineIsMultipleAudit = this.OnlineIsMultipleAudit
	site_pay_set.OnlineMultipleAuditTimes = this.OnlineMultipleAuditTimes
	site_pay_set.OnlineIsNormalAudit = this.OnlineIsNormalAudit
	site_pay_set.OnlineNormalAuditPercent = this.OnlineNormalAuditPercent
	site_pay_set.LineIsDeposit = this.LineIsDeposit
	site_pay_set.LineDiscountStandard = this.LineDiscountStandard
	site_pay_set.LineDiscountPercent = this.LineDiscountPercent
	site_pay_set.LineDepositMax = this.LineDepositMax
	site_pay_set.LineDepositMin = this.LineDepositMin
	site_pay_set.LineDiscountUp = this.LineDiscountUp
	site_pay_set.LineOtherDiscountStandard = this.LineOtherDiscountStandard
	site_pay_set.LineOtherDiscountPercent = this.LineOtherDiscountPercent
	site_pay_set.LineOtherDiscountUp = this.LineOtherDiscountUp
	site_pay_set.LineOtherDiscountUpDay = this.LineOtherDiscountUpDay
	site_pay_set.LineIsMultipleAudit = this.LineIsMultipleAudit
	site_pay_set.LineMultipleAuditTimes = this.LineMultipleAuditTimes
	site_pay_set.LineIsNormalAudit = this.LineIsNormalAudit
	site_pay_set.LineNormalAuditPercent = this.LineNormalAuditPercent
	site_pay_set.AuditAdminRate = this.AuditAdminRate
	site_pay_set.AuditRelaxQuota = this.AuditRelaxQuota
	sess.Cols("is_free,free_num,out_charge,once_quota_change_limmit," +
		"online_is_deposit_discount,online_is_deposit,online_discount_standard," +
		"online_discount_percent,online_deposit_max,online_deposit_min,online_discount_up," +
		"online_other_discount_standard,online_other_discount_percent,online_other_discount_up," +
		"online_other_discount_up_day,online_is_multiple_audit,online_multiple_audit_times," +
		"online_is_normal_audit,online_normal_audit_percent,line_is_deposit_discount," +
		"line_is_deposit,line_discount_standard,line_discount_percent,line_deposit_max," +
		"line_deposit_min,line_discount_up,line_other_discount_standard," +
		"line_other_discount_percent,line_other_discount_up,line_other_discount_up_day," +
		"line_is_multiple_audit,line_multiple_audit_times,line_is_normal_audit," +
		"line_normal_audit_percent,audit_relax_quota,audit_admin_rate")
	count, err := sess.Table(site_pay_set.TableName()).Where("id=?", this.Id).Update(site_pay_set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改公司支付设置名称
func (*PaymentSetBean) ChangeName(this *input.PaymentChangeName) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site_pay_set := new(schema.SitePaySet)
	site_pay_set.Title = this.Title
	count, err := sess.Table(site_pay_set.TableName()).
		Where("site_index_id=?", this.SiteIndexId).
		Where("site_id=?", this.SiteId).
		Where("id=?", this.Id).Cols("title").
		Update(site_pay_set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询会员层级中公司支付设定是否被使用
func (*PaymentSetBean) PaySetByMemberLevel(this *schema.MemberLevel) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.PaySetId != 0 {
		sess.Where("pay_set_id=?", this.PaySetId)
	}
	has, err := sess.Table(this.TableName()).Where("delete_time=?", 0).Get(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//删除一条支付设定
func (*PaymentSetBean) PaymentSetDelete(this *input.PaymentDelette) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	pay_set := new(schema.SitePaySet)
	pay_set.DeleteTime = time.Now().Unix()
	count, err := sess.Table(pay_set.TableName()).
		Where("site_index_id=?", this.SiteIndexId).
		Where("site_id=?", this.SiteId).
		Where("id=?", this.Id).
		Cols("delete_time").
		Update(pay_set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询一条公司支付设定
func (*PaymentSetBean) GetOnePaySet(this *input.PaymentSetOne) (
	*back.PaymentSet, bool, error) {
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
	pay_set := new(schema.SitePaySet)
	data := new(back.PaymentSet)
	has, err := sess.Table(pay_set.TableName()).Where("delete_time=?", 0).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//公共币种列表
func (*PaymentSetBean) PublicPaySetListCurrency(this *schema.PublicPaySet) (
	[]back.PublicPaySetBackC, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.PublicPaySetBackC
	public_pay := new(schema.PublicPaySet)
	err := sess.Table(public_pay.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
