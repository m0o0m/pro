package schema

import "global"

//代理退佣设定对应商品分类比例
type AgencyRebateProduct struct {
	AgencyId  int64   `xorm:"agency_id"`  //代理id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Rate      float64 `xorm:"rate"`       //比例
}

func (*AgencyRebateProduct) TableName() string {
	return global.TablePrefix + "agency_rebate_product"
}
