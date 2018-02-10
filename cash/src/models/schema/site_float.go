package schema

import "global"

//轮播图
type SiteFloat struct {
	Id          int64  `xorm:"id"`            //轮播id
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	ImgA        string `xorm:"img_a"`         //常规显示图片
	ImgB        string `xorm:"img_b"`         //鼠标覆盖事件显示图片
	Url         string `xorm:"url"`           //链接
	UrlInter    int64  `xorm:"url_inter"`     //内链
	IsBlank     int64  `xorm:"is_blank"`      //新开窗口 1-是 2-否
	IsSlide     int64  `xorm:"is_slide"`      //滑动效果 1-是 2-否
	IsClose     int64  `xorm:"is_close"`      //关闭按钮
	State       int64  `xorm:"state"`         //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort"`          //排序
	Ftype       int64  `xorm:"ftype"`         //类型 1-左 2-右
	DeleteTime  int64  `xorm:"delete_time"`   //删除时间
}

func (*SiteFloat) TableName() string {
	return global.TablePrefix + "site_float"
}
