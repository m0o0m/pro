package schema

import "global"

//多站点文案
type Iword struct {
	Id          int64  `xorm:"id"`            //文案id
	TopId       int64  `xorm:"top_id"`        //上级栏目
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	Title       string `xorm:"title"`         //标题
	TitleColor  string `xorm:"title_color"`   //标题颜色
	Content     string `xorm:"content"`       //内容
	Url         string `xorm:"url"`           //链接地址
	Img         string `xorm:"img"`           //图片路径
	State       int8   `xorm:"state"`         //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort"`          //排序
	From        int8   `xorm:"from"`          //'1-PC 2-WAP
	Itype       int64  `xorm:"itype"`         //类型代码
	TypeName    string `xorm:"type_name"`     //类型名称
	AddTime     int64  `xorm:"add_time"`      //操作时间
}

func (*Iword) TableName() string {
	return global.TablePrefix + "site_iword"
}
