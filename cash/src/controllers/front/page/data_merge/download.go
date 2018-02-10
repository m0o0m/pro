package data_merge

import (
	"framework/render"
)

type Download struct {
	Key
	CdnUrl string //cdn rui
}

//得到数据
func (m *Download) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//初始化头部尾部数据
	return m, nil
}

//得到页面
func (m *Download) GetPage() []string {
	return []string{DOWNLOAD}
}

//得到<视讯电子彩票体育的页面>
func (m *Download) GetSubPage() map[string]string {
	return nil
}
