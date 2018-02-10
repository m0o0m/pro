package function

import (
	"global"
	"models/schema"
)

type SitePaySetBean struct {
}

//根据站点Id和站点前台Id获取站点支付设定
func (*SitePaySetBean) GetPaySetInfo(siteId, siteIndexId string) (*schema.SitePaySet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paySet := new(schema.SitePaySet)
	sess.Table(paySet.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId)
	have, err := sess.Get(paySet)
	return paySet, have, err
}

//根据Id获取站点支付设定
func (*SitePaySetBean) GetPaySetInfoById(id int64) (*schema.SitePaySet, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	paySet := new(schema.SitePaySet)
	sess.Table(paySet.TableName()).
		Where("id=?", id)
	have, err := sess.Get(paySet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return paySet, have, err
	}
	return paySet, have, err
}
