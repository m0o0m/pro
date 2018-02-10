package back

//注单列表
type BetRecordList struct {
	OrderId        string  `xorm:"order_id" json:"orderId"`               // 注单号
	GameName       string  `xorm:"game_name" json:"gameName"`             // 游戏名
	BetTimeline    int64   `xorm:"bet_timeline" json:"betTimeline"`       // 投注时间戳
	SettleTimeline int64   `xorm:"settle_timeline" json:"settleTimeline"` // 结算时间戳
	Username       string  `xorm:"username" json:"username"`              // 平台用户名
	GameResult     string  `json:"gameResult"`                            // 游戏结果编码 (彩票，视讯，电子，捕鱼，体育的返回数据)
	BetAll         float64 `xorm:"bet_all" json:"betAll"`                 // 投注额度
	BetYx          float64 `xorm:"bet_yx" json:"betYx"`                   // 有效投注
	Win            float64 `xorm:"win" json:"win"`                        // 盈利
	Result         string  `xorm:"result" json:"result"`                  // 结果
	Extra          string  `xorm:"extra" json:"extra"`                    // 附加字段
	GUsername      string  `xorm:"g_username" json:"g_username"`          // 视讯账号
	SiteId         string  `xorm:"site_id" json:"site_id"`                // 站点id
	GameType       string  `xorm:"platform" json:"platform"`              // 游戏类型，电子，视讯，捕鱼
	platform       string  `xorm:"platform" json:"platform"`              // 视讯内部的平台
	IndexId        string  `xorm:"index_id" json:"index_id"`              // indexid
	AgentId        string  `xorm:"agent_id" json:"agent_id"`              // 代理id
	UaId           string  `xorm:"ua_id" json:"ua_id"`                    // 总代id
	ShId           string  `xorm:"sh_id" json:"sh_id"`                    // 股东id
	OtherBet       string  `xorm:"other_bet" json:"other_bet"`            // 附近投注(pt专用ProgressiveBet)
	OtherWin       string  `xorm:"sh_id" json:"sh_id"`                    // 附加盈利 (pt ProgressiveWin，ag彩金)
}

//交易记录返回列表
type TransactionRecord struct {
	BetTimeline int64   `json:"bet_timeline"` //投注时间
	OrderId     string  `json:"order_id"`     //注单号
	GameResult  string  `json:"game_result"`  //游戏结果编码 (彩票，视讯，电子，捕鱼，体育的返回数据)
	BetAll      float64 `json:"bet_all"`      //投注额度
	BetYx       float64 `json:"bet_yx"`       //有效投注
	Win         float64 `json:"win"`          //盈利
}

//会员现金流水返回列表
type MemberCashRecords struct {
	CreateTime   int64   `json:"create_time"`   //交易时间
	SourceType   int8    `json:"source_type"`   //流水类型
	Balance      float64 `json:"balance"`       //交易金额
	AfterBalance float64 `json:"after_balance"` //现有余额
	DisBalance   float64 `json:"dis_balance"`   //优惠金额
	Remark       string  `json:"remark"`        //备注
	//状态
	//流水项目
}

//wap 投注记录
type WapBetRecord struct {
	OrderId     string  `json:"order_id"`    //订单号
	BetTimeline int64   `json:"betTimeline"` //投注时间
	GameResult  string  `json:"gameResult"`  //游戏结果编码 (彩票，视讯，电子，捕鱼，体育的返回数据)
	BetAll      float64 `json:"betAll"`      //投注额度
	Currency    string  `json:"currency"`    //货币类型
	BetYx       float64 `json:"betYx"`       //有效投注
	Win         float64 `json:"win"`         //盈利
	Platform    string  `json:"platform"`    //视讯内部的平台
	VType       string  `json:"vType"`       //游戏标识 如bbin,bbin_dz,bbin_fc
}

