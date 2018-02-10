package site

//[控制器] [平台]下载地址管理

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/schema"
)

//下载地址管理
type DownloadController struct {
	controllers.BaseController
}

//下载列表查询
func (c *DownloadController) GetDownloadList(ctx echo.Context) error {
	data, err := siteDownBean.GetSiteDownList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//下载地址添加
func (c *DownloadController) PostDownloadAdd(ctx echo.Context) error {
	sitedown := new(input.SiteDownAdd)
	code := global.ValidRequest(sitedown, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var data schema.SiteDown
	data.IosUrl = sitedown.IosUrl
	data.AndroidUrl = sitedown.AndroidUrl
	data.Platform = sitedown.Platform
	data.Vers = sitedown.Vers
	data.State = 1
	count, err := siteDownBean.SiteDownAdd(data)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取单条下载详情
func (c *DownloadController) GetSiteDownOne(ctx echo.Context) error {
	sitedown := new(input.SiteDown)
	code := global.ValidRequest(sitedown, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := siteDownBean.GetInfo(sitedown)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has == false {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//下载地址修改
func (c *DownloadController) PutDownloadUpdate(ctx echo.Context) error {
	sitedown := new(input.SiteDownEdit)
	code := global.ValidRequest(sitedown, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteDownBean.UpdateInfo(sitedown)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(30289, ctx))
	}
	return ctx.NoContent(204)
}

//下载地址状态修改
func (c *DownloadController) PutDownloadStatusUpdate(ctx echo.Context) error {
	sitedown := new(input.SiteDownState)
	code := global.ValidRequestAdmin(sitedown, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteDownBean.UpdateStatus(sitedown)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count < 1 {
		return ctx.JSON(200, global.ReplyError(30289, ctx))
	}
	return ctx.NoContent(204)
}
