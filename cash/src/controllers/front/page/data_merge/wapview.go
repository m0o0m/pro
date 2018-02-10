package data_merge

import (
	"framework/render"
	"models/input"
)

type Wapview struct {
	Key
	WapDomain  string //手机域名
	SiteName   string //站点名称
	IosUrl     string //ios下载地址
	AndroidUrl string //安卓下载地址
	CdnUrl     string //cdn地址
}

//得到首页数据
func (m *Wapview) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//查询站点域名配置
	siteDomainList, err := siteDomainBean.GetThirdDomain(&input.GetSiteDomain{siteId, siteIndexId})
	if err != nil {
		return nil, err
	}
	//查询站点
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	m.WapDomain = siteDomainList.Domain[0]
	m.SiteName = siteName
	m.IosUrl = "itms-services://?action=download-manifest&url=android_download_url/" + siteId + "_" + siteIndexId + "/" + siteId + "_" + siteIndexId + ".plist"
	m.AndroidUrl = "android_download_url/" + siteId + "_" + siteIndexId + "/pkapp.apk"
	return m, err
}

//得到页面
func (m *Wapview) GetPage() []string {
	return []string{WAPVIEW}
}

//得到<视讯电子彩票体育的页面>
func (m *Wapview) GetSubPage() map[string]string {
	return nil
}
