//ip开关和ip白名单[admin]
package input

//ip开关表
type IpSetList struct {
	IpStart string `query:"ipStart" valid:"MaxSize(20);ErrorCode(50129)"` // 起始ip
	IpEnd   string `query:"ipEnd" valid:"MaxSize(20);ErrorCode(50130)"`   // 结束ip
}

//ip开关添加
type IpSetAdd struct {
	IpStart string `json:"ipStart" valid:"IP;RangeSize(7,15);ErrorCode(50129)"` // 起始ip
	IpEnd   string `json:"ipEnd" valid:"IP;RangeSize(7,15);ErrorCode(50130)"`   // 结束ip
	Type    string `json:"type" valid:"Required;MaxSize(10);ErrorCode(50114)"`  // 控制类型：1为代理后台，2为前台，3为wap端
	State   int8   `json:"state" valid:"Range(1,2);ErrorCode(30050)"`           // 状态：1为启用，2为停用
	Remark  string `json:"remark" valid:"MaxSize(255);ErrorCode(20019)"`        // 备注
}

//ip开关修改
type IpSetChange struct {
	Id      int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`
	IpStart string `json:"ipStart" valid:"IP;RangeSize(7,15);ErrorCode(50129)"` // 起始ip
	IpEnd   string `json:"ipEnd" valid:"IP;RangeSize(7,15);ErrorCode(50130)"`   // 结束ip
	Type    string `json:"type" valid:"Required;MaxSize(10);ErrorCode(50114)"`  // 控制类型：1为客户后阳台，2为代理后台，3为前台，4为wap端
	State   int8   `json:"state" valid:"Range(1,2);ErrorCode(30050)"`           // 状态：1为启用，2为停用
	Remark  string `json:"remark" valid:"MaxSize(255);ErrorCode(20019)"`        // 备注
}

//ip白名单列表
type WhiteList struct {
	SiteId string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"` // 站点id
	Ip     string `query:"ip" valid:"MaxSize(255);ErrorCode(50135)"`   //ip白名单:英文逗号分隔  填写多个
}

//添加白名单
type WhiteListAdd struct {
	Id     int64
	SiteId string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`   // 站点id
	Ip     string `json:"ip" valid:"Required;MaxSize(20);ErrorCode(50135)"`      //ip
	Remark string `json:"remark" valid:"Required;MaxSize(255);ErrorCode(20019)"` //备注
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`            //状态
}

//修改白名单
type WhiteListChange struct {
	Id     int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`
	SiteId string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`   // 站点id
	Ip     string `json:"ip" valid:"Required;MaxSize(200);ErrorCode(50135)"`     //ip
	Remark string `json:"remark" valid:"Required;MaxSize(255);ErrorCode(20019)"` //备注
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`            //状态
}
