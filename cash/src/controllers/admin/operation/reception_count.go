//[控制器] [平台]前台访问统计管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//前台访问统计管理
type ReceptionCountController struct {
	controllers.BaseController
}

//前台访问统计查询
func (c *ReceptionCountController) GetReceptionCount(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
