//[控制器] [平台]站点层级管理
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//站点层级管理
type SiteLevelController struct {
	controllers.BaseController
}

//站点层级详情列表(以层级为主)
func (c *SiteLevelController) GetSiteLevelList(ctx echo.Context) error {
	siteLevel := new(input.SiteList)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, count, err := siteLevelBean.ListSiteLevel(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//站点层级列表(以站点搜索)
func (c *SiteLevelController) GetSiteList(ctx echo.Context) error {
	siteLevel := new(input.SiteList)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, count, err := siteLevelBean.ListSite(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//站点层级移动修改
func (c *SiteLevelController) PutSiteLevelUpdate(ctx echo.Context) error {
	siteLevel := new(input.MoveSiteLevel)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code, count, err := siteLevelBean.MoveSiteLevel(siteLevel)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10135, ctx))
	}
	return ctx.NoContent(204)
}

//初始化站点层级设置，将未分层站点分到默认分层
func (c *SiteLevelController) PutSiteLevelAll(ctx echo.Context) error {
	code, err := siteLevelBean.InitSiteLevel()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	return ctx.NoContent(204)
}

//获取层级设定数据
func (c *SiteLevelController) GetSiteLevelInfo(ctx echo.Context) error {
	siteLevel := new(input.DelSiteLevel)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, count, err := siteLevelBean.DetailSiteLevel(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//站点层级数据添加
func (c *SiteLevelController) PostSiteLevelAdd(ctx echo.Context) error {
	siteLevel := new(input.AddSiteLevel)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteLevelBean.AddSiteLevel(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10130, ctx))
	}
	return ctx.NoContent(204)
}

//站点层级数据修改
func (c *SiteLevelController) PutLevelInfoUpdate(ctx echo.Context) error {
	siteLevel := new(input.EditSiteLevel)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteLevelBean.EditSiteLevel(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10131, ctx))
	}
	return ctx.NoContent(204)
}

//站点层级数据删除
func (c *SiteLevelController) DelSiteLevelInfo(ctx echo.Context) error {
	siteLevel := new(input.DelSiteLevel)
	code := global.ValidRequest(siteLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteLevelBean.DelSiteLevel(siteLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10132, ctx))
	}
	return ctx.NoContent(204)
}

//站点层级列表下拉框
func (c *SiteLevelController) GetSiteListDrop(ctx echo.Context) error {
	list, err := siteLevelBean.LevelSiteListDrop()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
