//[控制器] [平台]左下角广告管理
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//广告告管理
type AdvertController struct {
	controllers.BaseController
}

//广告列表查询
func (c *AdvertController) GetAdvertList(ctx echo.Context) error {
	advert := new(input.AdminAdvertList)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	//获取广告列表
	list, count, err := siteADBean.AdminAdvertList(advert, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, list, int64(len(list)), count, ctx))
}

//广告添加
func (c *AdvertController) PostAdvertAdd(ctx echo.Context) error {
	advert := new(input.AdminAdvertAdd)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteADBean.AdminAdvertAdd(advert)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(20254, ctx))
	}
	return ctx.NoContent(204)
}

//广告详情
func (c *AdvertController) GetAdvertInfo(ctx echo.Context) error {
	advert := new(input.AdminAdvertInfo)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	advertInfo, has, err := siteADBean.AdminAdvertInfo(advert.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(20256, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(advertInfo))
}

//广告修改
func (c *AdvertController) PutAdvertUpdate(ctx echo.Context) error {
	advert := new(input.AdminAdvertPut)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断广告id是否存在
	_, has, err := siteADBean.AdminAdvertInfo(advert.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30253, ctx))
	}
	//修改操作
	count, err := siteADBean.AdminAdvertPut(advert)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(20255, ctx))
	}
	return ctx.NoContent(204)
}

//广告状态修改
func (c *AdvertController) PutAdvertStatusUpdate(ctx echo.Context) error {
	advert := new(input.AdminAdvertState)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断广告id是否存在
	_, has, err := siteADBean.AdminAdvertInfo(advert.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30253, ctx))
	}
	//修改操作
	count, err := siteADBean.AdminAdvertState(advert)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(20257, ctx))
	}
	return ctx.NoContent(204)
}

//广告排序修改
func (c *AdvertController) PutAdvertSortUpdate(ctx echo.Context) error {
	advert := new(input.AdminAdvertSort)
	code := global.ValidRequestMember(advert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断广告id是否存在
	_, has, err := siteADBean.AdminAdvertInfo(advert.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30253, ctx))
	}
	//修改操作
	count, err := siteADBean.AdminAdvertSort(advert)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(20258, ctx))
	}
	return ctx.NoContent(204)
}

//广告图片上传
func (c *AdvertController) PutAdvertePicAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//广告图片列表拉取选择
func (c *AdvertController) PutAdvertePicUpdate(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
