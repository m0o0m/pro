package schema

import "global"

//出款稽核
type MemberAudit struct {
	Id            int64   `xorm:"id PK autoincr"`
	SiteId        string  `xorm:"site_id"`        //站点id
	SiteIndexId   string  `xorm:"site_index_id"`  //站点前台id
	MemberId      int64   `xorm:"member_id"`      //会员id
	Account       string  `xorm:"account"`        //会员账号
	BeginTime     int64   `xorm:"begin_time"`     //稽核开始时间
	EndTime       int64   `xorm:"end_time"`       //稽核结束时间(确认出款的时候将所有的没有稽核的修改为当前操作时间)
	NormalMoney   float64 `xorm:"normal_money"`   //常态稽核金额
	MultipleMoney float64 `xorm:"multiple_money"` //综合稽核金额
	Money         float64 `xorm:"money"`          //存款金额
	AdminMoney    float64 `xorm:"admin_money"`    //扣除的行政费用
	DepositMoney  float64 `xorm:"deposit_money"`  //优惠金额
	RelaxMoney    int64   `xorm:"relax_money"`    //放宽额度
	Status        int64   `xorm:"status"`         //'稽核状态  1未处理  2已处理',
}

func (*MemberAudit) TableName() string {
	return global.TablePrefix + "member_audit"
}
