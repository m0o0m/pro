package data_merge

import (
	"framework/render"
	"models/back"
	"models/input"
	"models/schema"
)

type DrawWrite struct {
	Key
	CdnUrl    string //
	SiteName  string
	LogoUrl   string
	UrlLink   string
	Datetime  string
	PaidType  map[string]int8
	TradeDate string
	OrderNum  string
	Poundage  *back.Poundage
	RedisKey  string
	PageType  int //页面类型 1存款 2取款 3额度转换 4客服
}

func (m *DrawWrite) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	list, err := webLogo.GetWebLogoList(&input.LogoInfoList{siteId, siteIndexId})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		m.LogoUrl = list[0].LogoUrl
	}
	siteInfo, _, err := siteInfoBean.GetSingleInfo(&schema.SiteInfo{siteId, siteIndexId, "", "", "", "", "", ""})
	if err != nil {
		return nil, err
	}
	m.UrlLink = siteInfo.UrlLink
	m.SiteName = siteName
	m.PageType = 2
	return m, nil
}

func (m *DrawWrite) GetPage() []string {
	return []string{WAP_DRAW_WRTIE, WAP_MEM_FOOTER}
}
