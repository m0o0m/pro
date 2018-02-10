package data_merge

import "framework/render"

type Login struct {
	Key
	CdnUrl  string //
	UrlLink string //
}

func (m *Login) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//站点客服链接
	siteinfo, has, err := SiteInfoBean.GetSingleSiteInfo(siteId, siteIndexId)
	if err != nil {
		m.UrlLink = ""
	}
	if has == false {
		m.UrlLink = ""
	} else {
		m.UrlLink = siteinfo.UrlLink
	}
	return m, nil
}

func (m *Login) GetPage() []string {
	return []string{WAP_LOGIN}
}
