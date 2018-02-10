package schema

import "global"

//会员自助申请 优惠活动表
type MemberAutoApplypro struct {
	Id               int64   `xorm:"'id' PK autoincr"`
	PromotionId      int64   `xorm:"promotion_id"`      //活动id
	PromotionTitle   string  `xorm:"promotion_title"`   //活动标题
	SiteId           string  `xorm:"site_id"`           //
	SiteIndexId      string  `xorm:"site_index_id"`     //
	MemberId         int64   `xorm:"member_id"`         //会员id
	Account          string  `xorm:"account"`           //用户名
	Status           int64   `xorm:"status"`            //（1
	ApplyMoney       float64 `xorm:"apply_money"`       //申请金额
	GiveMoney        float64 `xorm:"give_money"`        //
	Applyreason      string  `xorm:"applyreason"`       //申请理由
	Denyreason       string  `xorm:"denyreason"`        //拒绝理由
	AgreeRemark      string  `xorm:"agree_remark"`      //通过审核备注
	Createtime       int64   `xorm:"createtime"`        //申请时间
	Updatetime       int64   `xorm:"updatetime"`        //审核时间
	HandlerId        int64   `xorm:"handler_id"`        //操作者id
	HandlerName      string  `xorm:"handler_name"`      //操作者名称
	PromotionContent string  `xorm:"promotion_content"` //活动内容
	IsNormality      int64   `xorm:"is_normality"`      //(1:不参加常态稽核,2:参加常态稽核)
	IsComplex        int64   `xorm:"is_complex"`        //(1:不参加综合打码稽核,2:参加综合打码稽核)
	ComplexAudit     int64   `xorm:"complex_audit"`     //综合打码量
	AuditId          int64   `xorm:"audit_id"`          //稽核记录id
}

func (*MemberAutoApplypro) TableName() string {
	return global.TablePrefix + "member_auto_applypro"
}
