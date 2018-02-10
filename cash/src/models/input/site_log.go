package input

//操作日志
type SiteDoLog struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	StartTime   string `query:"startTime"`                                       //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Key         int8   `query:"key" valid:"Range(0,2);ErrorCode(50141)"`         //搜索条件(1登录名   2ip)
	Value       string `query:"value" valid:"MaxSize(30);ErrorCode(50142)"`      //搜索值
}

//登录日志
type SiteLoginLog struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	StartTime   string `query:"startTime"`                                       //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Key1        int8   `query:"accountId" valid:"Range(0,2);ErrorCode(50141)"`   //搜索条件（1.会员  2.管理员）
	Key2        int8   `query:"loginId" valid:"Range(0,2);ErrorCode(50141)"`     //搜索条件(1登录名   2ip)
	Value       string `query:"text" valid:"MaxSize(30);ErrorCode(50142)"`       //搜索值
}

//自动稽核
type AutoAudit struct {
	SiteId string `query:"siteId"` //站点id
	Device int8   `query:"device"` //设备
	Key    int8   `query:"key"`    //搜索条件(1登录名   2ip)
	Value  string `query:"value"`  //搜索值
}
