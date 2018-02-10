package cash

import (
	"controllers"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//额度统计
type QuotaController struct {
	controllers.BaseController
}

//额度统计列表
func (qtc *QuotaController) QuotaList(ctx echo.Context) error {
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取用户参数
	quota := new(input.QuotaCountList)
	code := global.ValidRequest(quota, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if quota.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", quota.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if quota.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", quota.EndTime, loc)
		times.EndTime = et.Unix()
	}
	data, err := quotaCountBean.QuotaCountList(quota, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//充值记录
func (qtc *QuotaController) RechargeRecord(ctx echo.Context) error {
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取用户参数
	quota := new(input.QuotaRecord)
	code := global.ValidRequest(quota, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	qtc.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	if quota.StartTime != "" {
		times.StartTime, _ = global.FormatTime2Timestamp2(quota.StartTime)
	}
	if quota.EndTime != "" {
		times.EndTime, _ = global.FormatTime2Timestamp2(quota.EndTime)
	}
	data, count, err := quotaCountBean.QuotaRecList(quota, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//额度记录列表
func (qtc *QuotaController) QuotaRecordList(ctx echo.Context) error {
	quota := new(input.QuotaRecordList)
	code := global.ValidRequest(quota, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	qtc.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	if quota.StartTime != "" {
		times.StartTime, _ = global.FormatTime2Timestamp2(quota.StartTime)
	}
	if quota.EndTime != "" {
		times.EndTime, _ = global.FormatTime2Timestamp2(quota.EndTime)
	}
	list, count, err := quotaCountBean.QuotaRecordList(quota, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list.QuotaRecord)), count, ctx))
}

//获取视讯下拉框
func (*QuotaController) GetPlatform(ctx echo.Context) error {
	list, err := quotaCountBean.GetPlatform()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//额度充值--银行卡
func (*QuotaController) PutBankCardRecharge(ctx echo.Context) error {
	bankCardRecharge := new(input.BankCardRecharge)
	code := global.ValidRequest(bankCardRecharge, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if bankCardRecharge.PayId <= 0 {
		return ctx.JSON(200, global.ReplyError(60019, ctx))
	}

	user := ctx.Get("user").(*global.RedisStruct)
	bankCardRecharge.SiteId = user.SiteId
	bankCardRecharge.OrderNum = uuid.NewV4().String()
	bankCardRecharge.Type = 2
	count, err := quotaCountBean.BankCardRechargeAdd(bankCardRecharge)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(70033, ctx))
	}
	return ctx.NoContent(204)
}
