package input

//入款银行设定列表
type BankInList struct {
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Status      int8   `query:"status" valid:"Range(0,2);ErrorCode(50085)"`   //状态
	BankId      int64  `query:"bankId" valid:"Min(0);ErrorCode(50101)"`       //银行id
	Account     string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"` //帐号
}

//增加一条入款银行设定
type BankInAdd struct {
	SiteId      string //用户站点ID
	SiteIndexId string
	Level       []string //会员层级
	BankId      int64    //银行id
	Account     string   //银行帐号
	OpenBank    string   //开户行
	Payee       string   //收款人
	StopBalance float64  //停用金额
	Remark      string   //备注
	Status      int      //状态
	PayTypeId   int64    // 支付类型id
	QrCode      string   // 二维码图片数据
}

//获取一条入款银行设定信息(get)
type OneBankPaySet struct {
	SiteId      string `query:"site_id"` //用户站点ID
	SiteIndexId string `query:"siteIndexId"`
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(50013)"` //id
}

//修改一条入款银行设定
type BankInUpdata struct {
	Id          int64
	SiteId      string //用户站点ID
	SiteIndexId string
	Level       []string //会员层级
	BankId      int64    //银行id
	Account     string   //银行帐号
	OpenBank    string   //开户行
	Payee       string   //收款人
	StopBalance float64  //停用金额
	Remark      string   //备注
	Status      int      //状态
	PayTypeId   int64    // 支付类型id
	QrCode      string   // 二维码图片数据
}

//修改入款银行状态
type UpdataStatus struct {
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(50013)"`         //id
	Status      int8   `json:"status" valid:"Required;Range(1,2);ErrorCode(50106)"` //状态
	SiteId      string `json:"siteId"`                                              //用户站点ID
	SiteIndexId string `json:"siteIndexId"`
}

//删除入款银行设定
type DeletePaySet struct {
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId"`
}

//存款记录
type DepositRecord struct {
	SetId       int64  `query:"setId"`       //id
	OrderNum    string `query:"orderNum"`    // 订单号
	StartTime   string `query:"startTime"`   //开始时间
	EndTime     string `query:"endTime"`     //结束时间
	SiteId      string `query:"siteId"`      //用户站点ID
	SiteIndexId string `query:"siteIndexUd"` //前台站点id
}

//适用层级
type ApplicationLevel struct {
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`
	SetId       int64  `query:"setId" valid:"Required;Min(0);ErrorCode(60025) "` //入款银行设定id
}

//传入第三方支付参数
type ThirdPayData struct {
	Amount       float64 //支付金额
	Order        string  //订单号
	PaidWay      int     //支付方式
	BusinessNum  string  //商户号
	MerchatId    int64   //商户id
	PayRedisKey  string  //保存第三方返回信息的key
	PaidCode     string  //bank code码简码
	ClientUserId int64   //客户id
	ClientName   string  //客户名
	ClientSecret string  //客户加密密码
	SiteId       string
	SiteIndexId  string
	CardMoney    float64 //卡面额
	CardNumber   string  //卡号
	CardPwd      string  //卡密
}
