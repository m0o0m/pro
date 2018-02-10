package data_merge

import "framework/render"

//修改密码
type BankCardAdd struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *BankCardAdd) GetData(siteId, siteIndexId string) (interface{}, error) {
	//m.CurrPage = 1
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*BankCardAdd) GetPage() []string {
	return []string{BANK_CARD_ADD}
}
