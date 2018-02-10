package schema

import "global"

//站点公司入款银行设定表
type SiteRebatePeriods struct {
	Id          int64  `xorm:"'id' PK autoincr"` //id
	SiteId      string `xorm:"site_id"`          // 站点id
	SiteIndexId string `xorm:"site_index_id"`    //站点前台id
	Title       string `xorm:"title"`            // 期数名称
	StartTime   int64  `xorm:"start_time"`       // 开始时间
	EndTime     int64  `xorm:"end_time"`         //结束时间
	Status      int64  `xorm:"status"`           // 退佣状态 0未退佣1已退佣
	DeleteTime  int64  `xorm:"delete_time"`      // 删除时间
}

func (*SiteRebatePeriods) TableName() string {
	return global.TablePrefix + "site_rebate_periods"
}
