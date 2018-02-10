package data_merge

import "framework/render"

//新版维护页
type Maintain struct {
	Key
	SiteName string
	CdnUrl   string
}

func (m *Maintain) GetData(siteId, siteIndexId string) (data interface{}, err error) {
	m.CdnUrl = render.CdnUrl
	//查询站点
	m.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return m, err
	}
	return m, err
}

func (*Maintain) GetPage() []string {
	return []string{MAINTAIN}
}

func (*Maintain) GetSubPage() map[string]string {
	return nil
}
