//[控制器] [平台]管理员管理
package rbac

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//管理员管理
type RbacAdminController struct {
	controllers.BaseController
}

//管理员列表
func (c *RbacAdminController) GetRbacAdminList(ctx echo.Context) error {
	admin := new(input.AdminList)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	list, count, err := adminBean.GetList(admin, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))

}

//管理员添加
func (c *RbacAdminController) PostRbacAdminAdd(ctx echo.Context) error {
	admin := new(input.AdminAdd)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看账号是否存在
	has, err := adminBean.GetAccount(admin.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//两次密码不一致
	if admin.Password != admin.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr(admin.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	admin.Password = md5Password
	//角色id不能小于5
	if admin.RoleId < 5 {
		return ctx.JSON(200, global.ReplyError(30077, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30080, ctx))
	}
	//查询该角色id是否存在
	has, err = adminBean.BeRoleId(admin.RoleId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	count, err := adminBean.Add(admin)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30096, ctx))
	}
	return ctx.NoContent(204)
}

//管理员修改
func (c *RbacAdminController) PutRbacAdminUpdate(ctx echo.Context) error {
	admin := new(input.AdminEditNew)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//根据id查看管理员是否存在
	has, err := adminBean.GetAccountById(admin.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	//两次密码不一致
	if admin.Password != admin.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	if admin.Password != "" {
		md5Password, err := global.MD5ByStr(admin.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		admin.Password = md5Password
	}
	count, err := adminBean.UpdateInfoNew(admin)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//获取账号信息(GET)
func (ac *RbacAdminController) EditAccount(ctx echo.Context) error {
	admin := new(input.AdminId)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := adminBean.GetInfo(admin)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30097, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//管理员状态修改
func (c *RbacAdminController) PutRbacAdminStatusUpdate(ctx echo.Context) error {
	admin := new(input.AdminStatus)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//id为1的超级管理员不能禁用
	if admin.Id == 1 {
		return ctx.JSON(200, global.ReplyError(30042, ctx))
	}
	//根据id查看管理员是否存在
	has, err := adminBean.GetAccountById(admin.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	count, err := adminBean.UpdateStatusNew(admin)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除账号(DELETE)
func (ac *RbacAdminController) DeleteAccountDelete(ctx echo.Context) error {
	admin := new(input.AdminId)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//id为1的超级管理员不能删除
	if admin.Id == 1 {
		return ctx.JSON(200, global.ReplyError(30042, ctx))
	}
	//根据id查看管理员是否存在
	has, err := adminBean.GetAccountById(admin.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	count, err := adminBean.Delete(admin)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//初始化密码
func (*RbacAdminController) PutInitPassword(ctx echo.Context) error {
	admin := new(input.InitPassword)
	code := global.ValidRequestAdmin(admin, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr("123456", global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	//查看管理员id是否存在
	has, err := adminBean.GetAccountById(admin.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30044, ctx))
	}
	count, err := adminBean.InitPassword(admin.Id, md5Password)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}
