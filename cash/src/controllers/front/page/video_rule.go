package page

import (
	"controllers/front/page/data_merge"
	"github.com/labstack/echo"
)

type VideoRuleController struct {
	PageBaseController
}

//wapview.html页面
func (c *VideoRuleController) VideoRule(ctx echo.Context) error {
	return c.Render(new(data_merge.VideoRule), ctx)
}
