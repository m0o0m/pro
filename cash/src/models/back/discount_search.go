//优惠查询
package back

//优惠列表
type DiscountSearchListBack struct {
	Id       string `xorm:"'id' PK" json:"id"`         //主键id
	IndexId  string `xorm:"index_id" json:"indexId"`   //前台id
	SiteName string `xorm:"site_name" json:"siteName"` //站点名称
}

//优惠查询统计
type DiscountList struct {
	Id          int64   `json:"id" xorm:"id"`
	SiteId      string  `json:"siteId" xorm:"site_id"`            //操作站点id
	SiteIndexId string  `json:"siteIndexId" xorm:"site_index_id"` //站点前台id
	AdminUser   string  `json:"adminUser" xorm:"admin_user"`      //操作者
	LevelId     string  `json:"levelId" xorm:"level_id"`          //层级,用逗号分隔
	StartTime   int64   `json:"startTime" xorm:"start_time"`      //开始时间
	EndTime     int64   `json:"endTime" xorm:"end_time"`          //结束时间
	CreateTime  int64   `json:"createTime" xorm:"create_time"`    //创建时间
	Event       string  `json:"event" xorm:"event"`               //事件
	NoPeopleNum int64   `json:"noPeopleNum" xorm:"no_people_num"` //冲销人数
	PeopleNum   int64   `json:"peopleNum" xorm:"people_num"`      //退水人数
	Money       float64 `json:"money" xorm:"money"`               //金额'
	Bet         float64 `json:"bet" xorm:"bet"`                   //综合打码倍数
}

//
type DiscountGetInfo struct {
	Id          int64   `json:"id" xorm:"id"`
	SiteId      string  `json:"siteId" xorm:"site_id"`            //操作站点id
	SiteIndexId string  `json:"siteIndexId" xorm:"site_index_id"` //站点前台id
	Account     string  `json:"account" xorm:"account"`           //会员账号
	StartTime   int64   `json:"startTime" xorm:"start_time"`      //开始时间
	EndTime     int64   `json:"endTime" xorm:"end_time"`          //结束时间
	MemberId    int64   `json:"memberId" xorm:"member_id"`        //所属会员id
	LevelId     string  `json:"levelId" xorm:"level_id"`          //所属层级id
	Betall      float64 `json:"betall" xorm:"betall"`             //有效总投注
	AllMoney    float64 `json:"allMoney" xorm:"all_money"`        //总返水金额
	SelfMoney   float64 `json:"selfMoney" xorm:"self_money"`      //自助反水金额
	RebateWater float64 `json:"rebateWater" xorm:"rebate_water"`  //本次退水金额
	Status      int8    `json:"status" xorm:"status"`             //是否已返佣 0未操作1已操作
	CreateTime  int64   `json:"createTime" xorm:"create_time"`    //创建时间
	ProductId   int64   `json:"productId" xorm:"product_id"`      //商品分类id
	ProductBet  float64 `json:"productBet" xorm:"product_bet"`    //对应商品有效投注额
	Rate        float64 `json:"rate" xorm:"rate"`                 //比例
	Money       float64 `json:"money" xorm:"money"`               //金额
}

//
type DiscountInfo struct {
	Id          int64             `json:"id" xorm:"id"`
	SiteId      string            `json:"site_id" xorm:"site_id"`             //操作站点id
	SiteIndexId string            `json:"site_index_id" xorm:"site_index_id"` //站点前台id
	Account     string            `json:"account" xorm:"account"`             //会员账号
	StartTime   int64             `json:"start_time" xorm:"start_time"`       //开始时间
	EndTime     int64             `json:"end_time" xorm:"end_time"`           //结束时间
	PeriodsId   int64             `json:"periods_id" xorm:"periods_id"`       //期数id
	MemberId    int64             `json:"member_id" xorm:"member_id"`         //所属会员id
	LevelId     string            `json:"level_id" xorm:"level_id"`           //所属层级id
	Betall      float64           `json:"betall" xorm:"betall"`               //有效总投注
	AllMoney    float64           `json:"all_money" xorm:"all_money"`         //总返水金额
	SelfMoney   float64           `json:"self_money" xorm:"self_money"`       //自助反水金额
	RebateWater float64           `json:"rebate_water" xorm:"rebate_water"`   //本次退水金额
	Status      int8              `json:"status" xorm:"status"`               //是否已返佣 0未操作1已操作
	CreateTime  int64             `json:"create_time" xorm:"create_time"`     //创建时间
	List        []DiscountProduct `json:"list"`                               //分类列表
}
type DiscountProduct struct {
	RecordId   int64   `json:"record_id" xorm:"record_id"`     //退水记录id
	ProductId  int64   `json:"product_id" xorm:"product_id"`   //商品分类id
	ProductBet float64 `json:"product_bet" xorm:"product_bet"` //对应商品有效投注额
	Rate       float64 `json:"rate" xorm:"rate"`               //比例
	Money      float64 `json:"money" xorm:"money"`             //金额
}
