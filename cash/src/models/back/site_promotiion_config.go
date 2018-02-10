package back

//自助优惠申请配置
type SitePromotionConfig struct {
	Id          int64  `xorm:"id" json:"id"`                     //主键id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	ProTitle    string `xorm:"pro_title" json:"proTitle"`        //申请标题
	ProContent  string `xorm:"pro_content" json:"proContent"`    //申请内容
	Createtime  int64  `xorm:"createtime" json:"createtime"`     //申请时间
	Updatetime  int64  `xorm:"updatetime" json:"updatetime"`     //更新时间
	Status      int8   `xorm:"status" json:"status"`             //1有效 2无效
	SiteName    string `xorm:"site_name" json:"siteName"`        //站点名称
}
