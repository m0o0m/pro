package schema

import "global"

//角色与权限
type RolePermission struct {
	RoleId       int64 `xorm:"role_id"`       //角色id
	PermissionId int64 `xorm:"permission_id"` //权限id
}

func (*RolePermission) TableName() string {
	return global.TablePrefix + "role_permission"
}
