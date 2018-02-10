//消息中心
package messagecenter

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
)

type MemberMessage struct {
	controllers.BaseController
}

//会员个人消息列表
func (m *MemberMessage) MessageList(ctx echo.Context) error {
	member := new(input.WapMemberMessageList)
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	member.MemberId = memberRedis.Id
	member.SiteId = memberRedis.Site
	member.SiteIndexId = memberRedis.SiteIndex
	listParams := new(global.ListParams)
	m.GetParam(listParams, ctx)
	//获取会员个人消息列表
	data, count, err := memberMessageBean.WapList(member, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//会员游戏公告列表
func (m *MemberMessage) NoticeList(ctx echo.Context) error {
	member := new(input.WapMemberNoticeList)
	//校验游戏公告类型
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	member.SiteId = memberRedis.Site
	listParams := new(global.ListParams)
	m.GetParam(listParams, ctx)
	//获取会员游戏公告列表
	data, count, err := siteNoticeBean.MemberNotice(member, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//游戏公告详情
func (m *MemberMessage) NoticeInfo(ctx echo.Context) error {
	//验证游戏公告Id参数是否合法
	member := new(input.WapMemberNoticeInfo)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, have, err := siteNoticeBean.NoticeInfo(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10207, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//个人消息详情
func (m *MemberMessage) MessageInfo(ctx echo.Context) error {
	//验证会员信息Id参数是否合法
	member := new(input.WapMemberMessageInfo)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	member.SiteId = memberRedis.Site
	member.SiteIndexId = memberRedis.SiteIndex
	member.MemberId = memberRedis.Id
	data, have, err := memberMessageBean.WapInfo(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10211, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//游戏公告删除
func (m *MemberMessage) NoticeDel(ctx echo.Context) error {
	//验证游戏公告Id参数是否合法
	member := new(input.WapMemberNoticeInfo)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断当前游戏公告Id是否存在
	have, err := siteNoticeBean.ExistNotice(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10207, ctx))
	}
	//删除该条游戏公告
	count, err := siteNoticeBean.NoticeDel(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10212, ctx))
	}
	return ctx.NoContent(204)
}

//个人消息删除
func (m *MemberMessage) MessageDel(ctx echo.Context) error {
	//验证会员Id参数是否合法
	member := new(input.WapMemberMessageDel)
	code := global.ValidRequestAdmin(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	member.SiteId = memberRedis.Site
	member.SiteIndexId = memberRedis.SiteIndex
	member.MemberId = memberRedis.Id
	//当前删除的消息是否存在
	memberInfo := new(input.WapMemberMessageInfo)
	memberInfo.Id = member.Id
	memberInfo.MemberId = member.MemberId
	memberInfo.SiteId = member.SiteId
	memberInfo.SiteIndexId = member.SiteIndexId
	_, have, err := memberMessageBean.WapInfo(memberInfo)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10211, ctx))
	}
	count, err := memberMessageBean.WapDel(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10212, ctx))
	}
	return ctx.NoContent(204)
}
