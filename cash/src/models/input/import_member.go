package input

//导入会员结构体
type ImportMember struct {
	SiteId         string  //站点id
	SiteIndexId    string  //站点前台id
	LevelId        string  //层级id
	FirstAgencyId  int64   //所属股东id
	SecondAgencyId int64   //所属总代id
	ThirdAgencyId  int64   //所属代理id
	Account        string  //会员帐号
	UserName       string  //会员真实名称
	Money          float64 //金额
	PayCard        int64   //银行id
	PayNum         string  //银行账号
	Password       string  //密码
}

//站点
type ImportSite struct {
	SiteId      string `query:"siteId" valid:"Required;ErrorCode(30261)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;ErrorCode(60038)"` //站点前台id
}

//代理
type ImportAgency struct {
	SiteId      string `query:"siteId" valid:"Required;ErrorCode(30261)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;ErrorCode(60038)"` //站点前台id
	AgencyId    int64  `query:"agencyId"`                                      //代理id
}
