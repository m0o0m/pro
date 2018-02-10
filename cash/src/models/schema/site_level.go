package schema

import "global"

//站点层级
type SiteLevel struct {
	Id        int64  `xorm:"id PK autoincr"` //主键id
	Lid       int64  `xorm:"lid"`            //层级编号
	LevelName string `xorm:"level_name"`     //层级名字
	SiteLevel string `xorm:"site_level"`     //包含站点
	Talk      string `xorm:"talk"`           //描述
	Remark    string `xorm:"remark"`         //备注
	DoTime    int64  `xorm:"do_time"`        //操作时间
	State     int8   `xorm:"state"`          //显示状态 1显示 2关闭
}

func (*SiteLevel) TableName() string {
	return global.TablePrefix + "site_level"
}
