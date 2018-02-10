package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type WebPopBean struct{}

//站点广告弹窗管理
func (*WebPopBean) GetWebPopList(this *input.AdvListBySite, listparam *global.ListParams) ([]back.WebPopList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	var infolist []back.WebPopList
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.Title != "" {
		sess.Where("title=?", this.Title)
	}
	sess.Where("delete_time=?", 0)
	sess.Where("type=?", 2)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	err := sess.Table(spv.TableName()).Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, 0, err
	}
	count, err := sess.Table(spv.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, count, err
	}
	return infolist, count, err
}

//站点广告弹窗管理详情
func (*WebPopBean) GetWebPopListDetail(this *input.AdvListBySiteDetail) (*back.WebPopList, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("delete_time=?", 0)
	info := new(back.WebPopList)
	has, err := sess.Table(spv.TableName()).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//站点广告 修改
func (*WebPopBean) PutWebPop(this *input.PopUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	count, err := sess.Table(spv.TableName()).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点广告 添加
func (*WebPopBean) PostWebPop(this *input.PopAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	this.AddTime = global.GetCurrentTime()
	spv := new(schema.SiteAdv)
	count, err := sess.Table(spv.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点广告 状态修改
func (*WebPopBean) UpdateWebPopStatus(this *input.UpdatePopStatus) (int64, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", this.Id)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	spv := new(schema.SiteAdv)
	has, err := sess.Get(spv)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, has, err
	}
	if !has {
		return 0, has, err
	}
	switch this.State {
	case 1:
		this.State = 2
	case 2:
		this.State = 1
	}
	data, err := sess.Table(spv.TableName()).Where(conds).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, has, err
	}
	return data, has, err
}
