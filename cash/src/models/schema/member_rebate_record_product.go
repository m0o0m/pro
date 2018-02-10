package schema

import "global"

//会员返佣记录对应商品金额表
type MemberRebateRecordProduct struct {
	RecordId   int64   `xorm:"record_id"`   //返佣记录id
	ProductId  int64   `xorm:"product_id"`  //商品分类Id
	ProductBet float64 `xorm:"product_bet"` //商品有效打码
	Money      float64 `xorm:"money"`       //金额(返佣金额)
}

func (m *MemberRebateRecordProduct) TableName() string {
	return global.TablePrefix + "member_rebate_record_product"
}
