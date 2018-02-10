package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//返点优惠设定
type MemberRetreatWaterSetController struct {
	controllers.BaseController
}

//优惠查询列表
func (mrwsc *MemberRetreatWaterSetController) SearchMemberRetreatWaterSet(ctx echo.Context) error {
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
	mrwsc.GetParam(listparam, ctx)
	list, count, err := memberRetreatWaterSetBean.SearchMemberRetreatWaterSet(memberRetreatWaterSet, listparam, sTime, eTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination2(listparam, list, int64(len(list)), count))
}

//获取单个列表详情
func (mrwsc *MemberRetreatWaterSetController) DetailMemberRetreatWaterSet(ctx echo.Context) error {
	memberRetreatWaterSet := new(input.GetOneRetreatWaterDetails)
	code := global.ValidRequest(memberRetreatWaterSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, count, err := memberRetreatWaterSetBean.DetailMemberRetreatWaterSet(memberRetreatWaterSet)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//返点优惠设定列表
func (mrwsc *MemberRetreatWaterSetController) ListMemberRetreatWaterSet(ctx echo.Context) error {
	memberRetreatWaterSet := new(input.RetreatWaterSetList)
	code := global.ValidRequest(memberRetreatWaterSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, count, err := memberRetreatWaterSetBean.ListMemberRetreatWaterSet(memberRetreatWaterSet)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	} else if count == 0 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//新增返点优惠设定
func (mrwsc *MemberRetreatWaterSetController) AddMemberRetreatWaterSet(ctx echo.Context) error {
	water := new(input.AddRetreatWaterSet)
	code := global.ValidRequest(water, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//商品占比不能大于百分之百
	for k := range water.Params {
		if water.Params[k].Rate > 100 {
			return ctx.JSON(200, global.ReplyError(30112, ctx))
		}
	}

	count, err := memberRetreatWaterSetBean.AddMemberRetreatWaterSet(water)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10105, ctx))
	}
	return ctx.NoContent(204)
}

//修改返点优惠设定
func (mrwsc *MemberRetreatWaterSetController) EditMemberRetreatWaterSet(ctx echo.Context) error {
	water := new(input.EditRetreatWaterSet)
	code := global.ValidRequest(water, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//商品占比不能大于百分之百
	for k := range water.Params {
		if water.Params[k].Rate > 100 {
			return ctx.JSON(200, global.ReplyError(30112, ctx))
		}
	}
	count, err := memberRetreatWaterSetBean.EditMemberRetreatWaterSet(water)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10106, ctx))
	}
	return ctx.NoContent(204)

}

//删除返点设定
func (mrwsc *MemberRetreatWaterSetController) DelMemberRetreatWaterSet(ctx echo.Context) error {
	member := new(input.DelRetreatWaterSet)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := memberRetreatWaterSetBean.DelMemberRetreatWaterSet(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10107, ctx)) //删除返点设定失败
	}
	return ctx.NoContent(204)
}
