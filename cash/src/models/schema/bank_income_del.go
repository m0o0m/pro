package schema

import "global"

//入款银行类型剔除表
type BankIncomeDel struct {
	SiteId      string `xorm:"site_id"`       // 站点id
	SiteIndexId string `xorm:"site_index_id"` // 站点前台id
	BankId      int64  `xorm:"bank_id"`       // 银行id
}

func (*BankIncomeDel) TableName() string {
	return global.TablePrefix + "bank_income_del"
}
