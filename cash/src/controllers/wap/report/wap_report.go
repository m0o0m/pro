package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//报表统计
type ReportController struct {
	controllers.BaseController
}

//报表统计
func (rc *ReportController) ReportStatistics(ctx echo.Context) error {
	wapReport := new(input.WapReport)
	code := global.ValidRequestMember(wapReport, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	times := new(global.Times)
	times.StartTime, times.EndTime, code = global.FormatDay2Timestamp(wapReport.StartTime, wapReport.EndTime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	wapReport.MemberId = member.Id
	list, err := betRecordInfo.WapReportStatistics(wapReport, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
