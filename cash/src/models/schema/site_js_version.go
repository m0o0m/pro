package schema

import (
	"global"
)

//js版本控制
type SiteJsVersion struct {
	Id    int64  `xorm:"'id' PK"` //主键id
	State int64  `xorm:"state"`   //状态 1启用 2停用
	Type  string `xorm:"type"`    //类别 1pc  2wap
	Vers  string `xorm:"vers"`    //版本号
}

func (*SiteJsVersion) TableName() string {
	return global.TablePrefix + "site_js_version"
}
