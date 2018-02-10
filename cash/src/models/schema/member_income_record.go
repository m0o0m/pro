package schema

import "global"

//会员存款记录
type MemberIncomeRecord struct {
	Id           int64   `xorm:"id"`
	SourceId     int64   `xorm:"source_id"`      //数据来源id
	SourceType   int8    `xorm:"source_type"`    //数据来源类型(1人工存取款 2公司入款 3线上入款...)
	SiteId       string  `xorm:"site_id"`        //站点id
	SiteIndexId  string  `xorm:"site_index_id"`  //站点前台id
	Status       int8    `xorm:"status"`         //状态(1表示该条存款已被取过 2表示该条存款未被取过)
	MemberId     int64   `xorm:"member_id"`      //会员id
	MemberName   string  `xorm:"member_name"`    //会员真实姓名
	DepositMoney float64 `xorm:"deposit_money"`  //存款金额
	CatmGive     float64 `xorm:"catm_give"`      //存款优惠
	AtmGive      float64 `xorm:"atm_give"`       //汇款优惠
	IsMultiple   int8    `xorm:"is_multiple"`    //是否有综合稽核
	IsNormal     int8    `xorm:"is_normal"`      //是否常态稽核
	MultipleTime int64   `xorm:"multiple_times"` //综合稽核倍数
	NormalTime   int64   `xorm:"normal_times"`   //常态稽核倍数
}

func (*MemberIncomeRecord) TableName() string {
	return global.TablePrefix + "member_income_record"
}
