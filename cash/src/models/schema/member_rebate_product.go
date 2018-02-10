package schema

import "global"

//会员返佣对应各商品分类比例
type MemberRebateProduct struct {
	SetId     int64   `xorm:"set_id"`     //返佣设定id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Rate      float64 `xorm:"rate"`       //比例
}

func (*MemberRebateProduct) TableName() string {
	return global.TablePrefix + "member_rebate_product"
}
