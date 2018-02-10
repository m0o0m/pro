package page

import (
	"controllers/front/page/data_merge"
	"github.com/labstack/echo"
)

type WapviewController struct {
	PageBaseController
}

//wapview.html页面
func (c *WapviewController) Wapview(ctx echo.Context) error {
	return c.Render(new(data_merge.Wapview), ctx)
}
