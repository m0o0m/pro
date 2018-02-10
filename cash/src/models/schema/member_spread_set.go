package schema

import "global"

//会员推广设定
type MemberSpreadSet struct {
	SiteId       string  `xorm:"site_id"`        //操作站点id
	SiteIndexId  string  `xorm:"site_index_id"`  //站点前台id
	IsOpen       int8    `xorm:"is_open"`        //是否开启会员推广
	IsIp         int8    `xorm:"is_ip"`          //是否限制ip
	IsMateAgency int8    `xorm:"is_mate_agency"` //是否匹配推广会员代理
	IsCode       int8    `xorm:"is_code"`        //返佣会员是否需要打码
	RankingNum   float64 `xorm:"ranking_num"`    //排行榜人数系数
	RankingMoney float64 `xorm:"ranking_money"`  //排行榜金额系数
}

func (*MemberSpreadSet) TableName() string {
	return global.TablePrefix + "member_spread_set"
}
