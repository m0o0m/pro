package data_merge

import "framework/render"

type Account struct {
	Key
	CdnUrl   string
	CurrPage int64  //当前页面
	UrlLink  string //客服链接
	IsSelf   int    //自助反水开关
}

func (m *Account) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CurrPage = 4
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

func (*Account) GetPage() []string {
	return []string{WAP_ACCOUNT, WAP_FOOTER}
}
