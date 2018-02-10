package back

//返回会员余额统计
type MemberClassifyBalance struct {
	ProductName string  ` xorm:"product_name" json:"product_name"  ` //是否启用会员注册
	StrMoney    float64 ` xorm:"str_money" json:"str_money" `        //是否需要邮箱
	EndMoney    float64 ` xorm:"end_money" json:"end_money" `        //是否需要微信
	UpdateTime  int64   ` xorm:"update_time" json:"update_time" `    //是否需要身份证
}

//返回会员统计
type MemberClassifyArr struct {
	Account     string  ` xorm:"account" json:"account"  `          //是否启用会员注册
	Mbalance    float64 ` xorm:"mbalance" json:"mbalance" `         //是否需要邮箱
	Status      int8    ` xorm:"status" json:"status" `             //是否需要微信
	Balance     float64 ` xorm:"balance" json:"balance" `           //是否需要身份证
	ProductName string  ` xorm:"product_name" json:"product_name" ` //是否需要身份证
}

//额度转换记录列表
type MemberBalanceConversionList struct {
	Id         int64   `json:"id"`       //记录id
	Account    string  `json:"account"`  //会员账号
	Platform   string  `json:"platform"` //交易平台
	ForType    int64   `json:"for_type"`
	Money      float64 `json:"money"`       //金额
	CreateTime int64   `json:"create_time"` //交易时间
	UpdateTime int64   `json:"update_time"` //确认时间
	Status     int8    `json:"status"`      //状态
	Remark     string  `json:"remark"`      //备注
}

//额度转换记录返回列表
type MemberBalanceConversionBackList struct {
	Id         int64   `json:"id"`         //记录id
	Account    string  `json:"account"`    //会员账号
	Platform   string  `json:"platform"`   //交易平台
	Money      float64 `json:"money"`      //金额
	CreateTime int64   `json:"createTime"` //交易时间
	UpdateTime int64   `json:"updateTime"` //确认时间
	Status     int8    `json:"status"`     //状态
	Remark     string  `json:"remark"`     //备注
	Type       string  `json:"type"`       //类型
}

type Proportion struct {
	Proportion float64 `json:"proportion"`
}

//wap 会员中心--会员余额刷新
type MemberBalance struct {
	AccountBalance float64 `json:"account_balance"` //账号余额
	GameBalance    float64 `json:"gameBalance"`     //视讯总余额
}

//余额
type Balance struct {
	Balance float64 `json:"balance"` //金额
}

//额度转换-获取各平台余额
type MemberPlatformBalance struct {
	Realname               string                   `json:"realname"`               //会员名称
	Account                string                   `json:"account"`                //会员账号
	AccountBalance         float64                  `json:"accountBalance"`         //账号余额
	GameBalance            float64                  `json:"gameBalance"`            //视讯总余额
	ProductClassifyBalance []ProductClassifyBalance `json:"productClassifyBalance"` //会员对应各分类下余额
}

//额度转换一键回归
type BalanceFlyback struct {
	MemberPlatformBalance
	SuccessNum int64 `json:"success_num"` //成功回归
	FailureNum int64 `json:"failure_num"` //失败回归
}

//会员对应各分类下余额
type ProductClassifyBalance struct {
	PlatformId int64   `xorm:"platform_id"  json:"platform_id"` //游戏平台id
	Platform   string  `xorm:"platform"  json:"platform"`       //游戏平台名称
	Balance    float64 `xorm:"balance"  json:"balance"`         //余额
}

//额度转换-单个平台余额刷新
type PlatformBalanceRefresh struct {
	AccountBalance         float64                  `json:"accountBalance"`         //账号余额
	GameBalance            float64                  `json:"gameBalance"`            //视讯总余额
	ProductClassifyBalance []ProductClassifyBalance `json:"productClassifyBalance"` //会员对应各分类下余额
}
