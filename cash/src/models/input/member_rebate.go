package input

//查询会员返佣列表
type RebateList struct {
	SiteId      string `query:"siteId" json:"siteId"`           //操作站点id
	SiteIndexId string `query:"siteIndexId" json:"siteIndexId"` //站点前台id
	Year        string `query:"year"`                           //日期（年）
	Month       string `query:"month" json:"month"`             //日期(月)
}

//查询会员返佣详情
type RebateDetails struct {
	SiteId      string `query:"site_id" json:"site_id"`                                       //操作站点id
	SiteIndexId string `query:"site_index_id" json:"site_index_id"`                           //站点前台id
	PeriodsId   int    `query:"periods_id" json:"periods_id" valid:"Min(1);ErrorCode(70012)"` //期数id
	SumPeople   int    `query:"sum_people"`                                                   //返佣总人数
	NoPeople    int    `query:"no_people"`                                                    //返佣冲销人数
}

//返佣冲销
type RebateWriteoff struct {
	RecordIds  []int64 `json:"record_ids"`                                      //需要冲销的记录id
	ClientType int64   `json:"client_type" valid:"Range(0,3);ErrorCode(70023)"` //客户端类型0pc 1wap 2android 3ios
}

//会员返佣统计
type MemberRebateCount struct {
	SiteId      string `query:"site_id" `                                          //操作站点id
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	STime       string `query:"s_time" valid:"MaxSize(10);ErrorCode(30155)"`       //起始日期
	ETime       string `query:"e_time" valid:"MaxSize(10);ErrorCode(30156)"`       //截止日期
	IsRebate    int    `query:"is_rebate" valid:"Required;Range(1,2)"`             //是否有返佣,1返佣,2没返佣
	AdminUser   string `query:"admin_user" `                                       //操作者
}

//会员返佣存入
type MemberRebateCommit struct {
	Event      string  `json:"event" valid:"Required;MinSize(1);MaxSize(30);ErrorCode(70016)"` //事件
	BetRate    int64   `json:"bet_rate" valid:"Required;ErrorCode(70017)"`                     //综合打码量(稽核倍率)
	Key        string  `json:"key" valid:"Required;ErrorCode(70018)"`                          //key,用来取出redis中缓存的数据
	MemberIds  []int64 `json:"member_id"`                                                      //将哪些会员的返佣存入到数据库
	ClientType int64   `json:"client_type" valid:"Range(0,3);ErrorCode(70023)"`                //客户端类型0pc 1wap 2android 3ios
}
