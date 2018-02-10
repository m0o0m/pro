package input

//优惠统计

//定时统计开关
type TimingCount struct {
	Open int64 `json:"open" valid:"Required;Range(1,2);ErrorCode(71024)"` //1开2关
}

//优惠统计-列表
type DiscountCountList struct {
	SiteId        string `query:"siteId" `                                                 //站点id
	Account       string `query:"account" `                                                //会员账号
	AgencyAccount string `query:"agencyAccount" `                                          //代理账号
	STime         string `query:"startTime" valid:"Required;MinSize(10);ErrorCode(30155)"` //开始时间
	ETime         string `query:"endTime" valid:"Required;MinSize(10);ErrorCode(30156)"`   //结束时间
}

//重新统计
type Recount struct {
	STime string `json:"startTime" valid:"Required;MinSize(10);ErrorCode(30155)"` //开始时间
	ETime string `json:"endTime" valid:"Required;MinSize(10);ErrorCode(30156)"`   //结束时间
}
