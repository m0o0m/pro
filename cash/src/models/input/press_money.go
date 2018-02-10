package input

//催款账单查询
type PressMoneyList struct {
	SiteId    string `query:"siteId"`    //站点
	StartTime string `query:"startTime"` //起始期数
	EndTime   string `query:"endTime"`   //结束期数
	Status    int8   `query:"status"`    //状态
}

//催款账单添加
type PressMoneyAdd struct {
	SiteId     []string `json:"siteId" valid:"Required;MaxSize(255);ErrorCode(60105)"`     //可传多个站点
	PayName    string   `json:"payName" valid:"Required;MaxSize(20);ErrorCode(20011)"`     //收款人姓名
	PayAddress string   `json:"payAddress" valid:"Required;MaxSize(100);ErrorCode(30062)"` //收款银行地址
	Bank       string   `json:"bank" valid:"Required;MaxSize(19);ErrorCode(50100)"`        //收款银行
	PayCard    string   `json:"payCard" valid:"Required;MaxSize(19);ErrorCode(60200)"`     //银行账号
	Remark     string   `json:"remark" valid:"MaxSize(50);ErrorCode(20019)"`               //备注

}

//获取单条详情
type PressMoneyOne struct {
	Id int64 `query:"id"` //序号id
}

//催单修改
type PressMoneyEdit struct {
	Id         int64  `json:"id" valid:"Min(1);ErrorCode(30041)"`                        //序号
	SiteId     string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`       //站点
	PayName    string `json:"payName" valid:"Required;MaxSize(20);ErrorCode(20011)"`     //收款人姓名
	PayAddress string `json:"payAddress" valid:"Required;MaxSize(100);ErrorCode(30062)"` //收款银行地址
	Bank       string `json:"bank" valid:"Required;MaxSize(19);ErrorCode(50100)"`        //收款银行
	PayCard    string `json:"payCard" valid:"Required;MaxSize(19);ErrorCode(60200)"`     //银行账号
	Remark     string `json:"remark" valid:"MaxSize(50);ErrorCode(20019)"`               //备注
}

//更新催单状态
type PressMoneyStatus struct {
	Id     int64 `json:"id" valid:"Min(1);ErrorCode(30041)"`         //序号id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //业主是否提交
}

//预缴款
type PrePayment struct {
	SiteId      string `json:"siteId" `                                                   //站点
	SiteIndexId string `json:"siteIndexId" `                                              //站点前台
	OrderNum    string `json:"orderNum" `                                                 //订单号
	AdminUser   string `json:"adminUser" `                                                //操作人
	Money       int64  `json:"money" valid:"Required;Min(10);ErrorCode(60228)"`           //预缴充值金额
	PayName     string `json:"payName" valid:"Required;MaxSize(20);ErrorCode(20011)"`     //收款人姓名
	PayAddress  string `json:"payAddress" valid:"Required;MaxSize(100);ErrorCode(30062)"` //开户银行地址
	Bank        string `json:"bank" valid:"Required;MaxSize(19);ErrorCode(50100)"`        //缴款银行类型
	PayCard     string `json:"payCard" valid:"Required;MaxSize(19);ErrorCode(60200)"`     //银行账号
	PayId       int8   `json:"payId"`                                                     //银行账号
}
