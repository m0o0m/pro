package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//红包样式
type RedPacketStyleController struct {
	controllers.BaseController
}

//添加或更新红包样式
func (*RedPacketStyleController) AddOrUpdate(ctx echo.Context) error {
	reqData := new(input.RedPacketStyleAdd)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	reqData.SiteId = user.SiteId
	count, err := redPacketStyleBean.AddOrUpdate(reqData)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 && reqData.Id == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	} else if count == 0 && reqData.Id != 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//红包样式列表
func (*RedPacketStyleController) FindList(ctx echo.Context) error {
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	redPacketStyles, err := redPacketStyleBean.List(user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(redPacketStyles, len(redPacketStyles)))
}

//红包样式下拉框
func (*RedPacketStyleController) FindListDrop(ctx echo.Context) error {
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	data, err := redPacketStyleBean.RedStyleDrop(user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//红包样式详情
func (*RedPacketStyleController) GetDetails(ctx echo.Context) error {
	reqData := new(input.RedPacketStyle)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	redPacketStyle, err := redPacketStyleBean.Details(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(redPacketStyle))
}

//删除红包样式
func (*RedPacketStyleController) Del(ctx echo.Context) error {
	reqData := new(input.RedPacketStyle)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	num, err := redPacketStyleBean.Del(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//红包样式下拉框
func (*RedPacketStyleController) FindListDropPicture(ctx echo.Context) error {
	reqData := new(input.RedPacketStyle)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := redPacketStyleBean.RedStyleDropPicture(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50171, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
