package schema

import "global"

//会员出款银行卡表
type MemberBank struct {
	Id            int64  `xorm:"'id' PK autoincr"`
	MemberId      int64  `xorm:"member_id"`           //会员Id
	BankId        int64  `xorm:"bank_id"`             //卡类型
	Card          string `xorm:"card"`                //卡号
	CardName      string `xorm:"card_name"`           //卡账号
	CardAddress   string `xorm:"card_address"`        //卡地址
	IsDefaultBank int8   `xorm:"is_default_bank"`     //是否该会员默认银行卡(1.是2.不是)
	CreateTime    int64  `xorm:"create_time created"` //创建时间
	DeleteTime    int64  `xorm:"delete_time"`         //删除时间
}

func (*MemberBank) TableName() string {
	return global.TablePrefix + "member_bank"
}
