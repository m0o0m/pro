package schema

import "global"

//站点公司入款银行对应会员层级中间表
type SiteOnlineIncomeMemberLevel struct {
	SetId   int64  `xorm:"set_id"`   // 入款银行设定id
	LevelId string `xorm:"level_id"` // 会员层级id
}

func (*SiteOnlineIncomeMemberLevel) TableName() string {
	return global.TablePrefix + "site_online_income_member_level"
}
