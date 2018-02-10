package back

//前台返回的数据struct
type OnlineBackIncome struct {
	PayId        int    `json:"payId"`        //编号
	PayName      string `json:"payName"`      //三方名称
	PayModels    string `json:"payModels"`    //三方模型
	DepositState int    `json:"DepositState"` //入款开关
	OutState     int    `json:"outState"`     //出款开关
	IpState      int    `json:"ipState"`      //白名单
	PayCode      string `json:"payCode"`      //扫码编码
	BankUrl      string `json:"bankUrl"`      //网银支付网关
	WechatUrl    string `json:"weChatUrl"`    //微信支付网关
	AliPayUrl    string `json:"aliPayUrl"`    //支付宝支付开关
	QqPayUrl     string `json:"qqPayUrl"`     //qq钱包
	TenPayUrl    string `json:"tenPayUrl"`    //
	VisaPayUrl   string `json:"visaPayUrl"`   //visa卡
	JdPayUrl     string `json:"jdPayUrl"`     //京东钱包
	BdPayUrl     string `json:"bdPayUrl"`     //百度钱包
	PayStatus    int    `json:"payStatus"`    //开关状态
}

//返回下拉框
type BackSelectThird struct {
	PayId   int    `xorm:"id" json:"payId"`      //编号
	PayName string `xorm:"title" json:"payName"` //三方名称
}

type BackSelectThirdJson struct {
	PayId   int    `json:"payId"`   //编号
	PayName string `json:"payName"` //三方名称
}
