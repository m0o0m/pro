package back

//支付设定列表返回
type OnlinePaidSetupList struct {
	Id            int64   `xorm:"id" json:"id"`                    //主键id
	ThirdName     string  `xorm:"title" json:"title"`              //第三方平台
	MerchatId     string  `json:"merchatId"`                       //商户id
	PaidLimit     float64 `json:"paidLimit"`                       //当日支付限额
	SuitableLevel string  `xorm:"fitfor_level" json:"fitforLevel"` //适用层级
	Status        int     `json:"status"`                          //状态
	PaidType      string  `xorm:"type_name" json:"typeName"`       //支付类型
	Remark        string  `json:"remark"`                          //备注
}

//随机取一条第三方
type GetOneThird struct {
	Id            int64   `xorm:"id" json:"id"`                       //支付设定id
	PaidPlatform  int     `xorm:"paid_platform" json:"paid_platform"` //第三方平台id
	ThirdName     string  `xorm:"title" json:"title"`                 //第三方平台
	MerchatId     string  `json:"merchat_id"`                         //商户id
	PaidLimit     float64 `json:"paid_limit"`                         //当日支付限额
	SuitableLevel string  `xorm:"fitfor_level" json:"fitfor_level"`   //适用层级
	Status        int     `json:"status"`                             //状态
	PaidType      string  `xorm:"type_name" json:"type_name"`         //支付类型
	Remark        string  `json:"remark"`                             //备注
	PaidCode      string  `json:"paid_code"`                          //支付码
}

//某个支付设定下面的存款记录
type OnePaidSetupBack struct {
	Id               int64   `xorm:"id" json:"id"`
	ThirdOrderNumber string  `xorm:"third_order_number" json:"third_order_number"` //第三方平台订单号`
	AmountDeposit    float64 `xorm:"amount_deposit" json:"amount_deposit"`         //存款金额
	ThirdPayTime     int64   `xorm:"third_pay_time" json:"third_pay_time" `        //第三方平台转账时间(存款时间)
	Remark           string  `xorm:"remark" json:"remark"`                         //备注
	Account          string  `xorm:"member_account" json:"account"`                //会员账号
}

//获取所有的支付类型的返回
type PaidTypeBack struct {
	Id           int    `json:"id"`
	PaidTypeName string `json:"paid_name"` //支付类型名称
}

//获取所有的支付类型和状态
type PaidTypeAndStatusBack struct {
	Id           int    `json:"id"`
	PaidTypeName string `json:"paid_name"`   //支付类型名称
	TypeStatus   int8   `json:"type_status"` //状态
}

//获取拒绝，取消出款原因
type OutRemark struct {
	OutRemark string `json:"outRemark"` //原因
}
