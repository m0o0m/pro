package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type PoundageBean struct{}

//获取手续费设定

func (*PoundageBean) PoundageGetOne(this *input.GetSiteList) (*back.Poundage, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SitePoundage)
	data := new(back.Poundage)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	have, err := sess.Table(periods.TableName()).Get(data)
	return data, have, err
}

//新增单条手续费设定
func (*PoundageBean) PeriodsAdd(this *input.PoundAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SitePoundage)
	count, err := sess.Table(periods.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改单条期数数据
func (*PoundageBean) PoundUpdate(this *input.Poundage) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SitePoundage)
	this.UpdateTime = time.Now().Unix()
	count, err := sess.Table(periods.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询设定是否存在
func (*PoundageBean) CheckSet(this *input.PoundAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SitePoundage)
	count, err := sess.Table(periods.TableName()).Where("site_id=?", this.SiteId).
		Where("site_index_id=?", this.SiteIndexId).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
