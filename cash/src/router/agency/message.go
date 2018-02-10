package agency

import (
	"controllers/agency/message"
	"github.com/labstack/echo"
	"router"
)

//消息中心
func MessageRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	//e := c.Group("")
	//
	//站点公告
	siteNotice := new(message.SiteNotice)
	e.GET("/systerm/notice/information", siteNotice.SiteNoticeList)   //查询站点公告列表
	e.GET("/notice/details", siteNotice.SiteNoticeDetails)            //查询站点公告详情
	e.POST("/notice/update", siteNotice.UpdateNotice)                 //修改公告详情 修改为已读
	e.DELETE("/notice/del", siteNotice.DelNotice)                     //删除公告
	e.GET("/systerm/notice/systermNotice", siteNotice.SiteNoticeList) //系统公告
	e.GET("/system/notice", siteNotice.SiteNoticeList)                //最新消息 公告管理
	e.GET("/add/type/drop", siteNotice.SiteNoticeTypeList)            //公告类型下拉
	//消息提醒
	memberMessage := new(message.MemberMessage)
	e.GET("/systerm/memberNews", memberMessage.MessageList) //查询消息列表
	e.GET("/message/details", memberMessage.MessageDetails) //查询消息详情
	e.POST("/message/update", memberMessage.UpdateMessage)  //修改消息详情 修改为已读
	e.DELETE("/message/del", memberMessage.DelMessage)      //删除消息

}
