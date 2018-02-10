package input

//电子样式设置
type PromotionSet struct {
	SiteId       string `json:"siteId"`       //
	SiteIndexId  string `json:"siteIndexId"`  //
	TitleBcolor  string `json:"titleBcolor"`  //导航背景颜色
	TitleColor   string `json:"titleColor"`   //导航选中颜色
	ButtonBcolor string `json:"buttonBcolor"` //按钮背景颜色
	ButtonColor  string `json:"buttonColor"`  //按钮选中颜色
	BborderColor string `json:"bborderColor"` //按钮边框颜色
	PopBcolor    string `json:"popBcolor"`    //背景颜色
}
type PromotionColorSet struct {
	SiteId      string `json:"site_id"`       //
	SiteIndexId string `json:"site_index_id"` //
	Bcolor      string `json:"bcolor"`        //电子内页主题色
}

//电子配置初始化
type DianZiInitialization struct {
	SiteId      string `json:"siteId"`
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点Id
}
