package data_merge

import (
	"encoding/json"
	"fmt"
	"models/schema"
)

//线路检测列表
type Detect struct {
	HeaderFooter
	Key
	InfoDomain []schema.InfoDomain //
}

//得到线路检测页面的数据
func (m *Detect) GetData(siteId, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	data, err := siteIwordBean.SiteDetectList(siteId, siteIndexId)
	m.InfoDomain = data
	js, _ := json.MarshalIndent(m, " ", "	")
	fmt.Println(string(js))
	return m, err
}

//得到页面
func (m *Detect) GetPage() []string {
	return []string{DETECT, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *Detect) GetSubPage() map[string]string {
	return nil
}
