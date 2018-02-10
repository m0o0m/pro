package back

//会员报表统计
type ReportForm struct {
	Title    string  `json:"title"`     //交易平台
	VType    string  `json:"v_type"`    //视讯内部的平台名
	Num      int8    `json:"num"`       //下注笔数
	BetAll   float64 `json:"bet_all"`   //总计打码
	BetValid float64 `json:"bet_valid"` //有效打码
	Win      float64 `json:"win"`       //盈利金额
}

//会员报表统计返回
type ReportFormBack struct {
	Project  string  `json:"project"`   //项目
	Num      int8    `json:"num"`       //下注笔数
	BetAll   float64 `json:"bet_all"`   //总计打码
	BetValid float64 `json:"bet_valid"` //有效打码
	Win      float64 `json:"win"`       //盈利金额
}

//总计小计
type SiteReportTotal struct {
	Num      int64   `json:"num"`      //下注笔数
	WinNum   int64   `json:"winNum"`   //赢的笔数
	BetAll   float64 `json:"betAll"`   //总计打码
	BetValid float64 `json:"betValid"` //有效打码
	Jack     float64 `json:"jack"`     //彩金
	//Win      float64 `json:"win"`       //盈利金额
}

//站点报表统计详情
type ReportFormDetailBack struct {
	Id         int64   `xorm:"id" json:"id"`                  // 主键id
	SiteId     string  `xorm:"site_id" json:"siteId"`         // 站点ID
	Num        int64   `xorm:"num" json:"num"`                // 总笔数
	BetAll     float64 `xorm:"bet_all" json:"betAll"`         // 总计打码
	BetValid   float64 `xorm:"bet_valid" json:"betValid"`     // 有效打码
	Win        float64 `xorm:"win" json:"win"`                // 结果
	WinNum     int64   `xorm:"win_num" json:"winNum"`         // 赢的笔数
	Jack       float64 `xorm:"jack" json:"jack"`              // 彩金
	VType      string  `xorm:"v_type" json:"vType"`           // 游戏类型
	DayTime    int64   `xorm:"day_time" json:"dayTime"`       // 标示统计哪天
	CreateTime int64   `xorm:"create_time" json:"createTime"` // 添加时间
}
