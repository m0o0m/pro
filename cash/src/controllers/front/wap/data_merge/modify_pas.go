package data_merge

import "framework/render"

//修改密码
type ModifyPas struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *ModifyPas) GetData(siteId, siteIndexId string) (interface{}, error) {
	//m.CurrPage = 1
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*ModifyPas) GetPage() []string {
	return []string{MODIFY_PAS}
}
