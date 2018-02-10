package schema

import (
	"global"
)

//视讯下载链接表
type SiteDown struct {
	Id         int64  `xorm:"'id' PK"`     //主键id
	IosUrl     string `xorm:"ios_url"`     //ios下载地址
	AndroidUrl string `xorm:"android_url"` //安卓下载地址
	State      int64  `xorm:"state"`       //状态 1启用 2停用
	Platform   string `xorm:"platform"`    //平台
	Vers       string `xorm:"vers"`        //版本号
	PcUrl      string `xorm:"pc_url"`      //pc下载地址
}

func (*SiteDown) TableName() string {
	return global.TablePrefix + "site_down"
}
