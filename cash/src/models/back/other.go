package back

//导航栏商品返回列表
type OtherProduct struct {
	TypeId      int64    `json:"type_id"` //商品类型id
	ProductName []string `json:"product_name"`
}

type Product struct {
	TypeId      int64  `json:"type_id"`      //商品类型id
	ProductName string `json:"product_name"` //商品名
}

//商品分类表;
type ProductVideo struct {
	Id          int64  `json:"id"`           //主键id
	ProductName string `json:"product_name"` //商品名
	VType       string `json:"v_type"`       //游戏类型
}

//支付类型数据;
type PaySetData struct {
	OnlineDepositMax float64 `json:"onlineDepositMax"` // 线上入款单次最高存款金額
	OnlineDepositMin float64 `json:"onlineDepositMin"` // 线上入款单次最低存款金額
	LineDepositMax   float64 `json:"lineDepositMax"`   // 公司入款单次最高存款金額
	LineDepositMin   float64 `json:"lineDepositMin"`   // 公司入款单次最低存款金額
}

//快速支付 支付类型数据;
type FastPaySetData struct {
	OnlineDepositMax float64 `json:"onlineDepositMax"` // 线上入款单次最高存款金額
	OnlineDepositMin float64 `json:"onlineDepositMin"` // 线上入款单次最低存款金額
}

//公司入款数据
type GetCompanyData struct {
	PaidType        map[string]int //支付类型
	OnlineIncomeSet []GetPayeeInfo //收款账号列表
	SiteIncomeBank  []SiteBank     //入款银行
	PaySet          PaySetData     //支付设定
}

//存款渲染数据
type GetIncomeData struct {
	Account        string              //会员账号
	SiteIncomeBank []SiteBank          //入款银行
	PaySet         PaySetData          //支付设定
	Income         map[int]*IncomeData //存款数据
}

//快速充值存款渲染数据
type GetFastIncomeData struct {
	Account string                  //会员账号
	PaySet  FastPaySetData          //支付设定
	Income  map[int]*FastIncomeData //存款数据
}

//线上、公司入款数据
type IncomeData struct {
	PaidType      int                //支付类型
	PaidName      string             //支付名
	OnlineIncome  []OnlineIncomeData //线上入款数据
	CompanyIncome []GetPayeeInfo     //公司入款数据
}

//快速充值入款数据
type FastIncomeData struct {
	PaidType     int                //支付类型
	OnlineIncome []OnlineIncomeData //线上入款数据
}

//线上入款数据
type OnlineIncomeData struct {
	Id   int64 `json:"id"`   //第三方支付设置id组[]PaidSetupId
	Sort int64 `json:"sort"` //排序
}

//获取所有的支付类型和对应的支付设定id
type PaidTypeAndPaySetBack struct {
	Id           int    `json:"id"`          //支付类型id
	PaidTypeName string `json:"paid_name"`   //支付类型名称
	TypeStatus   int8   `json:"type_status"` //状态
	SetId        int64  `json:"set_id"`      //支付设定
	Sort         int64  `json:"sort"`        //排序
}

//获取所有的支付类型和对应的支付设定id
type FastIncome struct {
	PaidType int   `json:"paid_type"` //支付类型id
	SetId    int64 `json:"set_id"`    //支付设定
	Sort     int64 `json:"sort"`      //排序
}

//获取网银在线的银行
type GetOnlineIncomeBank struct {
	PayId    int64  `json:"pay_id"`    //三方配置id
	BankName string `json:"bank_name"` //银行名称
	BankCode string `json:"bank_code"` //银行简码
}
