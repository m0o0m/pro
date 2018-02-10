package schema

import "global"

//会员退水记录
type MemberRetreatWaterRecord struct {
	Id          int64   `xorm:"id"`
	SiteId      string  `xorm:"site_id"`       //操作站点id
	SiteIndexId string  `xorm:"site_index_id"` //站点前台id
	Account     string  `xorm:"account"`       //会员账号
	StartTime   int64   `xorm:"start_time"`    //开始时间
	EndTime     int64   `xorm:"end_time"`      //结束时间
	PeriodsId   int64   `xorm:"periods_id"`    //期数id
	MemberId    int64   `xorm:"member_id"`     //所属会员id
	LevelId     string  `xorm:"level_id"`      //所属层级id
	Betall      float64 `xorm:"betall"`        //有效总投注
	AllMoney    float64 `xorm:"all_money"`     //总返水金额
	SelfMoney   float64 `xorm:"self_money"`    //自助反水金额
	RebateWater float64 `xorm:"rebate_water"`  //本次退水金额
	Status      int8    `xorm:"status"`        //是否已返佣 0未操作1已操作
	CreateTime  int64   `xorm:"create_time"`   //创建时间
}

func (*MemberRetreatWaterRecord) TableName() string {
	return global.TablePrefix + "member_retreat_water_record"
}
