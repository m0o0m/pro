package data_merge

import (
	"framework/render"
)

type ReturnWater struct {
	Key
	CdnUrl   string
	CurrPage int
	RData    []ReturnData
	IsSelf   int
	LevelId  string
	UrlLink  string
}

type ReturnData struct {
	ProductName string
	BetValid    float64
	ReturnMoney float64
}

func (m *ReturnWater) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	m.CurrPage = 0
	//获取剔除商品ID
	idDel, _ := siteProductBean.GetProductDel(siteId, siteIndexId)
	//获取站点商品
	productall, _ := siteProductBean.GetProductAll(idDel)
	var data ReturnData
	for _, v := range productall {
		data.ProductName = v.ProductName
		data.BetValid = 0
		data.ReturnMoney = 0
		m.RData = append(m.RData, data)
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
	return m, nil
}

func (*ReturnWater) GetPage() []string {
	return []string{WAP_RETURNWATER, WAP_FOOTER}
}
