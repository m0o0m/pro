//[控制器] [平台]绑定端口管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//绑定信息管理
type BindPortController struct {
	controllers.BaseController
}

//绑定端口列表查询
func (c *BindPortController) GetBindPortList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定端口添加
func (c *BindPortController) PostBindPortAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定端口修改
func (c *BindPortController) PutBindPortUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定端口状态修改
func (c *BindPortController) PutBindPortStatus(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定端口支付详情
func (c *BindPortController) GetBindPortPay(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定端口状态详情
func (c *BindPortController) GetBindPortStatus(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
