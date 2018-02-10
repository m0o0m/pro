package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//优惠查询
type MemberRetreatWaterController struct {
	controllers.BaseController
}

//优惠查询列表
func (mrwc *MemberRetreatWaterController) SearchMemberRetreatWater(ctx echo.Context) error {
	memberRetreatWaterSet := new(input.ListRetreatWater)
	code := global.ValidRequest(memberRetreatWaterSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var sTime int64
	var eTime int64
	if memberRetreatWaterSet.Year != "" && memberRetreatWaterSet.Month != "" {
		sTime, eTime, code = global.FormatMonth2Timestamp(memberRetreatWaterSet.Year + "-" + memberRetreatWaterSet.Month)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	mrwc.GetParam(listparam, ctx)
	list, count, err := memberRetreatWaterSetBean.SearchMemberRetreatWaterSet(memberRetreatWaterSet, listparam, sTime, eTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyPagination2(listparam, list, int64(len(list)), count))
}

//优惠查询明细
func (mrwc *MemberRetreatWaterController) DetailMemberRetreatWater(ctx echo.Context) error {
	memberRetreatWater := new(input.DetailRetreatWater)
	code := global.ValidRequest(memberRetreatWater, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, count, err := memberRetreatWaterBean.DetailMemberRetreatWater(memberRetreatWater)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//优惠查询冲销
func (mrwc *MemberRetreatWaterController) EditMemberRetreatWater(ctx echo.Context) error {
	water := new(input.EditRetreatWater)
	code := global.ValidRequest(water, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := memberRetreatWaterBean.EditMemberRetreatWater(water)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10108, ctx))
	}
	return ctx.NoContent(204)

}
