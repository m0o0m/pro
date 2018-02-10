package back

//H5动画设置查询
type SiteH5Set struct {
	Id            string `json:"id"`            //站点id
	IndexId       string `json:"indexId"`       //站点前台id
	H5StateSwitch int8   `json:"h5StateSwitch"` //动画状态1.开启2.关闭
}

//查看弹窗广告配置
type NoticePopupConfig struct {
	PopoverBgColor    string `json:"popoverBgColor"`    //站点弹窗广告背景颜色
	PopoverTitleColor string `json:"popoverTitleColor"` //站点弹窗广告标题颜色
	PopoverBarColor   string `json:"popoverBarColor"`   //站点弹窗广告标题栏颜色
}
