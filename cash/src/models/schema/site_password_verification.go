package schema

import "global"

//口令验证信息表
type SitePasswordVerification struct {
	SiteId     string `xorm:"'site_id' PK" json:"siteId"`    //站点
	Status     int    `xorm:"status" json:"status"`          //状态
	PassKey    string `xorm:"pass_key" json:"passKey"`       //密钥
	Account    string `xorm:"account" json:"account"`        //操作人
	UpdateTime int64  `xorm:"update_time" json:"updateTime"` //更新时间 0表示未启用
}

func (*SitePasswordVerification) TableName() string {
	return global.TablePrefix + "site_password_verification"
}
