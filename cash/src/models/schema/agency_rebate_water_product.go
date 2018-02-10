package schema

import "global"

//代理退水设定对应商品分类比例
type AgencyRebateWaterProduct struct {
	AgencyId  int64   `xorm:"agency_id"`  //代理id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Rate      float64 `xorm:"rate"`       //比例
}

func (*AgencyRebateWaterProduct) TableName() string {
	return global.TablePrefix + "agency_rebate_water_product"
}
