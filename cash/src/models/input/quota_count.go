package input

//额度统计列表
type QuotaCountList struct {
	SiteId      string `query:"siteId"`      //用户站点ID
	SiteIndexId string `query:"siteIndexId"` //用户前台id
	Account     string `query:"account" `    //帐号
	StartTime   string `query:"startTime"`   //开始时间
	EndTime     string `query:"endTime"`     //结束时间
}

//额度充值记录
type QuotaRecord struct {
	SiteId      string `query:"siteId"`      //用户站点ID
	SiteIndexId string `query:"siteIndexId"` //用户前台id
	Status      int8   `query:"status"`      //状态
	Type        int8   `query:"type"`        //类型
	StartTime   string `query:"startTime"`   //开始时间
	EndTime     string `query:"endTime"`     //结束时间
}

//修改充值记录
type SitePayRecordUpdate struct {
	Id        int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` // id
	AdminUser string // 提交者
	State     int8   `json:"state" valid:"Required;Range(1,2);ErrorCode(10081)"` // 状态1未支付2支付
	Remark    string `json:"remark"`                                             // 备注
}

//额度记录列表
type QuotaRecordList struct {
	SiteId      string `query:"siteId"`                                     //用户站点ID
	SiteIndexId string `query:"siteIndexId"`                                //用户前台id
	AdminName   string `query:"adminName"`                                  //会员帐号
	VdType      int64  `query:"vdType"`                                     //视讯类型
	CashType    int8   `query:"cashType"`                                   //交易类型
	DoType      int8   `query:"doType" valid:"Range(0,2);ErrorCode(60087)"` //操作类型 1存入 2 取出
	State       int8   `query:"state" valid:"Range(0,3);ErrorCode(30050)"`  //状态1待审核 2正常  3掉单
	StartTime   string `query:"startTime"`                                  //开始时间
	EndTime     string `query:"endTime"`                                    //结束时间
}

//银行卡额度充值
type BankCardRecharge struct {
	SiteId    string `json:"siteId" `                                               //站点
	OrderNum  string `json:"orderNum" `                                             //订单号
	AdminUser string `json:"adminUser" `                                            //操作人
	Type      int8   `json:"type" `                                                 //支付方式  //1 三方入款 2 公司入款
	Money     int64  `json:"money" valid:"Required;Min(10);ErrorCode(60228)"`       //预缴充值金额
	PayCard   string `json:"payCard" valid:"Required;MaxSize(19);ErrorCode(60200)"` //银行账号
	PayId     int8   `json:"payId"`                                                 //收款银行
}
