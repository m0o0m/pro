//[控制器] [平台]站点栏目管理
package rbac

import (
	"controllers"
	"global"
	"models/back"
	"models/input"
	"models/schema"

	"github.com/labstack/echo"
)

//站点栏目管理
type SiteColumnController struct {
	controllers.BaseController
}

//根据登录人角色获取菜单
func (c *SiteColumnController) Admin(ctx echo.Context) error {
	user := ctx.Get("admin").(*global.AdminRedisStruct)
	menu, count, err := roleMenuBean.GetMenuByRoleId(user.RoleId, "admin")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(50012, ctx))
	}
	var data []back.Trees
	if count > 1 {
		data = AndLevel(menu, 0)
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))
}

//站点栏目列表(admin)
func (c *SiteColumnController) GetAdminColumnList(ctx echo.Context) error {
	menu := new(schema.Menu)
	menu.Type = "admin"
	menu_info, err := menuBean.FindAllMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := AndLevel(menu_info, 0)
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点栏目列表(agency)
func (c *SiteColumnController) GetSiteColumnList(ctx echo.Context) error {
	menu := new(schema.Menu)
	menu.Type = "agency"
	menu_info, err := menuBean.FindAllMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := AndLevel(menu_info, 0)
	return ctx.JSON(200, global.ReplyItem(data))
}

//平台栏目添加
func (c *SiteColumnController) PostSiteColumnAdd(ctx echo.Context) error {
	menu := new(input.AddMenu)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
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

//菜单详情
func (amc *SiteColumnController) EditMenu(ctx echo.Context) error {
	menu := new(input.MenuInfo)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	menuinfo, ok, err := menuBean.GetOneMenuByName(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(menuinfo))
}

//平台栏目修改
func (c *SiteColumnController) PutSiteColumnUpdate(ctx echo.Context) error {
	menu := new(input.UpdataMenu)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
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

//平台栏目删除
func (c *SiteColumnController) DelSiteColumn(ctx echo.Context) error {
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

//菜单开启，禁用
func (amc *SiteColumnController) MenuOpenAndClose(ctx echo.Context) error {
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
	//进行禁用操作
	if menu.Status == 2 {
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
	//进行开启操作
	if menu.Status == 1 {
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

//平台栏目私有公有化
func (c *SiteColumnController) PutSiteColumnPrivate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//代理菜单下拉框
func (c *SiteColumnController) MoreMenuDrop(ctx echo.Context) error {
	menu := new(input.MenuInfo)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := menuBean.GetIdMenuName(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//平台栏目一键同步缓存
func (c *SiteColumnController) PutSiteColumnRedis(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
