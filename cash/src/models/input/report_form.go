package input

//会员报表统计
type ReportForm struct {
	SiteId      string //站点id
	SiteIndexId string //站点前台id
	StartTime   string `query:"start_time" valid:"Required;ErrorCode(30155)"` //开始时间
	EndTime     string `query:"end_time" valid:"Required;ErrorCode(30156)"`   //结束时间
	MemberId    int64  //会员id
}

//统计数据查询
type GetDataCenterList struct {
	SiteId        string `query:"siteId"`       // 站点id
	VType         string `query:"vType"`        // 游戏标识 如 bbin,bbin_dz,bbin_fc
	TimeZone      int8   `json:"timeZone"`      // 时区 1.非转时区 2.转时区
	Account       string `json:"account"`       // 会员账号
	AgencyAccount string `json:"agencyAccount"` // 代理账号
	SortBy        int    `query:"sortType"`     // 排序类型 1 总笔数、2 赢笔数、3 总投注、4 有效投注、5 派彩
	PageSize      int    `query:"pageSize"`     // 每页显示
	Page          int    `query:"page"`         // 页码
	StartTime     string `query:"startTime"`    //开始时间
	EndTime       string `query:"endTime"`      //结束时间
}

//wap报表统计
type WapReport struct {
	StartTime string `query:"startTime" valid:"Required;ErrorCode(30155)"` //开始时间
	EndTime   string `query:"endTime" valid:"Required;ErrorCode(30156)"`   //结束时间
	MemberId  int64  //会员id
}
