package schema

import "global"

//线上入款第三方平台列表
type OnlineIncomeThird struct {
	PayId        int64  `xorm:"id" json:"payId"`                   //接口过来的id值
	PayName      string `xorm:"title" json:"payName"`              // 第三方名称
	PayModels    string `xorm:"pay_models" json:"payModels"`       //三方模型
	DepositState int8   `xorm:"deposit_state" json:"DepositState"` //入款开关
	OutState     int8   `xorm:"out_state" json:"outState"`         //出款开关
	IpState      int8   `xorm:"ip_state" json:"ipState"`           //白名单
	PayCode      string `xorm:"pay_code" json:"payCode"`           //扫码编码网关
	BankUrl      string `xorm:"bank_url" json:"bankUrl"`           //网银支付网关
	WechatUrl    string `xorm:"wechat_url" json:"weChatUrl"`       //微信支付网关
	AliPayUrl    string `xorm:"ali_pay_url" json:"aliPayUrl"`      //支付宝支付网关
	PayStatus    int8   `xorm:"pay_status" json:"payStatus"`       //开关状态
	VisaPayUrl   string `xorm:"vias_pay_url" json:"visaPayUrl"`    //visa卡
	JdPayUrl     string `xorm:"jd_pay_url" json:"jdPayUrl"`        //京东钱包
	BdPayUrl     string `xorm:"bd_pay_url" json:"bdPayUrl"`        //百度钱包
	QqPayUrl     string `xorm:"qq_pay_url" json:"qqPayUrl"`        //qq钱包
	TenPayUrl    string `xorm:"ten_pay_url" json:"tenPayUrl"`
	CreateTime   int64  `xorm:"created"`     //本地添加时间
	DeleteTime   int64  `xorm:"delete_time"` //删除时间
}

func (*OnlineIncomeThird) TableName() string {
	return global.TablePrefix + "online_income_third"
}
