package note

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

type BetRecordInfoController struct {
	controllers.BaseController
}

//查询注单数据
func (not *BetRecordInfoController) GetBeRecordList(ctx echo.Context) error {
	BetRecordList := new(input.BetRecordList)
	code := global.ValidRequest(BetRecordList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//var user *global.RedisStruct
	BetRecordList.SiteId = user.SiteId
	listParams := new(global.ListParams)
	not.GetParam(listParams, ctx)
	//获取数据
	times := new(global.Times)
	if len(BetRecordList.StartTime) != 0 {
		times.StartTime, code = global.FormatTime2Timestamp2(BetRecordList.StartTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.StartTime = global.GetTimeStart(global.GetCurrentTime())
	}
	if len(BetRecordList.EndTime) != 0 {
		times.EndTime, code = global.FormatTime2Timestamp2(BetRecordList.EndTime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		times.EndTime = global.GetCurrentTime()
	}
	data, count, err := betRecordInfoBean.GetBetRecordList(BetRecordList, listParams, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))

}
