package back

//各个商品返佣的详情
type RebateDetail struct {
	Id         int                              `json:"id" xorm:"id"`
	Agency     string                           `json:"agency" xorm:"agency"`   //代理商账号
	Account    string                           `json:"account" xorm:"account"` //会员账号
	AllBet     float64                          `json:"all_bet" xorm:"all_bet"` //有效总投注
	Rebate     float64                          `json:"rebate" xorm:"rebate"`   //返佣小计
	Status     int                              `json:"status" xorm:"status"`   //返佣状态
	AllProduct *map[string]*RebateRecordProduct `json:"all_info" xorm:"-"`      //所有商品返佣金额详情
}

//总返佣详情
type SumRebateDetail struct {
	SumPeople  int                              `json:"sum_people"` //总人数
	NoPeople   int                              `json:"no_people"`  //冲销人数
	AllBet     float64                          `json:"all_bet"`    //总计有效总投注
	Rebate     float64                          `json:"rebate"`     //返佣总计
	AllProduct *map[string]*RebateRecordProduct //所有商品返佣金额详情
}

//商品打码返佣情况
type RebateRecordProduct struct {
	RecordId    int     `xorm:"record_id" json:"-"`               //返佣记录id
	ProductName string  `xorm:"product_name" json:"product_name"` //商品名称
	ProductId   int     `xorm:"product_id" json:"product_id"`     //商品id
	ProductBet  float64 `xorm:"product_bet" json:"product_bet"`   //商品有效打码
	Money       float64 `xorm:"money" json:"money"`               //金额(返佣金额)
}
