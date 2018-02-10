package schema

import "global"

//角色
type Role struct {
	Id         int64  `xorm:"'id' PK autoincr"`      //主键id
	RoleName   string `xorm:"role_name"`             //中文名
	RoleMark   string `xorm:"role_mark"`             //英文名
	IsOperate  int8   `xorm:"is_operate"`            //是否可以删除和禁用    0可以，1不可以
	Remark     string `xorm:"remark"`                //备注
	Status     int8   `xorm:"status"`                //状态
	CreateTime int64  `xorm:"'create_time' created"` //添加时间
	DeleteTime int64  `xorm:"delete_time"`           //删除时间
}

func (*Role) TableName() string {
	return global.TablePrefix + "role"
}
