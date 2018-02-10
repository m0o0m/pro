package back

//退佣查询列表
type Retirement struct {
	Id              int64   `json:"id"`
	SiteId          string  `json:"site_id"`          //操作站点id
	SiteIndexId     string  `json:"site_index_id"`    //站点前台id
	PeriodsId       int64   `json:"periods_id"`       //期数id
	AgencyId        int64   `json:"agency_id"`        //代理id
	AgencyAccount   string  `json:"agency_account"`   //代理账号
	BeforeJack      float64 `json:"before_jack"`      // 前期奖池
	NowJack         float64 `json:"now_jack"`         //当期奖池
	BeforeProfit    float64 `json:"before_profit"`    //前期赢利
	BeforeBetting   float64 `json:"before_betting"`   //前期赢利
	BeforeCost      float64 `json:"before_cost"`      //前期费用
	NowProfit       float64 `json:"now_profit"`       //当期赢利
	NowBetting      float64 `json:"now_betting"`      //当期有效投注
	NowCost         float64 `json:"now_cost"`         //当期费用
	EffectiveMember int64   `json:"effective_member"` //有效会员数(含前期)
	Rebate          float64 `json:"rebate"`           //本次退佣金额
	RebateWater     float64 `json:"rebate_water"`     //本次退水金额
	Status          int8    `json:"statue"`           //是否已退佣 0未操作1已操作
	CreateTime      int64   `json:"create_time"`      //创建时间
	Remark          string  `json:"remark"`           //备注
	Balance         float64 `json:"balance"`          //代理余额
	RecordId        int64   `xorm:"record_id"`        //退水记录id
	ProductId       int64   `xorm:"product_id"`       //商品分类id
	RebateRatio     float64 `xorm:"rebate_ratio"`     // 退佣比例
	RebateMoney     float64 `xorm:"rebate_money"`     // 退佣金额金额
	WaterRatio      float64 `xorm:"water_ratio"`      //退水比例
	WaterMoney      float64 `xorm:"water_money"`      //退水金额
}
type RetirementList struct {
	SiteId      string `json:"site_id"`       //操作站点id
	SiteIndexId string `json:"site_index_id"` //站点前台id
	PeriodsId   int64  `json:"periods_id"`    //期数id
	Title       string `json:"title"`         //期数名称
}
type GetProductList struct {
	Id    int64  `json:"id"`    //商品id
	Title string `json:"title"` //商品类型
}
type RetirementCheckList struct {
	Id              int64     `json:"id"`
	SiteId          string    `json:"site_id"`          //操作站点id
	SiteIndexId     string    `json:"site_index_id"`    //站点前台id
	PeriodsId       int64     `json:"periods_id"`       //期数id
	AgencyId        int64     `json:"agency_id"`        //代理id
	AgencyAccount   string    `json:"agency_account"`   //代理账号
	BeforeJack      float64   `json:"before_jack"`      // 前期奖池
	NowJack         float64   `json:"now_jack"`         //当期奖池
	BeforeProfit    float64   `json:"before_profit"`    //前期赢利
	BeforeBetting   float64   `json:"before_betting"`   //前期有效投注
	BeforeCost      float64   `json:"before_cost"`      //前期费用
	NowProfit       float64   `json:"now_profit"`       //当期赢利
	NowBetting      float64   `json:"now_betting"`      //当期有效投注
	NowCost         float64   `json:"now_cost"`         //当期费用
	EffectiveMember int64     `json:"effective_member"` //有效会员数(含前期)
	Rebate          float64   `json:"rebate"`           //本次退佣金额
	RebateWater     float64   `json:"rebate_water"`     //本次退水金额
	Status          int8      `json:"statue"`           //是否已退佣 0未操作1已操作
	CreateTime      int64     `json:"create_time"`      //创建时间
	Remark          string    `json:"remark"`           //备注
	Balance         float64   `json:"balance"`          //代理余额
	RebateRatio     []float64 `json:"rebate_retio"`     //返佣列表
	WaterRatio      []float64 `json:"water_ratio"`      //反水列表
}
