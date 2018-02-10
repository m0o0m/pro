//[控制器] [平台]注单管理
package note

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

//注单管理
type NoteBetController struct {
	controllers.BaseController
}

//注单列表查询
func (c *NoteBetController) GetNoteBetList(ctx echo.Context) error {
	BetRecordList := new(input.BetRecordList)
	code := global.ValidRequest(BetRecordList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
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

	//获取数据
	data, count, err := betRecordInfoBean.GetBetRecordList(BetRecordList, listParams, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))
}

//注单采集
func (c *NoteBetController) GetNoteBetInsetInto(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//额度转换记录
func (c *NoteBetController) GetVideoQuota(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//额度转换统计查询
func (c *NoteBetController) GetQuotaCount(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//视讯白名单查询
func (c *NoteBetController) GetGameWhite(ctx echo.Context) error {
	gameWhite := new(input.GameWhiteList)
	code := global.ValidRequest(gameWhite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	data, count, err := gameWhitelistBean.GetGameWhiteList(gameWhite, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//视讯白名单添加
func (c *NoteBetController) PostGameWhiteAdd(ctx echo.Context) error {
	gameadd := new(input.GameWhiteAdd)
	code := global.ValidRequest(gameadd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var data schema.GameWhiteList
	data.Ip = gameadd.Ip
	data.Remarks = gameadd.Remarks
	count, err := gameWhitelistBean.GameWhiteAdd(data)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//视讯白名单修改
func (c *NoteBetController) PutGameWhiteUpdate(ctx echo.Context) error {
	gamewhite := new(input.GameWhiteEdit)
	code := global.ValidRequest(gamewhite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := gameWhitelistBean.UpdateInfo(gamewhite)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//视讯白名单删除
func (c *NoteBetController) DelGameWhite(ctx echo.Context) error {
	gamewhite := new(input.GameWhiteDel)
	code := global.ValidRequest(gamewhite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := gameWhitelistBean.DelGameWhite(gamewhite.Id)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
