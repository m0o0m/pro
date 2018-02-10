//[控制器] [平台]报表辅助查询
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//报表辅助查询
type ReportOtherSerchController struct {
	controllers.BaseController
}

//报表辅助查询
func (c *ReportOtherSerchController) GetReportOtherSerch(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
