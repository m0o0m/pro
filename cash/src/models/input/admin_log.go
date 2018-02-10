package input

//后台人员操作日志查询 ,因为是条件匹配,所以均不是必传
type AdminLog struct {
	Account     string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"` //操作者账号
	StartTime   string `query:"startTime" valid:"ErrorCode(30155)"`           //开始时间
	EndTime     string `query:"endTime" valid:"ErrorCode(30156)"`             //结束时间
	OperatePath string `query:"url" valid:"MaxSize(50);ErrorCode(50140)"`     //操作路径
	Ip          string `query:"ip" valid:"MaxSize(10);ErrorCode(50135)"`      //操作者iP
	Type        int8   `query:"type" valid:"Range(0,4);ErrorCode(50114)"`     //操作类型1,增加2，删除，3，查看，4，修改
}
