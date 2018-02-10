package schema

import "global"

//站点商品分类剔除
type SiteProductDel struct {
	SiteId      string `xorm:"'site_id' PK"`       // 站点id
	SiteIndexId string `xorm:"'site_index_id' PK"` // 站点前台id
	ProductId   int64  `xorm:"product_id"`         // 商品id
}

func (*SiteProductDel) TableName() string {
	return global.TablePrefix + "site_product_del"
}
