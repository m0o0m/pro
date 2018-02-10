//[控制器] [平台]日志管理
package rbac

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//日志管理
type LogController struct {
	controllers.BaseController
}

//日志查询
func (c *LogController) GetLogList(ctx echo.Context) error {
	a_log := new(input.AdminLoginLog)
	code := global.ValidRequest(a_log, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if a_log.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", a_log.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if a_log.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02  15:04:05", a_log.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	data, count, err := adminLoginLogBean.AdminLoginLogList(a_log, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}
