package schema

import "global"

//自助返水查询
type MemberRetreatWaterRecordProductSelf struct {
	RecordId   int64   `xorm:"record_id"`   //退水记录id
	ProductId  int64   `xorm:"product_id"`  //商品分类id
	ProductBet float64 `xorm:"product_bet"` //对应商品有效投注额
	Rate       float64 `xorm:"rate"`        //比例
	Money      float64 `xorm:"money"`       //本次退水金额
}

func (*MemberRetreatWaterRecordProductSelf) TableName() string {
	return global.TablePrefix + "member_retreat_water_record_product_self"
}
