package data_merge

import "framework/render"

//修改密码
type BankCard struct {
	Key
	CdnUrl   string
	CurrPage int64 //当前页面
}

func (m *BankCard) GetData(siteId, siteIndexId string) (interface{}, error) {
	//m.CurrPage = 1
	m.CdnUrl = render.CdnUrl
	return m, nil
}

func (*BankCard) GetPage() []string {
	return []string{BANK_CARD}
}
