package input

//ThirdList 线上支付设定列表请求参数struct
type ThirdList struct {
	SiteId      string `query:"siteId" `
	SiteIndexId string `query:"siteIndexId"`
	Status      int    `query:"status"`
	ThirdId     int    `query:"thirdId"`
}

//随机选择一条符合条件的第三方
type GetOnePaid struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(60038)"`
	PaidType    int    `query:"paidType"`
	LevelId     string `query:"level_id" valid:"Required;MinSize(1);ErrorCode(60037)"` //层级
}

//根据id选择对应的第三方
type GetOnePaidById struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(60038)"`
	PaidType    int    `query:"paidType"`
	LevelId     string `query:"level_id" valid:"Required;MinSize(1);ErrorCode(60037)"` //层级
	Id          int64  `query:"id"`                                                    //OnlinePaidSetup 设定id
}

//NewThirdPayOnline 新增加线上支付设定请求struct
type NewThirdPayOnline struct {
	SiteId            string
	LevelId           string  `json:"level" valid:"Required;MinSize(1);ErrorCode(60037)"`                  //层级
	PaidDomain        string  `json:"paidDomain" valid:"Required;ErrorCode(60026)"`                        //支付域名
	Site              string  `json:"site" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`        //站点id
	SiteIndexId       string  `json:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60038)"` //站点前台id
	BackAddress       string  `json:"backAddress" valid:"Required;MinSize(1);ErrorCode(60027)"`            //返回地址
	MerchatId         string  `json:"merchatId" valid:"Required;MinSize(1);ErrorCode(60028)"`              //商户id
	PrviateKey        string  `json:"privateKey" valid:"Required;MinSize(1);ErrorCode(60029)"`             //私钥
	PublicKey         string  `json:"publicKey" valid:"Required;MinSize(1);ErrorCode(60030)"`              //公钥
	PaidLimit         float64 `json:"paidLimit" valid:"Required;ErrorCode(60031)"`                         //当日支付限额
	PaidPlatform      int     `json:"paidPlatform" valid:"Required;Min(1);ErrorCode(60032)"`               //支付平台
	SuitableEquipment int     `json:"suitableEquipment" valid:"Required;Range(0,2);ErrorCode(60033)"`      //适用设备
	PaidType          int     `json:"paidType" valid:"Required;Min(1);ErrorCode(60034)"`                   //支付类型
	PaidCode          string  `json:"paidCode" valid:"Required;MinSize(1);ErrorCode(60035)"`               //支付编码
	Staus             int     `json:"status" valid:"Required;Range(0,2);ErrorCode(60036)"`                 //是否启用
	MerUrl            string  `json:"merUrl"`                                                              //自填写支付网关
	IsApp             int     `json:"isApp"`
}

//DelThisPaidSetup 删除该支付设定请求struct
type DelThisPaidSetup struct {
	SiteId string
	Id     int `json:"id" valid:"Required;Min(1);ErrorCode(60025)"`
}

//StopThisPaidSetup 停用该支付设定请求struct
type StopThisPaidSetup struct {
	SiteId string
	Id     int `json:"id" valid:"Required;Min(1);ErrorCode(60025)"`
	Status int `json:"status" valid:"Required;Range(1,2);ErrorCode(60036)"`
}

