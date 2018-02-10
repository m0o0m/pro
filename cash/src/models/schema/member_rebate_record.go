package schema

import "global"

//会员返佣记录对应商品金额
type MemberRebateRecord struct {
	Id                   int64                        `xorm:"id PK autoincr"`
	SiteId               string                       `xorm:"site_id"`               //操作站点id
	SiteIndexId          string                       `xorm:"site_index_id"`         //站点前台id
	PeriodsId            int64                        `xorm:"periods_id"`            //总计id
	MemberId             int64                        `xorm:"member_id"`             //所属会员id
	Betting              float64                      `xorm:"betting"`               //有效总投注
	Rebate               float64                      `xorm:"rebate"`                //返佣金额
	Status               int8                         `xorm:"status"`                //是否已返佣 1返佣  2冲销
	CreateTime           int64                        `xorm:"'create_time' created"` //创建时间
	RebateRecordProducts []*MemberRebateRecordProduct `xorm:"-"`                     //返佣商品
}

func (*MemberRebateRecord) TableName() string {
	return global.TablePrefix + "member_rebate_record"
}
