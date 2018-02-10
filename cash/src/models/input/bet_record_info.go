package input

//注单列表
type BetRecordList struct {
	Platform    string `query:"platform"`      // 视讯内部的平台
	GameType    int8   `query:"game_type"`     // 游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育
	SiteId      string `query:"site_id"`       // 站点
	SiteIndexId string `query:"site_index_id"` // 站点前台
	Pankou      string `query:"pankou"`        // 盘口
	Sort        int    `query:"sort"`          // 排序方式 1desc从大到小 ,2asc从小到大
	SortType    int    `query:"sort_type"`     // 排序类型 1,账号2,注单号3,注单时间4,结算时间65,下注金额6,有效投注
	PageSize    int    `query:"page_size"`     // 每页显示
	Page        int    `query:"page"`          // 页码
	Check       int    `query:"check"`         // 结算方式
	Account     string `query:"account"`       // 平台用户名
	OrderId     string `query:"order_id"`      // 主单号
	GameName    string `query:"game_name"`     // 游戏名
	StartTime   string `query:"start_time"`    // 查询时间,开始
	EndTime     string `query:"end_time"`      // 查询时间,结束
	Qishu       string `query:"qishu"`         // 期数
}

//会员交易记录列表
type TransactionRecord struct {
	OrderId    string `query:"order_id"`                     //注单号
	BetTime    string `query:"bet_time"`                     //投注时间
	GameType   int8   `query:"game_type" valid:"Range(0,5)"` //游戏类型:1视讯  2电子 3捕鱼 4彩票 5体育
	PlatformId int64  `query:"platform_id"`                  //交易平台id
	MemberId   int64  //会员id
}

//会员现金流水
type MemberCashRecords struct {
	SourceType int8   `query:"source_type" valid:"Range(1,4);ErrorCode(30227)"` //流水项目1.存款 2.取款 3.额度转换  4.其他
	CreateTime string `query:"create_time"`                                     //交易时间
	MemberId   int64  //会员id
}

//wap 投注记录
type WapBetRecord struct {
	GameType int8   `query:"gameType" valid:"Range(1,5);ErrorCode(30226)"` //游戏类型，1视讯  2电子 3捕鱼 4彩票 5体育
	DateTime string `query:"dateTime"`                                     //时间
	MemberId int64  //会员id
}

//wap 现金流水(TODO:下注和派彩分离出去了，这里没做)
type WapMemberCashRecords struct {
	SourceType int8   `query:"sourceType" valid:"Range(1,4);ErrorCode(30227)"` //流水项目1.存款 2.取款 3.额度转换  4.其他
	DateTime   string `query:"dateTime"`                                       //时间
	MemberId   int64  //会员id
}

//wap会员交易记录列表
type RecordInfoList struct {
	VType       int64  `query:"vType" `        //游戏类型:1视讯  2电子 3捕鱼 4彩票 5体育
	MemberId    int64  `query:"member_id"`     //会员id
	GameName    string `query:"game_name"`     //具体游戏名称
	StartTime   int64  `query:"start_time"`    //开始时间
	EndTime     int64  `query:"end_time"`      //结束时间
	OrderNum    int64  `query:"order_num"`     //注单号
	GameOneType string `query:"game_one_type"` //彩票游戏下详细分类
	GameResult  string `query:"game_result"`   //查询体育串式 单式
	PageSize    int    `query:"pageSize"`      // 每页显示
	Page        int    `query:"page"`          // 页码
}
