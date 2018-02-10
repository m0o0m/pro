package schema

import "global"

//会员现金流水
type MemberCashRecord struct {
	ID            int64   `xorm:"id"`
	SiteId        string  `xorm:"site_id"`               //操作站点id
	SiteIndexId   string  `xorm:"site_index_id"`         //站点前台id
	MemberId      int64   `xorm:"member_id"`             //会员id
	UserName      string  `xorm:"user_name"`             //会员账号
	AgencyId      int64   `xorm:"agency_id"`             //会员所属代理id
	AgencyAccount string  `xorm:"agency_account"`        //会员所属代理账号
	SourceType    int64   `xorm:"source_type"`           //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	Type          int64   `xorm:"type"`                  //1.存入2.取出
	TradeNo       string  `xorm:"trade_no"`              //下单单号
	Balance       float64 `xorm:"balance"`               //金额
	DisBalance    float64 `xorm:"dis_balance"`           //优惠金额
	AfterBalance  float64 `xorm:"after_balance"`         //操作后余额
	Remark        string  `xorm:"remark"`                //备注
	ClientType    int64   `xorm:"client_type"`           //客户端类型1pc 2wap 3android 4ios
	CreateTime    int64   `xorm:"'create_time' created"` //添加时间
	DeleteTime    int64   `xorm:"'delete_time'"`         //删除时间
}

func (*MemberCashRecord) TableName() string {
	return global.TablePrefix + "member_cash_record"
}
