//[控制器] [平台]客户后台白名单控制管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//客户后台白名单控制管理
type CustomerIpCountController struct {
	controllers.BaseController
}

//客户后台白名单查询
func (c *ReceptionCountController) GetCustomerIpList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//客户白名单修改
func (c *ReceptionCountController) PutCustomerIpUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
