package input

//自助优惠申请列表查询
type SitePromotionConfig struct {
	SiteId      string `query:"siteId"`      //站点
	SiteIndexId string `query:"siteIndexId"` //子站点
}

//自助优惠申请添加
type SitePromotionConfigAdd struct {
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	ProTitle    string `json:"proTitle"`    //申请标题
	ProContent  string `json:"proContent"`  //申请内容
	Status      int8   `json:"status"`      //1有效 2无效
}

//自助优惠申请状态修改
type SitePromotionConfigStatus struct {
	Id          int64  `json:"id"`          //序号id
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	Status      int8   `json:"status"`      //1有效 2无效
}
