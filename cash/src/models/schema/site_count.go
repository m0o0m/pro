package schema

import "global"

//SiteCount 站点运营统计数据表
type SiteCount struct {
	SiteId            string  `xorm:"'site_id' PK"`        //站点id
	SiteIndexId       string  `xorm:"'site_index_id' PK"`  //站点前台id
	FirstAgencyCount  int64   `xorm:"first_agency_count"`  //股东数量
	SecondAgencyCount int64   `xorm:"second_agency_count"` //总代理数量
	ThirdAgencyCount  int64   `xorm:"third_agency_count"`  //代理数量
	MemberCount       int64   `xorm:"member_count"`        //会员数量
	OrderCount        int64   `xorm:"order_count"`         //有效订单数量
	TotalTurnover     float64 `xorm:"total_turnover"`      //总营业额
}

func (*SiteCount) TableName() string {
	return global.TablePrefix + "site_count"
}
