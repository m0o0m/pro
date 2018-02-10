//[控制器] [平台]绑定信息管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//绑定信息管理
type BindInfoController struct {
	controllers.BaseController
}

//绑定信息列表查询（ip和api）
func (c *BindInfoController) GetBindInfoList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定ip添加
func (c *BindInfoController) PostBindIpAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定ip修改
func (c *BindInfoController) PutBindIpUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定ip删除
func (c *BindInfoController) DelBindIp(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定api添加
func (c *BindInfoController) PostBindApiAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定api修改
func (c *BindInfoController) PutBindApiUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//绑定api删除
func (c *BindInfoController) DelBindApi(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
