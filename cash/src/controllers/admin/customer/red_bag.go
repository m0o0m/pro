//[控制器] [平台]红包数据补全
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//红包数据补全
type RedBagController struct {
	controllers.BaseController
}

//红包数据补全（强行修改某会员的打码或存款数据让其能抢红包）
func (c *RedBagController) PutRedBagSet(ctx echo.Context) error {
	redBag := new(input.RedBagData)
	code := global.ValidRequest(redBag, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, key, err := redBagBean.RedBagSet(redBag)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" {
		return ctx.JSON(200, global.ReplyError(10143, ctx)) //红包数据设置失败
	}
	return ctx.NoContent(204)
}
