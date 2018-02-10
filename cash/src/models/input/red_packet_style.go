package input

//添加(修改)红包皮肤
type RedPacketStyleAdd struct {
	Id       int    `json:"id"`
	SiteId   string `json:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`            //操作站点id
	BgPic    string `json:"bgPic" valid:"Required;MinSize(8);ErrorCode(71001)"`    //背景图片
	ClickPic string `json:"clickPic" valid:"Required;MinSize(8);ErrorCode(71002)"` //点击图片
	TimeCss  string `json:"timeCss" `                                              //倒计时样式
	ClickCss string `json:"clickCss"`                                              //点击图片样式
	Name     string `json:"name"  valid:"Required;ErrorCode(71005)"`               //样式名称
	Status   int    `json:"status"`                                                //状态 1启用,2禁用
}

//查询列表
type RedPacketStyleList struct {
	SiteId string `query:"siteId"` //操作站点id,可以不填
}

//查询详情
type RedPacketStyle struct {
	Id int64 `query:"id" json:"id"  valid:"Required;ErrorCode(71007)"` //皮肤id,必填
}
