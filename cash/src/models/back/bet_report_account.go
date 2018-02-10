package back

//统计数据总返回
type BetReportAccountAllBack struct {
	BetReportAccount      []BetReportAccount    `json:"betReportAccount"`
	BetReportAccountTotal BetReportAccountTotal `json:"betReportAccountTotal"`
}

//统计数据
type BetReportAccount struct {
	Id          int64   `xorm:"id" json:"id"`
	SiteId      string  `xorm:"site_id" json:"siteId"`            // 站点
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"` // 前台
	AgencyId    int64   `xorm:"agency_id" json:"agencyId"`        // 代理id
	MemberId    int64   `xorm:"member_id" json:"uid"`             // 用户id
	Account     string  `xorm:"account" json:"account"`           // 平台用户名
	Win         float64 `xorm:"win" json:"result"`                // 结果
	Num         int64   `xorm:"num" json:"num"`                   // 总笔数
	WinNum      int64   `xorm:"win_num" json:"winNum"`            // 赢笔数
	BetAll      float64 `xorm:"bet_all" json:"betAll"`            // 投注额度
	BetValid    float64 `xorm:"bet_valid" json:"betValid"`        // 有效投注
	DayType     int64   `xorm:"day_time" json:"dayType"`          // 统计时间 时间戳
	Jack        float64 `xorm:"jack" json:"jack"`                 // 投注额度
	Pc          float64 `xorm:"pc" json:"pc"`                     // 有效投注
}

//统计数据总计、小计
type BetReportAccountTotal struct {
	WinTotal      float64 `xorm:"SUM(win)" json:"winTotal"`            // 结果(总计)
	NumTotal      int64   `xorm:"SUM(num)" json:"numTotal"`            // 总笔数(总计)
	WinNumTotal   int64   `xorm:"SUM(win_num)" json:"winNumTotal"`     // 赢笔数(总计)
	BetAllTotal   float64 `xorm:"SUM(bet_all)" json:"betAllTotal"`     // 投注额度(总计)
	BetValidTotal float64 `xorm:"SUM(bet_valid)" json:"betValidTotal"` // 有效投注(总计)
	JackTotal     float64 `xorm:"SUM(jack)" json:"jackTotal"`          // 投注额度(总计)
	PcTotal       float64 `xorm:"-" json:"pcTotal"`                    // 派彩(总计)
	SmallWin      float64 `xorm:"-" json:"smallWin"`                   // 结果(小计)
	SmallNum      int64   `xorm:"-" json:"smallNum"`                   // 总笔数(小计)
	SmallWinNum   int64   `xorm:"-" json:"smallWinNum"`                // 赢笔数(小计)
	SmallBetAll   float64 `xorm:"-" json:"smallBetAll"`                // 投注额度(小计)
	SmallBetValid float64 `xorm:"-" json:"smallBetValid"`              // 有效投注(小计)
	SmallJack     float64 `xorm:"-" json:"smallJack"`                  // 投注额度(小计)
	SmallPc       float64 `xorm:"-" json:"smallPc"`                    // 派彩(小计)
}

//报表数据列表
type ReportList struct {
	SiteId      string  `xorm:"site_id" json:"siteId"`            // 站点
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"` // 前台
	AgencyId    int64   `xorm:"agency_id" json:"agencyId"`        // 代理id
	UaId        int64   `xorm:"ua_id" json:"uaId"`                // 总代id
	ShId        int64   `xorm:"sh_id" json:"shId"`                // 股东id
	Win         float64 `xorm:"win" json:"win"`                   // 盈利
	Num         int64   `xorm:"num" json:"num"`                   // 总笔数
	WinNum      int64   `xorm:"win_num" json:"winNum"`            // 赢笔数
	BetAll      float64 `xorm:"bet_all" json:"betAll"`            // 投注额度
	BetValid    float64 `xorm:"bet_valid" json:"betValid"`        // 有效投注
	Payout      float64 `xorm:"payout" json:"payout"`             // 总派彩
	Jack        float64 `xorm:"jack" json:"jack"`                 // 彩金
	Account     string  `xorm:"account" json:"account"`           // 账号
	GameType    int64   `xorm:"game_type" json:"gameType"`        // 游戏类型
	Platform    string  `xorm:"platform" json:"platform"`         // 视讯平台名
	VType       string  `xorm:"v_type" json:"vType"`              // 商品v_type
	ProductName string  `xorm:"product_name" json:"productName"`  // 商品名
	Name        string  `xorm:"name" json:"name"`                 // 商品名
}

