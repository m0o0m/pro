package schema

import "global"

//代理(指所有层级代理)商品佣金比例
type AgencyProductCommission struct {
	AgencyId   int64   `xorm:"'agency_id' PK"`  // 代理id(agency表主键)
	ProductId  int64   `xorm:"'product_id' PK"` // 商品分类id
	Commission float64 `xorm:"commission"`      // 比例,最大100%
}

func (*AgencyProductCommission) TableName() string {
	return global.TablePrefix + "agency_product_commission"
}
