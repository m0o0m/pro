package back

//优惠查询列表
type ListRetreatWater struct {
	Id          int64   `json:"id"`            //列表id
	AdminUser   string  `json:"admin_user"`    //操作者
	StartTime   int64   `json:"start_time"`    //开始时间
	EndTime     int64   `json:"end_time"`      //结束时间
	CreateTime  int64   `json:"create_time"`   //创建时间
	Event       string  `json:"event"`         //事件
	NoPeopleNum int64   `json:"no_people_num"` //冲销人数
	PeopleNum   int64   `json:"people_num"`    //退水人数
	Money       float64 `json:"money"`         //金额
	Bet         float64 `json:"bet"`           //综合打码倍数
}

//优惠查询明细记录+商品记录
type RetreatWaterRecordAndProduct struct {
	Id          int64   `json:"id"`           //id
	Account     string  `json:"account"`      //会员账号
	MemberId    int64   `json:"member_id"`    //所属会员id
	LevelId     string  `json:"level_id"`     //所属层级id
	Betall      float64 `json:"betall"`       //有效总投注
	AllMoney    float64 `json:"all_money"`    //总返水金额
	SelfMoney   float64 `json:"self_money"`   //自助反水金额
	RebateWater float64 `json:"rebate_water"` //本次退水金额
	Status      int8    `json:"status"`       //是否已返佣 0未操作1已操作
	CreateTime  int64   `json:"create_time"`  //创建时间
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //对应商品有效投注额
	Rate        float64 `json:"rate"`         //比例
	Money       float64 `json:"money"`        //金额
}

//优惠查询明细记录
type RetreatWaterRecord struct {
	Account     string                      `json:"account"`      //会员账号
	MemberId    int64                       `json:"member_id"`    //所属会员id
	LevelId     string                      `json:"level_id"`     //所属层级id
	Betall      float64                     `json:"betall"`       //有效总投注
	AllMoney    float64                     `json:"all_money"`    //总返水金额
	SelfMoney   float64                     `json:"self_money"`   //自助反水金额
	RebateWater float64                     `json:"rebate_water"` //本次退水金额
	Status      int8                        `json:"status"`       //是否已返佣 0未操作1已操作
	CreateTime  int64                       `json:"create_time"`  //创建时间
	Params      []RetreatWaterRecordProduct `json:"params"`       //查询明细商品记录
}

//优惠查询明细记录带id
type RetreatWaterRecord2 struct {
	Id          int64                       `json:"id"`           //明细id
	Account     string                      `json:"account"`      //会员账号
	MemberId    int64                       `json:"member_id"`    //所属会员id
	LevelId     string                      `json:"level_id"`     //所属层级id
	Betall      float64                     `json:"betall"`       //有效总投注
	AllMoney    float64                     `json:"all_money"`    //总返水金额
	SelfMoney   float64                     `json:"self_money"`   //自助反水金额
	RebateWater float64                     `json:"rebate_water"` //本次退水金额
	Status      int8                        `json:"status"`       //是否已返佣 0未操作1已操作
	CreateTime  int64                       `json:"create_time"`  //创建时间
	Params      []RetreatWaterRecordProduct `json:"params"`       //查询明细商品记录
}

//优惠查询明细商品记录
type RetreatWaterRecordProduct struct {
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	ProductBet  float64 `json:"product_bet"`  //对应商品有效投注额
	Rate        float64 `json:"rate"`         //比例
	Money       float64 `json:"money"`        //金额
}
