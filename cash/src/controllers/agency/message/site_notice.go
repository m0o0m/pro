package message

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strings"
)

//公告
type SiteNotice struct {
	controllers.BaseController
}

//查询所有站点公告 根据请求链接获取不同公告数据
func (c *SiteNotice) SiteNoticeList(ctx echo.Context) error {
	reqData := new(input.NoticeList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	showUrl := ctx.Request().RequestURI
	data := strings.Split(showUrl, "?")
	url := data[0]
	switch url {
	case "/api/systerm/notice/information":
		reqData.NoticeCate = []int64{4, 5, 6, 7}
	case "/api/systerm/notice/systermNotice":
		reqData.NoticeCate = []int64{2, 3}
	case "/api/system/notice":
		reqData.NoticeCate = []int64{1}
	default:
		reqData.NoticeCate = []int64{}
	}

	listParam := new(global.ListParams)
	//获取listParam的数据
	err := c.GetParam(listParam, ctx)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	noticeList, count, err := noticeBean.NoticeList(reqData, listParam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParam, noticeList, int64(len(noticeList)), count, ctx))
}

//获取公告类型下拉(暂时只有一个类型:普通公告)
func (c *SiteNotice) SiteNoticeTypeList(ctx echo.Context) error {
	var noticeType []back.NoticeTypeList
	var noticeTypeOne back.NoticeTypeList
	noticeTypeOne.NoticeType = "普通公告"
	noticeTypeOne.NoticeCate = 1
	noticeType = append(noticeType, noticeTypeOne)
	return ctx.JSON(200, global.ReplyItem(noticeType))
}

//查询单条公告详情
func (c *SiteNotice) SiteNoticeDetails(ctx echo.Context) error {
	reqData := new(input.Notice)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	notice, err := noticeBean.NoticeDetails(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(notice))
}

//查询指定站点的普通公告
func (c *SiteNotice) GetNormalNoticeBySiteId(siteId string) {

}

//修改公告详情状态
func (c *SiteNotice) UpdateNotice(ctx echo.Context) error {
	reqData := new(input.UpdateNotice)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := noticeBean.NoticeDetails(reqData.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data.NoticeAssign == "1" {
		return ctx.JSON(200, global.ReplyError(10249, ctx))
	}
	num, err := noticeBean.UpdateNotice(reqData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//删除公告
func (c *SiteNotice) DelNotice(ctx echo.Context) error {
	reqData := new(input.DelNotice)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	num, err := noticeBean.DelNotice(reqData.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
