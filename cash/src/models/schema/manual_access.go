package schema

import "global"

//人工存款与取款记录
type ManualAccess struct {
	Id                int64   `xorm:"'id' PK autoincr"`
	CreateTime        int64   `xorm:"create_time"`         // 创建时间
	SiteId            string  `xorm:"site_id"`             //站点id
	SiteIndexId       string  `xorm:"site_index_id"`       // 站点前台id
	MemberId          int64   `xorm:"member_id"`           // 会员id
	Account           string  `xorm:"account"`             // 会员账号
	AgencyId          int64   `xorm:"agency_id"`           // 会员所属代理id
	AgencyAccount     string  `xorm:"agency_account"`      // 会员所属代理账号
	Balance           float64 `xorm:"balance"`             // 会员存取款之前余额
	AccessType        int8    `xorm:"access_type"`         // 存款还是取款,1存款 2取款
	Money             float64 `xorm:"money"`               // 本次存取款金额
	IsDepositDiscount int8    `xorm:"is_deposit_discount"` // 是否有存款优惠
	DepositDiscount   float64 `xorm:"deposit_discount"`    //存款优惠
	IsRemitDiscount   int8    `xorm:"is_remit_discount"`   // 是否有汇款优惠
	RemitDiscount     float64 `xorm:"remit_discount"`      // 汇款优惠
	IsCodeCount       int8    `xorm:"is_code_count"`       // 是否有综合稽核
	CodeCount         int64   `xorm:"code_count"`          // 综合稽核打码量
	IsRoutineCheck    int8    `xorm:"is_routine_check"`    // 是否有常态性稽核
	DepositType       int8    `xorm:"deposit_type"`        // 存款项目(类型)1人工存入2存款优惠3负数额度归零4取消出款5返点优惠6活动优惠7其他15体育投注余额16额度掉单17体育投注余额18额度掉单
	IsWriteRebate     int8    `xorm:"is_write_rebate"`     //是否写入退佣
	Remark            string  `xorm:"remark"`              // 备注
	DoAgencyId        int64   `xorm:"do_agency_id"`        //操作人id(agency表主键)
	DoAgencyAccount   string  `xorm:"do_agency_account"`   // 操作人账号
}

func (*ManualAccess) TableName() string {
	return global.TablePrefix + "manual_access"
}
