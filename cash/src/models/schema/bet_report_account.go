package schema

import "global"

//统计表
type BetReportAccount struct {
	Id          int64   `xorm:"'id' PK autoincr"` // 主键id
	SiteId      string  `xorm:"site_id"`          // 站点
	SiteIndexId string  `xorm:"site_index_id"`    // 站点前台
	GameType    int8    `xorm:"game_type"`        // 游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育
	Account     string  `xorm:"account"`          // 平台用户名,会员账号
	AgencyId    int64   `xorm:"agency_id"`        // 代理id
	UaId        int64   `xorm:"ua_id"`            // 总代id
	ShId        int64   `xorm:"sh_id"`            // 股东id
	Num         int     `xorm:"num"`              // 下注笔数
	BetAll      float64 `xorm:"bet_all"`          // 投注额度
	BetValid    float64 `xorm:"bet_valid"`        // 有效投注
	Win         float64 `xorm:"win"`              // 盈利
	WinNum      int     `xorm:"win_num"`          // 赢的笔数
	Jack        float64 `xorm:"jack"`             // 彩金
	DayTime     int64   `xorm:"day_time"`         // 标示统计哪天 时间戳
	Platform    string  `xorm:"platform"`         // 视讯内部的平台
	CreateTime  int64   `xorm:"create_time"`      // 统计时间 时间戳
	VType       string  `xorm:"v_type"`           // 游戏标识 如bbin,bbin_dz,bbin_fc
}

func (*BetReportAccount) TableName() string {
	return global.TablePrefix + "bet_report_account"
}
