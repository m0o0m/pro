package input

//添加银行卡或者第三方 带.的是银行卡要填的,带=的是第三方要填的
type SitePayNameAdd struct {
	Id          int64  `json:"id"`                                                 //id有就更新,没有就添加
	PayName     string `json:"payName" valid:"Required;ErrorCode(30201)"`          //第三方名字,银行卡名字 . =
	PayType     int64  `json:"payType" `                                           //第三方支付类型 =
	State       int64  `json:"state" valid:"Required;Range(1,2);ErrorCode(10081)"` //状态1停用2启用.
	PayId       string `json:"payId" valid:"Required;ErrorCode(30200)"`            //商户ID,如果type类型是银行卡的话,就是银行卡号 . =
	PayKey      string `json:"payKey" `                                            //密钥 =
	FUrl        string `json:"fUrl" `                                              //支付域名 =
	Vircarddoin string `json:"vircarddoin" `                                       //国付宝终端号 =
	TerminalId  string `json:"terminalId" `                                        //宝付终端标号 =
	Type        int64  `json:"type" valid:"Required;Range(1,2);ErrorCode(30206)"`  //1第三方,2银行卡
	MyName      string `json:"myName" `                                            //收款人姓名 .
	Address     string `json:"address" `                                           //银行卡开户行地址 .
	Lid         int64  `json:"lid" `                                               //默认未分层,站点套餐id . =
}

//查询列表
type SitePayNameList struct {
	Type  int64 `query:"type" valid:"Required;Range(1,2);ErrorCode(30206)"` //1第三方,2银行卡
	State int64 `query:"state" valid:"Range(0,2);ErrorCode(10081)"`         //状态1停用2启用.
}
