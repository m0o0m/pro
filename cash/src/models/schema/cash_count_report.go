package schema

import "global"

type CashCountReport struct {
	ID             int64   `insert:"id;auto" xorm:"id PK autoincr"`
	SiteId         string  `insert:"site_id" xorm:"site_id"`                   //站点id
	SiteIndexId    string  `insert:"site_index_id" xorm:"site_index_id"`       //前台id
	IntoStyle      int64   `insert:"into_style" xorm:"into_style"`             //入款方式:0人工存入1公司入款2线上入款
	MemberId       int64   `insert:"member_id" xorm:"member_id"`               //'会员id',
	AgentId        int64   `insert:"agent_id" xorm:"agent_id"`                 //代理id
	SecondAgencyId int64   `insert:"second_agency_id" xorm:"second_agency_id"` //'总代id',
	FirstAgencyId  int64   `insert:"first_agency_id" xorm:"first_agency_id"`   //股东id',
	Account        string  `insert:"account" xorm:"account"`                   //账号
	DoTime         int64   `insert:"do_time" xorm:"do_time"`                   //操作时间
	Num            int64   `insert:"num" xorm:"num"`                           //笔数
	CashMoney      float64 `insert:"cash_money" xorm:"cash_money"`             //金额
	DayType        string  `insert:"day_type" xorm:"day_type"`                 //天标志
	DayTime        int64   `insert:"day_time" xorm:"day_time"`                 //统计哪天
}

func (m *CashCountReport) TableName() string {
	return global.TablePrefix + "cash_count_report"
}
