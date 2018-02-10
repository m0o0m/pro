package page

import (
	"controllers"
	"framework/render"
	"github.com/labstack/echo"
)

type PageBaseController struct {
	controllers.BaseController
}

func (m *PageBaseController) Render(pcPageData render.PcPageData, ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	var bytes []byte
	var ok bool
	var code int64
	if render.CacheSwitch == "on" {
		bytes, ok = render.GetCache(siteId, siteIndexId, pcPageData)
		if ok {
			return ctx.HTMLBlob(200, bytes)
		}
	}
	bytes, code = render.GenPcCache(siteId, siteIndexId, pcPageData)
	if code == 0 {
		return ctx.HTMLBlob(200, bytes)
	}
	return render.PageErr(code, ctx)
}

//不过缓存的调用此方法
func (m *PageBaseController) RenderNowData(pcPageData render.PcPageData, ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	var bytes []byte
	var code int64
	bytes, code = render.GenPcCache(siteId, siteIndexId, pcPageData)
	if code == 0 {
		return ctx.HTMLBlob(200, bytes)
	}
	return render.PageErr(code, ctx)
}
