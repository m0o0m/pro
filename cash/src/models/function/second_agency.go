package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type SecondAgencyBean struct{}

//获取某个总代基本资料
func (*SecondAgencyBean) BaseInfo(this *input.SecondAgencyInfo) (*back.SecondAgencyInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	data := new(back.SecondAgencyInfo)
	sess.Where(agency.TableName()+".site_id=?", this.SiteId)
	sess.Where(agency.TableName()+".id=?", this.Id)
	sess.Where(agency.TableName() + ".delete_time=0")
	sess.Where(agency.TableName() + ".level=3")
	if this.SiteIndexId != "" {
		sess.Where(agency.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	ok, err := sess.Table(agency.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}

//设置总代状态[启用/禁用]
func (*SecondAgencyBean) UpdateStatus(this *input.SecondAgencyInfo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	/*
		sess.Where("site_id = ?", this.SiteId)
		sess.Where("site_index_id = ?", this.SiteIndexId)
	*/
	has, err := sess.Table(agency.TableName()).Where("id = ?", this.Id).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, err
	}
	if !has {
		return 0, err
	}
	status := agency.Status
	if status == 1 {
		agency.Status = 2
	} else if status == 2 {
		agency.Status = 1
	}
	count, err := sess.Table(agency.TableName()).
		Where("id = ?", this.Id).
		Cols("status").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
