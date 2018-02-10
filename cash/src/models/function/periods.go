package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type PeriodsBean struct{}

//期数列表
func (*PeriodsBean) PeriodsList(this *input.PeriodsGet) ([]back.Periods, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.Periods
	periods := new(schema.SiteRebatePeriods)
	sess.Where("delete_time=?", 0)
	sess.Where("site_id=?", this.SiteId)
	err := sess.Table(periods.TableName()).Where("site_index_id=?", this.SiteIndexId).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获得单条期数数据
func (*PeriodsBean) PeriodsGetOne(this *input.PeriodsGetOne) ([]back.Periods, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.Periods
	periods := new(schema.SiteRebatePeriods)
	err := sess.Table(periods.TableName()).Where("id=?", this.PerId).
		Where("delete_time=?", 0).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//删除单条期数数据
func (*PeriodsBean) PeriodsDelete(this *input.PeriodsDeleteOne) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebatePeriods)
	periods.DeleteTime = global.GetCurrentTime()
	count, err := sess.Table(periods.TableName()).
		Where("id=?", this.PerId).
		Cols("delete_time").
		Update(periods)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//新增单条期数数据
func (*PeriodsBean) PeriodsAdd(this *input.PeriodsAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebatePeriods)
	count, err := sess.Table(periods.TableName()).Insert(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//修改单条期数数据
func (*PeriodsBean) PeriodsUpdate(this *input.PeriodsUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.SiteRebatePeriods)
	count, err := sess.Table(periods.TableName()).Where("id=?", this.Id).Update(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//退佣冲销
func (*PeriodsBean) Commission(this *input.Commission) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	periods := new(schema.SiteRebatePeriods)
	periods.Status = 0
	periods.SiteId = this.SiteId
	periods.SiteIndexId = this.SiteIndexId
	periods.Id = this.Id
	fmt.Println(periods)
	count, err := sess.Table(periods.TableName()).Where("id=?", this.Id).
		Cols("status").Update(periods)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//agency_rebate_record
	record := new(schema.AgencyRebateRecord)
	count, err = sess.Table(record.TableName()).Where("periods_id=?", this.Id).
		Where("site_id=?", this.SiteId).Where("site_index_id=?", this.SiteIndexId).Delete(record)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	sess.Commit()
	return count, err
}

//获取站点列表

func (*PeriodsBean) GetSiteList(SiteId string, SiteIndexId string) (data []back.SiteList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	periods := new(schema.Site)
	err = sess.Table(periods.TableName()).Where("id=?", SiteId).Find(&data)
	return
}
