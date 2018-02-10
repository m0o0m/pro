package schema

import "global"

//站点优惠活动
type SiteActivity struct {
	Id          int64  `xorm:"'id' PK autoincr"` //id
	TopId       int64  `xorm:"top_id"`           //上级栏目
	SiteId      string `xorm:"site_id"`          //站点id
	SiteIndexId string `xorm:"site_index_id"`    //站点前台id
	Title       string `xorm:"title"`            //标题
	Content     string `xorm:"content"`          //内容
	Img         string `xorm:"img"`              //标题图片路径
	State       int8   `xorm:"state"`            //状态 1启用  2关闭
	Sort        int64  `xorm:"sort"`             //排序
	From        int8   `xorm:"from"`             //'1-PC 2-WAP
	Itype       int64  `xorm:"itype"`            //类型代码
	TypeName    string `xorm:"type_name"`        //类型名称
	AddTime     int64  `xorm:"add_time"`         //操作时间
	DeleteTime  int64  `xorm:"delete_time"`      //删除时间   0未删除
}

func (*SiteActivity) TableName() string {
	return global.TablePrefix + "site_activity"
}
