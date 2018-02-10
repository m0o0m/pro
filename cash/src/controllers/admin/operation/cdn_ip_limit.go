//[控制器] [平台]CDN的Ip黑名单管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//CDN的Ip黑名单管理
type CdnIpLimitController struct {
	controllers.BaseController
}

//CDN的Ip黑名单列表查询
func (c *CdnIpLimitController) GetCdnIpLimitList(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//CDN的Ip黑名单添加
func (c *CdnIpLimitController) PostCdnIpLimitAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//CDN的Ip黑名单修改
func (c *CdnIpLimitController) PutCdnIpLimitUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//删除CDN的Ip黑名单
func (c *CdnIpLimitController) DelCdnIpLimit(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
