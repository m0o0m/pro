//[控制器] [平台]优惠统计
package report

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"sync"
)

//优惠统计管理
type DiscountCountController struct {
	controllers.BaseController
}

//优惠统计数据查询
func (c *DiscountCountController) GetDiscountCountList(ctx echo.Context) error {
	reqData := new(input.DiscountCountList)
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
	// DESCRIPTION:代理存在就查询查询代理表
	var agencyId int64
	if reqData.AgencyAccount != "" {
		var err error
		agencyId, err = agencyBean.GetAgencyIdByAccount(reqData.SiteId, reqData.AgencyAccount)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30298, ctx))
		}
	}
	// DESCRIPTION:查询哪些会员
	var memberIds []int64
	if reqData.AgencyAccount != "" || reqData.Account != "" {
		var err error
		memberIds, err = memberBean.GetMemberIdsByAgencyId(reqData.SiteId, agencyId, reqData.Account)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}

	// DESCRIPTION:查询
	retreatWaterRecordTotal, err := discountCountBean.GetList(reqData.SiteId, memberIds, &global.Times{StartTime: sTime, EndTime: eTime}, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for _, v := range *retreatWaterRecordTotal.Content {
		retreatWaterRecordTotal.SubNum += v.Num
		retreatWaterRecordTotal.SubDiscountMoney += v.DiscountMoney
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, retreatWaterRecordTotal, int64(len(*retreatWaterRecordTotal.Content)), retreatWaterRecordTotal.Num, ctx))
}

//重新统计
func (c *DiscountCountController) PostDiscountCountUpdate(ctx echo.Context) error {
	reqData := new(input.Recount)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	days, code := global.GetEveryDayTimes(reqData.STime, reqData.ETime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	wg := &sync.WaitGroup{}
	safeErr := global.NewSafeError()
	for i := range days {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			discountReports, err := discountCountBean.GetRetreatWater(days[i].StartTime, days[i].EndTime)
			if err != nil {
				safeErr.Push(err.Error())
				return
			}
			if !safeErr.IsValid() && len(discountReports) > 0 {
				err = discountCountBean.InsertOrUpdate(&discountReports)
				if err != nil {
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
func (c *DiscountCountController) PutDiscountCountSwitch(ctx echo.Context) error {
	reqData := new(input.TimingCount)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	code = discountCountBean.TimingSwitch(reqData.Open)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	return ctx.NoContent(204)
}
