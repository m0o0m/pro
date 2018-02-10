package back

//打码信息
type BetReport struct {
	Id          int     `json:"id" xorm:"id PK autoincr"`
	SiteId      string  `json:"site_id" xorm:"site_id"`
	SiteIndexId string  `json:"site_index_id" xorm:"site_index_id"`
	Account     string  `json:"account" xorm:"account"`
	AgencyId    int     `json:"agency_id" xorm:"agency_id"`     //代理ID',
	UaId        int     `json:"ua_id" xorm:"ua_id"`             //总代ID',
	ShId        int     `json:"sh_id" xorm:"sh_id"`             //股东ID',
	Num         int     `json:"num" xorm:"num"`                 //下注笔数',
	BetAll      float64 `json:"bet_all" xorm:"bet_all"`         //总计打码',
	BetValid    float64 `json:"bet_valid" xorm:"bet_valid"`     //有效打码',
	Win         float64 `json:"win" xorm:"win"`                 //结果,盈利',
	WinNum      int     `json:"win_num" xorm:"win_num"`         //赢的笔数',
	Jack        float64 `json:"jack" xorm:"jack"`               //彩金',
	DayTime     int     `json:"day_time" xorm:"day_time"`       //标示统计哪天 时间戳',
	Platform    string  `json:"platform" xorm:"platform"`       //视讯内部的平台',
	GameType    int     `json:"game_type" xorm:"game_type"`     //游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育',
	CreateTime  int     `json:"create_time" xorm:"create_time"` //统计时间 时间戳',
}

//本周打码报表总计
type WeekBetReportCount struct {
	Num      int     `json:"num" xorm:"num"`            //下注笔数',
	BetAll   float64 `json:"betAll" xorm:"bet_all"`     //总计打码',
	BetValid float64 `json:"betValid" xorm:"bet_valid"` //有效打码',
	Win      float64 `json:"win" xorm:"win"`            //结果,盈利',
	WinNum   int     `json:"winNum" xorm:"win_num"`     //赢的笔数',
	DayTime  string  `json:"dayTime" xorm:"day_time"`   //标示统计哪天 时间戳'
}

//本周打码报表
type ThisWeekBetReport struct {
	Num      int     `json:"num" xorm:"num"`            //下注笔数',
	BetAll   float64 `json:"betAll" xorm:"bet_all"`     //总计打码',
	BetValid float64 `json:"betValid" xorm:"bet_valid"` //有效打码',
	Win      float64 `json:"win" xorm:"win"`            //结果,盈利',
	WinNum   int     `json:"winNum" xorm:"win_num"`     //赢的笔数',
	DayTime  string  `json:"dayTime" xorm:"day_time"`   //标示统计哪天 时间戳',
	VType    string  `json:"vType" xorm:"v_type"`       //视讯内部的平台'
}
