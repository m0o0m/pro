//[控制器] [平台]域名管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//域名管理
type DomainController struct {
	controllers.BaseController
}

//域名列表查询
func (c *DomainController) GetDomainList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//域名添加
func (c *DomainController) PostDomainAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//域名修改
func (c *DomainController) PutDomainUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//域名绑定状态修改
func (c *DomainController) PutDomainStatusUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除域名
func (c *DomainController) DelDomain(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
