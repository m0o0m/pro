package schema

import "global"

//会员对应各分类下余额
type MemberProductClassifyBalance struct {
	MemberId   int64   `insert:"member_id" xorm:"member_id"`     //会员id
	PlatformId int64   `insert:"platform_id" xorm:"platform_id"` //交易平台id
	Balance    float64 `insert:"balance" xorm:"balance"`         //余额
}

func (*MemberProductClassifyBalance) TableName() string {
	return global.TablePrefix + "member_product_classify_balance"
}
