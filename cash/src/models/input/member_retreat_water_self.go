package input

//自助返水查询列表
type ListRetreatWaterSelf struct {
	SiteId       string //操作站点id
	SiteIndexId  string `json:"site_index_id" query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	StartTimeStr string `json:"start_time" query:"start_time"`                                           //开始时间
	EndTimeStr   string `json:"end_time" query:"end_time"`                                               //结束时间
	StartTime    int64  //开始时间
	EndTime      int64  //结束时间
	Account      string `json:"account" query:"account"`     //会员账号
	OrderNum     int64  `json:"order_num" query:"order_num"` //订单号
}

//自助返水查询明细
type DetailRetreatWaterSelf struct {
	SiteId      string //操作站点id
	SiteIndexId string //站点前台id
	Id          int64  `json:"id" query:"id" valid:"Required;ErrorCode(30041)"` //操作id
}

//wap查看单个的反水
type SingleReWater struct {
	Id        int64  `query:"id" valid:"Required;ErrorCode(30041)"`        //会员每日打码统计表id
	Vtype     string `query:"vType" valid:"Required;ErrorCode(60217)"`     //类型
	ProductId int64  `query:"productId" valid:"Required;ErrorCode(60218)"` //产品id
	DateTimes string `query:"dateTimes" valid:"Required;ErrorCode(60216)"` //日期
}

//wap一键查看所有的反水
type OneClickSeeReWater struct {
	NowDate string `query:"nowdate" valid:"Required;ErrorCode(60216)"` //今天的年月日
}

//一键领取所有的反水
type OneClickGetAllReWater struct {
	MemberId     int64   `json:"memberId" valid:"Required;ErrorCode()"`     //会员id
	SiteId       string  `json:"siteId" valid:"Required;ErrorCode()"`       //站点id
	SiteIndexId  string  `json:"siteIndexId" valid:"Required;ErrorCode()"`  //站点前台id
	BetValid     float64 `json:"betValid" valid:"Required;ErrorCode()"`     //有效投注额度
	RewaterTotal float64 `json:"rewaterTotal" valid:"Required;ErrorCode()"` //反水总额度
}
