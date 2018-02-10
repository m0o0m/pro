package wap

import (
	"controllers"
	"framework/render"
	"github.com/labstack/echo"
)

type WapBaseController struct {
	controllers.BaseController
}

func (m *WapBaseController) Render(wapPageData render.WapPageData, ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)

	if render.CacheSwitch == "on" {
		bytes, ok := render.GetCache(siteId, siteIndexId, wapPageData)
		if ok {
			return ctx.HTMLBlob(200, bytes)
		}
	}
	bytes, code := render.GenWapCache(siteId, siteIndexId, wapPageData)
	if code == 0 {
		return ctx.HTMLBlob(200, bytes)
	}
	return render.PageErr(code, ctx)
}
