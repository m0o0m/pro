//图文编辑
package website

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

type WebLogoController struct {
	controllers.BaseController
}

//附件管理
func (mc *WebLogoController) GetSiteThumb(ctx echo.Context) error {
	site_thumb := new(input.GetSiteThumb)
	code := global.ValidRequest(site_thumb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	site_thumb.SiteId = agency.SiteId
	listParams := new(global.ListParams)
	mc.GetParam(listParams, ctx)
	data, count, err := siteThumbBean.List(site_thumb, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

/*************************************************
*	站点logo图片管理							**
 *************************************************/
//站点logo图片列表
func (pc *WebLogoController) GetWebLogoList(ctx echo.Context) error {
	info_list := new(input.LogoInfoList)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webLogoBean.GetWebLogoList(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点logo管理修改
func (pc *WebLogoController) PutWebLogo(ctx echo.Context) error {
	info_list := new(input.UpdateLogoInfo)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	_, err := webLogoBean.GetWebLogo(info_list.Id, info_list.SiteId, info_list.SiteIndexId, 0)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data, err := webLogoBean.PutWebLogo(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90908, ctx))
	}
	return ctx.NoContent(204)
}

//站点logo管理修改路径
func (pc *WebLogoController) PutWebLogoWay(ctx echo.Context) error {
	info_list := new(input.UpdateLogoInfoWay)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	_, err := webLogoBean.GetWebLogo(info_list.Id, info_list.SiteId, info_list.SiteIndexId, 0)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data, err := webLogoBean.PutWebLogoWay(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90908, ctx))
	}
	return ctx.NoContent(204)
}

//站点logo管理新增
func (pc *WebLogoController) PostWebLogo(ctx echo.Context) error {
	info_list := new(input.PostLogoInfo)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	_, err := webLogoBean.GetWebLogo(0, info_list.SiteId, info_list.SiteIndexId, info_list.Type)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	data, err := webLogoBean.PostWebLogo(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90909, ctx))
	}

	return ctx.NoContent(204)
}

/************************************************************************
*	站点左右浮动														*
 ************************************************************************/
