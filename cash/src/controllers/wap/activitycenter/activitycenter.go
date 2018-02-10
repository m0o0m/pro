//活动中心
package activitycenter

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
)

type MemberActivity struct {
	controllers.BaseController
}

//会员活动列表
func (m *MemberActivity) WapActivityList(ctx echo.Context) error {
	activity := new(input.WapActivity)
	code := global.ValidRequestMember(activity, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	m.GetParam(listParams, ctx)
	//获取会员活动列表
	data, count, err := siteActivityBean.WapActivityList(activity, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//TODO目前图片放的是路径,需要获取图片的base64码。
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//会员活动详情
func (m *MemberActivity) WapActivityInfo(ctx echo.Context) error {
	//验证会员活动Id参数是否合法
	activity := new(input.WapActivityInfo)
	code := global.ValidRequestMember(activity, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	activity.Id = memberRedis.Id
	activity.SiteId = memberRedis.Site
	activity.SiteIndexId = memberRedis.SiteIndex
	//获取会员活动详情
	data, have, err := siteActivityBean.WapActivityInfo(activity)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10213, ctx))
	}
	//TODO目前图片放的是路径,需要获取图片的base64码。
	return ctx.JSON(200, global.ReplyItem(data))
}
