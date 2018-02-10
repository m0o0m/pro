package back

type SiteControllListBack struct {
	Id           string  `xorm:"id" json:"id"`                      //主键id
	IndexId      string  `xorm:"index_id" json:"indexId"`           //前台id
	SiteName     string  `xorm:"site_name" json:"siteName"`         //站点名称
	Status       int8    `xorm:"status" json:"status"`              //状态
	CreateTime   int64   `xorm:"create_time" json:"createTime"`     //创建时间
	OnlineTime   int64   `xorm:"online_time" json:"onlineTime"`     //上线时间
	PcDomain     string  `xorm:"pc_domain" json:"pcDomain"`         //pc域名  50
	VideoBalance float64 `xorm:"video_balance" json:"videoBalance"` //视讯余额
	ComboName    string  `xorm:"combo_name" json:"comboName"`       //套餐
	MoreSite     int8    `xorm:"-" json:"moreSite"`                 //多站点
}

//启用站点数，子站数，总站点数
type SiteListNumber struct {
	SiteNum  int64 `json:"siteNum"`  //站点数
	OpenSite int64 `json:"openSite"` //启用站点数
}

//启用站点数，子站数，总站点数
type SiteListNumberR struct {
	ChildNum int64 `json:"childNum"` //子站点数
	OpenSite int64 `json:"openSite"` //启用站点数
}