//站点左右浮动状态修改
func (pc *WebLogoController) PutWebFloatStatus(ctx echo.Context) error {
	infoStatus := new(input.FloatListStatus)
	code := global.ValidRequest(infoStatus, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	infoStatus.SiteId = agency.SiteId
	count, err := webFloatBean.PutWebFloatStatus(infoStatus)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点左右浮动状态
func (pc *WebLogoController) GetWebFloatStatus(ctx echo.Context) error {
	infoStatus := new(input.GetFloatListStatus)
	code := global.ValidRequest(infoStatus, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	infoStatus.SiteId = agency.SiteId
	data, err := webFloatBean.GetWebFloatListStatus(infoStatus)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//浮动图片删除
func (pc *WebLogoController) DeleteWebFloat(ctx echo.Context) error {
	fD := new(input.DeleteFloatPicture)
	code := global.ValidRequest(fD, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	fD.SiteId = agency.SiteId
	count, err := webFloatBean.DeleteWebFloatList(fD)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//站点左右浮动查询
func (pc *WebLogoController) GetWebFloatList(ctx echo.Context) error {
	info_list := new(input.FloatList)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webFloatBean.GetWebFloatList(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点左右浮动修改
func (pc *WebLogoController) PutWebFloatUpdate(ctx echo.Context) error {
	info_list := new(input.FloatUpdate)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	count, err := webFloatBean.PutWebFloat(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点浮动图片上传
func (pc *WebLogoController) PutWebFloatPicture(ctx echo.Context) error {
	info_list := new(input.FloatUpdatePicture)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	count, err := webFloatBean.PutWebFloatPicture(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点左右浮动新增
func (pc *WebLogoController) PostWebFloat(ctx echo.Context) error {
	info_list := new(input.FloatAdd)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webFloatBean.PostWebFloat(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90863, ctx))
	}

	return ctx.NoContent(204)
}

/***************************************************
*站点轮播图管理									****
 ***************************************************/
//站点轮播图查询
func (pc *WebLogoController) GetWebFlashList(ctx echo.Context) error {
	info_list := new(input.FlashList)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webFlashBean.GetWebFlashList(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//站点轮播图修改
func (pc *WebLogoController) PutWebFlashUpdate(ctx echo.Context) error {
	info_list := new(input.FlashUpdate)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	count, err := webFlashBean.PutWebFlash(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(90819, ctx))
	}
	return ctx.NoContent(204)
}

//站点轮播图新增
func (pc *WebLogoController) PostWebFlash(ctx echo.Context) error {
	info_list := new(input.FlashAdd)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webFlashBean.PostWebFlash(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90820, ctx))
	}
	return ctx.NoContent(204)
}

//站点轮播图状态更改
func (pc *WebLogoController) PutWebFlash(ctx echo.Context) error {
	infoList := new(input.FlashStatus)
	code := global.ValidRequest(infoList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	infoList.SiteId = agency.SiteId
	data, err := webFlashBean.PutWebFlashStatus(infoList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90820, ctx))
	}
	return ctx.NoContent(204)
}

//附件修改
func (c *WebLogoController) PutSitethumb(ctx echo.Context) error {
	site_thumb := new(input.SiteThumbEdit)
	code := global.ValidRequest(site_thumb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	site_thumb.SiteId = agency.SiteId
	count, err := siteThumbBean.SiteThumbUpdate(site_thumb)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//附件删除
func (c *WebLogoController) DelSitethumb(ctx echo.Context) error {
	site_thumb := new(input.SiteThumbDelete)
	code := global.ValidRequest(site_thumb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	site_thumb.SiteId = agency.SiteId
	count, err := siteThumbBean.SiteThumbDelete(site_thumb)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

/*************************************************
*	站点公告弹窗管理						******
 *************************************************/
//获取站点公告弹窗列表
func (pc *WebLogoController) GetWebAdvList(ctx echo.Context) error {
	info_list := new(input.AdvList)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, err := webAdvBean.GetWebAdvList(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取站点公告弹窗列表
func (pc *WebLogoController) GetWebAdvListDetail(ctx echo.Context) error {
	info_list := new(input.AdvListDetail)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	data, has, err := webAdvBean.GetWebAdvListDetail(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50155, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取站点公告弹窗 修改
func (pc *WebLogoController) PutWebAdvUpdate(ctx echo.Context) error {
	adv := new(input.AdvUpdate)
	code := global.ValidRequest(adv, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	adv.SiteId = agency.SiteId
	if adv.Id == 0 {
		return ctx.JSON(200, global.ReplyError(90232, ctx))
	}

	data, err := webAdvBean.PutWebAdv(adv)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90233, ctx))
	}

	return ctx.NoContent(204)
}

//获取站点公告弹窗 新增
func (pc *WebLogoController) PostWebAdv(ctx echo.Context) error {
	info_list := new(input.AdvAdd)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	info_list.AddTime = time.Now().Unix()
	data, err := webAdvBean.PostWebAdv(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90234, ctx))
	}
	return ctx.NoContent(204)
}

//获取站点公告弹窗 删除
func (pc *WebLogoController) DeleteAdv(ctx echo.Context) error {
	adv := new(input.UpdateDeleteTime)
	code := global.ValidRequest(adv, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := webAdvBean.DeleteWebAdv(adv)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90235, ctx))
	}
	return ctx.NoContent(204)
}

//获取站点公告弹窗 修改配置

func (pc *WebLogoController) PutConfig(ctx echo.Context) error {
	config := new(input.UpdateConfig)
	code := global.ValidRequest(config, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	config.SiteId = agency.SiteId
	data, err := webAdvBean.UpdateWebAdvConfig(config)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90236, ctx))
	}

	return ctx.NoContent(204)
}

//获取站点公告弹窗 获取详情
func (pc *WebLogoController) PutConfigDetail(ctx echo.Context) error {
	info_list := new(input.UpdateConfigDetail)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	info, has, err := webAdvBean.DetailWebAdvConfig(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

/*******************************************
*站点广告管理							****
 *******************************************/
//站点广告列表
func (pc *WebLogoController) GetWebPopList(ctx echo.Context) error {
	info_list := new(input.AdvListBySite)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	listparam := new(global.ListParams)
	//获取listparam的数据
	pc.GetParam(listparam, ctx)
	data, count, err := webPopBean.GetWebPopList(info_list, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//站点广告详情
func (pc *WebLogoController) GetWebPopListDetail(ctx echo.Context) error {
	info_list := new(input.AdvListBySiteDetail)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	info, has, err := webPopBean.GetWebPopListDetail(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(50156, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//站点广告管理 修改
func (pc *WebLogoController) PutWebPopUpdate(ctx echo.Context) error {
	info_list := new(input.PopUpdate)
	code := global.ValidRequest(info_list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	info_list.SiteId = agency.SiteId
	count, err := webPopBean.PutWebPop(info_list)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//站点广告管理 新增
func (pc *WebLogoController) PostWebPop(ctx echo.Context) error {
	pop := new(input.PopAdd)
	code := global.ValidRequest(pop, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	pop.SiteId = agency.SiteId
	pop.AddTime = time.Now().Unix()
	count, err := webPopBean.PostWebPop(pop)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//站点广告管理 状态修改

func (pc *WebLogoController) PutPopStatus(ctx echo.Context) error {
	pop := new(input.UpdatePopStatus)
	code := global.ValidRequest(pop, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	pop.SiteId = agency.SiteId
	count, has, err := webPopBean.UpdateWebPopStatus(pop)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50156, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//附件上传（附件添加）
func (c *WebLogoController) PostSitethumb(ctx echo.Context) error {
	return ctx.NoContent(204)
}
