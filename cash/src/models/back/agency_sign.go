package back

//登录返回字段
type AgencySign struct {
	Token       string `json:"token"`       //登录token
	Username    string `json:"username"`    //登录用户昵称
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //前台id
}

//平台管理员登录返回字段
type AdminSign struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Status  int8   `json:"status"`
	RoleId  int64  `json:"role_id"`
	Token   string `json:"token"`
}

//会员登录返回字段
type MemberSignBack struct {
	Id                 int64  `json:"id"`
	Account            string `json:"account"`            //会员账号
	Status             int8   `json:"status"`             //状态
	Token              string `json:"token"`              //令牌
	SiteId             string `json:"siteId"`             //站点id
	SiteIndexId        string `json:"siteIndexId"`        //站点前台id
	LevelId            string `json:"levelId"`            //层级id
	Type               string `json:"type"`               //什么前端类型
	ThirdAgencyId      int64  `json:"thirdAgencyId"`      //代理id
	ThirdAgencyAccount string `json:"thirdAgencyAccount"` //代理账号
	MemberId           int64  `json:"memberId"`           //会员id
}
