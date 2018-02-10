package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type WebAdvBean struct{}

//站点公告弹窗管理
func (*WebAdvBean) GetWebAdvList(this *input.AdvList) ([]back.WebAdvList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	var infolist []back.WebAdvList
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	if this.Type != 0 {
		sess.Where("type=?", this.Type)
	}
	sess.Where("delete_time=?", 0)
	sess.Where("type=?", 1)
	err := sess.Table(spv.TableName()).Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	return infolist, err
}

//站点公告弹窗管理详情
func (*WebAdvBean) GetWebAdvListDetail(this *input.AdvListDetail) (*back.WebAdvList, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	info := new(back.WebAdvList)
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("id=?", this.Id)
	sess.Where("delete_time=?", 0)
	has, err := sess.Table(spv.TableName()).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//站点公告弹窗修改
func (*WebAdvBean) PutWebAdv(this *input.AdvUpdate) (int64, error) {
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

//站点公告弹窗添加
func (*WebAdvBean) PostWebAdv(this *input.AdvAdd) (int64, error) {
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

//站点公告弹窗 删除
func (*WebAdvBean) DeleteWebAdv(this *input.UpdateDeleteTime) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	this.DeleteTime = time.Now().Unix()
	sess.Where("site_id=?", this.SiteId)
	sess.Where("site_index_id=?", this.SiteIndexId)
	sess.Where("id=?", this.Id)
	count, err := sess.Table(spv.TableName()).Cols("delete_time").Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点公告弹窗 配置修改
func (*WebAdvBean) UpdateWebAdvConfig(this *input.UpdateConfig) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.Site)
	sess.Where("id=?", this.SiteId)
	sess.Where("index_id=?", this.SiteIndexId)
	count, err := sess.Table(spv.TableName()).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//站点公告弹窗 配置获取
func (*WebAdvBean) DetailWebAdvConfig(this *input.UpdateConfigDetail) (*back.WebAdvConfigDetail, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.Site)
	info := new(back.WebAdvConfigDetail)
	sess.Where("id=?", this.SiteId)
	sess.Where("index_id=?", this.SiteIndexId)
	has, err := sess.Table(spv.TableName()).Get(info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, has, err
	}
	return info, has, err
}

//查询弹窗和左下广告
func (*WebAdvBean) GetList(siteId, siteIndexId string) (infoList []back.WebAdvList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spv := new(schema.SiteAdv)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.In("type", []int{1, 2})
	sess.Where("delete_time=?", 0)
	err = sess.Table(spv.TableName()).Find(&infoList)
	return
}
