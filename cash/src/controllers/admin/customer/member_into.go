//[控制器] [平台]导入会员
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//导入会员管理
type MemberIntoController struct {
	controllers.BaseController
}

//导入会员前期条件选择时条件变动
func (c *MemberIntoController) GetMemberIntoTerm(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//选择导入文件，导入会员
func (c *MemberIntoController) PostMemberIntoInsert(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