//wap 现金流水
type WapMemberCashRecord struct {
	Balance      float64 `json:"balance"`      //操作金额
	DisBalance   float64 `json:"disBalance"`   //优惠金额
	AfterBalance float64 `json:"afterBalance"` //余额
	Remark       string  `json:"remark"`       //备注
	CreateTime   int64   `json:"createTime"`   //添加时间
	Type         int8    `json:"type"`         //1  存入    2 取出
	SourceType   int8    `json:"sourceType"`   //流水项目1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11人工存入12会员返佣
}

//wap 报表统计返回
type WapReportStatisticsBacks struct {
	WapReportStatisticsInfo []WapReportStatisticsInfo `json:"dayReportStatisticsInfo"` //每天的内容详情
	WapReportStatisticsBack []WapReportStatistics     `json:"daysReportStatistics"`    //每天的总计
	WeekBetCount            int64                     `json:"weekBetCount"`            //本周注单量
	WeekBetAll              float64                   `json:"weekBetAll"`              //本周下注总额
	WeekBetYx               float64                   `json:"weekBetYx"`               //本周有效下注总额
	WeekWin                 float64                   `json:"weekWin"`                 //本周盈利金额
}

//wap 报表统计查询(每天的总计)
type WapReportStatistics struct {
	BetAll   float64 `json:"betAll"`   //下注总额
	BetYx    float64 `json:"betYx"`    //有效下注总额
	Win      float64 `json:"win"`      //盈利金额
	DateTime string  `json:"dateTime"` //投注时间
	BetCount int64   `json:"betCount"` //投注量
}

//wap 报表统计查询(每天的内容详情)
type WapReportStatisticsInfo struct {
	BetAll   float64 `json:"betAll"`   //下注总额
	BetYx    float64 `json:"betYx"`    //有效下注总额
	Win      float64 `json:"win"`      //盈利金额
	Platform string  `json:"platform"` //视讯内部的平台
	VType    string  `json:"vType"`    //游戏标识
	DateTime string  `json:"dateTime"` //投注时间
}

//wap 交易记录列表
type WapRecordList struct {
	Id             int64   ` json:"id"  xorm:"'id'"`                         // 主键id
	OrderId        string  ` json:"order_id"  xorm:"order_id"`               // 下单id
	Platform       string  ` json:"platform"  xorm:"platform"`               // 视讯内部的平台
	Account        string  ` json:"account"  xorm:"username"`                // 平台用户名
	GUsername      string  ` json:"g_username"  xorm:"g_username"`           // 视讯用户名
	Currency       string  ` json:"currency"  xorm:"currency"`               // 货币类型
	BetAll         float64 ` json:"bet_all"  xorm:"bet_all"`                 // 投注额度
	BetYx          float64 ` json:"bet_yx"  xorm:"bet_yx"`                   // 有效投注
	Win            float64 ` json:"win"  xorm:"win"`                         // 盈利
	OtherBet       float64 ` json:"other_bet"  xorm:"other_bet"`             // 附近投注(pt专用ProgressiveBet)
	OtherWin       float64 `json:"other_win" xorm:"other_win"`               //附加盈利
	BetTimeline    int64   ` json:"bet_timeline"  xorm:"bet_timeline"`       // 投注时间戳
	BetTime        string  ` json:"bet_time"  xorm:"bet_time"`               // 投注时间
	SettleTimeline int64   ` json:"settle_timeline"  xorm:"settle_timeline"` // 结算时间戳
	SettleTime     string  ` json:"settle_time"  xorm:"settle_time"`         // 结算时间
	GameType       int8    ` json:"game_type"  xorm:"game_type"`             // 类型，1视讯 2电子 3捕鱼 4彩票 5体育
	GameName       string  ` json:"game_name"  xorm:"game_name"`             // 游戏名
	GameResult     string  ` json:"game_result"  xorm:"game_result"`         // 游戏结果编码
	Extra          string  ` json:"extra"  xorm:"extra"`                     // 附加字段
	UpdateTime     int64   ` json:"update_time"  xorm:"update_time"`         // 时间戳
	Status         int8    ` json:"status"  xorm:"status"`                   // 状态
}
