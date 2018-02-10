package back

//入款银行列表
type BankInBack struct {
	Title       string  `xorm:"title" json:"title"`              //银行名称(bank)
	BankId      int64   `xorm:"bank_id" json:"bankId"`           // 银行id(bank_income_set)
	Account     string  `xorm:"account" json:"account"`          //收款账号
	OpenBank    string  `xorm:"open_bank" json:"openBank"`       // 开户行
	Payee       string  `xorm:"payee" json:"payee"`              // 收款人
	StopBalance float64 `xorm:"stop_balance" json:"stopBalance"` // 停用金额
	Status      int8    `xorm:"status" json:"status"`
	Remark      string  `xorm:"remark" json:"remark"`
	Id          int64   `xorm:"id" json:"id"`
	LevelNum    string  `xorm:"count(sales_site_bank_income_member_level.set_id)" json:"levelNum"` // 会员层级id
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"`
	SiteId      string  `xorm:"site_id" json:"siteId"`
}

//一条入款银行设定的返回结构体
type OneBankPaySetBack struct {
	OneBankPaySet *OneBankPaySet `json:"one_bank_pay_set"`
	LevelId       []string       `xorm:"level_id" json:"level_id"` // 会员层级id
}

//一条入款银行设定的返回结构体
type OneBankPaySet struct {
	Id          int64   `xorm:"'id' PK autoincr" json:"id"`
	CreateTime  int64   `xorm:"create_time created" json:"createTime"` // 创建时间
	SiteId      string  `xorm:"site_id" json:"siteId"`                 // 站点id
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"`      //站点前台id
	PayTypeId   int64   `xorm:"pay_type_id" json:"payTypeId"`          // 支付类型id
	BankId      int64   `xorm:"bank_id" json:"bankId"`                 // 银行id
	Account     string  `xorm:"account" json:"account"`                //收款账号
	OpenBank    string  `xorm:"open_bank" json:"openBank"`             // 开户行
	Payee       string  `xorm:"payee" json:"payee"`                    // 收款人
	StopBalance float64 `xorm:"stop_balance" json:"stopBalance"`       // 停用金额
	QrCode      string  `xorm:"qr_code" json:"qrCode"`                 // 二维码图片数据
	Remark      string  `xorm:"remark" json:"remark"`                  // 备注
	Status      int8    `xorm:"status" json:"status"`                  //状态,0禁用1正常
}

type LelvelId struct {
	LevelId string `json:"level_id"` //层级id
}

//存款记录
type DepositRecordBack struct {
	Id           int64   `xorm:"id" json:"id"`
	CreateTime   int64   `xorm:"'create_time'created" json:"create_time"` // 创建时间
	UserName     string  `xorm:"user_name" json:"user_name"`              // 会员账号
	OrderNum     string  `xorm:"order_num" json:"order_num"`              // 订单号
	DepositMoney float64 `xorm:"deposit_money" json:"deposit_money"`      // 存入金额
	Remark       string  `xorm:"remark" json:"remark"`                    // 备注
	Atotal       []AllTotal
}
type AllTotal struct {
	Subtotal float64 `json:"subtotal"`       //小计
	Total    float64 `xorm:"a" json:"total"` //总计
}

//适用层级
type ApplicationLevelBack struct {
	SetId       int64  `xorm:"set_id" json:"set_id"`
	LevelId     string `xorm:"level_id" json:"level_id"`
	Description string `xoem:"description" json:"description"`
}
