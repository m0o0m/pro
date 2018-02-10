package rbac

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//交易平台管理
type PlatformController struct {
	controllers.BaseController
}

// 交易平台列表
func (pc *PlatformController) Index(ctx echo.Context) error {
	platform := new(input.PlatformList)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	pc.GetParam(listparam, ctx)
	list, count, err := platformBean.GetPlatformList(platform, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

// 添加交易平台
func (pc *PlatformController) PlatformAdd(ctx echo.Context) error {
	platform := new(input.AddPlatform)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	has, err := platformBean.PlatformIsExist(platform.Platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(80002, ctx))
	}
	count, err := platformBean.AddPlatform(platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(80003, ctx))
	}
	return ctx.NoContent(204)
}

// 交易平台详情
func (pc *PlatformController) EditPlatform(ctx echo.Context) error {
	platform := new(input.OnePlatformInfo)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := platformBean.GetPlatformOne(platform.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

// 修改交易平台
func (pc *PlatformController) PlatformEdit(ctx echo.Context) error {
	platform := new(input.UpdatePlatform)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	has, err := platformBean.PlatformIsExist(platform.Platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(80002, ctx))
	}
	count, err := platformBean.UpdatePlatform(platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(500, global.ReplyError(80004, ctx))
	}
	return ctx.NoContent(204)
}

// 修改交易平台状态
func (pc *PlatformController) PlatformEditStatus(ctx echo.Context) error {
	platform := new(input.UpdatePlatformStatus)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := platformBean.UpdatePlatformStatus(platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(500, global.ReplyError(80004, ctx))
	}
	return ctx.NoContent(204)
}

// 删除交易平台
func (pc *PlatformController) PlatformState(ctx echo.Context) error {
	platform := new(input.DeletePlatform)
	code := global.ValidRequestAdmin(platform, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := platformBean.DeletePlatform(platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(500, global.ReplyError(80005, ctx))
	}
	return ctx.NoContent(204)
}
