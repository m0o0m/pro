package schema

import "global"

type PaidType struct {
	Id           int    `xorm:"id" json:"typeId"`               //主键id
	PaidTypeName string `xorm:"paid_type_name" json:"typeName"` //支付类型名称
	TypeStatus   int    `xorm:"type_status" json:"typeStatus"`  //状态
}

func (*PaidType) TableName() string {
	return global.TablePrefix + "paid_type"
}
