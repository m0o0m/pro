package back

//三方银行剔除列表
type BankListBack struct {
	Id       int64  `json:"bankId"`
	Title    string `json:"bankName"` //银行名称
	Status   int8   `json:"status"`   // 状态
	PayId    int64  `json:"payId"`
	BankCode string `json:"bankCode"`
}

//入款、出款银行剔除列表
type BankTwoBack struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`  //银行名称
	Status int8   `json:"status"` // 状态
}

//银行列表（admin）
type BankCardList struct {
	Id             int64  `xorm:"id" json:"id"`
	Title          string `xorm:"title" json:"title"`                     //银行名称
	IsIncome       int8   `xorm:"is_income" json:"isIncome"`              //入款银行是否可用
	IsOut          int8   `xorm:"is_out" json:"isOut"`                    // 可出款银行是否用
	Status         int8   `xorm:"status" json:"status"`                   // 状态
	BankWebsiteUrl string `xorm:"bank_website_url" json:"bankWebsiteUrl"` //银行官网
}

//银行列表下拉框
type BankCardListDrop struct {
	Id    int64  `xorm:"id" json:"id"`
	Title string `xorm:"title" json:"title"` //银行名称
}

//银行列表（admin）
type BankCardListDropOutDel struct {
	SiteId      string `xorm:"site_id"`       // 站点id
	SiteIndexId string `xorm:"site_index_id"` // 站点前台id
	BankId      int64  `xorm:"bank_id"`       // 银行id
}

//入款、出款银行列表(剔除后)
type SiteBank struct {
	Id    int64  `json:"id"`
	Title string `json:"title"` //银行名称
}
