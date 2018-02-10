package data_merge

import (
	"framework/render"
)

type StatisticsLast struct {
	Key
	CdnUrl string
}

func (m *StatisticsLast) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*StatisticsLast) GetPage() []string {
	return []string{WAP_StatisticsLast}
}
