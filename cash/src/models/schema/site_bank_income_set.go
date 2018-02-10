package schema

import "global"

//站点公司入款银行设定表
type SiteBankIncomeSet struct {
	Id          int64   `xorm:"'id' PK autoincr"`
	CreateTime  int64   `xorm:"create_time created"` // 创建时间
	SiteId      string  `xorm:"site_id"`             // 站点id
	SiteIndexId string  `xorm:"site_index_id"`       //站点前台id
	PayTypeId   int64   `xorm:"pay_type_id"`         // 支付类型id
	BankId      int64   `xorm:"bank_id"`             // 银行id
	Account     string  `xorm:"account"`             //收款账号
	OpenBank    string  `xorm:"open_bank"`           // 开户行
	Payee       string  `xorm:"payee"`               // 收款人
	StopBalance float64 `xorm:"stop_balance"`        // 停用金额
	QrCode      string  `xorm:"qr_code"`             // 二维码图片数据
	Remark      string  `xorm:"remark"`              // 备注
	Status      int8    `xorm:"status"`              //状态,0禁用1正常
	DeleteTime  int64   `xorm:"delete_time"`         //删除时间
	Sort        int64   `xorm:"sort" json:"sort"`    //排序，数字越大，排越前面
}

func (*SiteBankIncomeSet) TableName() string {
	return global.TablePrefix + "site_bank_income_set"
}
