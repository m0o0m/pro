package back

//催款账单
type SiteMoneyPress struct {
	Id         int64   `json:"id"`
	SiteId     string  `json:"siteId"`     //站点
	AddDate    int64   `json:"addDate"`    //添加时间
	UpdateDate int64   `json:"updateDate"` //更新时间
	Qishu      int64   `json:"qishu"`      //期数
	Money      float64 `json:"money"`      //应缴金额
	PayName    string  `json:"payName"`    //收款人姓名
	PayAddress string  `json:"payAddress"` //收款地址
	Bank       string  `json:"bank"`       //收款银行
	PayCard    string  `json:"payCard"`    //银行账号
	Remark     string  `json:"remark"`     //备注
	Status     int8    `json:"status"`     //1业主未提交,2业主已提交
	State      int8    `json:"state"`      //1已催款,2已完成催款
}

//催款账单site
type SiteMoneyPressSiteBack struct {
	SiteId string `xorm:"id" json:"siteId"` //站点
}
