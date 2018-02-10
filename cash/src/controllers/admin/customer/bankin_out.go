//[控制器] [平台]出入款管理
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//出入款管理
type BankInOutController struct {
	controllers.BaseController
}

//入款查询
func (c *BankInOutController) GetBankInRecord(ctx echo.Context) error {
	//获取用户参数
	in_out := new(input.InDeposit)
	code := global.ValidRequestAdmin(in_out, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("UTC")
	if in_out.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", in_out.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if in_out.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", in_out.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	//公司入款
	if in_out.InType == 1 {
		data, count, err := bankInOutBean.DepositByCompany(in_out, listparam, times)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	} else {
		//线上入款
		data, count, err := bankInOutBean.DepositByOnline(in_out, listparam, times)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	}
}

//出款查询
func (c *BankInOutController) GetBankOutRecord(ctx echo.Context) error {
	//获取用户参数
	in_out := new(input.OutDeposit)
	code := global.ValidRequestAdmin(in_out, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if in_out.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", in_out.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if in_out.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", in_out.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := bankInOutBean.OutDepositSearch(in_out, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}
