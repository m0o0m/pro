package schema

import "global"

//站点退佣设定对应商品分类比例
type SiteRebateProduct struct {
	SetId     int64   `xorm:"set_id"`       //退水记录id
	ProductId int64   `xorm:"product_id"`   //商品分类id
	Rebate    float64 `xorm:"rebate_radio"` //退佣比例
	Rewater   float64 `xorm:"water_radio"`  //退水比例
}

func (*SiteRebateProduct) TableName() string {
	return global.TablePrefix + "site_rebate_product"
}
