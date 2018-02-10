//[控制器] [平台]支付域名管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//支付域名管理
type PayDomainController struct {
	controllers.BaseController
}

//支付域名列表查询
func (c *PayDomainController) GetPayDomainList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//支付域名添加
func (c *PayDomainController) PostPayDomainAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//支付域名修改
func (c *PayDomainController) PutPayDomainUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//支付域名绑定状态修改
func (c *PayDomainController) PutPayDomainStatusUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除支付域名
func (c *PayDomainController) DelPayDomain(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
