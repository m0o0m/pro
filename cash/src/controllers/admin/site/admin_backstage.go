//[控制器] [平台]客户后台栏目管理
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//客户后台栏目管理
type AdminBackstargController struct {
	controllers.BaseController
}

//客户后台栏目列表查询
func (c *AdminBackstargController) GetAdminBackstargList(ctx echo.Context) error {
	menu := new(schema.Menu)
	menu.Type = "admin"
	menu_info, err := menuBean.FindAllMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := AndLevel(menu_info, 0)
	return ctx.JSON(200, global.ReplyItem(data))
}

//客户后台栏目添加
func (c *AdminBackstargController) PostAdminBackstargAdd(ctx echo.Context) error {
	menu := new(input.AddMenu)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	menu.Type = "admin"
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

//客户后台栏目修改
func (c *AdminBackstargController) PutAdminBackstargUpdate(ctx echo.Context) error {
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
	//修改
	count, err := menuBean.UpdataMenu(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//客户后台栏目删除
func (c *AdminBackstargController) PutAdminBackstargDel(ctx echo.Context) error {
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
	//查询该id的下级id,如果存在下级菜单不允许删除
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

//客户后台栏目开启禁用
func (c *AdminBackstargController) MenuOpenAndClose(ctx echo.Context) error {
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

//菜单下拉框
func (c *AdminBackstargController) GetMoreMenu(ctx echo.Context) error {
	menu := new(input.MenuInfo)
	code := global.ValidRequestAdmin(menu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	menu.Type = "admin"
	data, err := menuBean.GetIdMenuName(menu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//递归，菜单无限级目录树
func AndLevel(data []back.MenuListBack, parentid int64) []back.Trees {
	//递归调用当所有的循环没有完成的时候是没有进行child的存值操作
	var lend = 0
	var x = 0
	//这里是为了计算我存储数据的slice的长度
	for _, v := range data {
		if v.ParentId == parentid {
			lend = lend + 1
		}
	}
	//这里根据上面取得的长度定义slice
	var tree []back.Trees = make([]back.Trees, lend)
	if lend != 0 {
		for k, v := range data {
			//这里的k是不定的，所以需要定义另外的累加值进行累加计数
			//将计数累加放在这里会导致数组越界，因为没有满足条件，循环次数会超过上面定义的slice的长度
			if v.ParentId == parentid {
				k = x
				x = x + 1
				//满足条件赋值
				tree[k].MenuName = v.MenuName
				tree[k].Icon = v.Icon
				tree[k].Sort = v.Sort
				tree[k].Route = v.Route
				tree[k].Id = v.Id
				tree[k].Status = v.Status
				tree[k].Type = v.Type
				tree[k].Level = v.Level
				tree[k].LanguageKey = v.LanguageKey
				//下级菜单的个数不定所以这里更改id值和层级 循环再次调用自己
				child := AndLevel(data, v.Id)
				//将取出来的值赋值给子项
				tree[k].Children = child
			}
		}
	}
	return tree
}
