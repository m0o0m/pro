//出入款管理[admin]
package input

//入款管理
type InDeposit struct {
	InType      int8   `query:"inType" valid:"Range(0,2);ErrorCode(50114)"`      //入款类型
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(50058)"`      //站点
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台站点
	Account     string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"`    //帐号
	OrderNum    string `query:"orderNum" valid:"MaxSize(21);ErrorCode(50108)"`   //订单号
	StartTime   string `query:"startTime"`                                       //开始时间                                        //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Equipment   int8   `query:"equipment" valid:"Range(0,4);ErrorCode(60033)"`   //设备类型 1pc 2wap 3android 4ios
}

//出款管理
type OutDeposit struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(50058)"`      //站点
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台站点
	Account     string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"`    //帐号
	StartTime   string `query:"startTime"`                                       //开始时间                                        //开始时间
	EndTime     string `query:"endTime"`                                         //结束时间
	Equipment   int8   `query:"equipment" valid:"Range(0,4);ErrorCode(60033)"`   //设备类型 1pc 2wap 3android 4ios
}
