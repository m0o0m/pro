package schema

import "global"

//站点支持的第三方列表(至于银行卡只记录被剔除的数据)
type SiteOnlineIncomeSet struct {
	Id             int64   `xorm:"'id' PK autoincr"`
	SiteId         string  `xorm:"site_id"`         // 站点id
	SiteIndexId    string  `xorm:"site_index_id"`   // 站点前台id
	PayTypeId      int64   `xorm:"pay_type_id"`     // 支付类型id
	BankId         int64   `xorm:"bank_id"`         // 银行id
	PayDomain      string  `xorm:"pay_domain"`      // 支付域名
	CallbackDomain string  `xorm:"callback_domain"` // 回调地址
	ShopId         string  `xorm:"shop_id"`         // 商户id
	PublicKey      string  `xorm:"public_key"`      // 公钥
	PrivateKey     string  `xorm:"private_key"`     // 私钥
	StopBalance    float64 `xorm:"stop_balance"`    // 支付限额
	ThirdId        int64   `xorm:"third_id"`        // 第三方平台id
	Client         int8    `xorm:"client"`          // 适用设备 0全部 1pc端 2手机端
	Code           string  `xorm:"code"`            // 支付方式code码
}

func (*SiteOnlineIncomeSet) TableName() string {
	return global.TablePrefix + "site_online_income_set"
}
