//[控制器] [平台]集群管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//集群管理
type ColonyController struct {
	controllers.BaseController
}

//集群列表查询
func (c *ColonyController) GetColonyList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//集群添加
func (c *ColonyController) PostColonyAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//集群修改
func (c *ColonyController) PutColonyUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
