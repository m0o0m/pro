package schema

import "global"

//自助返水查询
type MemberRetreatWaterRecordSelf struct {
	Id          int64   `xorm:"id PK autoincr"`
	PeriodId    int64   `xorm:"period_id"`     //总计id
	SiteId      string  `xorm:"site_id"`       //操作站点id
	SiteIndexId string  `xorm:"site_index_id"` //站点前台id
	MemberId    string  `xorm:"member_id"`     //会员id
	Account     string  `xorm:"account"`       //会员账号
	AdminUser   string  `xorm:"admin_user"`    //操作人
	Betall      float64 `xorm:"betall"`        //总有效投注
	RebateWater float64 `xorm:"rebate_water"`  //本次退水金额
	CreateTime  int64   `xorm:"create_time"`   //创建时间就是反水时间
}

func (*MemberRetreatWaterRecordSelf) TableName() string {
	return global.TablePrefix + "member_retreat_water_record_self"
}
