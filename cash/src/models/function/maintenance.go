package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type MaintenanceBean struct{}

//查询列表
func (*MaintenanceBean) GetList(list []string) (data []back.DomainList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitedomain := new(schema.SiteDomain)
	err = sess.Table(sitedomain.TableName()).
		Select("`id`, `site_id`, `site_index_id`, `domain`,  `is_used`").
		In("site_index_id ", list).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取后台站点域名
func (*MaintenanceBean) GetDomainList(site back.ConditionList) (data []back.DomainInfoList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitedomain := new(schema.SiteDomain)
	siteinfo := new(schema.SiteInfo)
	if site.SiteId != "" {
		sess.Where("d.site_id=?", site.SiteId)
	}
	if site.SiteIndexId != "" {
		sess.Where("d.site_index_id=?", site.SiteIndexId)
	}
	sess.Where("d.delete_time=?", site.DeleteTime)
	err = sess.Table(sitedomain.TableName()).Alias("d").
		Join("LEFT", []string{siteinfo.TableName(), "i"}, "d.site_id=i.site_id and d.site_index_id=i.site_index_id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//获取视讯 电子下级菜单
func (*MaintenanceBean) GetInfoList(this *input.InfoList, typeId int64) (data []back.InfoList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)

	err = sess.Table(product.TableName()).Where("type_id=?", typeId).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}
