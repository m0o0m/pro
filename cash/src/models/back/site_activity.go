package back

//会员活动列表（Wap)
type WapActivityList struct {
	Id    int64  `xorm:"id" json:"id"`       //id
	Title string `xorm:"title" json:"title"` //标题
	Img   string `xorm:"img" json:"img"`     //标题图片路径
}

//单个会员活动详情
type WapActivityInfo struct {
	Id      int64  `xorm:"id" json:"id"`           //id
	Title   string `xorm:"title" json:"title"`     //标题
	Img     string `xorm:"img" json:"img"`         //标题图片路径
	Content string `xorm:"content" json:"content"` //内容
}

//优惠申请大厅活动数据
type Pro struct {
	ProId      int64  `xorm:"proId" json:"proId"`           //id
	ProTitle   string `xorm:"proTitle" json:"proTitle"`     //标题
	ProContent string `xorm:"proContent" json:"proContent"` //内容
}
