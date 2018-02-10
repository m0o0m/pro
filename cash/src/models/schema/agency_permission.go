package schema

import "global"

//子账号与权限
type AgencyPermission struct {
	AgencyId     int64 `xorm:"agency_id"`     //子账号id
	PermissionId int64 `xorm:"permission_id"` //权限id'
}

func (*AgencyPermission) TableName() string {
	return global.TablePrefix + "agency_permission"
}
