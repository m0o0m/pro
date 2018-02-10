//[控制器] [平台]CDN缓存管理
package operation

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//CDN缓存管理
type CdnCleanCountController struct {
	controllers.BaseController
}

//前台cdn缓存清除
func (c *ReceptionCountController) GetCdnClean(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
