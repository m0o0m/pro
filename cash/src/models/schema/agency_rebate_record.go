package schema

import "global"

//代理退佣退水记录
type AgencyRebateRecord struct {
	Id               int64   `xorm:"id"`
	SiteId           string  `xorm:"site_id"`            //操作站点id
	SiteIndexId      string  `xorm:"site_index_id"`      //站点前台id
	PeriodsId        int64   `xorm:"periods_id"`         //期数id
	AgencyId         int64   `xorm:"agency_id"`          //代理id
	AgencyAccount    string  `xorm:"agency_account"`     //代理账号
	BeforeJack       float64 `json:"before_jack"`        // 前期奖池
	NowJack          float64 `json:"now_jack"`           //当期奖池
	BeforeProfit     float64 `xorm:"before_profit"`      //前期赢利
	BeforeBetting    float64 `xorm:"before_betting"`     //前期赢利
	BeforeCost       float64 `xorm:"before_cost"`        //前期费用
	NowProfit        float64 `xorm:"now_profit"`         //当期赢利
	NowBetting       float64 `xorm:"now_betting"`        //当期有效投注
	NowCost          float64 `xorm:"now_cost"`           //当期费用
	EffectiveMember  int64   `xorm:"effective_member"`   //有效会员数(含前期)
	Rebate           float64 `xorm:"rebate"`             //本次退佣金额
	RebateWater      float64 `xorm:"rebate_water"`       //本次退水金额
	Status           int8    `xorm:"statue"`             //是否已退佣 0未操作1已操作
	CreateTime       int64   `xorm:"create_time"`        //创建时间
	Remark           string  `xorm:"remark"`             //备注
	Balance          float64 `xorm:"balance"`            //代理余额
	BankInFee        float64 `xorm:"bank_in_fee"`        //入款费用
	BankOutFee       float64 `xorm:"bank_out_fee"`       //出款费用
	CatmInFee        float64 `xorm:"catm_in_fee"`        //手动入款费用
	OtherFee         float64 `xorm:"other_fee"`          //其它优惠
	NowPayoffElect   float64 `xorm:"now_payoff_elect"`   //当期电子盈利
	NowPayoffVideo   float64 `xorm:"now_payoff_video"`   //当期视讯盈利
	NowPayoffLottery float64 `xorm:"now_payoff_lottery"` //当期彩票盈利
	NowPayoffSport   float64 `xorm:"now_payoff_sport"`   //当期体育盈利
	NowPayoffFish    float64 `xorm:"now_payoff_fish"`    //当期捕鱼费用
	DiscountFee      float64 `xorm:"discount_fee"`       //返水费用
}

func (*AgencyRebateRecord) TableName() string {
	return global.TablePrefix + "agency_rebate_record"
}
