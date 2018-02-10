//[控制器] [平台]数据中心管理
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//数据中心管理
type DataCenterController struct {
	controllers.BaseController
}

//统计数据查询
func (c *DataCenterController) GetDataCenterList(ctx echo.Context) error {
	betReportAccount := new(input.BetReportAccount)
	code := global.ValidRequest(betReportAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if betReportAccount.SiteId == "" {
		data := new(back.BetReportAccountAllBack)
		return ctx.JSON(200, global.ReplyCollections("data", data, "count", 0))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)

	times := new(global.Times)
	if len(betReportAccount.StartTime) != 0 {
		times.StartTime, code = global.FormatTime2Timestamp2(betReportAccount.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(betReportAccount.EndTime) != 0 {
		times.EndTime, code = global.FormatTime2Timestamp2(betReportAccount.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetCurrentTime()
	}
	//获取数据
	data, count, err := reportBean.GetCenter(betReportAccount, listParams, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data.BetReportAccount)), count, ctx))
}

//重新统计
func (c *DataCenterController) PostDataCenterUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//定时统计开关控制
func (c *DataCenterController) PutDataCenterSwitch(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
