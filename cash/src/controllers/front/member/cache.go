package member

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"strings"
	"sync"
)

type CacheController struct {
	controllers.BaseController
}

//刷新站点维护信息缓存
func (*CacheController) RefreshSiteModule(ctx echo.Context) error {
	siteModuleAll, err := siteModuleBean.GetModuleAll()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	global.SiteModuleCache = sync.Map{} //清除所有缓存
	for _, siteModule := range siteModuleAll {
		var siteIds []string //站点
		if siteModule.SiteIds == "0" || siteModule.SiteIds == "" {
			siteIds = append(siteIds, "all")
		} else {
			siteIds = append(siteIds, strings.Split(siteModule.SiteIds, ",")...)
		}
		var froms []string //来源
		switch siteModule.FType {
		case 2:
			froms = append(froms, "pc")
		case 3:
			froms = append(froms, "wap")
		case 4:
			froms = append(froms, "app")
		default:
			froms = append(froms, "pc", "wap", "app")
		}
		for _, siteId := range siteIds {
			for _, from := range froms {
				global.SiteModuleCache.Store(global.GenKey(siteId, from, siteModule.VType), siteModule) // 缓存数据
			}
		}
	}
	return ctx.NoContent(204)
}
