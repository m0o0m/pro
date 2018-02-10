package input

//财务报表-额度管理

//额度统计
type QuotaCount struct {
	SiteId  string `query:"site_id"` //站点
	Account string `query:"account"` //账号
	STime   string `query:"s_time"`  //开始时间
	ETime   string `query:"e_time"`  //结束时间
}
