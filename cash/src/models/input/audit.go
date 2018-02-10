package input

//稽核日志
type AuditLogGet struct {
	SiteId      string `query:"siteId" `      //站点id
	SiteIndexId string `query:"siteIndexId" ` //子站
	StartTime   string `query:"startTime" `   //日志开始时间条件
	EndTime     string `query:"endTime" `     //日志结束时间条件
	Account     string `query:"account" `     //会员账号
	PageSize    int    `query:"pageSize"`     //每页显示
	Page        int    `query:"page"`         //页码
}

//即时稽核
type AuditNow struct {
	SiteId  string `query:"siteId"`  //站点id
	Account string `query:"account"` //会员账号
}

//稽核日志(后台管理)
type AuditLogAdmin struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //子站
	StartTime   string `query:"startTime"`                                       //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Account     string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"`    //会员账号
}
