package back

//退佣列表
type OverRide struct {
	Id       int     `xorm:"id" int:"id"`
	Site     string  `xorm:"site_id" json:"site"`                 //站点id(site_id)
	Amount   float64 `xorm:"self_profit" json:"amount"`           // 自身盈利金额(self_profit)
	Member   string  `xorm:"effective_user" json:"member"`        // 有效会员数
	Pid      int     `xorm:"product_id" int:"product_id"`         // 分类id
	Percent  float64 `xorm:"rebate_radio" decimal:"rebate_radio"` // 退佣比例
	Percent1 float64 `xorm:"water_radio" decimal:"water_radio"`   // 比例
	Name     string  `xorm:"title" varchar:"title"`               //商品名称
	BetMoney int     `xorm:"valid_money" int:"valid_money"`       //有效会员投注金额
}
type Rebate struct {
	Percent float64 `xorm:"rebate_radio" decimal:"rebate_radio"` // 退佣比例
}
type Rewater struct {
	Percent1 float64 `xorm:"water_radio" decimal:"water_radio"` // 退水比例
}
type Arr struct {
	Name string `xorm:"title" varchar:"title"` //商品名称
}
type List struct {
	Id      int       `xorm:"id" json:"id"`                 //数据id
	Amount  float64   `xorm:"self_profit" json:"amount"`    // 自身盈利金额(self_profit)
	Member  string    `xorm:"effective_user" json:"member"` // 有效会员数
	Rebate  []float64 `json:"rebate"`                       //返佣列表
	Rewater []float64 `json:"rewater"`                      //返水列表
}
type Data struct {
	BetMoney int      `xorm:"valid_money" json:"bet_money"` //有效会员投注金额
	List     *List    `json:"list"`                         //详细数据列表
	Arr      []string `json:"arr"`                          //类型列表
}
type List1 struct {
	Amount  float64 `xorm:"self_profit" json:"amount"`    // 自身盈利金额(self_profit)
	Member  string  `xorm:"effective_user" json:"member"` // 有效会员数
	Rebate  []float64
	Rewater []float64
	Name    []string
}
