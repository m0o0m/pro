package input

type SecondAgency struct {
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	AccountName string `query:"accountName" valid:"MaxSize(10);ErrorCode(50010)"` //用户输入的查询帐号
	Isvague     int    `query:"isvague" valid:"Range(0,2);ErrorCode(50001)"`      // 是否模糊查询
	IsOnline    int    `query:"isOnline" valid:"Range(0,2);ErrorCode(50083)"`     //是否在线
	FormValue   int64  `query:"formValue" valid:"Min(0);ErrorCode(50084)"`        //站点id
	FirstId     int    `query:"firstId"`                                          //股东id
	Id          int    `query:"id"`                                               //开户人id
	Status      int    `query:"status" valid:"Range(0,2);ErrorCode(50085)"`       //状态  1：正常   2：禁用
	SecondId    int64  `query:"secondId"`                                         //总代id

}

//总代管理会员优惠设定查询
type SecondDiscountSet struct {
	UserId      int64  `query:"userId"`                                            //用户id
	AcountId    int64  `query:"acountId" valid:"Required;Min(1);ErrorCode(30242)"` //帐号id（修改谁传谁id）
	SiteIndexId string `query:"siteIndexId"`                                       //站点id
	SiteId      string `query:"siteId"`                                            //站点id

}

//修改、增加总代管理会员注册优惠设定
type SecondDiscountUpdata struct {
	SiteIndexId string  `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	AgencyId    int64   `json:"agencyId" valid:" Required;Min(0);ErrorCode(50013)"`       //代理id
	Offer       float64 `json:"offer" valid:" Required;ErrorCode(50086)"`                 //加入会员赠送优惠金额
	AddMosaic   int64   `json:"addMosaic" valid:" Required;Min(0);ErrorCode(50087)"`      //优惠打码倍数
	IsIp        int8    `json:"isIp" valid:" Required;Min(0);ErrorCode(50088)"`           //是否限制IP 1:是2:否
	SiteId      string  `json:"siteId"`                                                   //站点id
}

//股东id以及名称（下拉框）
type FirstIdNameBySite struct {
	SiteIndexId string `query:"siteIndexId"` //前台id
	SiteId      string `query:"siteId"`      //站点ID
}

//新增总代/代理
type AgencyAdd struct {
	SiteId          string //站点id
	SiteIndexId     string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`   //站点前台id
	Account         string `json:"account" valid:"RangeSize(6,12);ErrorCode(30009)"`  //账号
	Password        string `json:"password" valid:"RangeSize(6,12);ErrorCode(30010)"` //密码
	ConfirmPassword string `json:"confirmPassword" valid:"Required;ErrorCode(30011)"` //确认密码
	Username        string `json:"username" valid:"Required;ErrorCode(30014)"`        //名称
	Status          int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`        //状态
	RoleId          int64  `json:"roleId"`                                            //角色id
	Level           int8   `json:"level"`                                             //层级
	ParentId        int64  `json:"parentId"`                                          //上级id
}

//获取某个总代基本资料
type SecondAgencyInfo struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" query:"id" form:"id" valid:"Required;Min(1);ErrorCode(30041)"`
}
