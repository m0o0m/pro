package schema

import "global"

//轮播图
type SiteFlash struct {
	Id          int64  `xorm:"id"`            //轮播id
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	ImgTitle    string `xorm:"img_title"`     //标题
	ImgUrl      string `xorm:"img_url"`       //图片路径
	ImgLink     string `xorm:"img_link"`      //链接地址
	State       int8   `xorm:"state"`         //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort"`          //排序
	Ftype       int8   `xorm:"ftype"`         //类型 1-PC端 2-WAP端
}

func (*SiteFlash) TableName() string {
	return global.TablePrefix + "site_flash"
}
