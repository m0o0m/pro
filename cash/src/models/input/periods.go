package input

//期数管理
//获得期数列表
type PeriodsGet struct {
	SiteId      string `query:"siteId"`       // 站点id
	SiteIndexId string `query:"siteIndexId" ` //站点前台id
}

//获得单条期数数据
type PeriodsGetOne struct {
	PerId       int64  `query:"id" valid:"Max(11);ErrorCode(70012)"`                                    //期数id
	SiteId      string `query:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `query:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//(删除)
type PeriodsDeleteOne struct {
	PerId       int64  `query:"id" valid:"Max(11);ErrorCode(70012)"`                                    //期数id
	SiteId      string `query:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `query:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//新增单条数据
type PeriodsAdd struct {
	SiteId      string `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"MaxSize(30);ErrorCode(90401)"`                             // 期数名称
	StartTime   int64  `json:"start_time" valid:"Max(10);ErrorCode(90402)"`                            // 开始时间
	EndTime     int64  `json:"end_time" valid:"Max(10);ErrorCode(90403)"`                              //结束时间
	Status      int    `json:"status" valid:"Max(1);ErrorCode(90404)"`                                 // 退佣状态 0未退佣1已退佣
}

//修改单条数据
type PeriodsUpdate struct {
	Id          int64  `json:"'id'"  valid:"Max(11);ErrorCode(70012)"`                                 //期数id
	SiteId      string `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"MaxSize(30);ErrorCode(90401)"`                             // 期数名称
	StartTime   int64  `json:"start_time" valid:"Max(10);ErrorCode(90402)"`                            // 开始时间
	EndTime     int64  `json:"end_time" valid:"Max(10);ErrorCode(90403)"`                              //结束时间
	Status      int    `json:"status" valid:"Max(1);ErrorCode(90404)"`                                 // 退佣状态 0未退佣1已退佣
}

//获取站点列表
type GetSiteList struct {
	SiteId      string `query:"siteId"`      // 站点id
	SiteIndexId string `query:"siteIndexId"` //站点前台id
}

//退佣冲销
type Commission struct {
	Id          int64  `json:"'id'"  valid:"Max(11);ErrorCode(70012)"`                                 //期数id
	SiteId      string `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       // 站点id
	SiteIndexId string `json:"site_index_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	Status      int    `json:"status" valid:"Max(1);ErrorCode(90404)"`                                 // 退佣状态 0未退佣1已退佣
}
