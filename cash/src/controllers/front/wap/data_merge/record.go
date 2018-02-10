package data_merge

import (
	"framework/render"
	"models/back"
)

type Record struct {
	Key
	CdnUrl      string //
	SiteName    string
	ProductList []back.WapProductList
}

func (m *Record) GetData(siteId, siteIndexId string) (interface{}, error) {
	m.CdnUrl = render.CdnUrl
	siteName, err := siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return nil, err
	}
	//读取所有的游戏类型
	gameType, err := productBean.GetProductTypeList()
	if err != nil {
		return nil, err
	}

	data, err := productBean.GetProductList(siteId, siteIndexId)

	if err != nil {
		return nil, err
	}
	var wapProductList []back.WapProductList
	for _, va := range gameType {
		var wapProduct back.WapProductList
		for _, v := range data {
			if v.TypeId == va.Id {
				wapProduct.Id = va.Id
				wapProduct.Title = va.Title
				wapProduct.Children = append(wapProduct.Children, v)
			}
		}
		wapProductList = append(wapProductList, wapProduct)
	}
	m.ProductList = wapProductList
	//fmt.Printf("productlist: %+v\n",wapProductList)
	m.SiteName = siteName

	return m, nil
}

func (m *Record) GetPage() []string {
	return []string{WAP_RECORD}
}
