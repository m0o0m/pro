//[控制器] [平台]站点口令
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//站点口令管理
type SitePasswordController struct {
	controllers.BaseController
}

//站点口令列表查询
func (c *SitePasswordController) GetSitePassList(ctx echo.Context) error {
	sitePass := new(input.SitePassList)
	code := global.ValidRequestAdmin(sitePass, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	data, count, err := sitePassBean.SitePassList(sitePass, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//站点口令修改  有即修改，无即插入
func (c *SitePasswordController) PutSitePassUpdate(ctx echo.Context) error {
	sitePass := new(input.SitePassUpdate)
	code := global.ValidRequestAdmin(sitePass, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登录人信息
	admin := ctx.Get("admin").(*global.AdminRedisStruct)
	sitePass.Account = admin.Account
	//站点是否存在
	has, err := sitePassBean.IsExistSite(sitePass.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	count, err := sitePassBean.SitePassUpdate(sitePass)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//站点口令批量启用停用
func (*SitePasswordController) PutBatchDelChanges(ctx echo.Context) error {
	batchDelChanges := new(input.BatchDelChanges)
	code := global.ValidRequestAdmin(batchDelChanges, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(batchDelChanges.SiteId) == 0 {
		return ctx.JSON(200, global.ReplyError(10050, ctx))
	}

	//获取登录人信息
	admin := ctx.Get("admin").(*global.AdminRedisStruct)
	batchDelChanges.Account = admin.Account
	count, err := sitePassBean.BatchDelChanges(batchDelChanges)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
