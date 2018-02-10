package schema

import "global"

//平台账号
type Admin struct {
	Id           int64  `xorm:"'id' PK autoincr"`      // 主键id
	Account      string `xorm:"account"`               // 登录账号
	Password     string `xorm:"password"`              // 登录密码
	Remark       string `xorm:"remark"`                // 备注
	Status       int8   `xorm:"status"`                // 账号状态
	CreateTime   int64  `xorm:"'create_time' created"` // 创建时间
	DeleteTime   int64  `xorm:"delete_time"`           // 软删除时间(为0表示未删除)
	LoginKey     string `xorm:"login_key"`             //登录凭证
	RoleId       int64  `xorm:"role_id"`               //角色id
	LoginIp      string `xorm:"login_ip"`              //登录ip限制（没有则没限制）
	OnlineStatus int8   `xorm:"online_status"`         //在线状态 1:在线 2:离线
}

func (*Admin) TableName() string {
	return global.TablePrefix + "admin"
}
