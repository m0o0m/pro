package back

//自助返水查询列表
type ListRetreatWaterSelf struct {
	Id         int64   `json:"id"`         //列表id
	OrderNum   int64   `json:"orderNum"`   //订单号
	Account    string  `json:"account"`    //会员账号
	Betting    float64 `json:"betting"`    //有效总投注
	Money      float64 `json:"money"`      //返水金额
	CreateTime int64   `json:"createTime"` //创建时间
}

//自助返水查询明细记录+商品记录
type RetreatWaterRecordAndProductSelf struct {
	Id          int64   `json:"id"`           //id
	Account     string  `json:"account"`      //会员账号
	MemberId    int64   `json:"member_id"`    //所属会员id
	Betall      float64 `json:"betall"`       //有效总投注
	RebateWater float64 `json:"rebate_water"` //本次退水金额
	CreateTime  int64   `json:"create_time"`  //返水时间
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //商品返水小计
	Money       float64 `json:"money"`        //金额
}

//自助返水查询明细记录
type RetreatWaterRecordSelf struct {
	Id          int64                           `json:"id"`           //id
	Account     string                          `json:"account"`      //会员账号
	MemberId    int64                           `json:"member_id"`    //所属会员id
	Betall      float64                         `json:"betall"`       //有效总投注
	RebateWater float64                         `json:"rebate_water"` //本次退水金额
	CreateTime  int64                           `json:"create_time"`  //返水时间
	Params      []RetreatWaterRecordProductSelf `json:"params"`       //查询明细商品记录
}

//自助返水查询明细商品记录
type RetreatWaterRecordProductSelf struct {
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //商品返水小计
	Money       float64 `json:"money"`        //金额
}

//自助返水查询明细总计汇总
type RetreatWaterRecordTotalSelf struct {
	TotalNum         int64                           `json:"total_num"`          //总个数
	TotalBetall      float64                         `json:"total_betall"`       //有效总投注
	TotalRebateWater float64                         `json:"total_rebate_water"` //本次退水金额
	Params           []RetreatWaterRecordProductSelf `json:"params"`             //查询明细商品记录
}

//自助返水查询明细列表总计汇总
type RetreatWaterRecordListTotalSelf struct {
	List  []RetreatWaterRecordSelf    `json:"list"`  //明细
	Total RetreatWaterRecordTotalSelf `json:"total"` //总计
}

//wap端自助反水返回
type WapRetreatWaterRecord struct {
	ReWaterSetId       int64   `xorm:"set_id" json:"re_water_set_id"`    //优惠返点设定id
	ProductId          int64   `xorm:"product_id" json:"product_id"`     //产品id
	ProductName        string  `xorm:"product_name" json:"product_name"` //产品名称
	Rate               float64 `xorm:"rate" json:"rate"`                 //比率
	AlreadyRebateWater float64 `json:"already_rebate_water"`             //当日已经反水
	ValidBetting       float64 `json:"valid_betting"`                    //有效投注额度
	ReWaterQutoa       float64 `json:"re_water_qutoa"`                   //反水金额
}

//点击查看单个的反水额度
type ReWaterSingle struct {
	ValidBetting float64 `json:"valid_betting"` //有效投注额度
	ReWaterQutoa int     `json:"reWaterQutoa"`  //反水金额
}

//一键查看所有的反水金额
type OneClickSeeAll struct {
	ProductId    int64   `json:"product_id"`     //产品id
	ValidBetting float64 `json:"valid_betting"`  //有效投注额度
	ReWaterQutoa float64 `json:"re_water_qutoa"` //反水金额
}

//不同的返回的反水比例列表
type BackRate struct {
	Id           int64   `xorm:"id" json:"id"`                    //会员每日打码统计表id
	ProductName  string  `xorm:"product_name" json:"productName"` //产品名称
	Rate         float64 `xorm:"rate" json:"rate"`                //比例
	BetValid     float64 `xorm:"bet_valid" json:"betValid"`       //有效投注额度
	Vtype        string  `xorm:"v_type" json:"vtype"`             //类型
	ProductId    int64   `xorm:"product_id" json:"productId"`     //产品id
	ReWaterQutoa int     `json:"reWaterQutoa"`                    //反水金额
}

//总的返回
type ReWater struct {
	AlreadyReNowDate float64    `json:"alreadyReNowDate"` //今天已经退水额度
	BackList         []BackRate `json:"backList"`         //下面的列表数据
}

//返回比列
type ReMembetWaterRate struct {
	ProductId  int     `json:"product_id"`
	Rate       float64 `json:"rate"`
	DiscountUp int     `json:"discount_up"`
}

//自助反水缓存数据
type ReturnBetData struct {
	ProductId   int64
	ProductName string
	VType       string
	Rate        float64
	BetValid    float64
	RateMoney   float64
}

//将反水信息存入redis
type WaterData struct {
	Data  []ReturnBetData
	Count float64
}
