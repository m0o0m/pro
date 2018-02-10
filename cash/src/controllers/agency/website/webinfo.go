//站点资料编辑
package website

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
	"strings"
	"time"
)

type WebInfoController struct {
	controllers.BaseController
}

//网站基本信息-查询
func (c *WebInfoController) GetSiteInfo(ctx echo.Context) error {
	siteInfo := new(input.OrderModuleList)
	code := global.ValidRequest(siteInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	siteInfo.SiteId = user.SiteId
	data, _, err := webInfoBean.GetSiteInfo(siteInfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//网站基本信息-添加或修改pc
func (c *WebInfoController) PostSiteInfo(ctx echo.Context) error {
	info := new(input.PostSiteInfo)
	code := global.ValidRequest(info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	info.SiteId = user.SiteId
	count, err := webInfoBean.PostSiteInfo(info)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//网站基本信息-添加或修改wap
func (c *WebInfoController) PostSiteInfoWap(ctx echo.Context) error {
	info := new(input.PostSiteInfoWap)
	code := global.ValidRequest(info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := webInfoBean.PostSiteInfoWap(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10155, ctx))
	}
	return ctx.NoContent(204)
}

//电子管理-查询
func (c *WebInfoController) GetDzOrderList(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	data, err := webInfoBean.GetDzOrderList(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//电子管理-修改顺序
func (c *WebInfoController) PutDzOrderUpdate(ctx echo.Context) error {
	order := new(input.EditModuleOrder)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	count, err := webInfoBean.EditDzOrder(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10150, ctx))
	}
	return ctx.NoContent(204)
}

//体育管理-查询
func (c *WebInfoController) GetSpOrderList(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	data, err := webInfoBean.GetSpOrderList(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//体育管理-修改顺序
func (c *WebInfoController) PutSpOrderUpdate(ctx echo.Context) error {
	order := new(input.EditModuleOrder)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	count, err := webInfoBean.EditSpOrder(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10150, ctx))
	}
	return ctx.NoContent(204)
}

//彩票管理-查询
func (c *WebInfoController) GetFcOrderList(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	data, err := webInfoBean.GetFcOrderList(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//彩票管理-修改顺序
func (c *WebInfoController) PutFcOrderUpdate(ctx echo.Context) error {
	order := new(input.EditModuleOrder)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	count, err := webInfoBean.EditFcOrder(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10150, ctx))
	}
	return ctx.NoContent(204)
}

//彩票管理-重置排序
func (c *WebInfoController) PutFcOrderUpdateReset(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	_, err := webInfoBean.GetFcOrderFcReset(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//视讯管理-查询
func (c *WebInfoController) GetVideoOrderList(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	data, err := webInfoBean.GetVideoOrderList(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//视讯管理-修改顺序
func (c *WebInfoController) PutVideoOrderUpdate(ctx echo.Context) error {
	orderd := new(input.EditModuleOrder)
	code := global.ValidRequest(orderd, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	orderd.SiteId = user.SiteId
	count, err := webInfoBean.EditVideoOrder(orderd)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10150, ctx))
	}
	return ctx.NoContent(204)
}

//视讯管理-类型下拉框
func (c *WebInfoController) GetVideoOrderTypeList(ctx echo.Context) error {
	data, err := webInfoBean.GetVideoOrderTypeList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//视讯管理-风格下拉框
func (c *WebInfoController) GetVideoOrderStyleList(ctx echo.Context) error {
	style := new(input.StyleOrderModuleList)
	code := global.ValidRequest(style, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := webInfoBean.GetVideoOrderStyleList(style)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//视讯管理-风格配置使用
func (c *WebInfoController) PostVideoOrderStyleUse(ctx echo.Context) error {
	styleUse := new(input.PostVideoOrderStyleUse)
	code := global.ValidRequest(styleUse, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	styleUse.SiteId = user.SiteId
	count, err := webInfoBean.PostVideoOrderStyleUse(styleUse)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10151, ctx))
	}
	return ctx.NoContent(204)
}

//视讯管理-风格还原默认
func (c *WebInfoController) PutVideoOrderStyleUseUpdate(ctx echo.Context) error {
	styleUse := new(input.PutVideoOrderStyleUseUpdate)
	code := global.ValidRequest(styleUse, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	styleUse.SiteId = user.SiteId
	count, err := webInfoBean.PutVideoOrderStyleUseUpdate(styleUse)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10152, ctx))
	}
	return ctx.NoContent(204)
}

//模板管理-查询
func (c *WebInfoController) GetModuleOrderList(ctx echo.Context) error {
	order := new(input.OrderModuleList)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	data, err := webInfoBean.GetModuleOrderList(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//模板管理-修改顺序
func (c *WebInfoController) PutModuleOrderUpdate(ctx echo.Context) error {
	order := new(input.EditModuleOrder)
	code := global.ValidRequest(order, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	order.SiteId = user.SiteId
	count, err := webInfoBean.EditModuleOrder(order)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10150, ctx))
	}
	return ctx.NoContent(204)
}

//自助优惠申请配置查询
func (c *WebInfoController) GetSitePromotionConfig(ctx echo.Context) error {
	applyfor := new(input.SitePromotionConfig)
	code := global.ValidRequest(applyfor, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	applyfor.SiteId = agency.SiteId
	data, err := sitePromotionConfigBean.GetSitePromotionConfig(applyfor)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//自助优惠申请配置添加
func (c *WebInfoController) PostSitePromotionConfigAdd(ctx echo.Context) error {
	promotionconfig := new(input.SitePromotionConfigAdd)
	code := global.ValidRequest(promotionconfig, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var data schema.SitePromotionConfig
	agency := ctx.Get("user").(*global.RedisStruct)
	data.SiteId = agency.SiteId
	if promotionconfig.SiteIndexId == "" {
		data.SiteIndexId = agency.SiteIndexId
	} else {
		data.SiteIndexId = promotionconfig.SiteIndexId
	}
	data.SiteId = agency.SiteId
	data.ProTitle = promotionconfig.ProTitle
	data.ProContent = promotionconfig.ProContent
	data.Createtime = time.Now().Unix()
	data.Status = promotionconfig.Status
	count, err := sitePromotionConfigBean.Add(data)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//自助优惠申请配置状态修改
func (c *WebInfoController) PutSitePromotionConfigStatus(ctx echo.Context) error {
	promotionconfig := new(input.SitePromotionConfigStatus)
	code := global.ValidRequest(promotionconfig, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	promotionconfig.SiteId = user.SiteId
	count, err := sitePromotionConfigBean.UpdateStatus(promotionconfig)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//优惠活动宽度

//电子内页主题配置修改
func (c *WebInfoController) SetElectConfig(ctx echo.Context) error {
	promotionset := new(input.PromotionSet)
	code := global.ValidRequest(promotionset, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	promotionset.SiteId = user.SiteId
	count, err := webInfoBean.PutElectConfig(promotionset)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//电子内页主题配置查询
func (c *WebInfoController) GetElectConfig(ctx echo.Context) error {
	site := new(input.Site)
	code := global.ValidRequest(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	site.SiteId = user.SiteId
	data, has, err := webInfoBean.GetElectConfig(site)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	if len(data.Bcolor) > 0 {
		datalist := strings.Split(data.Bcolor, ",")
		data.TitleBcolor = datalist[0]
		data.TitleColor = datalist[1]
		data.ButtonBcolor = datalist[2]
		data.ButtonColor = datalist[3]
		data.BborderColor = datalist[4]
		data.PopBcolor = datalist[5]
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//电子内页主题配置初始化
func (c *WebInfoController) GetElectInitialization(ctx echo.Context) error {
	pm := new(input.DianZiInitialization)
	code := global.ValidRequest(pm, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	pm.SiteId = user.SiteId
	count, err := webInfoBean.InitializationDianZi(pm)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50189, ctx))
	}
	return ctx.NoContent(204)
}

//主动缓存单站点所有页面
// deprecated  因为前台和代理后台是不同的程序,内存不共享
func (*WebInfoController) GenPageCacheBySite(ctx echo.Context) error {

	return ctx.NoContent(204)
}

//单站点批量页面缓存 = 删除
// deprecated 因为前台和代理后台是不同的程序,内存不共享
func (*WebInfoController) DelPageCacheBySite(ctx echo.Context) error {

	return ctx.NoContent(204)
}

//主动缓存单站点所有页面---选择界面
// deprecated  因为前台和代理后台是不同的程序,内存不共享
func (*WebInfoController) GenPageCacheBySiteDrop(ctx echo.Context) error {

	return ctx.NoContent(204)
}
