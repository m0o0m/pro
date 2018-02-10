package schema

import "global"

//站点公司入款银行对应会员层级中间表
type SiteBankIncomeMemberLevel struct {
	SetId   int64  `xorm:"set_id PK autoincr"` // 入款银行设定id
	LevelId string `xorm:"level_id"`           // 会员层级id
}

func (*SiteBankIncomeMemberLevel) TableName() string {
	return global.TablePrefix + "site_bank_income_member_level"
}
