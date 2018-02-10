package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SiteActivityBean struct {
}

//会员活动中心(wap活动列表)
func (*SiteActivityBean) WapActivityList(this *input.WapActivity, listParams *global.ListParams) ([]*back.WapActivityList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	activity := new(schema.SiteActivity)
	sess.Where("delete_time=?", 0).
		//如果字段是mysql的关键字需要使用``或者表名.字段
		//Where(activity.TableName()+".from=?", 1).
		Where("`from`=?", 1).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId)
	conds := sess.Conds()
	listParams.Make(sess)
	data := make([]*back.WapActivityList, 0)
	sess.Table(activity.TableName())
	err := sess.Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(activity.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//活动Id主键是否存在
func (*SecondAgencyBean) WapActivityId(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	activity := new(schema.SiteActivity)
	sess.Table(activity.TableName())
	sess.Where("delete_time=?", 0)
	activity.Id = id
	have, err := sess.Exist(activity)
	return have, err
}

//单个活动详情信息
func (*SiteActivityBean) WapActivityInfo(this *input.WapActivityInfo) (*back.WapActivityInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	activity := new(schema.SiteActivity)
	sess.Table(activity.TableName())
	sess.Where("delete_time=?", 0).
		Where("`from`=?", 1).
		Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).
		Where("id=?", this.Id)
	data := new(back.WapActivityInfo)
	have, err := sess.Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, have, err
	}
	return data, have, err
}
