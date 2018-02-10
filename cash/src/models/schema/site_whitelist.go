package schema

import "global"

//站点ip白名单
type SiteWhitelist struct {
	Id     int64  `xorm:"id"`      //
	SiteId string `xorm:"site_id"` // 站点id
	Ip     string `xorm:"ip"`      //ip白名单:英文逗号分隔  填写多个
	State  int8   `xorm:"state"`   // 状态：1为启用，2为停用
	Remark string `xorm:"remark"`  // 备注
}

func (*SiteWhitelist) TableName() string {
	return global.TablePrefix + "site_whitelist"
}
