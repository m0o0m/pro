package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//期数管理
type PeriodsController struct {
	controllers.BaseController
}

//期数管理列表(get列表)
func (pc *PeriodsController) PeriodsList(ctx echo.Context) error {
	combo := new(input.PeriodsGet)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取listparam的数据

	data1 := pc.getSiteList(combo.SiteId, combo.SiteIndexId)
	list, err := periodsBean.PeriodsList(combo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("list", list, "data", data1))
}

//获取单条期数数据
func (pc *PeriodsController) PeriodsGetOne(ctx echo.Context) error {
	list := new(input.PeriodsGetOne)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := periodsBean.PeriodsGetOne(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10113, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//删除单条期数数据
func (pc *PeriodsController) PeriodsDelete(ctx echo.Context) error {
	list := new(input.PeriodsDeleteOne)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := periodsBean.PeriodsDelete(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10112, ctx))
	}
	return ctx.NoContent(204)
}

//增加单条期数数据

func (pc *PeriodsController) PeriodsAdd(ctx echo.Context) error {
	list := new(input.PeriodsAdd)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := periodsBean.PeriodsAdd(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data != 1 {
		return ctx.JSON(200, global.ReplyError(10111, ctx))
	}
	return ctx.NoContent(204)
}

//修改期数信息
func (pc *PeriodsController) PeriodsUpdate(ctx echo.Context) error {
	list := new(input.PeriodsUpdate)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := periodsBean.PeriodsUpdate(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data != 1 {
		return ctx.JSON(200, global.ReplyError(10110, ctx))
	}
	return ctx.NoContent(204)
}

//退佣冲销
func (pc *PeriodsController) Commission(ctx echo.Context) error {
	list := new(input.Commission)
	code := global.ValidRequest(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := periodsBean.Commission(list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10109, ctx))
	}
	return ctx.NoContent(204)
}

//获取站点列表
func (pc *PeriodsController) getSiteList(SiteId string, SiteIndexId string) []back.SiteList {
	list := new(input.GetSiteList)
	list.SiteId = SiteId
	list.SiteIndexId = SiteIndexId
	data, err := periodsBean.GetSiteList(list.SiteId, list.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return data
}
