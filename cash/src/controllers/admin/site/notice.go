//[控制器] [平台]公告管理
package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strings"
)

//公告管理
type NoticeController struct {
	controllers.BaseController
}

//公告列表查询
func (c *NoticeController) GetNoticeList(ctx echo.Context) error {
	siteNoticeList := new(input.SiteNoticeList)
	code := global.ValidRequestAdmin(siteNoticeList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	c.GetParam(listParams, ctx)
	data, count, err := noticeBean.SiteNoticeList(siteNoticeList, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//公告添加
func (c *NoticeController) PostNoticeAdd(ctx echo.Context) error {
	notice := new(input.SiteNoticeAdd)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断公告推送站点是否重复
	if notice.NoticeAssign != "1" {
		ss := strings.Split(notice.NoticeAssign, ",")
		var count int
		for i := 0; i < len(ss)-1; i++ {
			for j := i + 1; j < len(ss); j++ {
				if ss[i] == ss[j] {
					count += 1
				}
			}
		}
		if count > 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	count, err := noticeBean.Add(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10206, ctx))
	}
	return ctx.NoContent(204)
}

//公告修改
func (c *NoticeController) PutNoticeUpdate(ctx echo.Context) error {
	notice := new(input.SiteNoticeUpdate)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断公告是否存在
	have, err := noticeBean.ExistNotice(notice.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10207, ctx))
	}
	count, err := noticeBean.Update(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30270, ctx))
	}
	return ctx.NoContent(204)
}

//公告状态修改
func (c *NoticeController) PutNoticeStatusUpdate(ctx echo.Context) error {
	notice := new(input.SiteNoticeState)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断公告是否存在
	have, err := noticeBean.ExistNotice(notice.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10207, ctx))
	}
	//更新公告的状态
	count, err := noticeBean.State(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30269, ctx))
	}
	return ctx.NoContent(204)
}

//公告删除
func (c *NoticeController) PutNoticeDel(ctx echo.Context) error {
	notice := new(input.SiteNoticeDel)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	for index := range notice.Ids {
		//判断公告Id是否存在
		have, err := noticeBean.ExistNotice(notice.Ids[index])
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !have {
			return ctx.JSON(200, global.ReplyError(10207, ctx))
		}
	}
	count, err := noticeBean.Del(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}
