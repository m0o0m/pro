//[控制器] [平台]栏目管理
package rbac

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//权限管理
type ColumnController struct {
	controllers.BaseController
}

//权限列表
func (c *ColumnController) GetColumnList(ctx echo.Context) error {
	permission := new(input.PermissionList)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	list, count, err := permissionBean.GetLists(permission, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))

}

//权限添加
func (c *ColumnController) PostColumnAdd(ctx echo.Context) error {
	permission := new(input.PermissionAdd)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//一个模块下的权限名不能重复
	has, err := permissionBean.GetPermissionName(permission.Module, permission.PermissionName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30093, ctx))
	}
	//请求方法只能是GET,POST,PUT,DELETE
	if !(permission.Method == "GET" || permission.Method == "POST" ||
		permission.Method == "PUT" || permission.Method == "DELETE") {
		return ctx.JSON(200, global.ReplyError(30087, ctx))
	}
	//请求路由和请求方法的唯一校验
	has, err = permissionBean.GetRouteAndMethod(permission.Module, permission.Route, permission.Method)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30119, ctx))
	}
	count, err := permissionBean.Add(permission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30089, ctx))
	}
	return ctx.NoContent(204)
}

//权限详情
func (c *ColumnController) EditPermission(ctx echo.Context) error {
	permission := new(input.PermissionId)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := permissionBean.GetInfo(permission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30090, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//权限修改
func (c *ColumnController) PutColumnUpdate(ctx echo.Context) error {
	permission := new(input.PermissionUpdate)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询id是否存在
	has, err := permissionBean.BeOnePermission(permission.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30110, ctx))
	}
	//一个模块下的权限名不能重复
	has, err = permissionBean.GetPermissionNames(permission.Id, permission.Module, permission.PermissionName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30093, ctx))
	}
	//请求方法只能是GET,POST,PUT,DELETE
	if !(permission.Method == "GET" || permission.Method == "POST" ||
		permission.Method == "PUT" || permission.Method == "DELETE") {
		return ctx.JSON(200, global.ReplyError(30087, ctx))
	}
	//请求路由和请求方法的唯一校验
	has, err = permissionBean.GetRouteAndMethods(permission.Id, permission.Module, permission.Route, permission.Method)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30119, ctx))
	}
	count, err := permissionBean.ColumnUpdate(permission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//权限删除
func (c *ColumnController) DelColumn(ctx echo.Context) error {
	permission := new(input.ColumnDelete)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询id是否存在
	has, err := permissionBean.BeOnePermission(permission.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30110, ctx))
	}
	//删除
	count, err := permissionBean.ColumnDelete(permission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//修改权限状态
func (c *ColumnController) PermissionEditStauts(ctx echo.Context) error {
	permission := new(input.ColumnStatus)
	code := global.ValidRequestAdmin(permission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询id是否存在
	has, err := permissionBean.BeOnePermission(permission.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30110, ctx))
	}
	//修改状态
	count, err := permissionBean.ColumnStatus(permission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}
