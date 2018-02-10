package input

//退佣设定列表
type OverRide struct {
	Id          int64    `json:"id" valid:"Max(32);ErrorCode(90501)"`                                    //数据id
	SiteId      string   `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string   `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Amount      int64    `json:"self_profit" valid:"Max(32);ErrorCode(90502)"`                           // 自身盈利金额(self_profit)
	Member      int64    `json:"effective_user" valid:"Max(32);ErrorCode(90503)"`                        // 有效会员数
	List        OverList `json:"list"`                                                                   //返佣列表
}

//修改单条退佣
type OverRideUpdate struct {
	Id          int64      `json:"id" valid:"Max(32);ErrorCode(90501)"`                                    //数据id
	SiteId      string     `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string     `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Amount      int64      `json:"self_profit" valid:"Max(32);ErrorCode(90502)"`                           // 自身盈利金额(self_profit)
	Member      int64      `json:"effective_user" valid:"Max(32);ErrorCode(90503)"`                        // 有效会员数
	List        []OverList `json:"list"`                                                                   //返佣列表
}

//增加一条退佣记录
type OverRideAdd struct {
	SiteId      string        `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string        `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Amount      int64         `json:"self_profit" valid:"Required;ErrorCode(90502)"`                          // 自身盈利金额(self_profit)
	Member      int64         `json:"effective_user" valid:"Required;ErrorCode(90503)"`                       // 有效会员数
	List        []OverAddList `json:"list"`                                                                   //返佣列表
}

//获取一条记录
type OverRideGet struct {
	SiteId      string `query:"siteId"`      //站点id(site_id)
	SiteIndexId string `query:"siteIndexId"` //站点前台id
}

//新增退佣退水比例
type OverAddList struct {
	Rebate    float64 `json:"rebate_radio" valid:"Required;Min(1);Max(4);ErrorCode(90504)"` // 退佣比例
	Rewater   float64 `json:"water_radio" valid:"Required;Min(1);Max(4);ErrorCode(90505)"`  // 退水比例
	ProductId int64   `json:"product_id" valid:"Required;Min(1);Max(4);ErrorCode(90506)"`   //商品id
	SetId     int64   `json:"set_id" valid:"Required;Min(1);Max(4);ErrorCode(90507)"`       //设定id
}

//退佣退水比例
type OverList struct {
	Rebate    float64 `json:"rebate_radio" valid:"Required;Min(1);Max(4);ErrorCode(90504)"` // 退佣比例
	Rewater   float64 `json:"water_radio" valid:"Required;Min(1);Max(4);ErrorCode(90505)"`  // 退水比例
	ProductId int64   `json:"product_id" valid:"Required;Min(1);Max(4);ErrorCode(90506)"`   //商品id
	SetId     int64   `json:"set_id" valid:"Required;Min(1);Max(4);ErrorCode(90507)"`       //设定id
}

//
type OverRideDelet struct {
	SiteId      string `query:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `query:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(90507)"`                            //设定数据id
}
type OverGetOne struct {
	SiteId      string `query:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `query:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(90507)"`                            //设定数据id
}
type UpdateMoney struct {
	SiteId      string `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Money       int64  `json:"valid_money" valid:"Max(10);ErrorCode(90508)"`                           //有效会员投注金额
}
