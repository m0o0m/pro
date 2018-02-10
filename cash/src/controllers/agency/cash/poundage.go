package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//手续费设定
type PoundageController struct {
	controllers.BaseController
}

//手续费设定(get列表)
func (pc *PoundageController) Poundage(ctx echo.Context) error {
	combo := new(input.GetSiteList)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取listparam的数据
	data := pc.getSiteList(combo.SiteId, combo.SiteIndexId)
	list, has, err := poundageBean.PoundageGetOne(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10201, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10201, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("list", list, "data", data))
}

//增加手续费设定
func (pc *PoundageController) PoundAdd(ctx echo.Context) error {
	list := new(input.PoundAdd)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err1 := poundageBean.CheckSet(list)
	if err1 != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 1 {
		return ctx.JSON(200, global.ReplyError(10202, ctx))
	}
	count, err := poundageBean.PeriodsAdd(list)

	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10203, ctx))
	}
	return ctx.NoContent(204)
}

//修改手续费设定
func (pc *PoundageController) PoundUpdate(ctx echo.Context) error {
	list := new(input.Poundage)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := poundageBean.PoundUpdate(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data != 1 {
		return ctx.JSON(200, global.ReplyError(10204, ctx))
	}
	return ctx.NoContent(204)
}

//获取单站点手续费设定
func (pc *PoundageController) GetList(ctx echo.Context) error {
	list := new(input.GetSiteList)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := poundageBean.PoundageGetOne(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取站点列表
func (pc *PoundageController) getSiteList(SiteId string, SiteIndexId string) []back.SiteList {
	list := new(input.GetSiteList)
	list.SiteId = SiteId
	list.SiteIndexId = SiteIndexId
	data, err := periodsBean.GetSiteList(list.SiteId, list.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return data
}
