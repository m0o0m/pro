package input

//股东列表
type FirstAgency struct {
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	AccountName string `query:"accountName" valid:"MaxSize(10);ErrorCode(50010)"` //用户输入的查询帐号
	Isvague     int    `query:"isvague" valid:"Range(0,2);ErrorCode(50001)"`      // 是否模糊查询
	IsOnline    int    `query:"isOnline" valid:"Range(0,2);ErrorCode(50083)"`     //是否在线
	FormValue   string `query:"formValue" valid:"MaxSize(4);ErrorCode(50084)"`    //站点id
	FirstId     int    `query:"firstId"`                                          //股东id
	Id          int    `query:"id"`                                               //开户人id
	Status      int    `query:"status" valid:"Range(0,2);ErrorCode(50085)"`       //状态  1：正常   2：禁用
}

//股东管理会员优惠设定查询
type FirstDiscountSet struct {
	AcountId    int64  `query:"acount_id" valid:"Required;Min(1);ErrorCode(30041)"` //帐号id（修改谁传谁id）
	SiteIndexId string `query:"site_index_id"`                                      //站点id
	SiteId      string `query:"site_id"`                                            //站点id
}

//修改、增加股东管理会员注册优惠设定
type FirstDiscountUpdata struct {
	SiteIndexId string  `json:"site_index_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	AgencyId    int64   `json:"agency_id" valid:" Required;Min(0);ErrorCode(50013)"`        //代理id
	Offer       float64 `json:"offer" valid:" Required;ErrorCode(50086)"`                   //加入会员赠送优惠金额
	AddMosaic   int64   `json:"add_mosaic" valid:" Required;Min(0);ErrorCode(50087)"`       //优惠打码倍数
	IsIp        int8    `json:"is_ip" valid:" Required;Min(0);ErrorCode(50088)"`            //是否限制IP 1:是2:否
	SiteId      string  `json:"site_id"`                                                    //站点id
}

//新增股东
type FirstAgencyAdd struct {
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

//修改股东
type FirstAgencyEdit struct {
	Id              int64  `json:"id"`
	SiteId          string //站点id
	SiteIndexId     string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`      //站点前台id
	Password        string `json:"password" valid:"MaxSize(12);ErrorCode(30010)"`        //密码
	ConfirmPassword string `json:"confirmPassword" valid:"MaxSize(12);ErrorCode(30011)"` //确认密码
	Username        string `json:"username" valid:"Required;ErrorCode(30014)"`           //名称
}

//获取某个股东基本资料
type FirstAgencyInfo struct {
	SiteId      string //站点id
	SiteIndexId string `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" query:"id" form:"id" valid:"Required;Min(1);ErrorCode(30041)"`
}
