package input

//ConfirmCompany 确定一条公司入款请求数据struct
type ConfirmCompany struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `json:"id" valid:"Required;Min(1);ErrorCode(60055)"`
}

//CancleIncome 取消一条公司入款请求数据struct  或者不再提醒 取消status =2  不再提醒 status =3
type CancleIncome struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `json:"id" valid:"Required;Min(1);ErrorCode(60055)"`
	Status      int64 `json:"status" valid "Required;Min(1);"ErrorCode(60055)`
}

//CompanyIncomeList 公司入款列表请求struct
type CompanyIncomeList struct {
	SiteId         string
	SiteIndexId    string  `query:"siteIndexId"`    //前台站点
	AgencyAccount  string  `query:"agencyAccount"`  //代理账号
	Level          string  `query:"level"`          //层级
	Status         int8    `query:"status"`         //状态 1已确认2已取消,3未处理
	StartTime      string  `query:"startTime"`      //开始时间
	EndTime        string  `query:"endTime"`        //结束时间
	UpperLimit     float64 `query:"upperLimit"`     //金额上限
	LowerLimit     float64 `query:"lowerLimit"`     //金额下限
	ClientType     int8    `query:"clientType"`     //入款来源(1.pc2.wap)
	IsDiscount     int8    `query:"isDiscount"`     //是否有优惠（1.是  2.否）
	PaymentAccount string  `query:"paymentAccount"` //收款账号(账号的id)
	SelectBy       int     `query:"selectBy"`       //搜索条件(1.账号2.订单号)
	Conditions     string  `query:"conditions"`     //手输入条件
}

type SiteId struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId"` //前台站点
}

//添加一条公司入款记录
type AddCompanyIncome struct {
	BankId          int64   `json:"bank_id"`          //存入银行id
	BankName        string  `json:"bank_name"`        //存入银行名称
	DepositUsername string  `json:"deposit_username"` //存款人名称
	DepositMoney    float64 `json:"deposit_money"`    //存入金额
	SetId           int64   `json:"set_id"`           //入款银行设定id
	DepositMethod   int8    `json:"deposit_method"`   //存款类型
	Remark          string  `json:"remark"`           //备注验证码
	//SiteId          string  `json:"site_id"`               //操作站点id
	//SiteIndexId     string  `json:"site_index_id"`         //站点前台id
	//MemberId        int64   `json:"member_id"`             //会员id
	//Account  		  string  `json:"account"`   			 //会员账号
	//LevelId         string  `json:"level_id"`              //会员所属层级id
	//AgencyId        int64   `json:"agency_id"`             //会员所属代理id
	//AgencyAccount   string  `json:"agency_account"`        //会员所属代理账号
	//IsFirstDeposit  int     `json:"is_first_deposit"`      //是否首次存储
	//OrderNum        string  `json:"order_num"`             //订单号
	//DepositDiscount float64 `json:"deposit_discount"`      //存入优惠
	//OtherDiscount   float64 `json:"other_discount"`        //其他优惠
	//DepositCount    float64 `json:"deposit_count"`         //存入总额
	//ClientType      int8    `json:"client_type"`           //客户端类型 1pc 2wap 3android 4ios
	//Status          int8    `json:"status"`                //状态,0未处理1已确认2已取消
	//CreateTime      int64   `json:"'create_time' created"` //创建时间
}

//一条公司入款结果
type CompanyIncomeResult struct {
	OrderNum string `json:"order_num" ` //订单号
}

//一条公司出款结果
type CompanyOutcomeResult struct {
	OrderNum string `json:"order_num" ` //订单号
}

//添加一条公司入款记录
type AddOnlineIncome struct {
	SiteId          string  `json:"site_id"`               //操作站点id
	SiteIndexId     string  `json:"site_index_id"`         //站点前台id
	MemberId        int64   `json:"member_id"`             //会员id
	Account         string  `json:"account"`               //会员账号
	BankId          int64   `json:"bank_id"`               //存入银行id
	BankName        string  `json:"bank_name"`             //存入银行名称
	LevelId         string  `json:"level_id"`              //会员所属层级id
	AgencyId        int64   `json:"agency_id"`             //会员所属代理id
	AgencyAccount   string  `json:"agency_account"`        //会员所属代理账号
	IsFirstDeposit  int     `json:"is_first_deposit"`      //是否首次存储
	OrderNum        string  `json:"order_num"`             //订单号
	DepositUsername string  `json:"deposit_username"`      //存款人名称
	DepositMoney    float64 `json:"deposit_money"`         //存入金额
	DepositDiscount float64 `json:"deposit_discount"`      //存入优惠
	OtherDiscount   float64 `json:"other_discount"`        //其他优惠
	DepositCount    float64 `json:"deposit_count"`         //存入总额
	ClientType      int8    `json:"client_type"`           //客户端类型 1pc 2wap 3android 4ios
	SetId           int64   `json:"set_id"`                //入款银行设定id
	DepositMethod   int8    `json:"deposit_method"`        //存款类型
	Status          int8    `json:"status"`                //状态,0未处理1已确认2已取消
	CreateTime      int64   `json:"'create_time' created"` //创建时间
}
