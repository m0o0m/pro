package schema

import "global"

//权限
type Permission struct {
	Id             int64  `xorm:"'id' PK autoincr"`      //主键id
	Module         string `xorm:"module"`                //所属模块
	PermissionName string `xorm:"permission_name"`       //权限名称
	Route          string `xorm:"route"`                 //权限路由
	Type           string `xorm:"type"`                  //所属平台
	Method         string `xorm:"method"`                //路由请求方式
	Status         int8   `xorm:"status"`                //权限状态
	CreateTime     int64  `xorm:"'create_time' created"` //添加时间
	DeleteTime     int64  `xorm:"delete_time"`           //删除时间
}

func (*Permission) TableName() string {
	return global.TablePrefix + "permission"
}
