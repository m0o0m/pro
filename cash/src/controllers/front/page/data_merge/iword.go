package data_merge

import (
	"fmt"
	"framework/render"
	"models/back"
	"strings"
)

type IWord struct {
	HeaderFooter
	Key
	List []back.IwordList
	Id   int64
}

//得到文案页面的数据
func (m *IWord) GetData(siteId string, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	data, err := siteIwordBean.GetIWordList(siteId, siteIndexId, m.Id)
	m.List = data
	fmt.Printf("%+v\n", m)
	return m, err
}

//得到页面
func (m *IWord) GetPage() []string {
	return []string{IWORD, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *IWord) GetSubPage() map[string]string {
	return nil
}

func (m *IWord) GetPageCacheKey(siteId, siteIndexId string, pcPageData interface{}) string {
	pageData, _ := pcPageData.(render.PcPageData)
	var keys = []string{siteId, siteIndexId}
	pages := pageData.GetPage()
	keys = append(keys, pages...)
	subPageMap := pageData.GetSubPage()
	if len(subPageMap) != 0 {
		dz, ok := subPageMap[render.DzViewPath]
		if ok {
			keys = append(keys, dz)
		}
		fc, ok := subPageMap[render.FcViewPath]
		if ok {
			keys = append(keys, fc)
		}
		sp, ok := subPageMap[render.SpViewPath]
		if ok {
			keys = append(keys, sp)
		}
		vd, ok := subPageMap[render.VdViewPath]
		if ok {
			keys = append(keys, vd)
		}
	}
	return strings.Join(keys, "$") + "$" + fmt.Sprintf("%d", m.Id)
}
