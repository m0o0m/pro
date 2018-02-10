package function

import (
	"global"
	"time"

	"models/back"
	"models/input"
	"models/schema"
)

type SiteDomainBean struct{}

//获取该站点下面三个不同的站点域名集合
func (*SiteDomainBean) GetThirdDomain(selec *input.GetSiteDomain) (info back.DomainBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	site := new(schema.SiteDomain)
	sess.Where("site_id=?", selec.Site)
	sess.Where("site_index_id=?", selec.SiteIndex)
	sess.Where("delete_time=?", 0)

	var siteList []schema.SiteDomain
	err = sess.Table(site.TableName()).Find(&siteList)
	lenSite := len(siteList)
	for i := 0; i < lenSite; i++ {
		info.Domain = append(info.Domain, siteList[i].Domain)
	}
	return
}

//获取某个站点下面的域名个数
func (*SiteDomainBean) GetSiteDomain(siteId, indexId string) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	siteDomain := new(schema.SiteDomain)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", indexId)
	sess.Where("delete_time=?", 0)
	count, err = sess.Table(siteDomain.TableName()).Count()
	return
}

//Add  站点域名配置model层增加函数
func (*SiteDomainBean) Add(this *schema.SiteDomain) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.InsertOne(this)
	return
}

//判断是否主域名
func (*SiteDomainBean) IsMainDomain(siteId, siteIndexId string) (flag bool, err error) {
	sess := global.GetXorm()
	siteDomain := new(schema.SiteDomain)
	flag, err = sess.Table(siteDomain.TableName()).
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).Exist()
	return
}

//Delete 站点域名配置model层删除函数
func (*SiteDomainBean) Delete(this *schema.SiteDomain) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id > 0 {
		sess.Where("id=?", this.Id)
	}
	this.DeleteTime = time.Now().Unix()
	count, err = sess.Cols("delete_time").Update(this)
	return
}

//Edit 站点域名配置编辑
func (*SiteDomainBean) Edit(this *schema.SiteDomain) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id > 0 {
		sess.Where("id=?", this.Id)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	count, err = sess.Cols("domain").Update(this)
	return
}

//SiteDomainList  站点域名配置列表
func (*SiteDomainBean) DomainSiteList(this *input.DomainSiteList, listParam *global.ListParams) (infolist []back.SiteUpList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	domains := new(schema.SiteDomain)
	if this.Domain != "" {
		sess.Where("domain like ?", "%"+this.Domain+"%")
	}
	if this.Site != "" {
		sess.Where("site_id=?", this.Site)
	}
	if this.SiteIndex != "" {
		sess.Where("site_index_id=?", this.SiteIndex)
	}
	sess.Where("delete_time=?", 0)
	conds := sess.Conds()
	//获得分页记录
	listParam.Make(sess)
	err = sess.Table(domains.TableName()).Find(&infolist)
	count, err = sess.Table(domains.TableName()).Where(conds).Count()
	return
}

//GetOneSiteDomain  获取单个的站点配置
func (*SiteDomainBean) GetOneSiteDomain(this *schema.SiteDomain) (info back.InfoBack, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	if this.Id > 0 {
		sess.Where("id=?", this.Id)
	}
	flag, err = sess.Table(this.TableName()).
		Where("is_used=?", 2).Where("delete_time=?", 0).Get(&info)
	return
}

//判断域名是否已经使用
func (*SiteDomainBean) ExistDomain(this *schema.SiteDomain) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	domainSite := new(schema.SiteDomain)
	if this.Domain != "" {
		sess.Where("domain=?", this.Domain)
	}
	flag, err := sess.Get(domainSite)
	return flag, err
}

//判断域名是否已经使用
func (*SiteDomainBean) IsExistDomain(domain string) (bool, error) {
	if domain == "" {
		return true, nil
	}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sd := new(schema.SiteDomain)
	ok, err := sess.Where("domain=?", domain).Get(sd)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return ok, err
}

//根据域名获取站点信息
func (*SiteDomainBean) GetSiteByDomain(domain string) (info schema.SiteDomain, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("domain=?", domain).
		Where("delete_time=?", 0).
		Where("type=?", 1).
		Get(&info)
	return
}

//根据后台域名获取站点信息
func (*SiteDomainBean) GetSiteInfoByDomian(domian string) (info schema.SiteDomain, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("domain=?", domian)
	sess.Where("delete_time=?", 0)
	sess.In("type", []int8{2, 3})
	flag, err = sess.Get(&info)
	if err != nil {
		global.GlobalLogger.Error("GetSiteInfoByDomian error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, nil
}
