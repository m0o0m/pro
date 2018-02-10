//[控制器] [平台]入款统计
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"sync"
)

//入款统计管理
type DepositCountController struct {
	controllers.BaseController
}

//入款统计数据查询
func (c *DepositCountController) GetDepositCountList(ctx echo.Context) error {
	reqData := new(input.CashCountList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	c.GetParam(listParam, ctx)
	var sTime, eTime int64
	if reqData.STime != "" && reqData.ETime != "" {
		sTime, eTime, code = global.FormatDay2Timestamp(reqData.STime, reqData.ETime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	// DESCRIPTION:条件汇总有代理账号,查代理表
	var agencyId int64
	if reqData.AgencyAccount != "" {
		var err error
		agencyId, err = agencyBean.GetAgencyIdByAccount(reqData.SiteId, reqData.AgencyAccount)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30298, ctx))
		}
	}
	// DESCRIPTION:条件包含代理账号或者会员账号,查询会员表
	var memberIds []int64
	if reqData.AgencyAccount != "" || reqData.Account != "" {
		var err error
		memberIds, err = memberBean.GetMemberIdsByAgencyId(reqData.SiteId, agencyId, reqData.Account)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	// DESCRIPTION:查询入款统计表
	cashCountTotal, err := cashCountBean.GetList(reqData.SiteId, memberIds, reqData.IntoStyle, &global.Times{StartTime: sTime, EndTime: eTime}, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for _, v := range *cashCountTotal.Content {
		cashCountTotal.SubNum += v.Num
		cashCountTotal.SubCashMoney += v.CashMoney
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, cashCountTotal, int64(len(*cashCountTotal.Content)), cashCountTotal.Num, ctx))
}

//重新统计
func (c *DepositCountController) PostDepositCountUpdate(ctx echo.Context) error {
	reqData := new(input.CashRecount)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	days, code := global.GetEveryDayTimes(reqData.STime, reqData.ETime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(days) > 20 {
		return ctx.JSON(200, global.ReplyError(60233, ctx))
	}
	wg := &sync.WaitGroup{}
	safeErr := global.NewSafeError()
	sqlMutex := global.GetSiteCashCountReportMutex("all")
	for i := range days {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cashReports, err := cashCountBean.GetAllCash(days[i].StartTime, days[i].EndTime)
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				safeErr.Push(err.Error())
				return
			}
			if !safeErr.IsValid() && len(cashReports) > 0 {
				err = cashCountBean.InsertOrUpdate(&cashReports, sqlMutex)
				if err != nil {
					global.GlobalLogger.Error("err:%s", err.Error())
					safeErr.Push(err.Error())
				}
				return
			}
		}(i)
	}
	wg.Wait()
	if safeErr.IsValid() {
		global.GlobalLogger.Error("error:%s", safeErr.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//定时统计开关控制
func (c *DepositCountController) PutDepositCountSwitch(ctx echo.Context) error {
	reqData := new(input.TimingCashCount)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code = cashCountBean.TimingSwitch(reqData.Open)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	return ctx.NoContent(204)
}
