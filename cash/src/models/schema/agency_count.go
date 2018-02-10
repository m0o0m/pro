package schema

import "global"

//代理(指所有层级代理)人数统计
type AgencyCount struct {
	AgencyId    int64  `xorm:"'agency_id' PK"` // 代理id(agency表主键)
	SiteId      string `xorm:"site_id"`        //站点id
	SiteIndexId string `xorm:"site_index_id"`  // 所属前台id
	FirstId     int64  `xorm:"first_id"`       //所属股东id
	SecondId    int64  `xorm:"second_id"`      // 所属总代理id
	SecondCount int64  `xorm:"second_count"`   // 总代理数量
	ThirdCount  int64  `xorm:"third_count"`    // 代理数量
	MemberCount int64  `xorm:"member_count"`   // 推广会员数量
}

func (*AgencyCount) TableName() string {
	return global.TablePrefix + "agency_count"
}
