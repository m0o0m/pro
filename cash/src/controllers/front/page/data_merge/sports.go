package data_merge

import (
	"fmt"
	"framework/render"
	"strings"
)

type Sports struct {
	HeaderFooter
	Key
	Notice      string         //公告内容
	LogoUrl     string         //logo地址
	SiteName    string         //站点名称
	SportsOrder map[string]int //体育平台排序
	SportLen    int
}

//得体育页面的数据
func (m *Sports) GetData(siteId, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	m.Header.PageType = 4 //体育页面
	//查询体育公告
	notices, _ := noticeBean.GetNoticeBySiteId(siteId)
	//logo
	logUrl, _ := siteOperateBean.Logo(siteId, siteIndexId, 1)
	//查询站点
	siteName, _ := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)

	if len(notices) > 0 {
		for k, _ := range notices {
			m.Notice += ";" + notices[k].NoticeContent
		}
		m.Notice = m.Notice[1:]
	}

	m.LogoUrl = logUrl
	m.SiteName = siteName

	//获取体育平台和排名
	sport_order, err := siteOrderModule.SportModule(siteId, siteIndexId)
	m.SportsOrder = make(map[string]int)
	Order := strings.Split(sport_order, ",")
	if len(Order) > 1 {
		for _, v := range Order {
			str := strings.Split(v, "_")
			m.SportsOrder[str[0]] = 1
		}
	} else {
		m.SportsOrder = nil
	}
	m.SportLen = len(m.SportsOrder)
	fmt.Println(m)
	return m, err
}

func (*Sports) GetPage() []string {
	return []string{SPORTS, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

func (*Sports) GetSubPage() map[string]string {
	return map[string]string{
		render.SpViewPath: SP,
	}
}
