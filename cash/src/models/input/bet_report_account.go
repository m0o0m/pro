package input

//注单列表
type BetReportAccount struct {
	SiteId      string `query:"siteId"`      // 站点
	SiteIndexId string `query:"siteIndexId"` // 站点前台
	VType       string `query:"vType"`       // 游戏标识 如bbin,bbin_dz,bbin_fc
	TimeZone    int8   `query:"timeZone"`    // 时区
	UserAccount string `query:"userAccount"` // 平台用户名
	Account     string `query:"account"`     // 平台用户名
	StartTime   string `query:"startTime"`   // 查询时间,开始
	EndTime     string `query:"endTime"`     // 查询时间,结束
	PageSize    int    `query:"pageSize"`    // 每页显示
	Page        int    `query:"page"`        // 页码
	Sort        int    `query:"sort"`        // 排序方式 1desc从大到小 ,2asc从小到大
	SortBy      int    `query:"sortType"`    // 排序类型 1,账号2,总笔数3,赢笔数4,投注额度5,有效投注6,统计时间
}

//优惠统计-存入
type StoreBetReportAccount struct {
	SiteId      string
	SiteIndexId string
	MemberIds   []int64 `json:"member_ids"`                                                     //会员账号数组
	Event       string  `json:"event" valid:"Required;MinSize(1);MaxSize(30);ErrorCode(70016)"` //事件名称
	Bet         float64 `json:"bet" valid:"Required;ErrorCode(70017)"`                          //综合打码倍数
	Key         string  `json:"key" valid:"Required;ErrorCode(70018)"`                          //存在redis里的统计数据
}

//优惠统计
type CountBetReportAccount struct {
	SiteId       string
	SiteIndexId  string
	StartTimeStr string   `query:"start_time" json:"start_time" valid:"Required;MinSize(10);MaxSize(10);ErrorCode(30155)"` //开始时间
	EndTimeStr   string   `query:"end_time" json:"end_time" valid:"Required;MinSize(10);MaxSize(10);ErrorCode(30156)"`     //结束时间
	StartTime    int64    //开始时间
	EndTime      int64    //结束时间
	System       int8     `query:"system" json:"system" valid:"Required;Range(1,2)"` //1全部 层级查会员 2会员 直接查询
	LevelId      []string `query:"level_id" json:"level_id"`                         //会员层级
	Account      []string `query:"account" json:"account"`                           //会员账号
	ReturnWater  int8     `query:"return_water" json:"return_water"`                 //1有优惠
}

type ReportList struct {
	SiteId      string   `json:"siteId"`      // 站点
	SiteIndexId string   `json:"siteIndexId"` // 站点
	AgencyId    int64    `json:"agencyId"`    //代理ID
	Rtype       int8     `json:"rtype"`       //报表类型	1：总报表，2：代理报表，3：会员报表
	Username    string   `json:"username"`    //报表查询账号
	VType       []string `json:"vType"`       //商品v_type
	StartTime   string   `json:"startTime"`   // 查询时间,开始
	EndTime     string   `json:"endTime"`     // 查询时间,结束
	TimeZone    int8     `json:"timeZone"`    // 时区 1.非转时区 2.转时区
}

type RepSearch struct {
	SiteId      string `query:"siteId"`      // 站点
	SiteIndexId string `query:"siteIndexId"` // 站点
}

type ReportClick struct {
	SiteId      string `query:"site_id"`       // 站点
	SiteIndexId string `query:"site_index_id"` // 站点
	AgencyId    int64  `query:"agency_id"`     // 代理id
	UaId        int64  `query:"ua_id"`         // 总代id
	ShId        int64  `query:"sh_id"`         // 股东id
	Select      string `query:"select"`        // 报表层级(all:站点sh:股东ua:总代理at:代理)
	VType       string `query:"v_type"`        // 商品v_type
	StartTime   string `query:"start_time"`    // 查询时间,开始
	EndTime     string `query:"end_time"`      // 查询时间,结束
}

//查询账单列表
type BillList struct {
	SiteId string `query:"siteId"`
	Status int8   `query:"status"`
	Year   string `query:"year" valid:"MaxSize(4);ErrorCode(60203)"`  // 年份
	Qishu  string `query:"qishu" valid:"MaxSize(2);ErrorCode(60204)"` // 期数
}

//查询账单列表
type BillListBatch struct {
	Id string `json:"id" valid:"MinSize(1);ErrorCode(30041)"` //批量下发id
}

//查询账单数据
type ReportBills struct {
	SiteId    []string `json:"siteId"`
	VType     []string `json:"vType"`
	ComboId   int64    `json:"comboId"`   // 套餐id
	StartTime string   `json:"startTime"` // 查询时间,开始
	EndTime   string   `json:"endTime"`   // 查询时间,结束
	TimeZone  int8     `json:"timeZone"`  // 时区 1.非转时区 2.转时区
}

//账单导出
type ReportExports struct {
	SiteId    string `json:"site_id"`
	TableBtml string `json:"table_html"` //文件内容
	Filename  string `json:"filename"`   // 文件名
	StartTime string `json:"start_time"` // 查询时间,开始
	EndTime   string `json:"end_time"`   // 查询时间,结束
}

//账单添加
type BillsAdd struct {
	Year       string `json:"year" valid:"Required;Length(4);ErrorCode(60203)"`  // 年份
	Qishu      string `json:"qishu" valid:"Required;Length(2);ErrorCode(60204)"` // 期数
	SiteId     string `json:"site_id"`                                           // 站点id
	ReportData string `json:"report_data"`                                       //文件内容
}

//账单修改
type SiteBillUpdate struct {
	Id         int64  `json:"id"`          // 账单id
	Qishu      string `json:"qishu"`       // 期数
	SiteId     string `json:"site_id"`     // 站点id
	ReportData string `json:"report_data"` // 文件内容
	Status     int8   `json:"status"`      // 推送状态 1已下发 2未下发 3删除
}
