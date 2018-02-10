package function

import (
	"errors"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type NoticeBean struct {
}

//修改公告内容
func (m *NoticeBean) UpdateNotice(upData *input.UpdateNotice) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	noticeSchema := new(schema.SiteNotice)
	noticeSchema.NoticeContent = upData.Content
	noticeSchema.NoticeTitle = upData.Title
	sess.Cols("notice_content,notice_title").
		Where("id = ?", upData.Id)
	if upData.SiteId != "" {
		sess.Where("notice_assign=?", upData.SiteId)
	}
	count, err := sess.Update(noticeSchema)
	if err != nil {
		return count, err
	}
	return count, err
}

//删除公告
func (m *NoticeBean) DelNotice(id int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	noticeSchema := new(schema.SiteNotice)
	noticeSchema.ID = id
	count, err := sess.Delete(noticeSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//公告详情
func (m *NoticeBean) NoticeDetails(id int64) (back.NoticeList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	noticeSchema := new(schema.SiteNotice)
	var notice back.NoticeList
	has, err := sess.Table(noticeSchema.TableName()).
		Where("id = ?", id).
		Get(&notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return notice, err
	}
	if !has {
		err = errors.New("get notice failure")
	}
	return notice, err
}

//公告列表 noticeCate 1,站点,2系统
func (m *NoticeBean) NoticeList(this *input.NoticeList, listParams *global.ListParams) (
	noticeList []back.NoticeList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("notice_assign=?", 1)
	if this.SiteId != "" {
		sess.Or("notice_assign=?", this.SiteId)
	}
	sess.Select("id,notice_title,notice_date,notice_state,notice_content")
	sess.In("notice_cate", this.NoticeCate)
	if this.NoticeContent != "" {
		sess.Where("notice_content like ?", "%"+this.NoticeContent+"%")
	}
	sess.Where("delete_time=?", 0)
	sess.Desc("notice_date")
	noticeSchema := new(schema.SiteNotice)
	listParams.Make(sess)
	conds := sess.Conds()
	err = sess.Table(noticeSchema.TableName()).
		Find(&noticeList)
	if err != nil {
		return
	}
	count, err = sess.Table(noticeSchema.TableName()).Where(conds).Count()
	return
}

//获取站点公告列表
func (m *NoticeBean) SiteNoticeList(this *input.SiteNoticeList, listParams *global.ListParams) (siteNoticeList []*back.SiteNoticeList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	noticeSchema := new(schema.SiteNotice)
	listParams.Make(sess)
	sess.Where("delete_time=?", 0)
	if this.StartTime != "" {
		loc, _ := time.LoadLocation("Local")
		start, _ := time.ParseInLocation("2006-01-02 15:04:05", this.StartTime, loc)
		sess.Where("notice_date>=?", start.Unix())
	}
	if this.EndTime != "" {
		loc, _ := time.LoadLocation("Local")
		end, _ := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime, loc)
		sess.Where("notice_date<=?", end.Unix())
	}
	if this.NoticeContent != "" { //公告内容的模糊搜索
		sess.Where("notice_content like ?", "%"+this.NoticeContent+"%")
	}
	conds := sess.Conds()
	err = sess.Table(noticeSchema.TableName()).Find(&siteNoticeList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteNoticeList, count, err
	}
	count, err = sess.Table(noticeSchema.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return siteNoticeList, count, err
	}
	return siteNoticeList, count, err
}

//站点添加公告
func (m *NoticeBean) Add(this *input.SiteNoticeAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := &schema.SiteNotice{
		NoticeCate:    this.NoticeCate,
		NoticeContent: this.NoticeContent,
		NoticeState:   this.NoticeState,
		NoticeTitle:   this.NoticeTitle,
		NoticeAssign:  this.NoticeAssign,
	}
	count, err := sess.Insert(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点公告Id是否存在
func (m *NoticeBean) ExistNotice(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	notice.ID = id
	sess.Where("delete_time=?", 0)
	have, err := sess.Table(notice.TableName()).Exist(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return have, err
	}
	return have, err
}

//站点修改公告的状态(启用和关闭)
func (m *NoticeBean) State(this *input.SiteNoticeState) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	notice.NoticeState = this.Status
	count, err := sess.Where("id=?", this.Id).Cols("notice_state").Update(notice)
	return count, err
}

//更新站点公告的信息
func (m *NoticeBean) Update(this *input.SiteNoticeUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	notice := new(schema.SiteNotice)
	notice.NoticeAssign = this.NoticeAssign
	notice.NoticeCate = this.NoticeCate
	notice.NoticeTitle = this.NoticeTitle
	notice.NoticeContent = this.NoticeContent

	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	sess.Cols("notice_assign", "notice_cate", "notice_title", "notice_content")
	count, err := sess.Update(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//批量删除站点公告信息
func (m *NoticeBean) Del(this *input.SiteNoticeDel) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	notice.DeleteTime = time.Now().Unix()
	sess.Where("delete_time=?", 0)
	count, err := sess.In("id", this.Ids).Cols("delete_time").Update(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//根据站点信息获取公告
func (m *NoticeBean) GetNoticeBySiteId(siteId string) (noticeContent []*back.SiteNoticeContent, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	err = sess.Table(notice.TableName()).
		Select("notice_content"). //公告内容
		Where("delete_time = ?", 0).
		Where("notice_cate = ?", 1).        //普通公告
		Where("notice_assign = ?", siteId). //推送站点
		Find(&noticeContent)
	return
}

//根据站点信息获取公告和时间
func (m *NoticeBean) GetNoticeAndDateBySiteId(siteId string) (noticeContent []*back.NoticeData, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	err = sess.Table(notice.TableName()).
		Select("notice_content,notice_date"). //公告内容,时间
		Where("delete_time = ?", 0).
		Where("notice_cate = ?", 1).        //普通公告
		Where("notice_assign = ?", siteId). //推送站点
		Find(&noticeContent)
	return
}

//根据站点信息和公告类型获取列表数据
func (m *NoticeBean) MemberNotice(this *input.WapMemberNoticeList, listParams *global.ListParams) ([]*back.MemberNoticeList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	sess.Where("notice_cate=?", this.Mtype)
	sess.Where("notice_assign like ? or notice_assign=? ", "%"+this.SiteId+"%", 1)
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	sess.Table(notice.TableName())
	data := make([]*back.MemberNoticeList, 0)
	err := sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(notice.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//根据Id获取公告详情
func (m *NoticeBean) NoticeInfo(this *input.WapMemberNoticeInfo) (*back.MemberNoticeInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	sess.Table(notice.TableName())
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", this.Id)
	data := new(back.MemberNoticeInfo)
	have, err := sess.Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, have, err
	}
	return data, have, err

}

//根据Id删除公告
func (m *NoticeBean) NoticeDel(this *input.WapMemberNoticeInfo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	sess.Table(notice.TableName())
	sess.Where("delete_time=?", 0)
	sess.Where("id=?", this.Id)
	notice.DeleteTime = time.Now().Unix()
	count, err := sess.Cols("delete_time").Update(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//获取公告列表
func (m *NoticeBean) GetNoticeList(siteId string, Id int64, times *global.Times, listparams *global.ListParams) (data []back.WapSiteNoticeList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	sess.Where("notice_assign=?", siteId).Or("notice_assign=?", 1)
	if Id != 0 {
		sess.Where("notice_cate=?", Id)
	} else {
		sess.In("notice_cate", 4, 5, 6, 7)
	}

	if times != nil {
		times.Make("notice_date", sess)
	}
	sess.Where("delete_time=?", 0)
	sess.OrderBy("id desc")
	conds := sess.Conds()
	listparams.Make(sess)
	sess.Table(notice.TableName()).
		Select("id,notice_title,notice_content,notice_cate,notice_date,notice_assign,notice_state").
		Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(notice.TableName()).Where(conds).Count()
	return
}

//获取wap公告列表
func (m *NoticeBean) GetWapNoticeList(siteId string, Id int64, times *global.Times) (data []back.WapSiteNoticeList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	notice := new(schema.SiteNotice)
	sess.Where("notice_assign=?", siteId).Or("notice_assign=?", 1)
	if Id != 0 {
		sess.Where("notice_cate=?", Id)
	} else {
		sess.In("notice_cate", 4, 5, 6, 7)
	}

	if times != nil {
		times.Make("notice_date", sess)
	}
	sess.Where("delete_time=?", 0)
	sess.OrderBy("id desc")
	sess.Table(notice.TableName()).
		Select("id,notice_title,notice_content,notice_cate,notice_date,notice_assign,notice_state").
		Find(&data)
	if err != nil {
		return
	}
	return
}
