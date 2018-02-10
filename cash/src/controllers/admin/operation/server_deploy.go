//[控制器] [平台]服务器部署管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//服务器部署管理
type ServerDeployController struct {
	controllers.BaseController
}

//服务器部署列表查询
func (c *ServerDeployController) GetServerDeployList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//服务器部署添加
func (c *ServerDeployController) PostServerDeployAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//服务器部署修改
func (c *ServerDeployController) PutServerDeployUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除站点部署
func (c *ServerDeployController) DelServerDeploy(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
