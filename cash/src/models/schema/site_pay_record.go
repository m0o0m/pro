package schema

import "global"

//站点额度购买充值记录
type SitePayRecord struct {
	Id          int64   `xorm:"id PK autoincr"`
	SiteId      string  `xorm:"site_id"`           //站点ID
	SiteIndexId string  `xorm:"site_index_id"`     //多站点ID
	OrderNum    string  `xorm:"order_num"`         // 订单号
	AdminUser   string  `xorm:"admin_user"`        // 提交者
	Money       float64 `xorm:"money"`             // 交易额度
	Type        int8    `xorm:"type"`              // 支付方式  1第三方入款，2公司入款
	Bank        string  `xorm:"bank"`              // 银行类型
	DoTime      int64   `xorm:"'do_time' updated"` // 操作时间
	UpdateTime  int64   `xorm:"update_time"`       // 支付时间
	State       int8    `xorm:"state"`             // 状态1未支付2支付
	Remark      string  `xorm:"remark"`            // 备注
	PayId       int8    `xorm:"pay_id"`
}

func (*SitePayRecord) TableName() string {
	return global.TablePrefix + "site_pay_record"
}
