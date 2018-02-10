package schema

import "global"

//会员返佣期数(总计)
type MemberRebateCommission struct {
	Id          int64   `xorm:"id PK autoincr"`
	SiteId      string  `xorm:"site_id"`               //操作站点id
	SiteIndexId string  `xorm:"site_index_id"`         //站点前台id
	AdminUser   string  `xorm:"admin_user"`            //操作者
	Event       string  `xorm:"event"`                 //事件
	StartTime   int64   `xorm:"start_time"`            //开始时间
	EndTime     int64   `xorm:"end_time"`              //结束时间
	CreateTime  int64   `xorm:"'create_time' created"` //创建时间
	NoPeopleNum int64   `xorm:"no_people_num"`         //冲销人数
	PeopleNum   int64   `xorm:"people_num"`            //退水人数
	Money       float64 `xorm:"money"`                 //金额
	Bet         int     `xorm:"bet"`                   //综合打码倍数
	TotalBet    float64 `xorm:"total_bet"`             //总有效打码
}

func (*MemberRebateCommission) TableName() string {
	return global.TablePrefix + "member_rebate_commission"
}
