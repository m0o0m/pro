//[控制器] [平台]权限组管理
package rbac

import (
	"controllers"
	"global"
	"models/function"
	"models/input"

	"github.com/labstack/echo"

	"models/back"
)

//权限组管理
type RbacGroupController struct {
	controllers.BaseController
}

//权限组列表(角色列表)
func (c *RbacGroupController) GetRbacGroupList(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	list, count, err := roleBean.GetList(listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//权限组添加
func (c *RbacGroupController) PostRbacGroupAdd(ctx echo.Context) error {
	role := new(input.RoleAdd)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看角色名是否存在
	has, err := function.IsRoleName(role.RoleName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30075, ctx))
	}
	count, err := roleBean.Add(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30076, ctx))
	}
	return ctx.NoContent(204)
}

//角色详情
func (rc *RbacGroupController) EditRole(ctx echo.Context) error {
	role := new(input.RoleId)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := roleBean.GetInfo(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//权限组修改
func (c *RbacGroupController) PutRbacGroupUpdate(ctx echo.Context) error {
	role := new(input.RoleEditNew)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色为开户人，股东，总代，代理,超级管理员的都不能修改
	if role.Id == 1 || role.Id == 2 || role.Id == 3 || role.Id == 4 || role.Id == 5 {
		return ctx.JSON(200, global.ReplyError(30094, ctx))
	}
	//查询角色是否存在 以及是否能修改
	r, has, err := roleBean.RoleIsOperateNew(role.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	if r.IsOperate == 2 {
		return ctx.JSON(200, global.ReplyError(30094, ctx))
	}
	//查看角色名是否存在
	has, err = function.IsRoleNames(role.Id, role.RoleName)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30075, ctx))
	}
	count, err := roleBean.UpdateInfoNew(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//修改角色状态
func (rc *RbacGroupController) RoleEditStauts(ctx echo.Context) error {
	role := new(input.RoleStatus)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色为开户人，股东，总代，代理的都不能修改
	if role.Id == 1 || role.Id == 2 || role.Id == 3 || role.Id == 4 || role.Id == 5 {
		return ctx.JSON(200, global.ReplyError(30094, ctx))
	}
	//查询角色是否存在，存在的话是否能被修改
	r, has, err := roleBean.RoleIsOperateNew(role.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	if r.IsOperate == 2 {
		return ctx.JSON(200, global.ReplyError(30094, ctx))
	}
	if role.Status == 1 {
		has, err := roleBean.BeOneRole(role.Id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if has {
			return ctx.JSON(200, global.ReplyError(50146, ctx))
		}
	}
	count, err := roleBean.UpdateStatusNew(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除角色
func (rc *RbacGroupController) RoleState(ctx echo.Context) error {
	role := new(input.RoleId)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色为开户人，股东，总代，代理的都不能删除
	if role.Id == 1 || role.Id == 2 || role.Id == 3 || role.Id == 4 || role.Id == 5 {
		return ctx.JSON(200, global.ReplyError(30095, ctx))
	}
	//查询角色是否存在，存在的话是否能被删除
	r, has, err := roleBean.RoleIsOperateNew(role.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	if r.IsOperate == 2 {
		return ctx.JSON(200, global.ReplyError(30094, ctx))
	}
	//角色如果被使用，则不能被删除
	has, err = roleBean.BeOneRole(role.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50145, ctx))
	}
	count, err := roleBean.Delete(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//角色权限列表
func (rc *RbacGroupController) RolePermission(ctx echo.Context) error {
	//获取角色权限
	role := new(input.PromissRoleId)
	code := global.ValidRequestAdmin(role, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限返回列表
	var rolePermisionBack back.RolePermissionBack
	var p back.Permissions
	//权限返回列表
	var permissions []back.Permissions
	//获取module
	module, err := permissionBean.GetModules(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var mdl []string
	for m := range module {
		mdl = append(mdl, module[m].Module)
	}
	list, err := permissionBean.GetListByModule(mdl)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var permission back.Permission
	for k := range module {
		p.Module = module[k].Module
		var ps []back.Permission
		for i := range list {
			if list[i].Module == module[k].Module {
				permission.Id = list[i].Id
				permission.CreateTime = list[i].CreateTime
				permission.PermissionName = list[i].PermissionName
				permission.Route = list[i].Route
				permission.Method = list[i].Method
				permission.Status = list[i].Status
				permission.Type = list[i].Type
				permission.IsPermission = list[i].IsPermission
				ps = append(ps, permission)
			}
		}
		p.Permission = ps
		permissions = append(permissions, p)
	}
	data, err := roleBean.GetPermission(role)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(data) == 0 {
		//没有在角色权限表赋值
		role, err := roleBean.GetRoleInfo(role.Id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		rolePermisionBack.RoleName = role.RoleName
		rolePermisionBack.IsOperate = role.IsOperate
		rolePermisionBack.Status = role.Status
	} else {
		//已经在角色权限表赋值
		for k := range data {
			rolePermisionBack.RoleName = data[k].RoleName
			rolePermisionBack.IsOperate = data[k].IsOperate
			rolePermisionBack.Status = data[k].Status
		}
	}
	for k := range data {
		for i := range permissions {
			for j := range permissions[i].Permission {
				if data[k].PermissionName == permissions[i].Permission[j].PermissionName {
					permissions[i].Permission[j].IsPermission = 1
				}
			}
		}
	}
	rolePermisionBack.Permissions = permissions
	return ctx.JSON(200, global.ReplyItem(rolePermisionBack))
}

//设置角色权限
func (rc *RbacGroupController) RolePermissionDo(ctx echo.Context) error {
	rolePermission := new(input.RolePermission)
	code := global.ValidRequestAdmin(rolePermission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看角色id是否存在
	has, err := function.RoleIsExist(rolePermission.RoleId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	//查看权限id是否存在
	if len(rolePermission.PermissionId) > 0 {
		has, err = function.IsPermissionById(rolePermission.PermissionId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30110, ctx))
		}
	}
	//当角色id>5时才可以修改状态
	if rolePermission.RoleId > 5 {
		if rolePermission.Status != 1 && rolePermission.Status != 2 {
			return ctx.JSON(200, global.ReplyError(30050, ctx))
		}
	}
	count, err := roleBean.UpdatePermission(rolePermission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//设置角色菜单
func (rc *RbacGroupController) RoleMenuDo(ctx echo.Context) error {
	roleMenu := new(input.RoleMenu)
	code := global.ValidRequestAdmin(roleMenu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看角色id是否存在
	has, err := function.RoleIsExist(roleMenu.RoleId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	//查看菜单id是否存在
	has, err = function.IsMenuById(roleMenu.MenuId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30111, ctx))
	}
	count, err := roleBean.UpdateMenu(roleMenu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30084, ctx))
	}
	return ctx.NoContent(204)
}

//平台账号中的角色下拉框
func (rc *RbacGroupController) RoleList(ctx echo.Context) error {
	info, err := roleBean.GetRoleList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//角色菜单列表
func (rc *RbacGroupController) RoleMenu(ctx echo.Context) error {
	var roleMuneBack back.RoleMenus
	roleMenu := new(input.RoleId)
	code := global.ValidRequestAdmin(roleMenu, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取角色信息
	role, _, err := roleBean.GetInfo(roleMenu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	roleMuneBack.RoleName = role.RoleName
	//获取菜单列表
	menuList, err := menuBean.FindMenu(role.RoleMark)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := AndLevel(menuList, 0)
	//根据角色id获取菜单信息
	list, err := roleBean.GetMenuByRoleId(roleMenu)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k := range list {
		for i := range data {
			//第一层菜单
			if list[k].MenuId == data[i].Id {
				data[i].IsMenu = 1
			}
			for j := range data[i].Children {
				//第二层菜单
				if list[k].MenuId == data[i].Children[j].Id {
					data[i].Children[j].IsMenu = 1
				}
				for l := range data[i].Children[j].Children {
					//第三层菜单
					if list[k].MenuId == data[i].Children[j].Children[l].Id {
						data[i].Children[j].Children[l].IsMenu = 1
					}
				}
			}
		}
	}
	roleMuneBack.MenuList = data
	return ctx.JSON(200, global.ReplyItem(roleMuneBack))
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
	var tree = make([]back.Trees, lend)
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
