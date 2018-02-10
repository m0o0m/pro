package schema

import "global"

//子帐号权限细项
type DetailsMember struct {
	ChildId    int64  `xorm:"child_id"`    //子帐号id
	ChildPower string `xorm:"child_power"` // 子帐号的资料细项权限，用逗号分隔
	ChildSite  string `xorm:"child_site"`  // 子帐号能操作的子站点，用逗号分隔
}

func (m *DetailsMember) TableName() string {
	return global.TablePrefix + "details_member"
}
