package data_merge

import (
	"framework/render"
)

type Isoview struct {
	Key
	CdnUrl string //cdn地址
}

//得到数据
func (m *Isoview) GetData(siteId, siteIndexId string) (interface{}, error) {
	//初始化头部尾部数据
	m.CdnUrl = render.CdnUrl
	return m, nil
}

//得到页面
func (m *Isoview) GetPage() []string {
	return []string{ISOVIEW}
}

//得到<视讯电子彩票体育的页面>
func (m *Isoview) GetSubPage() map[string]string {
	return nil
}
