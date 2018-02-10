package data_merge

import (
	"framework/render"
	"models/schema"
)

type Carry struct {
	Key
	CdnUrl   string //
	UrlLink  string //客服链接
	PageType int    //页面类型 1存款 2取款 3额度转换 4客服
}

func (m *Carry) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl

	//客服链接
	siteInfo, _, err := siteInfoBean.GetSingleInfo(&schema.SiteInfo{siteId, siteIndexId, "", "", "", "", "", ""})
	if err != nil {
		return nil, err
	}
	m.UrlLink = siteInfo.UrlLink
	m.PageType = 1
	return m, nil
}

func (m *Carry) GetPage() []string {
	return []string{WAP_CARRY, WAP_MEM_FOOTER}
}
