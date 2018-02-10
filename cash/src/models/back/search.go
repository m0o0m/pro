package back

type SearchBack struct {
	AgencyId      int64  `xorm:"agency_id" json:"agency_id"`         // 代理id(agency表主键)
	Account       string `xorm:"account" json:"account"`             //登录账号
	SiteIndexId   string `xorm:"site_index_id" json:"site_index_id"` // 所属前台id
	FirstId       int64  `xorm:"first_id" json:"first_id"`           //所属股东id
	SecondId      int64  `xorm:"second_id" json:"second_id"`         // 所属总代理id
	SecondCount   int64  `xorm:"second_count" json:"second_count"`   // 总代理数量
	ThirdCount    int64  `xorm:"third_count" json:"third_count"`     // 代理数量
	MemberCount   int64  `xorm:"member_count" json:"member_count"`   // 推广会员数量
	FirstAccount  string `json:"first_account"`                      //股东帐号
	SecondAccount string `json:"second_account"`                     //总代帐号
	ThirdAccount  string `json:"third_account"`                      //代理帐号
	MemberAccount string `json:"member_account"`                     //会员帐号
	SiteId        string `xorm:"site_id" json:"site_id"`
}
