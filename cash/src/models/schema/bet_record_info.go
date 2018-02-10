package schema

import "global"

//注单表
type BetRecordInfo struct {
	Id             int64   `xorm:"'id' PK autoincr"` // 主键id
	OrderId        string  `xorm:"order_id"`         // 订单号
	Platform       string  `xorm:"platform"`         // 视讯内部的平台
	Account        string  `xorm:"username"`         // 平台用户名
	GUsername      string  `xorm:"g_username"`       // 视讯用户名
	Currency       string  `xorm:"currency"`         // 货币类型
	SiteId         string  `xorm:"site_id"`          // 站点
	SiteIndexId    string  `xorm:"index_id"`         // 站点前台
	AgencyId       int64   `xorm:"agent_id"`         // 代理id
	UaId           int64   `xorm:"ua_id"`            // 总代id
	ShId           int64   `xorm:"sh_id"`            // 股东id
	BetAll         float64 `xorm:"bet_all"`          // 投注额度
	BetYx          float64 `xorm:"bet_yx"`           // 有效投注
	Win            float64 `xorm:"win"`              // 盈利
	OtherBet       float64 `xorm:"other_bet"`        // 附近投注(pt专用ProgressiveBet)
	OtherWin       float64 `xorm:"other_win"`        // 附加盈利 (pt ProgressiveWin，ag彩金)
	BetTimeline    int64   `xorm:"bet_timeline"`     // 投注时间戳
	BetTime        int64   `xorm:"bet_time"`         // 投注时间
	SettleTimeline int64   `xorm:"settle_timeline"`  // 结算时间戳
	SettleTime     int64   `xorm:"settle_time"`      // 结算时间
	GameType       int8    `xorm:"game_type"`        // 游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育
	GameName       string  `xorm:"game_name"`        // 游戏名
	GameResult     string  `xorm:"game_result"`      // 游戏结果编码
	Extra          string  `xorm:"extra"`            // 附加字段
	UpdateTime     int64   `xorm:"update_time"`      // 时间戳
	Status         int8    `xorm:"status"`           // 状态
}

func (*BetRecordInfo) TableName() string {
	return global.TablePrefix + "bet_record_info"
}
