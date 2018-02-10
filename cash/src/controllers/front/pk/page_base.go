package pk

import (
	"controllers"
	"framework/render"
	"github.com/labstack/echo"
)

type PageBaseController struct {
	controllers.BaseController
}

func (m *PageBaseController) Render(lotteryPageData render.LotteryPageData, ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	var bytes []byte
	var code int64
	bytes, code = render.GenLottery(siteId, siteIndexId, lotteryPageData)
	if code == 0 {
		return ctx.HTMLBlob(200, bytes)
	}
	return render.PageErr(code, ctx)
}
