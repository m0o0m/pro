package input

//退佣统计查询
type CheckList struct {
	SiteId        string  `json:"site_id"  valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`         //操作站点id
	SiteIndexId   string  `json:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"`    //站点前台id
	PeriodsId     int64   `json:"periods_id"  valid:"Required;Min(1);Max(11);ErrorCode(70012)"`             //期数id
	AgencyAccount string  `json:"agency_account"  valid:"Required;MinSize(1);MaxSize(20);ErrorCode(10028)"` //代理账号
	Rebate        float64 `json:"rebate"  valid:"Required;Max(23);ErrorCode(90201)"`                        //本次退佣金额
	StartNum      int64   `json:"start_num"  valid:"Required;Min(1);Max(10);ErrorCode(70014)"`              //有效会员最小
	EndNum        int64   `json:"end_num"  valid:"Required;Min(1);Max(10);ErrorCode(70014)"`                //最大有效会员
}
