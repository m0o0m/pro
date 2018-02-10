package schema

import "global"

//代理退佣记录对应商品分类
type AgencyRebateRecordProduct struct {
	RecordId    int64   `xorm:"record_id"`    //退水记录id
	ProductId   int64   `xorm:"product_id"`   //商品分类id
	RebateRatio float64 `xorm:"rebate_ratio"` // 退佣比例
	RebateMoney float64 `xorm:"rebate_money"` // 退佣金额金额
	WaterRatio  float64 `xorm:"water_ratio"`  //退水比例
	WaterMoney  float64 `xorm:"water_money"`  //退水金额
}

func (*AgencyRebateRecordProduct) TableName() string {
	return global.TablePrefix + "agency_rebate_record_product"
}
