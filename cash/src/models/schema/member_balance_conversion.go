package schema

import "global"

//会员额度转换
type MemberBalanceConversion struct {
	Id          int64   `insert:"id;auto" xorm:"id PK autoincr"`
	SiteId      string  `insert:"site_id" xorm:"site_id"`                   //操作站点id
	SiteIndexId string  `insert:"site_index_id" xorm:"site_index_id"`       //站点前台id
	MemberId    int64   `insert:"member_id" xorm:"member_id"`               //会员id
	Account     string  `insert:"account" xorm:"account"`                   //会员账号
	AgencyId    int64   `insert:"agency_id" xorm:"agency_id"`               //会员所属代理id
	Money       float64 `insert:"money" xorm:"money"`                       //金额
	FromType    int64   `insert:"from_type" xorm:"from_type"`               //转入类型,0为系统金额,1000的时候是站点视讯额度修改但不是此类型
	ForType     int64   `insert:"for_type" xorm:"for_type"`                 //转出类型,0为系统金额,1000的时候是站点视讯额度修改但不是此类型
	CreateTime  int64   `insert:"create_time" xorm:"'create_time' created"` //操作时间
	UpdateTime  int64   `insert:"update_time" xorm:"update_time"`           //转换确认时间
	Status      int8    `insert:"status" xorm:"status"`                     //状态 1：成功  2：失败（默认是失败）
	DoUserId    int64   `insert:"do_user_id" xorm:"do_user_id"`             //操作人
	DoUserType  int8    `insert:"do_user_type" xorm:"do_user_type"`         //操作人类型1平台管理员2会员
	Remark      string  `insert:"remark" xorm:"remark"`                     //备注
	TradeNo     string  `insert:"trade_no" xorm:"trade_no"`                 //订单号
}

func (*MemberBalanceConversion) TableName() string {
	return global.TablePrefix + "member_balance_conversion"
}
