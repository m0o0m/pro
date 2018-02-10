package schema

import "global"

//站点详情
type SiteInfo struct {
	SiteId      string `xorm:"'site_id' PK"`
	SiteIndexId string `xorm:"'site_index_id' PK"`
	Remark      string `xorm:"remark"`
	Qq          string `xorm:"qq"`
	Wechat      string `xorm:"wechat"`
	Phone       string `xorm:"phone"`
	Email       string `xorm:"email"`
	UrlLink     string `xorm:"url_link"` //客服连接地址
}

func (*SiteInfo) TableName() string {
	return global.TablePrefix + "site_info"
}
