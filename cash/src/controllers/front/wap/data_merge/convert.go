package data_merge

import "framework/render"

type Convert struct {
	Key
	CdnUrl   string
	UrlLink  string
	PageType int //页面类型 1存款 2取款 3额度转换 4客服
}

func (m *Convert) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	m.PageType = 3
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

func (*Convert) GetPage() []string {
	return []string{WAP_CONVERT, WAP_MEM_FOOTER}
}
