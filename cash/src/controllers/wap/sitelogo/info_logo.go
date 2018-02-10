package sitelogo

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strings"
)

//首页轮播图，公告，logo
type InfoLogoController struct {
	controllers.BaseController
}

//首页logo,公告，轮播图
func (c *InfoLogoController) HomePageLogo(ctx echo.Context) error {
	site := new(input.Home)
	code := global.ValidRequestMember(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取logo、公告
	data, err := infoLogoBean.HomePageOneLogo(site.SiteId, site.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//4大模块
	orderNumbers, err := infoLogoBean.GetOrderNumberModuleBySite(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	modules := strings.Split(orderNumbers.Module, ",")
	products, err := infoLogoBean.ProductList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	productMap := make(map[string]*back.ProductMapping)
	for k, _ := range products {
		productMap[products[k].VType] = &back.ProductMapping{
			Name:  products[k].ProductName,
			VType: products[k].VType,
		}
	}

	f := func(vTypes []string) (mappings []*back.ProductMapping) {
		for k, _ := range vTypes {
			mapping, ok := productMap[vTypes[k]]
			if ok {
				mappings = append(mappings, mapping)
			} else {
				global.GlobalLogger.Error("err:%s not exist", vTypes[k])
			}
		}
		return
	}
	var orderModules []back.OrderModuleBySite
	for _, v := range modules {
		var orderModule back.OrderModuleBySite
		switch v {
		case "video_module":
			orderModule.Module = back.ProductMapping{"视讯直播", "video"}
			orderModule.SubModule = f(strings.Split(orderNumbers.VideoModule, ","))
		case "fc_module":
			orderModule.Module = back.ProductMapping{"彩票游戏", "fc"}
			orderModule.SubModule = f(strings.Split(orderNumbers.FcModule, ","))
		case "dz_module":
			orderModule.Module = back.ProductMapping{"电子游艺", "dz"}
			orderModule.SubModule = f(strings.Split(orderNumbers.DzModule, ","))
		case "sp_module":
			orderModule.Module = back.ProductMapping{"体育赛事", "sp"}
			orderModule.SubModule = f(strings.Split(orderNumbers.SpModule, ","))
		default:
			global.GlobalLogger.Error("%s not in (video_module,fc_module,dz_module,sp_module)", v)
			continue
		}
		orderModules = append(orderModules, orderModule)
	}
	data.HomePageProductAndTypeBack = orderModules
	return ctx.JSON(200, global.ReplyItem(data))
}
