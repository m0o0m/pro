//子帐号
package input

//子帐号列表
type ChildAccountList struct {
	SiteId    string `query:"siteId" valid:"MaxSize(4);ErrorCode(50058)"`   //站点
	Account   string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"` //帐号
	Name      string `query:"name" valid:"MaxSize(20);ErrorCode(50110)"`    //名称
	IsLogin   int8   `query:"isLogin" valid:"Range(0,2);ErrorCode(50083)"`  //在线
	Type      int8   `query:"type" valid:"Range(0,2);ErrorCode(50114)"`     //类型(0.所有的1.不包括子账号的2.子账号)
	CType     int8   `query:"cType" valid:"Range(0,2);ErrorCode(50114)"`    //搜索条件(0.所有的1.账号2.名称)
	Value     string `query:"value" `                                       //搜索值
	StartTime string `query:"startTime"`                                    //开始时间
	EndTime   string `query:"endTime"`                                      //结束时间
}

//获取一条子帐号
type OneChildAccount struct {
	Id     int64  `query:"id" valid:"Required;Min(1);ErrorCode(30041)"`         //子帐号id
	SiteId string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
}

//子帐号修改
type ChildAccountChange struct {
	SiteId     string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
	Id         int64  `json:"id" valid:"Required;Min(1);ErrorCode(10002)"`         //子帐号id
	Password   string `json:"password" valid:"MaxSize(50);ErrorCode(20005)"`       //密码
	RePassword string `json:"rePassword" valid:"MaxSize(50);ErrorCode(20007)"`     //重复密码
	Name       string `json:"name" valid:"MaxSize(20);ErrorCode(50110)"`           //名称
}

//子帐号状态修改
type ChildAccountStatus struct {
	SiteId string `json:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`  //站点id
	Id     int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //子帐号id
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(10081)"`  //状态
}

//子站点下拉
type SiteIndexIdList struct {
	SiteId string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
}
