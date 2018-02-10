package data_merge

import "framework/render"

type Info struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *Info) GetData(siteId, siteIndexId string) (interface{}, error) {
	//m.CurrPage = 1
	m.CdnUrl = render.CdnUrl

	return m, nil
}

func (*Info) GetPage() []string {
	return []string{WAP_INFO}
}
