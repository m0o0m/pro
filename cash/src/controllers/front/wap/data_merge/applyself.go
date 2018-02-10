package data_merge

import (
	"framework/render"
)

type ApplySelf struct {
	Key
	CdnUrl  string
	UrlLink string
}

func (m *ApplySelf) GetData(siteId, siteIndexId string) (interface{}, error) {
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

func (*ApplySelf) GetPage() []string {
	return []string{WAP_APPLYSELF, WAP_FOOTER}
}
