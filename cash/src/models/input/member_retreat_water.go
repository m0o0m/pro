package input

//优惠查询列表
type ListRetreatWater struct {
	SiteId       string //操作站点id
	SiteIndexId  string `json:"siteIndexId" query:"siteIndexId"` //站点前台id
	StartTimeStr string `json:"startTime" query:"startTime"`     //开始时间
	EndTimeStr   string `json:"endTime" query:"endTime"`         //结束时间
	StartTime    int64
	EndTime      int64
	Year         string `query:"year"`               //日期（年）
	Month        string `query:"month" json:"month"` //日期(月)
}

//优惠查询明细
type DetailRetreatWater struct {
	SiteId      string //操作站点id
	SiteIndexId string //站点前台id
	Id          int64  `json:"id" query:"id" valid:"Required;ErrorCode(30041)"` //操作id
}

//优惠查询冲销
type EditRetreatWater struct {
	SiteId      string  //操作站点id
	SiteIndexId string  //站点前台id
	Id          []int64 `json:"id" query:"id" valid:"Required;ErrorCode(30041)"` //操作id组
}
