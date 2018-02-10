package back

//会员银行
type MemberBanksListBack struct {
	Id            int64  `xorm:"id" json:"id"`
	Card          string `xorm:"card" json:"card"`                     //卡号
	Title         string `xorm:"title" json:"title"`                   //银行名称
	IsDefaultBank int8   `xorm:"is_default_bank" json:"isDefaultBank"` //是否该会员默认银行卡(1.是2.不是)
}

//银行下拉框
type BankDropBack struct {
	Id     int64  `xorm:"id" json:"id"`
	Title  string `xorm:"title" json:"title"` //银行名称
	Status int8   `xorm:"status" json:"status"`
}

//银行卡详情
type MemberBankCardDetailsBack struct {
	Id            int64  `xorm:"id" json:"id"`
	BankName      string `xorm:"title" json:"bankName"`
	Card          string `xorm:"card" json:"card"`                     //卡号
	CardName      string `xorm:"card_name" json:"cardName"`            //卡账号
	CardAddress   string `xorm:"card_address" json:"cardAddress"`      //卡地址
	IsDefaultBank int8   `xorm:"is_default_bank" json:"isDefaultBank"` //是否该会员默认银行卡(1.是2.不是)
}

//添加入款订单成功
type AddBankIn struct {
	OrderNumber string `json:"orderNumber"` //订单号
}

//添加出订单成功
type AddBankOut struct {
	OrderNumber string `json:"orderNumber"` //订单号
}

//会员银行列表
type MemberBanksList struct {
	Id            int64  `xorm:"id" json:"id"`
	Card          string `xorm:"card" json:"card"`                     //卡号
	Title         string `xorm:"title" json:"title"`                   //银行名称
	IsDefaultBank int8   `xorm:"is_default_bank" json:"isDefaultBank"` //是否该会员默认银行卡(1.是2.不是)
	CardName      string `xorm:"card_name" json:"card_name"`           //出款人姓名
}
