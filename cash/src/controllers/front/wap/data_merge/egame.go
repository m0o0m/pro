package data_merge

import (
	"framework/render"
)

type EGame struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *EGame) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl

	return m, nil
}

func (*EGame) GetPage() []string {
	return []string{WAP_EGAME}
}
