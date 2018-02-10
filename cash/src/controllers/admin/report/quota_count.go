//[控制器] [平台]额度管理
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//额度管理
type QuotaCountController struct {
	controllers.BaseController
}

//额度统计列表
func (qtc *QuotaCountController) QuotaList(ctx echo.Context) error {
	//获取用户参数
	quota := new(input.QuotaCountList)
	code := global.ValidRequestAdmin(quota, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var sTime, eTime int64
	if quota.StartTime != "" {
		sTime, code = global.FormatTime2Timestamp2(quota.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	if quota.EndTime != "" {
		eTime, code = global.FormatTime2Timestamp2(quota.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	data, err := quotaBean.QuotaCountList(quota, &global.Times{StartTime: sTime, EndTime: eTime})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//额度记录列表
func (qtc *QuotaCountController) QuotaRecordList(ctx echo.Context) error {
	quota := new(input.QuotaRecordList)
	code := global.ValidRequest(quota, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	qtc.GetParam(listparam, ctx)
	var sTime, eTime int64
	if quota.StartTime != "" {
		sTime, code = global.FormatTime2Timestamp2(quota.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	if quota.EndTime != "" {
		eTime, code = global.FormatTime2Timestamp2(quota.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	list, count, err := quotaBean.QuotaRecordList(quota, listparam, &global.Times{StartTime: sTime, EndTime: eTime})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list.QuotaRecord)), count, ctx))
}

//获取视讯下拉框
func (*QuotaCountController) GetPlatform(ctx echo.Context) error {
	list, err := quotaBean.GetPlatform()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//充值记录
func (qtc *QuotaCountController) RechargeRecord(ctx echo.Context) error {
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
	loc, _ := time.LoadLocation("Local")
	if quota.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", quota.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if quota.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", quota.EndTime, loc)
		times.EndTime = et.Unix()
	}
	data, count, err := quotaBean.QuotaRecList(quota, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//充值记录修改
func (*QuotaCountController) RechargeRecordUpdate(ctx echo.Context) error {
	reqData := new(input.SitePayRecordUpdate)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("admin").(*global.AdminRedisStruct)
	reqData.AdminUser = user.Account
	num, err := quotaBean.QuotaRecordUpdate(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//额度记录状态修改
func (c *QuotaCountController) PutQuotaRecordUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//入款管理（第三方和银行一起，type区分 添加）
func (c *QuotaCountController) PostQuotaSetAddOrUpdate(ctx echo.Context) error {
	reqData := new(input.SitePayNameAdd)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if reqData.Type == 1 {
		//第三方
		if reqData.PayType == 0 {
			return ctx.JSON(200, global.ReplyError(30202, ctx))
		}
		if reqData.FUrl == "" {
			return ctx.JSON(200, global.ReplyError(30203, ctx))
		}
	} else {
		//银行卡
		if reqData.MyName == "" {
			return ctx.JSON(200, global.ReplyError(30204, ctx))
		}
		if reqData.Address == "" {
			return ctx.JSON(200, global.ReplyError(30205, ctx))
		}
	}
	count, err := quotaBean.AddOrUpdatePayName(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//入款管理（第三方和银行一起，type区分 查询）
func (c *QuotaCountController) GetQuotaSetList(ctx echo.Context) error {
	reqData := new(input.SitePayNameList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	results, err := quotaBean.GetPayNameList(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(results, len(results)))
}

//支付类型下拉框
func (c *QuotaCountController) ThirdTypeDrop(ctx echo.Context) error {
	results, err := quotaBean.ThirdTypeDrop()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(results))
}
