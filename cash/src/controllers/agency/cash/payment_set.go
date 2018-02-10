package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

type PaymentSetController struct {
	controllers.BaseController
}

//公共币种列表
func (psc *PaymentSetController) PublicPaySetList(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	psc.GetParam(listparam, ctx)
	public_pay := new(schema.PublicPaySet)
	data, count, err := paymentSetBean.PublicPaySetList(public_pay, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//公共币种列表
func (psc *PaymentSetController) PublicPaySetListCurrency(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	public_pay := new(schema.PublicPaySet)
	data, err := paymentSetBean.PublicPaySetListCurrency(public_pay)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//查询一条公共支付参数设定
func (psc *PaymentSetController) OnePublicPaymentSet(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_set := new(input.OnePublicPaySet)
	code := global.ValidRequest(pay_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := paymentSetBean.OnePublicPaySet(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//增加一条公司自定设置支付参数
func (psc *PaymentSetController) PaymentSetAdd(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_setting := new(input.PaymentSetAdd)
	code := global.ValidRequest(pay_setting, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询名称是否已被使用
	_, has, err := paymentSetBean.OnePaymentSet(pay_setting)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50111, ctx))
	}
	//查询币种信息
	pay_set := new(input.OnePublicPaySet)
	pay_set.Id = pay_setting.Id
	data, has, err := paymentSetBean.OnePublicPaySet(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50109, ctx))
	}
	site_pay_set := new(schema.SitePaySet)
	site_pay_set.SiteId = user.SiteId
	site_pay_set.SiteIndexId = pay_setting.SiteIndexId
	site_pay_set.Title = pay_setting.Title
	site_pay_set.IsFree = data.IsFree
	site_pay_set.FreeNum = data.FreeNum
	site_pay_set.OutCharge = data.OutCharge
	site_pay_set.OnceQuotaChangeLimmit = data.OnceQuotaChangeLimmit
	site_pay_set.OnlineIsDepositDiscount = data.OnlineIsDepositDiscount
	site_pay_set.OnlineIsDeposit = data.OnlineIsDeposit
	site_pay_set.OnlineDiscountStandard = data.OnlineDiscountStandard
	site_pay_set.OnlineDiscountPercent = data.OnlineDiscountPercent
	site_pay_set.OnlineDepositMax = data.OnlineDepositMax
	site_pay_set.OnlineDepositMin = data.OnlineDepositMin
	site_pay_set.OnlineDiscountUp = data.OnlineDiscountUp
	site_pay_set.OnlineOtherDiscountStandard = data.OnlineOtherDiscountStandard
	site_pay_set.OnlineOtherDiscountPercent = data.OnlineOtherDiscountPercent
	site_pay_set.OnlineOtherDiscountUp = data.OnlineOtherDiscountUp
	site_pay_set.OnlineOtherDiscountUpDay = data.OnlineOtherDiscountUpDay
	site_pay_set.OnlineIsMultipleAudit = data.OnlineIsMultipleAudit
	site_pay_set.OnlineMultipleAuditTimes = data.OnlineMultipleAuditTimes
	site_pay_set.OnlineIsNormalAudit = data.OnlineIsNormalAudit
	site_pay_set.OnlineNormalAuditPercent = data.OnlineNormalAuditPercent
	site_pay_set.LineIsDeposit = data.LineIsDeposit
	site_pay_set.LineDiscountStandard = data.LineDiscountStandard
	site_pay_set.LineDiscountPercent = data.LineDiscountPercent
	site_pay_set.LineDepositMax = data.LineDepositMax
	site_pay_set.LineDepositMin = data.LineDepositMin
	site_pay_set.LineDiscountUp = data.LineDiscountUp
	site_pay_set.LineOtherDiscountStandard = data.LineOtherDiscountStandard
	site_pay_set.LineOtherDiscountPercent = data.LineOtherDiscountPercent
	site_pay_set.LineOtherDiscountUp = data.LineOtherDiscountUp
	site_pay_set.LineOtherDiscountUpDay = data.LineOtherDiscountUpDay
	site_pay_set.LineIsMultipleAudit = data.LineIsMultipleAudit
	site_pay_set.LineMultipleAuditTimes = data.LineMultipleAuditTimes
	site_pay_set.LineIsNormalAudit = data.LineIsNormalAudit
	site_pay_set.LineNormalAuditPercent = data.LineNormalAuditPercent
	site_pay_set.AuditAdminRate = data.AuditAdminRate
	site_pay_set.AuditRelaxQuota = data.AuditRelaxQuota
	count, err := paymentSetBean.AddPaySet(site_pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//公司支付设定列表
func (psc *PaymentSetController) PaymentSetList(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_list := new(input.PaymentSetList)
	pay_list.SiteId = user.SiteId
	listparam := new(global.ListParams)
	//获取listparam的数据
	psc.GetParam(listparam, ctx)
	data, count, err := paymentSetBean.PaymentSetList(pay_list, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//公司支付设定设置
func (psc *PaymentSetController) PaymentSetUp(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_set := new(input.PaymentSetUp)
	code := global.ValidRequest(pay_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := paymentSetBean.PaymentSetUp(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(50177, ctx))
	}
	return ctx.NoContent(204)
}

//公司支付设定修改（修改名称）
func (psc *PaymentSetController) PaymentSetChange(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_set := new(input.PaymentChangeName)
	code := global.ValidRequest(pay_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	pay_set.SiteId = user.SiteId
	count, err := paymentSetBean.ChangeName(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//公司支付设定删除
func (psc *PaymentSetController) PaymentSetDelete(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_set := new(input.PaymentDelette)
	code := global.ValidRequest(pay_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	pay_set.SiteId = user.SiteId
	//查询该支付设定是否已被使用
	mem_level := new(schema.MemberLevel)
	mem_level.PaySetId = pay_set.Id
	has, err := paymentSetBean.PaySetByMemberLevel(mem_level)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50049, ctx))
	}
	//删除
	count, err := paymentSetBean.PaymentSetDelete(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//查询一条公司支付设定
func (psc *PaymentSetController) PaymentSetOne(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_set := new(input.PaymentSetOne)
	code := global.ValidRequest(pay_set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := paymentSetBean.GetOnePaySet(pay_set)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
