package schema

import "global"

//红包样式表
type RedPacketStyle struct {
	Id       int    `xorm:"'id' PK autoincr"`
	SiteId   string `xorm:"site_id"`   //站点id
	BgPic    string `xorm:"bg_pic"`    //背景图片
	ClickPic string `xorm:"click_pic"` //点击图片
	TimeCss  string `xorm:"time_css"`  //倒计时样式
	ClickCss string `xorm:"click_css"` //点击图片样式
	Name     string `xorm:"name"`      //样式名称
	Status   int    `xorm:"status"`    //状态1启用,2禁用
}

func (*RedPacketStyle) TableName() string {
	return global.TablePrefix + "red_packet_style"
}
