package data_merge

import (
	"framework/render"
	"global"
	"models/back"
)

//维护页
type Maintenance struct {
	Key
	SiteName string
	CdnUrl   string
	Content  string     //维护原因
	WebHome  int        //1整站维护2单商品维护
	Products []*Product //多个商品维护情况
}

//商品
type Product struct {
	Name          string //商品名
	IsMaintenance int    //1正常2,维护
	Content       string
	Router        string
	VType         string //
}

func (m *Maintenance) GetData(siteId, siteIndexId string) (data interface{}, err error) {
	m.CdnUrl = render.CdnUrl
	//查询站点
	m.SiteName, err = siteOperateBean.GetSiteNameBySiteIndexId(siteId, siteIndexId)
	if err != nil {
		return
	}
	//如果是单商品维护
	if m.WebHome != 1 {
		data, err := productBean.GetProductList(siteId, siteIndexId)
		if err != nil {
			return m, err
		}
		for _, v := range data {
			product := new(Product)
			product.Name = v.ProductName
			product.IsMaintenance = 1
			product.VType = v.VType
			//product_type数据库中差一列,用来区分路由,这里就只能写死了
			switch v.TypeId {
			case 1:
				product.Router = "livetop"
			case 2:
				product.Router = "egame"
			case 4:
				product.Router = "lottery"
			case 5:
				product.Router = "sports"
			}
			siteModuleInf, ok := global.SiteModuleCache.Load(global.GenKey("all", "pc", v.VType))
		OK:
			if ok {
				product.IsMaintenance = 2
				siteModule, _ := siteModuleInf.(*back.SiteModule)
				product.Content = siteModule.Content
				m.Products = append(m.Products, product)
				continue
			}
			siteModuleInf, ok = global.SiteModuleCache.Load(global.GenKey(siteId, "pc", v.VType))
			if ok {
				goto OK
			}
			m.Products = append(m.Products, product)
		}
	}
	return m, nil
}

func (*Maintenance) GetPage() []string {
	return []string{MAINTENANCE}
}

func (*Maintenance) GetSubPage() map[string]string {
	return nil
}
