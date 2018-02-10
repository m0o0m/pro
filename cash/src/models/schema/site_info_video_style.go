package schema

import "global"

//视讯模板视讯风格分类
type SiteInfoVideoStyle struct {
	Id     int64  `xorm:"id"`     //视讯模板id
	Name   string `xorm:"name"`   //类型名字
	Pid    int8   `xorm:"pid"`    //父级
	Aid    int8   `xorm:"aid"`    //中间关联字段
	Style  int8   `xorm:"style"`  //视讯样式id
	Remark string `xorm:"remark"` //备注预留字段
	Status int8   `xorm:"status"` //视讯模板开关：1、开启  2、关闭
}

func (*SiteInfoVideoStyle) TableName() string {
	return global.TablePrefix + "site_info_video_style"
}
