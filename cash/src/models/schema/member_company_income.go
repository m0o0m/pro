package schema

import "global"

//公司入款记录
type MemberCompanyIncome struct {
	Id              int64   `xorm:"'id' PK autoincr"`
	CreateTime      int64   `xorm:"'create_time' created"` // 创建时间
	SiteId          string  `xorm:"site_id"`               //操作站点id
	SiteIndexId     string  `xorm:"site_index_id"`         // 站点前台id
	MemberId        int64   `xorm:"member_id"`             // 会员id
	Account         string  `xorm:"account"`               // 会员账号
	BankId          int64   `xorm:"bank_id"`               // 存入银行id
	BankName        string  `xorm:"bank_name"`             // 存入银行名称
	LevelId         string  `xorm:"level_id"`              // 会员所属层级id
	AgencyId        int64   `xorm:"agency_id"`             // 会员所属代理id
	AgencyAccount   string  `xorm:"agency_account"`        // 会员所属代理账号
	IsFirstDeposit  int     `xorm:"is_first_deposit"`      //是否首次存储
	OrderNum        string  `xorm:"order_num"`             // 订单号
	DepositUsername string  `xorm:"deposit_username"`      // 存款人名称
	DepositMoney    float64 `xorm:"deposit_money"`         // 存入金额
	DepositDiscount float64 `xorm:"deposit_discount"`      // 存入优惠
	OtherDiscount   float64 `xorm:"other_discount"`        // 其他优惠
	DepositCount    float64 `xorm:"deposit_count"`         // 存入总额
	Remark          string  `xorm:"remark"`                // 备注
	ClientType      int8    `xorm:"client_type"`           // 客户端类型 1pc 2wap 3android 4ios
	Status          int8    `xorm:"status"`                // 状态,0未处理1已确认2已取消
	DoAgencyId      int64   `xorm:"do_agency_id"`          // 操作人id(agency表主键)
	DoAgencyRemark  string  `xorm:"do_agency_remark"`      // 操作人备注
	UpdateTime      int64   `xorm:"update_time"`           // 操作人操作时间
	SetId           int64   `xorm:"set_id"`                //入款银行设定id
	DepositMethod   int8    `xorm:"deposit_method"`        //存款类型
	DepositTime     int64   `xorm:"deposit_time"`          // 存款时间
}

func (*MemberCompanyIncome) TableName() string {
	return global.TablePrefix + "member_company_income"
}
