package schema

import "global"

//总后台登录日志
type AdminLoginLog struct {
	Id          int64  `xorm:"id"`           // 主键id
	Device      int8   `xorm:"device"`       // 登录设备
	Account     string `xorm:"account"`      // 登录账号
	LoginResult int8   `xorm:"login_result"` // 登录是否成功(1.成功2.失败)
	LoginTime   int64  `xorm:"login_time"`   // 登入时间(年月日时分秒)
	LoginIp     string `xorm:"login_ip"`     // 登录ip
	LoginRole   int64  `xorm:"login_role"`   // 登录角色
}

func (*AdminLoginLog) TableName() string {
	return global.TablePrefix + "admin_login_log"
}
