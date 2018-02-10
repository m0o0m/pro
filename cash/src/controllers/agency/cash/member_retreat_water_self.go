package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//自助返水查询
type MemberRetreatWaterSelfController struct {
	controllers.BaseController
}

//自助返水查询列表
func (mrwsc *MemberRetreatWaterSelfController) SearchMemberRetreatWaterSelf(ctx echo.Context) error {
	memberRetreatWaterSelf := new(input.ListRetreatWaterSelf)
	code := global.ValidRequest(memberRetreatWaterSelf, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var err error
	if memberRetreatWaterSelf.StartTimeStr != "" && memberRetreatWaterSelf.EndTimeStr != "" {
		memberRetreatWaterSelf.StartTime, memberRetreatWaterSelf.EndTime, code = global.FormatDay2Timestamp(memberRetreatWaterSelf.StartTimeStr, memberRetreatWaterSelf.EndTimeStr)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	mrwsc.GetParam(listparam, ctx)
	list, count, err := memberRetreatWaterSelfBean.SearchMemberRetreatWaterSelf(memberRetreatWaterSelf, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination2(listparam, list, int64(len(list)), count))
}

//自助返水查询明细
func (mrwsc *MemberRetreatWaterSelfController) DetailMemberRetreatWaterSelf(ctx echo.Context) error {
	memberRetreatWaterSelf := new(input.DetailRetreatWaterSelf)
	code := global.ValidRequest(memberRetreatWaterSelf, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, count, err := memberRetreatWaterSelfBean.DetailMemberRetreatWaterSelf(memberRetreatWaterSelf)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}
