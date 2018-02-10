package member

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

type ReportFormController struct {
	controllers.BaseController
}

//报表统计
func (rfc *ReportFormController) Report(ctx echo.Context) error {
	reportForm := new(input.ReportForm)
	code := global.ValidRequest(reportForm, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	if len(reportForm.StartTime) != 0 {
		times.StartTime, code = global.FormatDay2Timestamp2(reportForm.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(reportForm.EndTime) != 0 {
		times.EndTime, code = global.FormatDay2Timestamp2(reportForm.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetCurrentTime()
	}
	//时间区间最多为7天
	diff := times.EndTime - times.StartTime
	day := diff / 3600 / 24
	if day-1 > 7 {
		return ctx.JSON(200, global.ReplyError(30184, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	reportForm.MemberId = member.Id
	reportForm.SiteId = member.Site
	reportForm.SiteIndexId = member.SiteIndex
	//根据站点查询哪些商品被剔除
	productIds, err := otherBean.GetProductDel(reportForm.SiteId, reportForm.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	list, err := reportFormBean.MemberReport(reportForm, productIds, times)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//(未做)本周报表
func (rfc *ReportFormController) ThisWeek(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyError(10000, ctx))
}

//(未做)上周报表
func (rfc *ReportFormController) LastWeek(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyError(10000, ctx))
}
