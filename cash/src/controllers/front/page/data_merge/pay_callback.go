package data_merge

import "framework/render"

type PayCallback struct {
	Key
	NewHtml string //第三方返回的支付页面源码
	CdnUrl  string //cdn rui
}

//得到数据
func (m *PayCallback) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	return m, nil
}

//得到页面
func (m *PayCallback) GetPage() []string {
	return []string{PAY_CALLBACK}
}

//得到<视讯电子彩票体育的页面>
func (m *PayCallback) GetSubPage() map[string]string {
	return nil
}