//ChangeThisPaidSetup 修改该支付设定的请求struct
type ChangeThisPaidSetup struct {
	SiteId            string
	Site              string  `json:"site" valid:"MaxSize(4);ErrorCode(60105)"`              //站点id
	SiteIndexId       string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(60038)"`       //站点前台id
	Id                int     `json:"id" valid:"Required;Min(1);ErrorCode(60025)"`           //支付设定Id
	MerchatId         string  `json:"merchatId" valid:"Required;ErrorCode(60028)"`           //商户号
	PayId             int     `json:"payId" valid:"Min(1);ErrorCode(60220)"`                 //三方类型
	PayType           int     `json:"payType" valid:"Min(1);ErrorCode(60034)"`               //支付类型
	IsApp             int     `json:"isApp" valid:"Range(1,2);ErrorCode(60221)"`             //是否跳转app
	MerUrl            string  `json:"merUrl" valid:"MinSize(1);ErrorCode(60222)"`            //自填写支付网关
	Status            int     `json:"status" valid:"Range(0,1);ErrorCode(60036)"`            //状态(1.启用0.停用)
	FitforLevel       string  `json:"fitforLevel" valid:"MinSize(1);ErrorCode(30113)"`       //适用层级
	PaidLimit         float64 `json:"paidLimit"`                                             //当日支付限额
	PaidDomain        string  `json:"paidDomain" valid:"MinSize(1);ErrorCode(30203)"`        //支付域名
	BackAddress       string  `json:"backAddress" valid:"MinSize(1);ErrorCode(60027)"`       //返回地址
	PrivateKey        string  `json:"privateKey" valid:"MinSize(1);ErrorCode(60029)"`        //密匙
	PublicKey         string  `json:"publicKey" valid:"MinSize(1);ErrorCode(60030)"`         //公匙
	SuitableEquipment int     `json:"suitableEquipment" valid:"Range(0,2);ErrorCode(30189)"` //适用设备
	Remark            string  `json:"remark"`                                                //备注
	PaidCode          string  `json:"paidCode" valid:"MinSize(1);ErrorCode(60035)"`          //支付code
}

//GetInfoSetupDeposit 某个支付设定下面的存款记录请求struct
type GetInfoSetupDeposit struct {
	SiteId      string
	PaidSetup   int    `query:"paidSetup" valid:"Required;Min(1);ErrorCode(60025)"` //线上支付设定id
	OrderNumber string `query:"orderNumber"`                                        //订单号
	StartTime   string `query:"startTime"`                                          //开始时间
	EndTime     string `query:"endTime"`                                            //结束时间
}

//GetInfoSetup 获取该支付设定请求struct
type GetInfoSetup struct {
	SiteId string
	Id     int `query:"id" valid:"Required;Min(1);ErrorCode(60025)"`
}

//GetBank 获取某个支付类型下面所支持的银行卡请求struct
type GetBank struct {
	SiteId   string
	PaidType int `query:"paidTpye" valid:"Required;Min(1);ErrorCode(60041)"`
}

//获取单个的线上支付设定解析用的struct
type OnlinePaidSetParse struct {
	MerId      string `json:"merId"`      //id
	MerchantId string `json:"merchantId"` //商户号
	PayId      int    `json:"payId"`      //三方配置id
	PayType    int    `json:"payType"`    //支付方式
	PrivateKey string `json:"privateKey"` //私钥
	PublicKey  string `json:"publicKey"`  //公钥
	NotifyUrl  string `json:"notifyUrl"`  //回调地址
	PayStatus  int    `json:"payStatus"`  //状态
	PayCode    string `json:"payCode"`    //支付编码
	MerUrl     string `json:"merUrl"`     //支付网关
	IsApp      int    `json:"isApp"`      //是否跳转(0不跳转，1.跳转)
	LevelId    string `json:"levelId"`    //支持层级
}

//新增加支付设定之后线上的返回
type AddSetupParse struct {
	AgentLine     string `json:"agentLine"`    //代理
	SubAgenctLine string `json:"subAgentLine"` //子代理线
	PayId         int    `json:"payId"`        //配置id
	PayType       int    `json:"payType"`      //支付方式
	MerchantId    string `json:"merchantId"`   //商户id
	PrivateKey    string `json:"privateKey"`   //私钥
	PublicKey     string `json:"publicKey"`    //公钥
	NotifyUrls    string `json:"notifyUrl"`    //返回地址
	LevelId       string `json:"levelId"`      //层级
	Code          string `json:"code"`         //支付编码
	MerUrl        string `json:"merUrl"`       //自填写的网关
}
