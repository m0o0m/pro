package back

//会员基本资料
type MemberSelfInfoBack struct {
	Id            int64  `xorm:"id" json:"id"`
	SiteId        string `xorm:"site_id" json:"siteId"`                //站点Id
	SiteIndexId   string `xorm:"site_index_id" json:"siteIndexId"`     //站点前台Id
	Account       string `xorm:"account" json:"account"`               //登陆账号
	RealName      string `xorm:"realname" json:"realName"`             //真实姓名
	LastLoginTime int64  `xorm:"last_login_time" json:"lastLoginTime"` //上次登录时间
	CreateTime    int64  `xorm:"create_time" json:"createTime"`        //创建时间
	Mobile        string `xorm:"mobile" json:"mobile"`                 //手机号码
	Email         string `xorm:"email" json:"email"`                   //邮箱
	Birthday      int64  `xorm:"birthday" json:"birthday"`             //生日
}

//会员中心主页
type MemberHomePageBack struct {
	Id           int64   `xorm:"id" json:"id"`
	Account      string  `xorm:"account" json:"account"`             //登陆账号
	Realname     string  `xorm:"realname" json:"realname"`           //真实姓名
	Balance      float64 `xorm:"balance" json:"balance"`             //账号余额
	SBalance     float64 `xorm:"a" json:"sBalance"`                  //余额
	IsSelfRebate int8    `xorm:"is_self_rebate" json:"isSelfRebate"` //是否开启自动返水功能。(1.开启2.未开启)
}
