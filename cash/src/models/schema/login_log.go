package schema

import "global"

//登录日志
type LoginLog struct {
	Id          int64  `xorm:"'id' PK autoincr"`
	Domains     string `xorm:"domains"`       //登录域名
	Device      int8   `xorm:"device"`        //录设备（1.pc2.wap3.ios4.android）
	Account     string `xorm:"account"`       //登录账号
	LoginResult int8   `xorm:"login_result"`  //登录是否成功(1.成功2.失败)
	LoginTime   int64  `xorm:"login_time"`    //登入时间(年月日时分秒)
	LoginIp     string `xorm:"login_ip"`      //登录ip
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	LoginRole   int8   `xorm:"login_role"`    //登录角色1.会员2.管理员
}

func (*LoginLog) TableName() string {
	return global.TablePrefix + "login_log"
}
