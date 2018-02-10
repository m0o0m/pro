package input

type ThirdAgency struct {
	SiteId      string `query:"siteId"` //用户站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	AccountName string `query:"accountName" valid:"MaxSize(10);ErrorCode(50010)"` //用户输入的查询帐号
	Isvague     int    `query:"isvague" valid:"Range(0,2);ErrorCode(50001)"`      // 是否模糊查询
	IsOnline    int    `query:"isOnline" valid:"Range(0,2);ErrorCode(50083)"`     //是否在线
	FormValue   int    `query:"formValue" valid:"Min(0);ErrorCode(50084)"`        //站点id
	FirstId     int    `query:"firstId"`                                          //股东id
	Id          int    `query:"id"`                                               //开户人id
	Status      int    `query:"status" valid:"Range(0,2);ErrorCode(50085)"`       //状态  1：正常   2：禁用
	SecondId    int64  `query:"secondId"`                                         //总代id
	ThirdId     int64  `query:"thirdId"`                                          //代理id
}

//代理管理会员优惠设定查询
type ThirdDiscountSet struct {
	AccountId   int64  `query:"accountId" valid:"Required;Min(1);ErrorCode(30041)"` //帐号id（修改谁传谁id）
	SiteIndexId string `query:"siteIndexId"`                                        //站点id
	SiteId      string `query:"siteId"`                                             //站点id
}

//修改、增加代理管理会员注册优惠设定
type ThirdDiscountUpdata struct {
	SiteIndexId string  `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	AgencyId    int64   `json:"agencyId" valid:" Required;Min(0);ErrorCode(50013)"`       //代理id
	Offer       float64 `json:"offer" valid:" Required;ErrorCode(50086)"`                 //加入会员赠送优惠金额
	AddMosaic   int64   `json:"addMosaic" valid:" Required;Min(0);ErrorCode(50087)"`      //优惠打码倍数
	IsIp        int8    `json:"isIp" valid:" Required;Min(0);ErrorCode(50088)"`           //是否限制IP 1:是2:否
	SiteId      string  `json:"siteId"`                                                   //站点id
}

//总代id以及名称（下拉框）
type SecondIdNameBySite struct {
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台id
	FirstId     int64  `query:"firstId" valid:"Min(0);ErrorCode(10050)"`         //股东id
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(50013)"`      //站点ID
}

//会员下拉框
type MemberBankDropIn struct {
	SiteId      string `query:"siteId"`                                          //站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041)"`              //会员id
}

//获取代理详细资料
type ThirdInformation struct {
	SiteId      string `query:"siteId"`                                          //站点ID
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041)"`              //代理id
}

//修改代理详细资料
type ThirdInformationUpdata struct {
	SiteId      string   `json:"siteId"`                                                  //站点ID
	SiteIndexId string   `json:"siteIndexId"`                                             //前台id
	AgencyId    int64    ` json:"agencyId" valid:"Required;Min(1);ErrorCode(10050)"`      //代理id
	ChName      string   ` json:"chName" valid:"Required;MaxSize(50);ErrorCode(50089)"`   //中文名称  50
	UsName      string   ` json:"usName" valid:"Required;MaxSize(50);ErrorCode(50090)"`   //英文名称  50
	Card        string   ` json:"card" valid:"Required;MaxSize(18);ErrorCode(50091)"`     //身份证号 固定20
	Phone       string   ` json:"phone" valid:"Required;MaxSize(11);ErrorCode(50092)"`    //手机   固定16
	QQ          string   ` json:"qq" valid:"Required;MaxSize(11);ErrorCode(50093)"`       //qq 固定12
	Email       string   ` json:"email" valid:"Required;MaxSize(50);ErrorCode(50094)"`    //邮箱  50
	ProvinceId  int64    ` json:"provinceId" valid:"Required;Min(1);ErrorCode(50095)"`    //省id
	CityId      int64    ` json:"cityId" valid:"Required;Min(1);ErrorCode(50096)"`        //市id
	AreaId      int64    ` json:"areaId" valid:"Required;Min(1);ErrorCode(50097)"`        //区id
	SpreadId    string   ` json:"spreadId" valid:"Required;MaxSize(20);ErrorCode(50164)"` //推广id
	Remark      string   ` json:"remark" valid:"Required;MinSize(1);ErrorCode(50098)"`    //备注  255
	Ids         string   `json:"ids" valid:"MaxSize(20);ErrorCode(30041)"`                //代理银行id
	BankIds     string   `json:"bankId" valid:"MaxSize(20);ErrorCode(30059)"`             //卡类型
	Cards       []string `json:"cards" valid:"MaxSize(20);ErrorCode(30060)"`              //19卡号
	CardAddress []string `json:"cardAddress" valid:"MaxSize(20);ErrorCode(30062)"`        //50卡开户行
	DoIds       string   `json:"doIds" valid:"MaxSize(20);ErrorCode(30118)"`              //主键id
	Domain      []string `json:"domain" valid:"MaxSize(20);ErrorCode(30068)"`             //推广域名（唯一判断）
}

//获取某个代理基本资料
type ThirdAgencyInfo struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" query:"id" form:"id" valid:"Required;ErrorCode(30041)"`
}
