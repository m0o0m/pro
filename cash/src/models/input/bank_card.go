package input

//入款银行（开启、剔除列表）
type InComeList struct {
	SiteId      string `query:"site_id"`                                                  //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //用户前台id
	IsIncome    int8   `query:"isIncome"`                                                 //入款银行是否可用
	BankName    string `query:"bankName" valid:"MaxSize(20);ErrorCode(50100)"`            //银行名称
}

//开启、剔除
type OpenAndRejectBank struct {
	SiteId      string `json:"siteId"`                                     //用户站点ID
	SiteIndexId string `json:"siteIndexId"`                                //用户前台id
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(50085)"` //状态  1:剔除   2：开启
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(30041)"`         //银行id
}

//开启、剔除
type AgencyBankOutByDrop struct {
	SiteId      string `query:"siteId"`                                      //用户站点ID
	SiteIndexId string `query:"siteIndexId"`                                 //用户前台id
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(30041)"` //id
}

//出款银行（开启、剔除列表）
type IsOutList struct {
	SiteId      string `query:"siteId"`                                                   //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //用户前台id
	Isout       int8   `query:"isOut" `                                                   //出款银行是否可用
	BankName    string `query:"bankName"`                                                 //银行名称
}

//第三方银行（开启、剔除列表）
type IsThirdList struct {
	SiteId      string `query:"siteId"`       //用户站点ID
	SiteIndexId string `query:"siteIndexId" ` //用户前台id
	PaidType    int    `query:"paidTpye"`
}

//银行卡列表(admin)
type BankCardList struct {
	BankName string `query:"bankName"  valid:"MaxSize(20);ErrorCode(50100)"` //银行名称
}

//添加银行卡(admin)
type BankCardAdd struct {
	Title string `json:"title" valid:"Required;MaxSize(20);ErrorCode(50100)"` //银行名称
	//Icon           string `json:"icon" valid:"Required;Base64;ErrorCode(50125)"`          //银行图标 todo 前端传base64格式
	IsIncome int8 `json:"isIncome" valid:"Range(1,2);ErrorCode(50099)"` //入款银行是否可用
	IsOut    int8 `json:"isOut" valid:"Range(1,2);ErrorCode(50126)"`    //出款银行是否可用
	//PayTypeId      int64  `json:"pay_type_id" valid:"Required;Min(1);ErrorCode(60034)"`   //支付类型id
	Status         int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`           //状态
	BankWebsiteUrl string `json:"bankWebsiteUrl" valid:"MaxSize(100);ErrorCode(50127)"` //银行官网
}

//银行卡修改
type BankCardChange struct {
	Id    int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`         //银行id
	Title string `json:"title" valid:"Required;MaxSize(20);ErrorCode(50100)"` //银行名称
	//Icon           string `json:"icon" valid:"Required;Base64;ErrorCode(50125)"`          //银行图标
	IsIncome int8 `json:"isIncome" valid:"Range(1,2);ErrorCode(50099)"` //入款银行是否可用
	IsOut    int8 `json:"isOut" valid:"Range(1,2);ErrorCode(50126)"`    //出款银行是否可用
	Status   int8 `json:"status" valid:"Range(1,2);ErrorCode(30050)"`   //状态
	//PayTypeId      int64  `json:"pay_type_id" valid:"Required;Min(1);ErrorCode(60034)"`   //支付类型id
	BankWebsiteUrl string `json:"bankWebsiteUrl" valid:"MaxSize(200);ErrorCode(50127)"` //银行官网
}

//银行卡状态修改
type BankCardStatus struct {
	Id     int64 `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //银行id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"`  //状态
}

//银行卡删除
type BankCardDelete struct {
	Id int64 `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //银行id
}
