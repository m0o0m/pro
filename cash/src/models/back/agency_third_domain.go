package back

type AgencyThirdDomain struct {
	Id          int64  `json:"id"` //id
	Account     string `xorm:"account" json:"agencyCount"`
	AgencyId    int64  `json:"agencyId"`    //代理id
	Domain      string `json:"domain"`      //推广域名（唯一判断）
	MemberCount int64  `json:"memberCount"` //会员总数
	CreateTime  int64  `json:"createTime"`  //添加时间
}
