package page

import (
	"models/function"
)

var (
	noticeBean         = new(function.NoticeBean)         //公告
	siteBean           = new(function.SiteOperateBean)    //站点信息
	registerStatusBean = new(function.RegisterStatusBean) //注册状态判断
	redPacketSetBean   = new(function.RedPacketSetBean)   //红包
	memberBean         = new(function.MemberBean)         //会员
	member_level_bean  = new(function.MemberLevelBean)    //会员层级
)
