package schema

import "global"

//角色与菜单
type RoleMenu struct {
	RoleId int64 `xorm:"role_id"` //角色id
	MenuId int64 `xorm:"menu_id"` //菜单id'
}

func (*RoleMenu) TableName() string {
	return global.TablePrefix + "role_menu"
}
