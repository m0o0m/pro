package data_merge

import (
	"fmt"
	"framework/render"
	"models/back"
	"strings"
)

type MemberIndex struct {
	HeaderFooter
	Key
	Notice     string            //公告内容
	MemberPage int               //页面属性
	IsSelf     int               //自助反水开关
	IncomeWord map[string]string //入款文案
	IsOther    int64             //其他页面跳转到现金记录
}

//得到首页数据
func (m *MemberIndex) GetData(siteId, siteIndexId string) (interface{}, error) {
	//初始化头部尾部数据
	err := m.initHeaderFooter(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	if m.MemberPage == 2 || m.MemberPage == 9 { //公司入款页面和第三方入款页面
		contents := make([]back.IndexContentWord, 0)
		m.IncomeWord = make(map[string]string)
		contents, err = siteIwordBean.IWordContentList(siteId, siteIndexId)
		for _, v := range contents { //类型7以上的都为入款文案
			m.IncomeWord[v.TypeName] = v.Content
		}
	}
	m.Notice = m.Header.Notice //头部有就直接赋上去
	return m, err
}

//得到页面
func (m *MemberIndex) GetPage() []string {
	return []string{MEMBERINDEX, MEMBERACCOUNT, MEMBERCOMPANY, WITHDRAWALINDEX, MEMBERCOVERT, MEMBERRECORD, MEMBERREPORT, MEMBERSPREAD, MEMBERNEWS, MEMBERTHIRD, MEMBERCOMPLETE, MEMBERDRAWWRITE, MEMBERHEADER, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *MemberIndex) GetSubPage() map[string]string {
	return nil
}

//pc得到页面名称组合,作为缓存的key
func (m *MemberIndex) GetPageCacheKey(siteId, siteIndexId string, pageData interface{}) string {
	pcPageData, _ := pageData.(render.PcPageData)
	var keys = []string{siteId, siteIndexId}
	pages := pcPageData.GetPage()
	keys = append(keys, pages...)
	subPageMap := pcPageData.GetSubPage()
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
	return strings.Join(keys, "$") + "$" + fmt.Sprintf("%d", m.MemberPage)
}
