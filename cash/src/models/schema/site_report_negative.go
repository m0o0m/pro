package schema

import "global"

//报表上月负数
type SiteReportNegative struct {
	Id          int64  `xorm:"id PK autoincr"`
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	Years       string `xorm:"years"`         //年月
	State       int8   `xorm:"state"`         //状态 1累计 2清零
	UpdateDate  int64  `xorm:"update_date"`   //更新时间
}

func (*SiteReportNegative) TableName() string {
	return global.TablePrefix + "site_report_negative"
}
