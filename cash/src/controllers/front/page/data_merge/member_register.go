package data_merge

import (
	"models/schema"
)

type Register struct {
	Register schema.SiteMemberRegisterSet //
	HeaderFooter
	Key
	Agreement string
	AgencyId  int64
	IsReg     int8
}

//得到会员注册页面的数据
func (m *Register) GetData(siteId string, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	m.Header.PageType = 7 //会员注册页面
	data, has, err := memberRegister.GetOneSet(siteId, siteIndexId)
	if err != nil {
		m.IsReg = 2
	}
	if !has {
		m.IsReg = 2
	} else {
		m.IsReg = data.IsReg
	}

	agencyInfo, has, err := AgencyBean.GetAgencyId(siteId, siteIndexId)

	if err != nil {
		return nil, err
	}
	if has {
		m.AgencyId = agencyInfo.Id
	}
	m.Register = data
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
	return []string{ZHUCE, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *Register) GetSubPage() map[string]string {
	return nil
}
