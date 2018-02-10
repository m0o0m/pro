package data_merge

import (
	"framework/render"
	"models/input"
)

type Applypro struct {
	Key
	CdnUrl      string //cdn
	SiteName    string //站点名称
	ProTitles   []*Pro //活动标题
	SiteId      string
	SiteIndexId string
}

type Pro struct {
	ProId      int64  //id
	ProTitle   string //标题
	ProContent string //内容
}

//得到数据
func (m *Applypro) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	//查询站点
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}

	//查询活动列表
	list, err := sitePromotionConfig.GetSitePromotionConfig(&input.SitePromotionConfig{siteId, siteIndexId})
	if err != nil {
		return nil, err
	}
	for k, _ := range list {
		m.ProTitles = append(m.ProTitles, &Pro{list[k].Id, list[k].ProTitle, list[k].ProContent})
	}
	m.SiteName = siteName
	m.SiteId = siteId
	m.SiteIndexId = siteIndexId
	return m, err
}

//得到页面
func (m *Applypro) GetPage() []string {
	return []string{APPLYPRO}
}

//得到<视讯电子彩票体育的页面>
func (m *Applypro) GetSubPage() map[string]string {
	return nil
}
