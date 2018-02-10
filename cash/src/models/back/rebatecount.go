package back

//退佣统计
type RebateCountList struct {
	Id          int64  `xorm:"id" json:"id"`                       //期数id
	SiteId      string `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	Title       string `xorm:"title" json:"title"`                 //期数名称
}
type RebateCountList1 struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	AgencyId    int64  `xorm:"agency_id" json:"agency_id" `        //代理
}
type Rebatec struct {
	Id            int64  `json:"id"`             //设定id
	SiteId        string `json:"site_id"`        //站点id
	SiteIndexId   string `json:"site_index_id"`  //站点前台id
	EffectiveUser int64  `json:"effective_user"` //有效会员数量
	ValidMoney    int64  `json:"valid_money"`    //有效会员投注金额
}
type Charge struct {
	SiteId             string  `xorm:"site_id" json:"site_id"`                            //操作站点id
	SiteIndexId        string  `xorm:"site_index_id" json:"site_index_id"`                //站点前台id
	IncomPoundageRatio float64 `xorm:"income_poundage_ratio" json:"incom_poundage_ratio"` //入款手续费比例
	IncomPoundageUp    int     `xorm:"income_poundage_up" json:"income_poundage_up"`      //入款手续费上限
	OutPoundageRatio   float64 `xorm:"out_poundage_ratio" json:"out_poundage_ratio"`      //出款手续费比例
	OutPoundageUp      int     `xorm:"out_poundage_up" json:"out_poundage_up"`            //出款手续费上限
	IsDeliveryModel    int8    `xorm:"is_delivery_model" json:"is_delivery_model"`        //是否开启交收模式
}
type AgencyInfo struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	AgencyId    int    `xorm:"id" json:"agency_id"`                //代理id
	UserName    string `xorm:"username" json:"username"`           //代理名称
	Account     string `xorm:"account" json:"account"`             //代理账号
}
type CashRecord struct {
	Id          int64   `xorm:"id" json:"id"`
	SiteId      string  `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	MemberId    int64   `xorm:"member_id" json:"member_id"`         //会员id
	UserName    string  `xorm:"user_name" json:"user_name"`         //会员账号
	AgencyId    int64   `xorm:"agency_id" json:"agency_id"`         //会员所属代理id
	SourceType  int     `xorm:"source_type" json:"source_type"`     //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	Balance     float64 `xorm:"balance" json:"balance"`             //金额
	DisBalance  float64 `xorm:"dis_balance" json:"dis_balance"`     //优惠金额
}

type ReportAll struct {
	Win      float64 `xorm:"win" json:"win"`             // 盈利
	Num      int64   `xorm:"num" json:"num"`             // 总笔数
	WinNum   int64   `xorm:"win_num" json:"win_num"`     // 赢笔数
	BetAll   float64 `xorm:"bet_all" json:"bet_all"`     // 投注额度
	BetValid float64 `xorm:"bet_valid" json:"bet_valid"` // 有效打码
	Jack     float64 `xorm:"jack" json:"jack"`           // 彩金
	DayTime  int64   `xorm:"day_time" json:"day_time"`   //统计时间
	GameType int64   `xorm:"game_type" json:"game_type"` // 游戏类型
	Platform string  `xorm:"platform" json:"platform"`   // 视讯平台名
	VType    string  `xorm:"v_type" json:"v_type"`       // 商品v_type
	AgencyId int64   `xorm:"agency_id" json:"agency_id"` //代理id
}

