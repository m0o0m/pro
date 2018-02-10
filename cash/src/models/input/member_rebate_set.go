package input

//会员返佣设定:修改,添加
type MemberRebateSetAdd struct {
	Id          int64  `json:"id"`
	SiteId      string `json:"site_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //操作站点id
	SiteIndexId string `json:"site_index_id" valid:"Required"`                       //站点前台id
	ValidMoney  int64  `json:"valid_money" valid:"Required;ErrorCode(70002)"`        //有效总投注
	DiscountUp  int64  `json:"discount_up" valid:"Required;ErrorCode(70003)"`        //优惠上限
	DeleteTime  int64  `json:"delete_time"`                                          //删除时间:软删除
}

//删除返佣设定:删除,查询
type MemberRebateSetDel struct {
	Id int64 `json:"id" query:"id" valid:"Required;ErrorCode(70008)"` //软删除的id
}

//查询返佣设定列表
type MemberRebateSetList struct {
	SiteId      string `query:"site_id"`                                           //操作站点id
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
}
