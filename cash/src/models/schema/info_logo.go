package schema

import "global"

//多站点文案
type InfoLogo struct {
	Id          int64  `xorm:"id"`            //文案id
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	Title       string `xorm:"title"`         //logo名称
	LogoUrl     string `xorm:"logo_url"`      //logo地址
	Type        int8   `xorm:"type"`          //文案类型
	State       int8   `xorm:"state"`         //状态1启用 2停用
	Form        int64  `xorm:"form"`          // 1 pc 2 wap
}

func (*InfoLogo) TableName() string {
	return global.TablePrefix + "info_logo"
}
