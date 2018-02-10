package schema

import "global"

//ip段控制
type BanIp struct {
	Id      int64  `xorm:"id"`
	IpStart string `xorm:"ip_start"` // 起始ip
	IpEnd   string `xorm:"ip_end"`   // 结束ip
	Type    string `xorm:"type"`     // 控制类型：1为代理后台，2为前台，3为wap端
	State   int8   `xorm:"state"`    // 状态：1为启用，2为停用
	Remark  string `xorm:"remark"`   // 备注
}

func (*BanIp) TableName() string {
	return global.TablePrefix + "ban_ip"
}
