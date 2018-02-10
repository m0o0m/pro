package schema

import "global"

//即时稽核
type MemberAuditRecord struct {
	Id            int64   `xorm:"id" json:"id"`                         //会员id
	SiteId        string  `xorm:"site_id" json:"site_id"`               //站点id
	SiteIndexId   string  `xorm:"site_index_id" json:"site_index_id"`   //前台id
	MemberId      int64   `xorm:"member_id" json:"member_id"`           //会员id
	Account       string  `xorm:"account" json:"account"`               //账号
	BeginTime     int64   `xorm:"begin_time" json:"begin_time"`         //稽核开始时间
	EndTime       int64   `xorm:"end_time" json:"end_time"`             //稽核结束时间
	NormalMoney   float64 `xorm:"normal_money" json:"normal_money"`     //常态稽核
	MultipleMoney float64 `xorm:"multiple_money" json:"multiple_money"` //综合稽核
	Money         float64 `xorm:"money" json:"money"`                   //存款金额
	AdminMoney    float64 `xorm:"admin_money" json:"admin_money"`       //扣除的行政费用
	DepositMoney  float64 `xorm:"deposit_money" json:"deposit_money"`   //优惠金额
	RelaxMoney    int64   `xorm:"relax_money" json:"relax_money"`       //放宽额度
	Status        int     `xorm:"status" json:"status"`                 //状态
}

func (*MemberAuditRecord) TableName() string {
	return global.TablePrefix + "member_audit"
}
