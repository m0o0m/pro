//代理管理
package input

//代理管理列表
type AgencyManageList struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(50058)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台站点id
	Status      int8   `query:"status" valid:"Range(0,2);ErrorCode(30050)"`      //状态
	Type        int8   `query:"type" valid:"Range(0,4);ErrorCode(50114)"`        //代理类型
	SpreadId    int8   `query:"spreadId" valid:"Range(1,3);ErrorCode(30118)"`    //搜索条件1.账号2.代理账号3.推广id
	Name        string `query:"name" valid:"MaxSize(20);ErrorCode(30118)"`       //搜索值
}

//代理占成比
type AgencyOccupationRatio struct {
	Id int64 `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
}

//站点管理-代理列表
type SiteAgencyList struct {
	SiteId string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
}

//站点管理-代理详情
type SiteAgencyInfo struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(30117)"`
}

//站点管理-代理修改
type SiteAgencyEdit struct {
	Id          int64  `json:"id"  valid:"Required;ErrorCode(30117)"`
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Username    string `json:"userName" valid:"Required;MaxSize(50);ErrorCode(10027)"`   //代理名称
	Account     string `json:"account" valid:"Required;MaxSize(30);ErrorCode(10028)"`    //账号
	Remark      string `json:"remark"`                                                   //备注
}

//站点管理-代理增加
type SiteAgencyAdd struct {
	SiteId        string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`         //站点id
	SiteIndexId   string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`    //站点前台id
	Username      string `json:"userName" valid:"Required;MaxSize(50);ErrorCode(10027)"`      //代理名称
	ParentAccount string `json:"parentAccount" valid:"Required;MaxSize(30);ErrorCode(30211)"` //父级账号
	Account       string `json:"account" valid:"Required;MaxSize(30);ErrorCode(10028)"`       //账号
	Remark        string `json:"remark"`                                                      //备注
}
