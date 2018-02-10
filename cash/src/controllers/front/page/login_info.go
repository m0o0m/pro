package page

import (
	"controllers/front/page/data_merge"
	"github.com/labstack/echo"
)

type LoginInfoController struct {
	PageBaseController
}

//login_info.html页面
func (c *LoginInfoController) LoginInfo(ctx echo.Context) error {
	return c.Render(new(data_merge.LoginInfo), ctx)
}
