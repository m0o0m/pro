package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type FirstAgencyBean struct{}

//获取某个股东基本资料
func (*FirstAgencyBean) BaseInfo(this *input.FirstAgencyInfo) (*back.FirstAgencyInfo, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	data := new(back.FirstAgencyInfo)
	sess.Where(agency.TableName()+".site_id=?", this.SiteId)
	sess.Where(agency.TableName()+".id=?", this.Id)
	sess.Where(agency.TableName() + ".delete_time=0")
	sess.Where(agency.TableName() + ".level=2")
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
