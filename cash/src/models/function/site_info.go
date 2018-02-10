package function

import (
	"global"
	"models/schema"
)

type SiteInfoBean struct{}

//增加
func (*SiteInfoBean) AddInfo(siteInfo *schema.SiteInfo) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.InsertOne(&siteInfo)
	return
}

//修改
func (*SiteInfoBean) EditInfo(siteInfo *schema.SiteInfo) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("id=?", siteInfo.SiteId)
	sess.Where("site_index_id=?", siteInfo.SiteIndexId)
	count, err = sess.Update(&siteInfo)
	return
}

//获取
func (*SiteInfoBean) GetSingleInfo(siteInfo *schema.SiteInfo) (info schema.SiteInfo, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("site_id=?", siteInfo.SiteId)
	sess.Where("site_index_id=?", siteInfo.SiteIndexId)
	flag, err = sess.Get(&info)
	return
}

//获取单条
func (*SiteInfoBean) GetSingleSiteInfo(siteId, siteIndexId string) (info schema.SiteInfo, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	flag, err = sess.Get(&info)
	return
}
