package rebate

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/function"
	"models/input"
)

//会员推广
type SpreadController struct {
	controllers.BaseController
}

//会员推广设定
func (c *SpreadController) SpreadDo(ctx echo.Context) error {
	data := &input.SpreadSet{}
	code := global.ValidRequest(data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	data.SiteId = user.SiteId
	count, err := spreadBean.AddSpreadSet(data)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(70000, ctx))
	}
	return ctx.NoContent(204)
}

//修改会员推广设定
func (c *SpreadController) SpreadDoSubmit(ctx echo.Context) error {
	data := &input.SpreadEdit{}
	code := global.ValidRequest(data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	data.SiteId = user.SiteId
	count, err := spreadBean.UpdateSpreadSet(data)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(70001, ctx))
	}
	return ctx.NoContent(204)
}

//查询会员推广设置
func (c *SpreadController) SpreadList(ctx echo.Context) error {
	reqData := new(input.SiteId)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	result, err := spreadBean.FindSpreadSet(reqData.SiteId, reqData.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if result.SiteId == "" && result.SiteIndexId == "" {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(result))
}

//查询会员推广信息
func (c *SpreadController) SpreadInfo(ctx echo.Context) error {
	reqData := &input.SpreadInfo{}
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	memberBean := &function.MemberBean{}
	ids, spreadNum, count, err := memberBean.GetSpreadNum(reqData, listParam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	members, err := memberBean.GetSpreadInfo(ids)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	for i, v := range members {
		members[i].Number = spreadNum[v.Id]
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, members, int64(len(members)), count, ctx))
}
