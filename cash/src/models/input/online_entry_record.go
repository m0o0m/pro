package input

//CancleDeposit 取消一条线上入款的请求struct
type CancleDeposit struct {
	SiteId string
	Id     int64 `json:"id" valid:"Required;Min(1);ErrorCode(60019)"`
}

//ConfirmDeposit 确定一条线上入款的请求struct
type ConfirmDeposit struct {
	SiteId      string
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(60026)"` //前台站点Id
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(60019)"`
}

//拒绝，取消出款原因
type OutRemark struct {
	SiteId string
	Id     int64 `query:"id" valid:"Required;Min(1);ErrorCode(60019)"`
}

//OnlineDepositList 线上入款列表的请求参数struct
type OnlineDepositList struct {
	SiteId          string
	SiteIndexId     string  `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(60026)"`   //前台站点Id
	AgencyId        int64   `query:"agencyId"`                                          //代理账号id
	Level           string  `query:"level"`                                             //层级
	Status          int8    `query:"status" valid:"Range(0,4);ErrorCode(60036)"`        //状态(1.未支付2.已经支付3.已取消4已确认)
	SourceDeposit   int8    `query:"sourceDeposit" valid:"Range(0,2);ErrorCode(60040)"` //入款来源(1.pc 2.wap)
	ThirdId         string  `query:"thirdId"`                                           //入款商户(第三方支付平台id)
	IsDiscount      int8    `query:"isDiscount"`                                        //是否有优惠（1.是  2.否）
	SelectBy        int8    `query:"selectBy"`                                          //根据什么进行查询(1.账号 2.订单号)
	Conditions      string  `query:"conditions"`                                        //搜索条件
	StartTime       string  `query:"startTime"`                                         //开始时间
	EndTime         string  `query:"endTime"`                                           //结束时间
	UpperLimitMoney float64 `query:"upperLimitMoney"`                                   //金额上限
	LowerLimitMoney float64 `query:"lowerLimitMoney"`                                   //金额下限
}