type LastRebate struct {
	AgencyId         int64   `json:"agency_id"`          //代理id
	SiteId           string  `json:"site_id"`            //操作站点id
	SiteIndexId      string  `json:"site_index_id"`      //站点前台id
	Statue           int8    `json:"statue"`             //是否已退佣 0未操作1已操作
	EffectiveMember  int64   `json:"effective_member"`   //有效会员数
	BeforeJack       float64 `json:"before_jack"`        //前期奖池
	BeforeCost       float64 `json:"before_cost"`        //前期费用
	BeforeBetting    float64 `json:"before_betting"`     //前期有效投注
	BeforeProfit     float64 `json:"before_profit"`      //前期赢利
	NowPayoffElect   float64 `json:"now_payoff_elect"`   //前期电子赢利
	NowPayoffVideo   float64 `json:"now_payoff_video"`   //前期视讯赢利
	NowPayoffLottery float64 `json:"now_payoff_lottery"` //前期彩票赢利
	NowPayoffSport   float64 `json:"now_payoff_sport"`   //前期体育赢利
	NowPayoffFish    float64 `json:"now_payoff_fish"`    //前期捕鱼赢利
}
type AllFee struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	AgencyId    int64   `xorm:"agency_id" json:"agency_id"`         //会员所属代理id
	SourceType  int     ` json:"source_type"`                       //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8取消出款
	Balance     float64 ` json:"balance"`                           //金额
	DisBalance  float64 ` json:"dis_balance"`                       //优惠金额
}
type ReportData struct {
	Num              int64   `json:"num"`                //下注笔数
	ValidMember      int64   `json:"valid_member"`       //有效会员
	WinNum           int64   `json:"win_num"`            //赢的笔数
	Win              float64 `json:"win"`                //盈利
	Jack             float64 `json:"jack"`               //彩金
	BetAll           float64 `json:"bet_all"`            //总计打码
	BetValid         float64 `json:"bet_valid"`          //有效打码
	TotalBack        float64 `json:"total_back"`         //退佣总计
	TotalWater       float64 `json:"total_water"`        //退水总计
	IsGrand          int64   `json:"is_grand"`           //是否累计标识
	Status           int64   `json:"status"`             //退佣数据累计
	TotalCost        float64 `json:"total_cost"`         //总计退佣退水费用
	NowPayoffElect   float64 `json:"now_payoff_elect"`   //当期 电子 盈利
	NowPayoffVideo   float64 `json:"now_payoff_video"`   //当期 视讯 盈利
	NowPayoffLottery float64 `json:"now_payoff_lottery"` //当期 彩票 盈利
	NowPayoffSport   float64 `json:"now_payoff_sport"`   //当期 体育 盈利
	NowPayoffFish    float64 `json:"now_payoff_fish"`    //当期 捕鱼 盈利
	SiteId           string  `json:"site_id"`            //操作站点id
	SiteIndexId      string  `json:"site_index_id"`      //站点前台id
	PeriodsId        int     `json:"periods_id"`         //期数id
	AgencyId         int64   `json:"agency_id"`          //代理id
	AgencyAccount    string  `json:"agency_account"`     //代理账号
	BeforeJack       float64 `json:"before_jack"`        // 前期奖池
	NowJack          float64 `json:"now_jack"`           //当期奖池
	BeforeProfit     float64 `json:"before_profit"`      //前期赢利
	BeforeBetting    float64 `json:"before_betting"`     //前期赢利
	BeforeCost       float64 `json:"before_cost"`        //前期费用
	NowProfit        float64 `json:"now_profit"`         //当期赢利
	NowBetting       float64 `json:"now_betting"`        //当期有效投注
	NowCost          float64 `json:"now_cost"`           //当期费用
	EffectiveMember  int64   `json:"effective_member"`   //有效会员数(含前期)
	Rebate           float64 `json:"rebate"`             //本次退佣金额
	RebateWater      float64 `json:"rebate_water"`       //本次退水金额
	Remark           string  `json:"remark"`             //备注
	Balance          float64 `json:"balance"`            //代理余额
	BankInFee        float64 `json:"bank_in_fee"`        //入款费用
	BankOutFee       float64 `json:"bank_out_fee"`       //出款费用
	CatmInFee        float64 `json:"catm_in_fee"`        //手动入款费用
	OtherFee         float64 `json:"other_fee"`          //其它优惠
	DiscountFee      float64 `json:"discount_fee"`       //返水费用
}
type RebateProduct struct {
	ProductId   int64   `json:"product_id"`   //分类id
	RebateRadio float64 `json:"rebare_radio"` //退佣比例
	WaterRadio  float64 `json:"water_radio"`  //退水比例
}
type PoundageRateList struct {
	VideoRate   float64 `json:"video_rate"`   //视讯退佣比例
	ElectRate   float64 `json:"elect_rate"`   //电子退佣比例
	SportRate   float64 `json:"sport_rate"`   //体育退佣比例
	LotteryRate float64 `json:"lottery_rate"` //彩票退佣比例
	FishRate    float64 `json:"fish_rate"`    //捕鱼退佣比例
}
type PoundageWaterList struct {
	VideoWater   float64 `json:"video_water"`   //视讯退佣比例
	ElectWater   float64 `json:"elect_water"`   //电子退佣比例
	SportWater   float64 `json:"sport_water"`   //体育退佣比例
	LotteryWater float64 `json:"lottery_water"` //彩票退佣比例
	FishWater    float64 `json:"fish_water"`    //捕鱼退佣比例
}
type HelperData struct {
	Negative   float64 `json:"negative"`    //负数退佣
	TotalMoney float64 `json:"total_money"` //剩余费用
	AddStatus  int64   `json:"add_status"`  //累计标识
	ReturnCash float64 `json:"return_cash"` //用户当期退佣
}
type RebateListIn struct {
	SiteId           string  `xorm:"site_id" json:"site_id"`                       //操作站点id
	SiteIndexId      string  `xorm:"site_index_id" json:"site_index_id"`           //站点前台id
	PeriodsId        int64   `xorm:"periods_id" json:"periods_id"`                 //期数id
	AgencyId         int64   `xorm:"agency_id" json:"agency_id"`                   //代理id
	AgencyAccount    string  `xorm:"agency_account" json:"agency_account"`         //代理账号
	BeforeJack       float64 `xorm:"before_jack" json:"before_jack"`               // 前期奖池
	NowJack          float64 `xorm:"now_jack" json:"now_jack"`                     //当期奖池
	BeforeProfit     float64 `xorm:"before_profit" json:"before_profit"`           //前期赢利
	BeforeBetting    float64 `xorm:"before_betting" json:"before_betting"`         //前期赢利
	BeforeCost       float64 `xorm:"before_cost" json:"before_cost"`               //前期费用
	NowProfit        float64 `xorm:"now_profit" json:"now_profit"`                 //当期赢利
	NowBetting       float64 `xorm:"now_betting" json:"now_betting"`               //当期有效投注
	NowCost          float64 `xorm:"now_cost" json:"now_cost"`                     //当期费用
	EffectiveMember  int64   `xorm:"effective_member" json:"effective_member"`     //有效会员数(含前期)
	Rebate           float64 `xorm:"rebate" json:"rebate"`                         //本次退佣金额
	RebateWater      float64 `xorm:"rebate_water" json:"rebate_water"`             //本次退水金额
	Status           int8    `xorm:"statue" json:"statue"`                         //是否已退佣 0未操作1已操作
	CreateTime       int64   `xorm:"create_time" json:"create_time"`               //创建时间
	Remark           string  `xorm:"remark" json:"remark"`                         //备注
	Balance          float64 `xorm:"balance" json:"balance"`                       //代理余额
	BankInFee        float64 `xorm:"bank_in_fee" json:"bank_in_fee"`               //入款费用
	BankOutFee       float64 `xorm:"bank_out_fee" json:"bank_out_fee"`             //出款费用
	CatmInFee        float64 `xorm:"catm_in_fee" json:"catm_in_fee"`               //手动入款费用
	OtherFee         float64 `xorm:"other_fee" json:"other_fee"`                   //其它优惠
	NowPayoffElect   float64 `xorm:"now_payoff_elect" json:"now_payoff_elect"`     //当期电子盈利
	NowPayoffVideo   float64 `xorm:"now_payoff_video" json:"now_payoff_video"`     //当期视讯盈利
	NowPayoffLottery float64 `xorm:"now_payoff_lottery" json:"now_payoff_lottery"` //当期彩票盈利
	NowPayoffSport   float64 `xorm:"now_payoff_sport" json:"now_payoff_sport"`     //当期体育盈利
	NowPayoffFish    float64 `xorm:"now_payoff_fish" json:"now_payoff_fish"`       //当期捕鱼费用
	DiscountFee      float64 `xorm:"discount_fee" json:"discount_fee"`             //返水费用
}

//wap现金表数据插入
type WapCashRecord struct {
	SiteId        string  ` json:"site_id"`        //操作站点id
	SiteIndexId   string  `json:"site_index_id"`   //站点前台id
	MemberId      int64   ` json:"member_id"`      //会员id
	UserName      string  `json:"user_name"`       //会员账号
	AgencyId      int64   ` json:"agency_id"`      //会员所属代理id
	AgencyAccount string  ` json:"agency_account"` //会员所属代理id
	SourceType    int     `json:"source_type"`     //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	AfterBalance  float64 `json:"after_balance"`   //操作后余额
	Balance       float64 `json:"balance"`         //金额
	DisBalance    float64 `json:"dis_balance"`     //优惠金额
	ClientType    int64   `json:"client_type"`     //客户端类型
	TradeNo       string  `json:"trade_no"`        //订单号
}
