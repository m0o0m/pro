package schema

import "global"

//代理退水记录对应商品分类
type AgencyRebateWaterRecordProduct struct {
	RecordId  int64   `xorm:"record_id"`  //退水记录id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Money     float64 `xorm:"money"`      //金额
}

func (*AgencyRebateWaterRecordProduct) TableName() string {
	return global.TablePrefix + "agency_rebate_water_record_product"
}
