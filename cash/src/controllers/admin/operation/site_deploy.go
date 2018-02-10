//[控制器] [平台]站点部署管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//站点部署管理
type SiteDeployController struct {
	controllers.BaseController
}

//站点部署列表查询
func (c *SiteDeployController) GetSiteDeloyList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点部署添加
func (c *SiteDeployController) PostSiteDeployAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点部署修改
func (c *SiteDeployController) PutSiteDeployUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//站点部署文件导入
func (c *SiteDeployController) PutSiteDeployInto(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除站点部署
func (c *SiteDeployController) DelSiteDeploy(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
