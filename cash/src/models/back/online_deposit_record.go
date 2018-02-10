package back

//OnlineListDeposit 线上入款列表
type OnlineListDeposit struct {
	Id                   int64   `json:"id"`
	ThirdOrderNumber     string  `json:"thirdOrderNumber"`     //订单号
	SourceDeposit        int     `json:"sourceDeposit"`        //入款来源
	CreateTime           int64   `json:"createTime"`           //操作时间
	AgencyAccount        string  `json:"agencyAccount"`        //代理账号
	MemberAccount        string  `json:"memberAccount"`        //存款人账号
	ThirdPayTime         int64   `json:"thirdPayTime"`         //存款时间
	AmountDeposit        float64 `json:"amountDeposit"`        //存款金额
	DepositDiscount      float64 `json:"depositDiscount"`      //存款优惠
	OtherDepositDiscount float64 `json:"otherDepositDiscount"` //其他优惠
	Status               int     `json:"status"`               //状态1.未支付2.已经支付3.已取消4已确认
	Title                string  `json:"title"`                //支付方式名称
	IsFirstDeposit       int     `json:"isFirstDeposit"`       //是否首次存款
	Account              string  `json:"account"`              //操作者账号
}

//OnlineDepositL 线上入款返回列表
type OnlineDepositBack struct {
	TotalMoney               float64             `json:"totalMoney"`               //总计存入金额
	TotalDepositDiscount     float64             `json:"totalDepositDiscount"`     //总计存款优惠
	TotalOtherDiscount       float64             `json:"totalOtherDiscount"`       //总计其他优惠
	TotalDeposit             float64             `json:"totalDeposit"`             //总计存入总金额
	PageTotalMoney           float64             `json:"pageTotalMoney"`           //小计存入金额
	PageTotalDepositDiscount float64             `json:"pageTotalDepositDiscount"` //小计存款优惠
	PageTotalOtherDiscount   float64             `json:"pageTotalOtherDiscount"`   //小计其他优惠
	PageTotalDeposit         float64             `json:"pageTotalDeposit"`         //小计存入总金额
	TotalCount               int64               `json:"totalCount"`               //总计笔数
	PageCount                int                 `json:"pageCount"`                //小计笔数
	AllData                  []OnlineListDeposit //本页面的所有数据
}

//公司入款列表
type CompenyIncomeList struct {
	Id              int64   `json:"id"`
	OrderNum        string  `json:"orderNum"`                   //订单号
	ClientType      int     `json:"clientType"`                 //入款来源:系统
	UpdateTime      int64   `json:"updateTime"`                 //操作时间
	AgencyAccount   string  `json:"agencyAccount"`              //经销商账号
	Account         string  `json:"account"`                    //存款人账号(会员账号)
	CreateTime      int64   `json:"createTime"`                 //提交时间
	DepositTime     int64   `json:"depositTime"`                // 存款时间
	LevelId         string  `xorm:"level_id" json:"levelId"`    // 会员所属层级id
	DepositMoney    float64 `json:"depositMoney"`               //存款金额
	DepositDiscount float64 `json:"depositDiscount"`            //存款优惠
	OtherDiscount   float64 `json:"otherDiscount"`              //其他优惠
	DepositCount    float64 `json:"depositCount"`               //存入总额
	Payee           string  `json:"payee"`                      //入款银行卡主
	BankAccount     string  `xorm:"account" json:"bankAccount"` //入款银行卡号
	Remark          string  `json:"remark"`                     //备注
	Status          int     `json:"status"`                     //状态
	IsFirstDeposit  int     `json:"isFirstDeposit"`             //是否首次存款
	DoAgencyId      int64   `json:"doAgencyId"`                 //操作者id
}

//公司入款列表返回的数据
type CompenyIncomeBackList struct {
	Id              int64   `xorm:"id" json:"id"`
	OrderNum        string  `xorm:"order_num" json:"orderNum"`               //订单号
	ClientType      string  `xorm:"client_type" json:"clientType"`           //入款来源:系统
	UpdateTime      int64   `xorm:"update_time" json:"updateTime"`           //操作时间
	LevelId         string  `xorm:"level_id" json:"levelId"`                 // 会员所属层级id
	AgencyAccount   string  `xorm:"agency_account" json:"agencyAccount"`     //经销商账号
	Account         string  `xorm:"account" json:"account"`                  //存款人账号(会员账号)
	CreateTime      int64   `xorm:"create_time" json:"createTime"`           //提交时间
	DepositTime     int64   `xorm:"deposit_time" json:"depositTime"`         // 存款时间
	DepositMoney    float64 `xorm:"deposit_money" json:"depositMoney"`       //存款金额
	DepositDiscount float64 `xorm:"deposit_discount" json:"depositDiscount"` //存款优惠
	OtherDiscount   float64 `xorm:"other_discount" json:"otherDiscount"`     //其他优惠
	DepositCount    float64 `xorm:"deposit_count" json:"depositCount"`       //存入总额
	Remark          string  `xorm:"remark" json:"remark"`                    //备注
	Payee           string  `xorm:"payee" json:"payee"`                      //入款银行卡主
	BankAccount     string  `xorm:"account" json:"bankAccount"`              //入款银行卡号
	Status          int     `xorm:"status" json:"status"`                    //状态
	IsFirstDeposit  string  `xorm:"is_first_deposit" json:"isFirstDeposit"`  //是否首次存款
	OperateName     string  `xorm:"operate_name" json:"operateName"`         //操作者名称
}

//公司入款列表返回的数据以及小计，总计
type CompenyIncomeBackLists struct {
	TotalMoney               float64                 `json:"totalMoney"`               //总计存入金额
	TotalDepositDiscount     float64                 `json:"totalDepositDiscount"`     //总计存款优惠
	TotalOtherDiscount       float64                 `json:"totalOtherDiscount"`       //总计其他优惠
	TotalDeposit             float64                 `json:"totalDeposit"`             //总计存入总金额
	PageTotalMoney           float64                 `json:"pageTotalMoney"`           //小计存入金额
	PageTotalDepositDiscount float64                 `json:"pageTotalDepositDiscount"` //小计存款优惠
	PageTotalOtherDiscount   float64                 `json:"pageTotalOtherDiscount"`   //小计其他优惠
	PageTotalDeposit         float64                 `json:"pageTotalDeposit"`         //小计存入总金额
	TotalCount               int64                   `json:"totalCount"`               //总计笔数
	PageCount                int                     `json:"pageCount"`                //小计笔数
	AllData                  []CompenyIncomeBackList //本页面的所有数据
}
