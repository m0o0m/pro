package back

//第三方或者银行卡的列表
type SitePayNameList struct {
	ID           int64  `json:"id"  xorm:"id"`                      //
	PayName      string `json:"payName"  xorm:"pay_name" `          // 第三方名字,银行卡名字
	PayType      int64  `json:"payType"  xorm:"pay_type" `          //第三方支付类型
	State        int64  `json:"state"  xorm:"state" `               //状态1停用2启用
	PayId        string `json:"payId"  xorm:"pay_id" `              //商户ID,如果type类型是银行卡的话,就是银行卡号
	PayKey       string `json:"payKey"  xorm:"pay_key" `            //密钥
	FUrl         string `json:"fUrl"  xorm:"f_url" `                //支付域名
	Vircarddoin  string `json:"vircarddoin"  xorm:"vircarddoin" `   //国付宝终端号
	TerminalId   string `json:"terminalId"  xorm:"terminal_id" `    //宝付终端标号
	Type         int64  `json:"type"  xorm:"type" `                 //1第三方,2银行卡
	MyName       string `json:"myName"  xorm:"my_name" `            //收款人姓名
	Address      string `json:"address"  xorm:"address" `           //银行卡开户行地址
	Lid          int64  `json:"lid"  xorm:"lid" `                   //默认未分层,站点套餐id
	LevelName    string `xorm:"level_name" json:"levelName"`        //站点层级名称
	PaidTypeName string `xorm:"paid_type_name" json:"paidTypeName"` //第三方支付类型
}

//站点层级下拉框
type SiteLevelDropBack struct {
	Id        int64  `xorm:"id" json:"id"`                //主键id
	LevelName string `xorm:"level_name" json:"levelName"` //层级名字
}

//支付类型下拉框
type ThirdTypeDropBack struct {
	Id           int    `xorm:"id" json:"typeId"`               //主键id
	PaidTypeName string `xorm:"paid_type_name" json:"typeName"` //支付类型名称
}
