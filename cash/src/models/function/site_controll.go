package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type SiteControllBean struct{}

//站点列表
func (*SiteControllBean) SiteManageList(this *input.SiteManageList, listparams *global.ListParams) ([]back.SiteControllListBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	agency := new(schema.Agency)
	sd := new(schema.SiteDomain)
	cb := new(schema.Combo)
	data := make([]back.SiteControllListBack, 0)
	var (
		count int64
		err   error
	)
	if this.Id != 0 {
		sess.Where(site.TableName()+".agency_id=?", this.Id)
	}
	if this.SiteId != "" {
		sess.Where(site.TableName()+".id=?", this.SiteId)
	}
	if this.Status != 0 {
		sess.Where(site.TableName()+".status=?", this.Status)
	}
	if this.SiteName != "" {
		sess.Where(site.TableName()+".site_name=?", this.SiteName)
	}
	if this.SiteDomain != "" {
		sess.Where(sd.TableName()+".domain=?", this.SiteDomain)
	}
	sess.Where(site.TableName() + ".delete_time=0")
	conds := sess.Conds()
	listparams.Make(sess)
	//只显示默认站点
	err = sess.Table(site.TableName()).Where(site.TableName()+".is_default=1").
		Join("LEFT", sd.TableName(), site.TableName()+".id="+sd.TableName()+
			".site_id AND "+site.TableName()+".index_id="+sd.TableName()+
			".site_index_id AND "+sd.TableName()+".is_default=1").
		Join("LEFT", agency.TableName(), site.TableName()+".agency_id="+
			agency.TableName()+".id").
		Join("LEFT", cb.TableName(), site.TableName()+".combo_id="+
			cb.TableName()+".id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(site.TableName()).Where(site.TableName()+".is_default=1").
		Join("LEFT", sd.TableName(), site.TableName()+".id="+sd.TableName()+
			".site_id AND "+site.TableName()+".index_id="+sd.TableName()+
			".site_index_id AND "+sd.TableName()+".is_default=1").
		Join("LEFT", agency.TableName(), site.TableName()+".agency_id="+
			agency.TableName()+".id").
		Join("LEFT", cb.TableName(), site.TableName()+".combo_id="+
			cb.TableName()+".id").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//站点列表
func (*SiteControllBean) SiteManageListDataA() ([]schema.Site, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	var datA []schema.Site
	err := sess.Table(site.TableName()).
		Where("delete_time=?", 0).
		Find(&datA)
	return datA, err
}

//查询站点在agency表中是否被使用
func (*SiteControllBean) ManageBySiteSiteIndexID(this *input.SiteManageStatus) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	agency := new(schema.Agency)
	has, err = sess.Table(agency.TableName()).Where("delete_time=?", 0).
		Get(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//查询站点在site表中是否被使用
func (*SiteControllBean) ManageBySiteSiteIndexIdStatus(this *input.SiteManageStatus) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("index_id=?", this.SiteIndexId)
	}
	site := new(schema.Site)
	has, err = sess.Where("delete_time=?", 0).
		Get(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//站点状态修改
func (*SiteControllBean) SiteManageStatus(this *input.SiteManageStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.Status = this.Status
	count, err = sess.Table(site.TableName()).Where("id=?", this.SiteId).
		Where("index_id=?", this.SiteIndexId).Cols("status").Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//删除站点
func (*SiteControllBean) SiteManageDelete(this *input.SiteManageStatus) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.Site)
	site.DeleteTime = time.Now().Unix()
	count, err = sess.Table(site.TableName()).Where("id=?", this.SiteId).
		Where("index_id=?", this.SiteIndexId).Cols("delete_time").
		Update(site)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询一条站点
func (*SiteControllBean) BeOneSite(this *input.SiteManageStatus) (has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	site := new(schema.Site)
	has, err = sess.Table(site.TableName()).Where("delete_time=?", 0).Get(site)
	return
}
