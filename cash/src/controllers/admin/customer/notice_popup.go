//[控制器] [平台]公告弹窗
package customer

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
)

//公告弹窗管理
type NoticePopupController struct {
	controllers.BaseController
}

//获取站点公告弹窗设置列表
func (c *NoticePopupController) GetNoticePopupSet(ctx echo.Context) error {
	notice := new(input.GetNotice)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := noticePopupBean.GetNoticePopupSet(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyItem(list))
}

//获取站点公告弹窗设置详情
func (c *NoticePopupController) GetNoticePopupSetInfo(ctx echo.Context) error {
	notice := new(input.NoticeInfo)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, has, err := noticePopupBean.GetNoticePopupSetInfo(notice.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30223, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//添加公告弹窗
func (c *NoticePopupController) PostNoticePopupAdd(ctx echo.Context) error {
	notice := new(input.NoticeAdd)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看站点是否存在
	has, err := noticePopupBean.IsExistSite(notice.SiteId, notice.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	notice.Type = 1
	count, err := noticePopupBean.NoticeAdd(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30192, ctx))
	}
	return ctx.NoContent(204)
}

//公告弹窗列表
func (c *NoticePopupController) NoticePopupList(ctx echo.Context) error {
	site := new(input.NoticePopupList)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	list, count, err := noticePopupBean.SiteList(site, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//站点公告弹窗设置修改
func (c *NoticePopupController) PutNoticePopupSet(ctx echo.Context) error {
	notice := new(input.NoticeEdit)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	_, has, err := noticePopupBean.GetNoticePopupSetInfo(notice.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30224, ctx))
	}
	count, err := noticePopupBean.EditNoticePopupSet(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30277, ctx))
	}
	return ctx.NoContent(204)
}

//站点公告弹窗设置删除
func (c *NoticePopupController) DelNoticePopupSet(ctx echo.Context) error {
	notice := new(input.NoticePopupDel)
	code := global.ValidRequestAdmin(notice, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	_, has, err := noticePopupBean.GetNoticePopupSetInfo(notice.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30224, ctx))
	}
	count, err := noticePopupBean.DelNoticePopupSet(notice)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30220, ctx))
	}
	return ctx.NoContent(204)
}

//站点公告弹窗配置查询
func (c *NoticePopupController) GetNoticePopupConfig(ctx echo.Context) error {
	site := new(input.GetNoticePopupSet)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, _, err := noticePopupBean.GetNoticePopupConfig(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyItem(data))
}

//站点公告弹窗配置修改
func (c *NoticePopupController) PutNoticePopupConfig(ctx echo.Context) error {
	site := new(input.PutNoticePopupSet)
	code := global.ValidRequestAdmin(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断存不存在
	info, flag, err := siteOperateBean.GetSingleSite(site.SiteIndexId, site.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag || info.Id == "" {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//修改
	count, err := noticePopupBean.PutNoticePopupConfig(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30276, ctx))
	}
	return ctx.NoContent(204)
}

//H5动画设置查询
func (c *NoticePopupController) GetSiteH5Set(ctx echo.Context) error {
	siteH5Set := new(input.SiteH5Set)
	code := global.ValidRequestAdmin(siteH5Set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	list, count, err := noticePopupBean.GetSiteH5Set(siteH5Set, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//h5动画设置修改
func (c *NoticePopupController) PutSiteH5Set(ctx echo.Context) error {
	putSiteH5Set := new(input.PutSiteH5Set)
	code := global.ValidRequestAdmin(putSiteH5Set, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	count, err := noticePopupBean.PutSiteH5Set(putSiteH5Set)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30275, ctx))
	}
	return ctx.NoContent(204)
}

//富文本内容接收
func (c *NoticePopupController) Editor(ctx echo.Context) error {
	CKEditorFuncNum := ctx.FormValue("CKEditorFuncNum")
	img, err := ctx.FormFile("upload")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	imgFile, err := global.ReadByte(img)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, "<script type='text/javascript'>window.parent.CKEDITOR.tools.callFunction("+CKEditorFuncNum+",'"+string(imgFile)+"','success');</script>")
}
