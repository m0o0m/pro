package back

//会员推广设定
type SpreadSet struct {
	SiteId       string  `xorm:"site_id" json:"siteId"`              //站点id
	SiteIndexId  string  `xorm:"site_index_id" json:"siteIndexId"`   //前台站点ID
	IsOpen       int8    `xorm:"is_open" json:"isOpen"`              //是否开启会员推广
	IsIp         int8    `xorm:"is_ip" json:"isIp"`                  //是否过滤ip
	IsMateAgency int8    `xorm:"is_mate_agency" json:"isMateAgency"` //是否匹配推广会员代理
	IsCode       int8    `xorm:"is_code" json:"isCode"`              //返佣会员是否需要打码
	RankingNum   int64   `xorm:"ranking_num"   json:"rankingNum"`    //排行榜人数系数
	RankingMoney float64 `xorm:"ranking_money"  json:"rankingMoney"` //排行榜金额系数
}

//会员返佣设定
type MemberRebateSet struct {
	Id           int64                            `xorm:"id" json:"id"`
	SiteId       string                           `xorm:"site_id" json:"site_id"`             //操作站点id
	SiteIndexId  string                           `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	ValidMoney   int64                            `xorm:"valid_money" json:"valid_money"`     //有效总投注
	DiscountUp   int64                            `xorm:"discount_up" json:"discount_up"`     //优惠上限
	DeleteTime   int64                            `xorm:"delete_time" json:"delete_time"`     //删除时间
	ProductRates *map[string]*MemberRebateProduct `xorm:"-" json:"product_rate"`              //会员返佣比例
}

//会员返佣对应各商品分类比例
type MemberRebateProduct struct {
	SetId       int64   `xorm:"set_id" json:"set_id"`             //返佣设定id
	ProductId   int64   `xorm:"product_id" json:"product_id"`     //商品Id
	ProductName string  `xorm:"product_name" json:"product_name"` //商品名称
	VType       string  `xorm:"v_type" json:"v_type"`             //商品标识
	Rate        float64 `xorm:"rate" json:"rate"`                 //比例
}

//会员返佣期数(总计) 插入
type MemberRebateCommission struct {
	Id              int64                       `json:"id" xorm:"id PK autoincr"`
	SiteId          string                      `json:"siteId" xorm:"site_id"`                   //操作站点id
	SiteIndexId     string                      `json:"siteIndexId" xorm:"site_index_id"`        //站点前台id
	AdminUser       string                      `json:"adminUser" xorm:"admin_user"`             //操作者
	Event           string                      `json:"event" xorm:"event"`                      //事件
	StartTime       int64                       `json:"startTime" xorm:"start_time"`             //开始时间
	EndTime         int64                       `json:"endTime" xorm:"end_time"`                 //结束时间
	CreateTime      int64                       `json:"createTime" xorm:"'create_time' created"` //创建时间
	NoPeopleNum     int64                       `json:"noPeopleNum" xorm:"no_people_num"`        //冲销人数
	PeopleNum       int64                       `json:"peopleNum" xorm:"people_num"`             //退水人数
	Money           float64                     `json:"money" xorm:"money"`                      //金额
	Bet             int64                       `json:"bet" xorm:"bet"`                          //综合打码倍数
	TotalBet        float64                     `json:"totalBet" xorm:"total_bet"`               //总有效打码
	RebateRecordMap *map[int64]*PreRebateRecord `json:"rebateRecordMap" xorm:"-"`                //返佣记录(详情)
}

//预返佣记录  用来存入redis
type PreRebateRecord struct {
	Id         int64                               `xorm:"id" json:"id"`                  //所属会员id
	Agency     string                              `xorm:"agency" json:"agency"`          //代理商
	Account    string                              `xorm:"account" json:"account"`        //所属会员
	AllBet     float64                             `xorm:"all_bet" json:"allBet"`         //总计有效
	Rebate     float64                             `xorm:"rebate" json:"rebate"`          //返佣小计
	AllProduct *map[string]*PreRebateRecordProduct `xorm:"all_product" json:"allProduct"` //所有商品返佣金额
}

//预返佣记录  用来返回前端
type PreRebateRecordShow struct {
	Id         int64                     `xorm:"id" json:"id"`                   //所属会员id
	Agency     string                    `xorm:"agency" json:"agency"`           //代理商
	Account    string                    `xorm:"account" json:"account"`         //所属会员
	AllBet     float64                   `xorm:"all_bet" json:"all_bet"`         //总计有效
	Rebate     float64                   `xorm:"rebate" json:"rebate"`           //返佣小计
	AllProduct []*PreRebateRecordProduct `xorm:"all_product" json:"all_product"` //所有商品返佣金额
}

//预返佣总计(最下面一排的统计信息) 方便计算
type PreRebateRecordTotal struct {
	People     int64                               `xorm:"people" json:"people"`           //人数
	AllBet     float64                             `xorm:"all_bet" json:"all_bet"`         //总计有效
	Rebate     float64                             `xorm:"rebate" json:"rebate"`           //返佣小计
	AllProduct *map[string]*PreRebateRecordProduct `xorm:"all_product" json:"all_product"` //所有商品返佣金额总计
}

//预返佣总计(最下面一排的统计信息) 方便展示
type PreRebateRecordTotalShow struct {
	People     int64                     `xorm:"people" json:"people"`           //人数
	AllBet     float64                   `xorm:"all_bet" json:"all_bet"`         //总计有效
	Rebate     float64                   `xorm:"rebate" json:"rebate"`           //返佣小计
	AllProduct []*PreRebateRecordProduct `xorm:"all_product" json:"all_product"` //所有商品返佣金额总计
}

//预返佣商品金额
type PreRebateRecordProduct struct {
	//RecordId    int     `xorm:"-" json:"-"`               //返佣记录id
	ProductName string  `xorm:"product_name" json:"product_name"` //商品名称
	VType       string  `xorm:"v_type" json:"v_type"`
	ProductId   int64   `xorm:"product_id" json:"product_id"`   //商品id
	ProductBet  float64 `xorm:"product_bet" json:"product_bet"` //商品有效打码
	Money       float64 `xorm:"-" json:"money"`                 //金额(应该返佣金额)
	SpreadId    int64   `xorm:"spread_id" json:"-"`             //推广id
}

//返佣有效会员
type ValidMember struct {
	Id      int64  `xorm:"id" json:"id"`           //会员id
	Agency  string `xorm:"agency" json:"agency"`   //代理账号
	Account string `xorm:"account" json:"account"` //会员账号
}

//返佣或者冲销的会员
type RebateAuditMember struct {
	Id            int64   `xorm:"id"`             //会员id
	Account       string  `xorm:"account"`        //会员账号
	AgencyId      int64   `xorm:"agency_id"`      //代理id
	AgencyAccount string  `xorm:"agency_account"` //代理账号
	Balance       float64 `xorm:"balance"`        //剩余金额
	Amount        float64 `xorm:"-"`              //操作金额
}
