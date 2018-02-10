//[控制器] [平台]日志管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//代理管理
type SiteLogController struct {
	controllers.BaseController
}

//登录日志查询
func (c *SiteLogController) GetSiteLoginLog(ctx echo.Context) error {
	siteLoginLog := new(input.SiteLoginLog)
	code := global.ValidRequestAdmin(siteLoginLog, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if siteLoginLog.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", siteLoginLog.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if siteLoginLog.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", siteLoginLog.EndTime, loc)
		times.EndTime = et.Unix()
	}
	list, count, err := siteLogBean.SiteLoginLog(siteLoginLog, listParam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//操作日志查询
func (c *SiteLogController) GetSiteDoLog(ctx echo.Context) error {
	siteDoLog := new(input.SiteDoLog)
	code := global.ValidRequestAdmin(siteDoLog, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if siteDoLog.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", siteDoLog.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if siteDoLog.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", siteDoLog.EndTime, loc)
		times.EndTime = et.Unix()

	}
	list, count, err := siteLogBean.SiteDoLog(siteDoLog, listParam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}

//自动稽核
func (c *SiteLogController) GetSiteAutoAudit(ctx echo.Context) error {
	siteLoginLog := new(input.AutoAudit)
	code := global.ValidRequestAdmin(siteLoginLog, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listParam的数据
	c.GetParam(listParam, ctx)
	list, count, err := siteLogBean.AutoAudit(siteLoginLog, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, list, int64(len(list)), count, ctx))
}
