package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type NoticePopupBean struct{}

//H5动画设置查询
func (*NoticePopupBean) GetSiteH5Set(this *input.SiteH5Set, listParams *global.ListParams) (list []back.SiteH5Set, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	if this.SiteId != "" {
		sess.Where("id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("index_id=?", this.SiteIndexId)
	}
	if this.Status != 0 {
		sess.Where("h5_state_switch = ?", this.Status)
	}
	conds := sess.Conds()
	listParams.Make(sess)
	err = sess.Table(site.TableName()).Find(&list)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	count, err = sess.Table(site.TableName()).Where(conds).Count()
	return
}

//H5动画设置修改
func (*NoticePopupBean) PutSiteH5Set(this *input.PutSiteH5Set) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.H5StateSwitch = this.Status

	if len(this.Site) != 0 {
		for i := range this.Site {
			if i == 0 {
				sess.Where("id = ? AND index_id=?", this.Site[i].SiteId, this.Site[i].SiteIndexId)
			} else {
				sess.Or("id = ? AND index_id=?", this.Site[i].SiteId, this.Site[i].SiteIndexId)
			}

		}
	}
	count, err = sess.Cols("h5_state_switch").Update(site)
	if err != nil {
		return count, err
	}
	return count, err
}

//查看弹窗广告配置
func (*NoticePopupBean) GetNoticePopupConfig(this *input.GetNoticePopupSet) (data *back.NoticePopupConfig, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	data = new(back.NoticePopupConfig)
	has, err = sess.Table(site.TableName()).Where("id = ?", this.SiteId).
		Where("index_id=?", this.SiteIndexId).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}

//修改弹窗广告配置
func (*NoticePopupBean) PutNoticePopupConfig(this *input.PutNoticePopupSet) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.PopoverBgColor = this.PopoverBgColor
	site.PopoverBarColor = this.PopoverBarColor
	site.PopoverTitleColor = this.PopoverTitleColor
	count, err = sess.Where("id = ?", this.SiteId).Where("index_id=?", this.SiteIndexId).
		Cols("popover_bg_color,popover_bar_color,popover_title_color").Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//公告弹窗列表
func (*NoticePopupBean) SiteList(this *input.NoticePopupList, listParams *global.ListParams) (sites []back.BackSiteDrop, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	if this.SiteId != "" {
		sess.Where("id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("index_id=?", this.SiteIndexId)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	listParams.Make(sess)
	err = sess.Table(site.TableName()).Find(&sites)
	if err != nil {
		return
	}
	count, err = sess.Table(site.TableName()).Where(conds).Count()
	return
}

//获取站点公告弹窗设置
func (*NoticePopupBean) GetNoticePopupSet(this *input.GetNotice) (notices []back.Notice, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAdv := new(schema.SiteAdv)
	err = sess.Table(siteAdv.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Where("delete_time=0").Find(&notices)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return notices, err
	}
	return notices, err
}

//获取站点广告详情
func (*NoticePopupBean) GetNoticePopupSetInfo(id int64) (notice *back.NoticeInfo, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAdv := new(schema.SiteAdv)
	notice = new(back.NoticeInfo)
	has, err = sess.Table(siteAdv.TableName()).Where("id=?", id).
		Where("delete_time=0").Get(notice)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return notice, has, err
	}
	return notice, has, err
}

//编辑站点公告弹窗设置
func (*NoticePopupBean) EditNoticePopupSet(this *input.NoticeEdit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAdv := new(schema.SiteAdv)
	siteAdv.Title = this.Title
	siteAdv.Type = this.Type
	siteAdv.Content = this.Content
	count, err = sess.Where("id=?", this.Id).Cols("title,type,content").
		Update(siteAdv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除站点公告弹窗设置
func (*NoticePopupBean) DelNoticePopupSet(this *input.NoticePopupDel) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAdv := new(schema.SiteAdv)
	siteAdv.DeleteTime = time.Now().Unix()
	count, err = sess.Where("id=?", this.Id).Cols("delete_time").Update(siteAdv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//增加公告
func (*NoticePopupBean) NoticeAdd(this *input.NoticeAdd) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteAdv := new(schema.SiteAdv)
	siteAdv.SiteId = this.SiteId
	siteAdv.SiteIndexId = this.SiteIndexId
	siteAdv.Title = this.Title
	siteAdv.Type = this.Type
	siteAdv.Content = this.Content
	siteAdv.AddTime = time.Now().Unix()
	siteAdv.State = 2
	count, err = sess.InsertOne(siteAdv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查看站点是否存在
func (*NoticePopupBean) IsExistSite(siteId, siteIndexId string) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	has, err = sess.Where("id=?", siteId).Where("index_id=?", siteIndexId).
		Where("delete_time=0").Get(site)
	if err != nil {
		return has, err
	}
	return has, err
}
