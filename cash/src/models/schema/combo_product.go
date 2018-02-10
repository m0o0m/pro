package schema

import "global"

//套餐和商品的占成
type ComboProduct struct {
	ComboId    int64   `xorm:"combo_id"`    //套餐id
	PlatformId int64   `xorm:"platform_id"` //交易平台id
	ProductId  int64   `xorm:"product_id"`  //商品id
	Proportion float64 `xorm:"proportion"`  //占比(小数点后保留两位)
}

func (*ComboProduct) TableName() string {
	return global.TablePrefix + "combo_product"
}
