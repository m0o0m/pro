package schema

import "global"

//公共银行列表
type Bank struct {
	Id             int64  `xorm:"'id' PK autoincr"`
	Title          string `xorm:"title"`               //银行名称
	Icon           string `xorm:"icon"`                //银行图标
	PayTypeId      int64  `xorm:"pay_type_id"`         // 支付类型id
	IsIncome       int8   `xorm:"is_income"`           //入款银行是否可用
	IsOut          int8   `xorm:"is_out"`              // 可出款银行是否用
	IsThird        int8   `xorm:"is_third"`            // 第三方平台是否可用
	BankWebsiteUrl string `xorm:"bank_website_url"`    //银行官网链接
	Status         int8   `xorm:"status"`              // 状态
	CreateTime     int64  `xorm:"create_time created"` // 创建时间
	DeleteTime     int64  `xorm:"delete_time"`         // 删除时间
}

func (*Bank) TableName() string {
	return global.TablePrefix + "bank"
}
