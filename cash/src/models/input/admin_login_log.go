package input

//总后台登录
type AdminLoginLog struct {
	Account   string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"` //帐号
	Ip        string `query:"ip" valid:"MaxSize(20);ErrorCode(50135)"`      //ip
	RoleId    int64  `query:"roleId" valid:"Min(0);ErrorCode(30077)"`       //角色id
	Device    int8   `query:"device" valid:"Range(0,5);ErrorCode(60033)"`   //设备
	StartTime string `query:"startTime" valid:"ErrorCode(30155)"`           //开始时间
	EndTime   string `query:"endTime" valid:"ErrorCode(30156)"`             //结束时间
}
