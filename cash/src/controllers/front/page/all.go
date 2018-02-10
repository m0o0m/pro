package page

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
)

//缓存操作相关
type CacheController struct {
	controllers.BaseController
}

//缓存站点皮肤信息
func (*CacheController) CacheTheme(ctx echo.Context) error {
	themes, err := siteBean.GetThemeAll()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for _, v := range themes {
		global.ThemeCache.Store(v.SiteId+"$"+v.SiteIndexId, v.ThemeName)
	}
	return ctx.NoContent(204)
}