//点击金额加载数据
type ReportClick struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` // 前台
	AgencyId    int64   `xorm:"agency_id" json:"agency_id"`         // 代理id
	UaId        int64   `xorm:"ua_id" json:"ua_id"`                 // 总代id
	ShId        int64   `xorm:"sh_id" json:"sh_id"`                 // 股东id
	Select      string  `xorm:"select" json:"select"`               // 报表层级(all:站点sh:股东ua:总代理at:代理)
	VType       string  `xorm:"v_type" json:"v_type"`               // 商品v_type
	Win         float64 `xorm:"win" json:"win"`                     // 盈利
	Num         int64   `xorm:"num" json:"num"`                     // 总笔数
	WinNum      int64   `xorm:"win_num" json:"win_num"`             // 赢笔数
	BetAll      float64 `xorm:"bet_all" json:"bet_all"`             // 投注额度
	BetValid    float64 `xorm:"bet_valid" json:"bet_valid"`         // 有效投注
	Payout      float64 `xorm:"payout" json:"payout"`               // 总派彩
	Jack        float64 `xorm:"jack" json:"jack"`                   // 彩金
	Account     string  `xorm:"account" json:"account"`             // 平台用户名
	Name        string  `xorm:"name" json:"name"`                   // 名称
}

//账单查询
type ReportBills struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` // 前台
	Win         float64 `xorm:"win" json:"win"`                     // 盈利
	GameType    int64   `xorm:"game_type" json:"game_type"`         // 游戏类型
	PlatformId  int64   `xorm:"platform_id" json:"platform_id"`     // 视讯平台id
	ProductName string  `xorm:"product_name" json:"product_name"`   // 商品名
	ProductId   int64   `xorm:"product_id" json:"product_id"`       // 商品Id
}

//账单查询返回数据
type ReportBillsRetrun struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` // 前台
	SiteName    string  `xorm:"site_name" json:"site_name"`         //站点名称
	ComboName   string  `xorm:"combo_name" json:"combo_name"`       //套餐id
	Win         float64 `xorm:"win" json:"win"`                     // 盈利
	GameType    int64   `xorm:"game_type" json:"game_type"`         // 游戏类型
	PlatformId  int64   `xorm:"platform_id" json:"platform_id"`     // 视讯平台id
	ProductName string  `xorm:"product_name" json:"product_name"`   // 商品名
	ProductId   int64   `xorm:"product_id" json:"product_id"`       // 商品Id
	Proportion  float64 `xorm:"proportion" json:"proportion"`       // 条件比例
}

//报表查询excel数据
type ReportInfo struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` // 前台
	SiteName    string `xorm:"site_name" json:"site_name"`         // 站点名称
	ComboName   string `xorm:"combo_name" json:"combo_name"`       // 套餐名
	List        []ReportInfoList

	Html string `json:"html"` // excel表格模板
}

//报表查询excel数据
type ReportInfoList struct {
	Win         float64 `xorm:"win" json:"win"`                   // 盈利
	VType       string  `xorm:"v_type" json:"v_type"`             // 游戏类型
	ProductName string  `xorm:"product_name" json:"product_name"` // 商品名
	ProductId   int64   `xorm:"product_id" json:"product_id"`     // 商品Id
}

