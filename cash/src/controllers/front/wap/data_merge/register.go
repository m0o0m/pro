package data_merge

import (
	"framework/render"
	"models/schema"
)

type Register struct {
	Key
	Register  schema.SiteMemberRegisterSet //
	CdnUrl    string
	UrlLink   string
	Agreement string
	IsReg     int8
}

//得到会员注册页面的数据
func (m *Register) GetData(siteId string, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	data, has, err := memberRegister.GetOneSet(siteId, siteIndexId)
	if err != nil {
		m.IsReg = 2
	}
	if !has {
		m.IsReg = 2
	} else {
		m.IsReg = data.IsReg
		m.Register = data
	}
	//站点客服链接
	siteinfo, has, err := SiteInfoBean.GetSingleSiteInfo(siteId, siteIndexId)
	if err != nil {
		m.UrlLink = ""
	}
	if has == false {
		m.UrlLink = ""
	} else {
		m.UrlLink = siteinfo.UrlLink
	}
	//开户协议
	Iword, err := siteIwordBean.GetAgreeMent(siteId, siteIndexId)

	if err != nil {
		m.Agreement = ""
	} else {
		m.Agreement = Iword.Content
	}
	return m, err
}

//得到页面
func (m *Register) GetPage() []string {
	return []string{WAP_REG}
}
