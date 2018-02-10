package schema

import "global"

//代理出款银行卡
type AgencyThirdBank struct {
	Id          int64  `xorm:"'id' PK autoincr"` //主键id
	AgencyId    int64  `xorm:"agency_id"`        //代理id
	BankId      int64  `xorm:"bank_id"`          //卡类型
	Card        string `xorm:"card"`             //19卡号
	CardName    string `xorm:"card_name"`        //20银行卡账号
	CardAddress string `xorm:"card_address"`     //50卡开户行
}

func (*AgencyThirdBank) TableName() string {
	return global.TablePrefix + "agency_third_bank"
}
