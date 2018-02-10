//[控制器] [平台]手机端域名管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//手机端域名管理
type WapDomainController struct {
	controllers.BaseController
}

//手机端域名列表查询
func (c *WapDomainController) GetWapDomainList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//手机端域名添加
func (c *WapDomainController) PostWapDomainAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//手机端域名修改
func (c *WapDomainController) PutWapDomainUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//手机端域名绑定状态修改
func (c *WapDomainController) PutWapDomainStatusUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除手机端域名
func (c *WapDomainController) DelWapDomain(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
