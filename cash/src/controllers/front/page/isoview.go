package page

import (
	"controllers/front/page/data_merge"
	"github.com/labstack/echo"
)

type IsoviewController struct {
	PageBaseController
}

//wapview.html页面
func (c *IsoviewController) Isoview(ctx echo.Context) error {
	return c.Render(new(data_merge.Isoview), ctx)
}
