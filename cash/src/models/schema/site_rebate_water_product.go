package schema

import "global"

//站点退水设定对应商品分类比例
type SiteRebateWaterProduct struct {
	SetId     int64   `xorm:"set_id"`     //退水记录id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Rate      float64 `xorm:"rate"`       //比例
}

func (*SiteRebateWaterProduct) TableName() string {
	return global.TablePrefix + "site_rebate_water_product"
}
