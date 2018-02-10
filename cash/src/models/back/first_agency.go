package back

//股东列表 （返回前端）
type FirstAgencyBack struct {
	AgencyId    int64  `xorm:"agency_id" json:"agencyId"`        // 代理id(agency表主键)
	Username    string `xorm:"username" json:"username"`         //代理名称
	IsLogin     int8   `xorm:"is_login" json:"isLogin"`          //登录时存的key
	Status      int8   `xorm:"status" json:"status"`             //状态 1正常2禁用
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` // 所属前台id
	SecondCount int64  `xorm:"second_count" json:"secondCount"`  // 总代理数量
	ThirdCount  int64  `xorm:"third_count" json:"thirdCount"`    // 代理数量
	MemberCount int64  `xorm:"member_count" json:"memberCount"`  // 推广会员数量
	Account     string `xorm:"account" json:"account"`           //登录账号
	CreateTime  int64  `xorm:"create_time" json:"createTime"`
}

//查询股东对会员注册优惠设定
type FirstDiscountSetBack struct {
	SiteIndexId string  `xorm:"'site_index_id'" json:"site_index_id"` //站点前台id
	AgencyId    int64   `xorm:"'agency_id'" json:"agency_id"`         //代理id
	Offer       float64 `xorm:"offer" json:"offer"`                   //加入会员赠送优惠金额
	AddMosaic   int64   `xorm:"add_mosaic" json:"add_mosaic"`         //优惠打码倍数
	IsIp        int8    `xorm:"is_ip" json:"is_ip"`                   //是否限制IP 1:是2:否
}

//查询某个股东基本资料
type FirstAgencyInfo struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`  //账号
	Username string `json:"username"` //名称
}
