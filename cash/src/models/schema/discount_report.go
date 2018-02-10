package schema

import "global"

//统计每日优惠
type DiscountReport struct {
	Id             int     `insert:"id;auto" xorm:"id PK autoincr"`            //
	SiteId         string  `insert:"site_id" xorm:"site_id" `                  //站点id
	SiteIndexId    string  `insert:"site_index_id" xorm:"site_index_id" `      //站点前台id
	MemberId       int64   `insert:"member_id" xorm:"member_id" `              //会员id
	AgentId        int64   `insert:"agent_id" xorm:"agent_id" `                //代理id
	SecondAgencyId int64   `insert:"second_agency_id" xorm:"second_agency_id"` //总代id
	FirstAgencyId  int64   `insert:"first_agency_id" xorm:"first_agency_id"`   //股东id
	Account        string  `insert:"account" xorm:"account" `                  //账号
	DoTime         int64   `insert:"do_time" xorm:"do_time"`                   //操作时间
	Num            int64   `insert:"num" xorm:"num" `                          //总笔数
	DiscountMoney  float64 `insert:"discount_money" xorm:"discount_money" `    //优惠金额
	DayType        string  `insert:"day_type" xorm:"day_type" `                //天的标记
	DayTime        int64   `insert:"day_time" xorm:"day_time" `                //统计哪天
}

func (m *DiscountReport) TableName() string {
	return global.TablePrefix + "discount_report"
}
