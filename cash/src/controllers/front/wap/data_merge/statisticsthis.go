package data_merge

import (
	"framework/render"
)

type StatisticsThis struct {
	Key
	CdnUrl string
}

func (m *StatisticsThis) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*StatisticsThis) GetPage() []string {
	return []string{WAP_StatisticsThis}
}
