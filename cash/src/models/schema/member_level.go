package schema

import "global"

//站点会员层级表
type MemberLevel struct {
	SiteId          string  `xorm:"'site_id' PK"`          //站点Id
	SiteIndexId     string  `xorm:"'site_index_id' PK"`    //站点前台Id
	LevelId         string  `xorm:"'level_id' PK"`         //会员层级名称
	PaySetId        int64   `xorm:"pay_set_id" `           //支付设置Id
	Description     string  `xorm:"description" `          //层级描述
	DepositNum      int64   `xorm:"deposit_num" `          //取款次数
	DepositCount    float64 `xorm:"deposit_count" `        //取款总额
	StartTime       int64   `xorm:"start_time" `           //会员加入开始时间
	IsDefault       int8    `xorm:"is_default" `           //是否为默认层级
	EndTime         int64   `xorm:"end_time" `             //会员加入结束时间
	DepositNumber   int64   `xorm:"deposit_number"`        //存款次数(统计)
	DepositTotal    float64 `xorm:"deposit_total"`         //存款总额(统计)
	WithdrwalNumber int64   `xorm:"withdrawal_number"`     //提款次数(统计)
	Remark          string  `xorm:"remark" `               //备注
	CreateTime      int64   `xorm:"'create_time' created"` //创建时间
	DeleteTime      int64   `xorm:"delete_time" `          //删除时间
	IsSelfRebate    int8    `xorm:"is_self_rebate" `       //是否开启自动返水功能。(1.开启2.未开启)
	Count           int64   `xorm:"count"`                 //会员数量
}

func (*MemberLevel) TableName() string {
	return global.TablePrefix + "member_level"
}
