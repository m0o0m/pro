package data_merge

import "framework/render"

//修改密码
type ModifyInfo struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *ModifyInfo) GetData(siteId, siteIndexId string) (interface{}, error) {
	//m.CurrPage = 1
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*ModifyInfo) GetPage() []string {
	return []string{MODIFY_INFO}
}
