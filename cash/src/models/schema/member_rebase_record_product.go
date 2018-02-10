package schema

import "global"

//会员返佣记录对应商品金额
type MemberRebaseRecordProduct struct {
	RecordId  int64   `xorm:"record_id"`  //返佣记录id
	ProductId int64   `xorm:"product_id"` //商品分类id
	Money     float64 `xorm:"money"`      //金额
}

func (*MemberRebaseRecordProduct) TableName() string {
	return global.TablePrefix + "member_rebase_record_product"
}
