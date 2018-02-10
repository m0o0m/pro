//[控制器] [平台]cdn的js控制
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//cdn的js版本管理
type CdnJsController struct {
	controllers.BaseController
}

//js版本列表查询
func (c *CdnJsController) GetCdnJsList(ctx echo.Context) error {
	data, err := siteJsVersionBean.GetSiteJsVersionList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//js版本修改
func (c *CdnJsController) PutCdnJsUpdate(ctx echo.Context) error {
	sitejs := new(input.SiteJsVersion)
	code := global.ValidRequest(sitejs, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := siteJsVersionBean.UpdateInfo(sitejs)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//js文件夹添加
func (c *CdnJsController) PostCdnJsAdd(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//js文件差异详情查询
func (c *CdnJsController) GetCdnJsDiff(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}

//js文件夹删除
func (c *CdnJsController) PutCdnJsDel(ctx echo.Context) error {
	return ctx.JSON(200, global.ReplyItem(""))
}