//报表查询excel数据
type ReportExport struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` // 前台
	SiteName    string `xorm:"site_name" json:"site_name"`         // 站点名称
	ComboName   string `xorm:"combo_name" json:"combo_name"`       // 套餐名
	List        []ReportExportList

	Html string `xorm:"html" json:"html"` // 账单html
}

//报表查询excel详细数据
type ReportExportList struct {
	Win         float64 `xorm:"win" json:"win"`                   // 盈利
	VType       string  `xorm:"v_type" json:"v_type"`             // 游戏类型
	ProductName string  `xorm:"product_name" json:"product_name"` // 商品名
	ProductId   int64   `xorm:"product_id" json:"product_id"`     // 商品Id
	Proportion  float64 `xorm:"proportion" json:"proportion"`     // 占成比

	Negative float64 `xorm:"negative" json:"negative"` // 报表上月负数
	Ought    float64 `xorm:"ought" json:"ought"`       // 报表应收金额
}

//优惠统计
type CountBetReportAccount struct {
	Id          int64   `json:"id"`           //id
	Account     string  `json:"account"`      //会员账号
	MemberId    int64   `json:"member_id"`    //所属会员id
	LevelId     string  `json:"level_id"`     //层级
	Betall      float64 `json:"betall"`       //有效总投注
	AllMoney    float64 `json:"all_money"`    //总返水金额
	SelfMoney   float64 `json:"self_money"`   //自助返水
	RebateWater float64 `json:"rebate_water"` //返水小计
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //商品投注额
	Money       float64 `json:"money"`        //商品返水额
}

//优惠统计会员有效打码量
type BetValidBetReportAccountList struct {
	Account     string  `json:"account"`      //会员账号
	MemberId    int64   `json:"member_id"`    //所属会员id
	LevelId     string  `json:"level_id"`     //层级
	BetValid    float64 `json:"bet_valid"`    //有效总投注
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品分类名
}

//优惠统计商品
type CountBetReportAccountProduct struct {
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //商品投注额
	Rate        float64 `json:"rate"`         //比例
	Money       float64 `json:"money"`        //商品返水额
}

//优惠统计小计
type CountBetReportAccountSubtotal struct {
	Account     string                         `json:"account"`      //会员账号
	MemberId    int64                          `json:"member_id"`    //所属会员id
	LevelId     string                         `json:"level_id"`     //层级
	Betall      float64                        `json:"betall"`       //有效总投注
	AllMoney    float64                        `json:"all_money"`    //总返水金额
	SelfMoney   float64                        `json:"self_money"`   //自助返水
	RebateWater float64                        `json:"rebate_water"` //返水小计
	Products    []CountBetReportAccountProduct `json:"products"`     //各个商品投注额返水额的集合
}

//优惠统计小计map
type CountBetReportAccountSubtotalMap struct {
	Account     string                                   `json:"account"`      //会员账号
	MemberId    int64                                    `json:"member_id"`    //所属会员id
	LevelId     string                                   `json:"level_id"`     //层级
	Betall      float64                                  `json:"betall"`       //有效总投注
	AllMoney    float64                                  `json:"all_money"`    //总返水金额
	SelfMoney   float64                                  `json:"self_money"`   //自助返水
	RebateWater float64                                  `json:"rebate_water"` //返水小计
	Products    map[string]*CountBetReportAccountProduct `json:"products"`     //各个商品投注额返水额的集合
}

//优惠统计总计 已退水或未退水或已冲销
type CountBetReportAccountTotalMap struct {
	PeopleNum int64                                        `json:"people_num"` //退水总人数
	Money     float64                                      `json:"money"`      //退水总金额
	BetTotal  float64                                      `json:"bet_total"`  //退水有效总投注
	Subtotal  map[string]*CountBetReportAccountSubtotalMap `json:"subtotal"`   //每个会员对应一条返水小计记录
}

//优惠统计总计 已退水或未退水或已冲销
type CountBetReportAccountTotal struct {
	PeopleNum int64                           `json:"people_num"` //退水总人数
	Money     float64                         `json:"money"`      //退水总金额
	BetTotal  float64                         `json:"bet_total"`  //退水有效总投注
	Subtotal  []CountBetReportAccountSubtotal `json:"subtotal"`   //每个会员对应一条返水小计记录
}

//优惠统计总计Map 已退水+未退水
type CountAllBetReportAccountTotalMap struct {
	StartTime    int64                         `json:"start_time"`     //开始时间
	EndTime      int64                         `json:"end_time"`       //结束时间
	PeopleNumAll int64                         `json:"people_num_all"` //退水总人数
	MoneyAll     float64                       `json:"money_all"`      //退水总金额
	BetTotalAll  float64                       `json:"bet_total_all"`  //退水有效总投注
	HadRetreat   CountBetReportAccountTotalMap `json:"had_retreat"`    //已退水列表
	UnRetreat    CountBetReportAccountTotalMap `json:"un_retreat"`     //未退水列表
}

//优惠统计查询明细id
type SearchIdBetReportAccount struct {
	Id       int64  `json:"id"`        //id
	Account  string `json:"account"`   //会员账号
	MemberId int64  `json:"member_id"` //所属会员id
}

//账单数据
type Bills struct {
	ReportData string `xorm:"report_data" json:"report_data"` // 帐单数据
	StartTime  int64  `xorm:"report_data" json:"start_time"`  // 开始时间
	EndTime    int64  `xorm:"report_data" json:"end_time"`    // 结束时间
	Qishu      string `xorm:"qishu" json:"qishu"`             // 期数
}
