package back

//登录日志返回
type AdminLoginLogBack struct {
	Id          int64  `xorm:"id" json:"id"`                    // 主键id
	Device      int8   `xorm:"device" json:"device"`            // 登录设备
	Account     string `xorm:"account" json:"account"`          // 登录账号
	LoginResult int8   `xorm:"login_result" json:"loginResult"` // 登录是否成功(1.成功2.失败)
	LoginTime   int64  `xorm:"login_time" json:"loginTime"`     // 登入时间(年月日时分秒)
	LoginIp     string `xorm:"login_ip" json:"loginIp"`         // 登录ip
	RoleName    string `xorm:"role_name" json:"roleName"`       // 登录角色
	RoleMark    string `xorm:"role_mark" json:"roleMark"`       //角色标识(英文)
}
