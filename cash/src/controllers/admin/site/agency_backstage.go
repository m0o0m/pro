//[控制器] [平台]代理后台栏目管理
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

//代理后台栏目管理
type AgencyBackstargController struct {
	controllers.BaseController
}

//代理后台栏目列表查询
func (c *AgencyBackstargController) GetAgencyBackstargList(ctx echo.Context) error {
	menu := new(schema.Menu)
	menu.Type = "agency"
	menu_info, err := menuBean.FindAllMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := AndLevel(menu_info, 0)
	return ctx.JSON(200, global.ReplyItem(data))
}

//代理后台栏目添加
func (c *AgencyBackstargController) PostAgencyBackstargAdd(ctx echo.Context) error {
	menu := new(input.AddMenu)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	menu.Type = "agency"
	//menuinfo := new(input.MenuInfo)
	//menuinfo.MenuName = menu.MenuName
	//menuinfo.VType = menu.VType
	//检验菜单名称是否存在
	//_, has, err := menuBean.GetOneMenuByName(menuinfo)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if has {
	//	return ctx.JSON(200, global.ReplyError(50051, ctx))
	//}
	count, err := menuBean.AddMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//代理后台栏目修改
func (c *AgencyBackstargController) PutAgencyBackstargUpdate(ctx echo.Context) error {
	menu := new(input.UpdataMenu)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询该栏目是否存在
	has, err := menuBean.BeOneMenu(menu.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30111, ctx))
	}
	count, err := menuBean.UpdataMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//代理后台栏目删除
func (c *AgencyBackstargController) PutAgencyBackstargDel(ctx echo.Context) error {
	menu := new(input.MenuOpenClose)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询该栏目是否存在
	has, err := menuBean.BeOneMenu(menu.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30111, ctx))
	}
	//查询该id的下级id
	data, err := menuBean.GetNextIdById(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data != nil {
		return ctx.JSON(200, global.ReplyError(50080, ctx))
	}
	count, err := menuBean.MenuDeleteTime(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(6000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//代理后台栏目开启禁用
func (c *AgencyBackstargController) MenuOpenAndCloseAgency(ctx echo.Context) error {
	menu := new(input.MenuOpenClose)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询该栏目是否存在
	has, err := menuBean.BeOneMenu(menu.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30111, ctx))
	}
	//菜单状态为1，进行禁用操作
	if menu.Status == 1 {
		//查询该id的下级id
		data, err := menuBean.GetNextIdById(menu)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if len(data) > 0 {
			return ctx.JSON(200, global.ReplyError(50079, ctx))
		}
		count, err := menuBean.CloseMenu(menu)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(6000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(50173, ctx))
		}
	}
	//菜单状态为2，进行开启操作
	if menu.Status == 2 {
		count, err := menuBean.OpenMenu(menu)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(6000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(50173, ctx))
		}
	}
	return ctx.NoContent(204)
}

//代理菜单下拉框
func (c *AgencyBackstargController) GetMoreMenuAgency(ctx echo.Context) error {
	menu := new(input.MenuInfo)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	menu.Type = "agency"
	data, err := menuBean.GetIdMenuName(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
