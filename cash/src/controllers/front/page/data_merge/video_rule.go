package data_merge

import (
	"framework/render"
)

type VideoRule struct {
	CdnUrl string //cdn地址
	Key
}

//得到首页数据
func (m *VideoRule) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//初始化头部尾部数据
	return m, nil
}

//得到页面
func (m *VideoRule) GetPage() []string {
	return []string{VIDEO_RULE}
}

//得到<视讯电子彩票体育的页面>
func (m *VideoRule) GetSubPage() map[string]string {
	return nil
}
