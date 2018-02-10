package data_merge

import (
	"framework/render"
	"models/input"
)

type LoginInfo struct {
	HeaderFooter
	Key
	WebAdv   string //弹窗广告
	SiteName string //站点名称
	CdnUrl   string //cdn rui
}

//得到login_info页面的数据
func (m *LoginInfo) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//查询广告
	advList, err := webAdv.GetWebAdvList(&input.AdvList{siteId, siteIndexId, 1})
	if err != nil {
		return nil, err
	}
	//查询站点
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	if len(advList) > 0 {
		for k, _ := range advList {
			m.WebAdv += "|" + advList[k].Content
		}
		m.WebAdv = m.WebAdv[1:]
	}
	m.SiteName = siteName
	return m, err
}

//得到页面
func (m *LoginInfo) GetPage() []string {
	return []string{LOGIN_INFO}
}

//得到<视讯电子彩票体育的页面>
func (m *LoginInfo) GetSubPage() map[string]string {
	return nil
}
