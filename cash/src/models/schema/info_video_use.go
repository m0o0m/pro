package schema

import "global"

//登录日志
type InfoVideoUse struct {
	Id          int64  `xorm:"'id' PK autoincr"`
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台
	Style       int64  `xorm:"style"`         //0默认风格
	Type        int64  `xorm:"type"`          //模板
	DoTime      int64  `xorm:"do_time"`       //操作时间
	State       int8   `xorm:"state"`         //状态1关闭2启用
	Remark      string `xorm:"remark"`        //备注
	Video       string `xorm:"video"`         //正在使用的视讯
}

func (*InfoVideoUse) TableName() string {
	return global.TablePrefix + "info_video_use"
}
