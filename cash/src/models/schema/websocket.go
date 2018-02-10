package schema

//监听传送过来的数据结构体
type Listening struct {
	SiteId      string //站点id
	SiteIndexId string //站点前台id
	Types       int    //(1.公司入款2.线上入款3.出款)
}

//推送数据的结构体
//公司入款结构
type CompanyIncomeStruct struct {
	Id                int64   `xorm:"id" json:"id"`                             //主键id
	OrderNumber       string  `xorm:"order_num" json:"order_number"`            //订单号
	System            int     `xorm:"client_type" json:"system"`                //系统
	OperateTime       int64   `xorm:"create_time" json:"operate_time"`          //操作时间
	LevelName         string  `xorm:"level_name" json:"level_name"`             //层级
	Agency            string  `xorm:"agency_name" json:"agency"`                //经销商
	MemberAccount     string  `xorm:"user_name" json:"member_account"`          //会员账号
	Deposit           float64 `xorm:"deposit_money" json:"deposit"`             //存入金额
	DepositDiscount   float64 `xorm:"deposit_discount" json:"deposit_discount"` //存款优惠
	OtherDiscount     float64 `xorm:"other_discount" json:"other_discount"`     //其他优惠
	DepositTotal      float64 `xorm:"deposit_count" json:"deposit_total"`       //存入总额
	MemberBankAccount string  `xorm:"bank_name" json:"member_bank_account"`     //会员银行账号名称
	DepositMehtod     int     //存款方式
	DepositBank       string  //存入银行账户
	Status            int     `xorm:"status" json:"status"`                     //状态
	IsFirstDeposit    int     `xorm:"is_first_deposit" json:"is_first_deposit"` //是否首次存储
	OperateName       string  `xorm:"operate_name"`                             //操作者名称
}

//线上入款
type OnlineIncome struct {
	Id             int64   `xorm:"id" json:"id"`                             //主键id
	OrderNumber    string  `xorm:"third_order_number" json:"order_number"`   //订单号
	LocalCreat     int64   `xorm:"create_time" json:"local_creat"`           //系统
	OperateTime    int64   `xorm:"third_pay_time" `                          //操作时间
	LevelName      string  `xorm:"level" json:"level_name"`                  //会员所属层级
	MemberAccount  string  `xorm:"member_account" json:"member_account"`     //会员账号
	Deposite       float64 `xorm:"amount_deposit" json:"deposite"`           //存入金额
	Status         int     `xorm:"status" json:"status"`                     //状态
	IsFirstDeposit int     `xorm:"is_first_deposit" json:"is_first_deposit"` //是否首次存储
	DepositMethod  string  `xorm:"title" json:"deposit_method"`              //入款方式
	OperateName    string  `xorm:"operate" json:"operate_name"`              //操作者
}

//出款管理结构体

type MoneyManagementList struct {
	Id                    int64   `xorm:"id"`
	Site                  string  `xorm:"site_id" json:"site"`                            //站点
	SiteIndexId           string  `xorm:"site_index_id" json:"site_index_id"`             //站点前台id
	LevelName             string  `xorm:"level_id" json:"level_name"`                     //会员所属层级
	Agency                string  `xorm:"agency_name" json:"agency"`                      //经销商名称
	MemberAccount         string  `xorm:"user_name" json:"member_account"`                //会员账号
	IsFirstOut            int     `xorm:"is_first" json:"is_first_out"`                   //是否首次出款(1,是2.不是)
	OutMoney              float64 `xorm:"outward_num" json:"out_money"`                   //提现额度
	Charge                float64 `xorm:"charge" json:"charge"`                           //手续费
	PreferentialTreatment float64 `xorm:"favourable_money" json:"preferential_treatment"` //优惠金额
	AdministrativeFee     float64 `xorm:"expenese_money" json:"administrative_fee"`       //行政费
	RealyOut              float64 `xorm:"outward_money" json:"realy_out"`                 //实际出款金额
	AfterOperateBalance   float64 `xorm:"balance" json:"after_operate_balance"`           //账户余额(操作后)
	DedutExpenses         float64 `xorm:"favourable_out" json:"dedut_expenses"`           //是否有优惠扣除
	OutDate               int64   `xorm:"out_time" json:"out_date"`                       //出款日期
	IsAudit               int     //是否稽核
	Status                int     `xorm:"out_status"` //状态
	AutomaticOut          int     //自动出款
	IsGiveup              int     //是否下发
	IsThirdGiveup         int     //三方下发状态
	OperateName           string  `xorm:"do_agency_account" json:"operate_name"` //操作者
	Remark                string  `xorm:"out_remark" json:"remark"`              //备注
}
