package schema

import "global"

//自助优惠申请struct
type SelfHelpApplyfor struct {
	Id            int64   `xorm:"id"`             //主键id
	SiteId        string  `xorm:"site_id"`        //站点id
	SiteIndexId   string  `xorm:"site_index_id"`  //站点前台id
	Status        int8    `xorm:"status"`         //状态3.待审核2.审核不通过1.审核通过
	ApplyAccount  string  `xorm:"apply_account"`  //申请者账号
	ApplyName     string  `xorm:"apply_name"`     //申请者姓名
	ApplyTime     int64   `xorm:"apply_time"`     //申请时间(年月日)
	ApplyReason   string  `xorm:"apply_reason"`   //申请原因
	ApplyMoney    float64 `xorm:"apply_money"`    //申请金额
	GiveawayMoney float64 `xorm:"giveaway_money"` //赠送金额
	AuditTime     int64   `xorm:"audit_time"`     //审核时间
	OperateId     int64   `xorm:"operate_id"`     //操作者id
	ActivityClass int8    `xorm:"activity_class"` //活动类别
	RejectReason  string  `xorm:"reject_reason"`  //拒绝原因
	ThroughRemark string  `xorm:"through_remark"` //通过备注
}

func (*SelfHelpApplyfor) TableName() string {
	return global.TablePrefix + "self_help_applyfor"
}
