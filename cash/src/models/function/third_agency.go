package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type ThirdAgencyBean struct{}

//获取某个代理基本资料
func (*ThirdAgencyBean) BaseInfo(this *input.ThirdAgencyInfo) (*back.ThirdAgencyInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()

	agency := new(schema.Agency)
	data := new(back.ThirdAgencyInfo)
	sess.Where(agency.TableName()+".site_id=?", this.SiteId)
	sess.Where(agency.TableName()+".id=?", this.Id)
	sess.Where(agency.TableName() + ".delete_time=0")
	sess.Where(agency.TableName() + ".level=4")
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
func (*ThirdAgencyBean) UpdateStatus(this *input.ThirdAgencyInfo) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	/*
		sess.Where("site_id = ?", this.SiteId)
		sess.Where("site_index_id = ?", this.SiteIndexId)
	*/
	_, err := sess.Table(agency.TableName()).Where("id = ?", this.Id).Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
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

//根据代理账号获取代理信息
func (*ThirdAgencyBean) AgencyNameInfo(Account, SiteId, SiteIndexId string) (*back.ThirdAgencyInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	data := new(back.ThirdAgencyInfo)
	agency := new(schema.Agency)
	data = new(back.ThirdAgencyInfo)
	sess.Where("site_id=?", SiteId)
	sess.Where("account=?", Account)
	sess.Where("delete_time=0")
	sess.Where("level=4")
	if SiteIndexId != "" {
		sess.Where("site_index_id=?", SiteIndexId)
	}
	ok, err := sess.Table(agency.TableName()).Get(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, ok, err
	}
	return data, ok, err
}
