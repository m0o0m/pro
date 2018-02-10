package schema

import "global"

//线上入款列表
type OnlineEntryRecord struct {
	Id                   int64   `xorm:"'id' PK autoincr"`
	ThirdOrderNumber     string  `xorm:"third_order_number"`     //第三方平台订单号
	SourceDeposit        int     `xorm:"source_deposit"`         //入款来源(1.pc 2.wap)
	LocalOperateTime     int64   `xorm:"create_time"`            //本地平台操作时间(本平台点击跳转到三方平台付账时候的时间,也是前台查询时间的搜索条件)
	ThirdPayTime         int64   `xorm:"third_pay_time"`         //第三方平台转账时间(存款时间)
	AmountDeposit        float64 `xorm:"amount_deposit"`         //存款金额
	MemberAccount        string  `xorm:"member_account"`         //存款会员账号
	MemberId             int64   `xorm:"member_id"`              //存款会员id
	Level                string  `xorm:"level"`                  //存款会员层级
	SiteIndexId          string  `xorm:"site_index_id"`          //前台站点id
	SiteId               string  `xorm:"site_id"`                //站点Id
	AgencyId             int64   `xorm:"agency_id"`              //所属代理id
	AgencyAccount        string  `xorm:"agency_account"`         //代理账号
	Status               int     `xorm:"status"`                 //状态(1.未支付2.已经支付3.已取消)
	IsDiscount           int     `xorm:"is_discount"`            //是否有优惠(1.有优惠2.没有优惠)
	DepositDiscount      float64 `xorm:"deposit_discount"`       //存款优惠金额
	OtherDepositDiscount float64 `xorm:"other_deposit_discount"` //其他优惠金额
	OperateId            int     `xorm:"operate_id"`             //操作者id
	Remark               string  `xorm:"remark"`                 //备注
	ThirdId              int     `xorm:"third_id"`               //第三方支付平台id
	PaidType             int     `xorm:"paid_type"`              //线上支付类型
	PaidSetupId          int     `xorm:"paid_setup_id"`          //线上支付设定id
	IsFirstDeposit       int     `xorm:"is_first_deposit"`       //是否首次充值(1.不是2.是)
}

func (*OnlineEntryRecord) TableName() string {
	return global.TablePrefix + "online_entry_record"
}
