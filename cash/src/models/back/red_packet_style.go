package back

//list
type RedPacketStyleList struct {
	Id   int    `xorm:"id" json:"id"`     //id
	Name string `xorm:"name" json:"name"` //样式名称
}

//红包样式下拉框
type RedStyleDrop struct {
	Id       int    `xorm:"id" json:"id"`
	SiteId   string `xorm:"site_id" json:"siteId"`     //站点id
	BgPic    string `xorm:"bg_pic" json:"bgPic"`       //背景图片
	ClickPic string `xorm:"click_pic" json:"clickPic"` //点击图片
	Name     string `xorm:"name" json:"name"`          //样式名称
}

//红包样式图片
type RedStyleDropPicture struct {
	BgPic    string `xorm:"bg_pic" json:"bgPic"`       //背景图片
	ClickPic string `xorm:"click_pic" json:"clickPic"` //点击图片
}

//红包详情
type RedPacketStyle struct {
	Id       int    `xorm:"id" json:"id"`
	SiteId   string `xorm:"site_id" json:"siteId"`     //站点id
	BgPic    string `xorm:"bg_pic" json:"bgPic"`       //背景图片
	ClickPic string `xorm:"click_pic" json:"clickPic"` //点击图片
	TimeCss  string `xorm:"time_css" json:"timeCss"`   //倒计时样式
	ClickCss string `xorm:"click_css" json:"clickCss"` //点击图片样式
	Name     string `xorm:"name" json:"name"`          //样式名称
	Status   int    `xorm:"status" json:"status"`      //状态1启用,2禁用
}
