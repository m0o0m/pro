//[控制器] [平台] 权限管理（manage中已做，移到此位置）
package rbac

import (
	"controllers"
	"models/function"
)

//角色权限管理
type RbacController struct {
	controllers.BaseController
}

var (
	permissionBean    = new(function.PermissionBean)    //权限
	roleBean          = new(function.RoleBean)          //角色
	menuBean          = new(function.MenuBean)          //菜单
	adminLogBean      = new(function.AdminLogBean)      //操作日志
	roleMenuBean      = new(function.RoleMenuBean)      //角色菜单
	adminBean         = new(function.AdminBean)         //管理员
	adminLoginLogBean = new(function.AdminLoginLogBean) //登录日志
	comboBeen         = new(function.ComboBeen)         //套餐管理
	siteOperateBean   = new(function.SiteOperateBean)   //站点
	productBean       = new(function.ProductBean)       //商品
	productTypeBean   = new(function.ProductTypeBean)   //商品类型
	platformBean      = new(function.PlatformBean)      //会员管理
	reportBean        = new(function.ReportBean)        //数据统计
)
