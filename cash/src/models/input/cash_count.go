package input

//优惠统计

//定时入款统计开关
type TimingCashCount struct {
	Open int64 `json:"open" valid:"Required;Range(1,2);ErrorCode(71024)"` //1开2关
}

// 入款-列表
type CashCountList struct {
	SiteId        string `query:"siteId" `                                                 //站点id
	Account       string `query:"account" `                                                //会员账号
	AgencyAccount string `query:"agencyAccount" `                                          //代理账号
	IntoStyle     int64  `query:"intoStyle" `                                              //入款方式
	STime         string `query:"startTime" valid:"Required;MinSize(10);ErrorCode(30155)"` //开始时间
	ETime         string `query:"endTime"  valid:"Required;MaxSize(10);ErrorCode(30156)"`  //结束时间
}

//入款重新统计
type CashRecount struct {
	STime string `json:"startTime" valid:"Required;MinSize(10);ErrorCode(30155)"` //开始时间
	ETime string `json:"endTime" valid:"Required;MaxSize(10);ErrorCode(30156)"`   //结束时间
}

//出入账目汇总查询条件
type CashCollect struct {
	SiteId string `query:"siteId" `   //站点id
	STime  string `query:"startTime"` //开始时间
	ETime  string `query:"endTime"`   //结束时间
}

//导出excel
type CashCollectExport struct {
	Key string `json:"key" valid:"Required;Length(37);ErrorCode()"` //云端缓存的数据
}
