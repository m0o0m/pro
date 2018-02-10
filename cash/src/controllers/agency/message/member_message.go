package message

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//会员消息
type MemberMessage struct {
	controllers.BaseController
}

//查询所有会员消息
func (c *MemberMessage) MessageList(ctx echo.Context) error {
	reqData := new(input.MessageList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)

	listParam := new(global.ListParams)
	//获取listParam的数据
	err := c.GetParam(listParam, ctx)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	noticeList, count, err := memberMessageBean.MessageList(user.SiteId, user.SiteIndexId, user.Id, listParam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, noticeList, int64(len(noticeList)), count, ctx))

}

//查询单条会员消息详情
func (c *MemberMessage) MessageDetails(ctx echo.Context) error {
	reqData := new(input.Message)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	notice, err := memberMessageBean.MessageDetails(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(notice))
}

//修改单条消息详情为已读
func (c *MemberMessage) UpdateMessage(ctx echo.Context) error {
	reqData := new(input.Message)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	num, err := memberMessageBean.UpdateMessage(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//删除单条消息
func (c *MemberMessage) DelMessage(ctx echo.Context) error {
	reqData := new(input.Message)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	num, err := memberMessageBean.DelMessage(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
