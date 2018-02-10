package back

//股东、总代、代理列表 （返回前端）
type ThirdAgencyBack struct {
	AgencyId      int64  `xorm:"'agency_id' PK" json:"agencyId"`   // 代理id(agency表主键)
	Username      string `xorm:"username" json:"username"`         //代理名称
	IsLogin       int8   `xorm:"is_login" json:"isLogin"`          //登录时存的key
	Status        int8   `xorm:"status" json:"status"`             //状态 1正常2禁用
	SiteIndexId   string `xorm:"site_index_id" json:"siteIndexId"` // 所属前台id
	MemberCount   int64  `xorm:"member_count" json:"memberCount"`  // 推广会员数量
	FirstId       int64  `xorm:"first_id" json:"firstId"`          //所属股东id
	SecondId      int64  `xorm:"second_id" json:"seconId"`         // 所属总代理id
	FirstAccount  string `json:"firstAccount"`                     //所属股东帐号
	SecondAccount string `json:"secondAccount"`                    //所属总代帐号
	Account       string `json:"account"`                          //登录账号
	CreateTime    int64  `json:"createTime"`                       //创建时间
}

//代理数量，在线人数，启用人数，禁用人数
type AgentNumberPerson struct {
	TotalNum  int64 `json:"totalNum"`  //总人数
	OnlineNum int64 `json:"onlineNum"` //在线人数
	OpenNum   int64 `json:"openNum"`   //启用人数
	CloseNum  int64 `json:"closeNum"`  //禁用人数
}

//查询股东对会员注册优惠设定
type ThirdDiscountSetBack struct {
	SiteIndexId string  `xorm:"'site_index_id'" json:"siteIndexId"` //站点前台id
	AgencyId    int64   `xorm:"'agency_id'" json:"agencyId"`        //代理id
	Offer       float64 `xorm:"offer" json:"offer"`                 //加入会员赠送优惠金额
	AddMosaic   int64   `xorm:"add_mosaic" json:"addMosaic"`        //优惠打码倍数
	IsIp        int8    `xorm:"is_ip" json:"isIp"`                  //是否限制IP 1:是2:否
}

//某个site_index_id下的所有
type SecondIdNameBack struct {
	Id       int64  `xorm:"id" json:"id"`             //主键id
	Account  string `xorm:"account" json:"account"`   //帐号
	Username string `xorm:"username" json:"username"` //代理名称
}

//third_agency  获取详细资料
type ThirdInformationBack struct {
	AgencyId   int64  `xorm:"'agency_id'" json:"agencyId"`   //代理id
	ChName     string `xorm:"ch_name" json:"chName"`         //中文名称  50
	UsName     string `xorm:"us_name" json:"usName"`         //英文名称  50
	Card       string `xorm:"card" json:"card"`              //身份证号 固定20
	Phone      string `xorm:"phone" json:"phone"`            //手机   固定16
	QQ         string `xorm:"qq" json:"qq"`                  //qq 固定12
	Email      string `xorm:"email" json:"email"`            //邮箱  50
	ProvinceId int64  `xorm:"province_id" json:"provinceId"` //省id
	CityId     int64  `xorm:"city_id" json:"cityId"`         //市id
	AreaId     int64  `xorm:"area_id" json:"areaId"`         //区id
	SpreadId   string `xorm:"spread_id" json:"spreadId"`     //推广id
	Remark     string `xorm:"remark" json:"remark"`          //备注  255
}

//代理银行信息
type ThirdAgencyInfoByBank struct {
	Id          int64  `xorm:"id" json:"id"`                    //主键id
	AgencyId    int64  `xorm:"agency_id" json:"agencyId"`       //代理id
	BankId      int64  `xorm:"bank_id" json:"bankId"`           //卡类型
	Card        string `xorm:"card" json:"card"`                //19卡号
	CardAddress string `xorm:"card_address" json:"cardAddress"` //50卡开户行
}

//代理域名
type ThirdAgencyInfoByDomain struct {
	Id       int64  `xorm:"'id' PK autoincr"` //主键id
	AgencyId int64  `xorm:"agency_id"`        //代理id
	Domain   string `xorm:"domain"`           //推广域名（唯一判断）
}

//查询某个股东基本资料
type ThirdAgencyInfo struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`  //账号
	Username string `json:"username"` //名称
}

//代理管理页面展示出当前所有代理的数量，启用的数量，停用的数量
type AgencyNumber struct {
	AllNumber    int64 `json:"allNumber"`
	EnableNumber int64 `json:"enableNumber"`
	UnableNumber int64 `json:"unableNumber"`
}
