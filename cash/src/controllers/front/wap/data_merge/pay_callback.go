package data_merge

import (
	"framework/render"
)

type PayCallback struct {
	Key
	CdnUrl  string
	NewHtml string
}

func (m *PayCallback) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*PayCallback) GetPage() []string {
	return []string{WAP_PAY_CALLBACK}
}
