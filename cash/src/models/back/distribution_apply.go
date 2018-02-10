package back

//代理申请返回列表
type AgentIndex struct {
	Id          int    `json:"id"`
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	Account     string `json:"account"`     //帐号
	Email       string `json:"email"`       //邮箱
	Qq          string `json:"qq"`          //QQ
	Skype       string `json:"skype"`       //网络电话
	Wechat      string `json:"wechat"`      //微信
	Status      int8   `json:"status"`      //申请状态 1已添加账号2未处理
	CreateTime  int64  `json:"createTime"`  //申请时间
	ZhName      string `json:"zhName"`      //中文昵称
}

//代理注册设定返回
type SiteAgencyRegisterSet struct {
	SiteId                string `json:"siteId"`                //站点id
	SiteIndexId           string `json:"siteIndexId"`           //站点前台id
	RegisterProxy         int8   `json:"registerProxy"`         //是否启用代理注册(1.启用 2.不启用)
	ChineseNickname       int8   `json:"chineseNickname"`       //是否需要中文昵称(1.需要 2.不需要)
	EnglishNickname       int8   `json:"englishNickname"`       //是否需要英文昵称(1.需要 2.不需要)
	NeedCard              int8   `json:"needCard"`              //是否需要证件号(1.需要 2.不需要)
	NeedEmail             int8   `json:"needEmail"`             //是否需要邮箱(1.需要 2.不需要)
	NeedQq                int8   `json:"needQq"`                //是否需要qq(1.需要 2.不需要)
	NeedPhone             int8   `json:"needPhone"`             //是否需要手机号(1.需要 2.不需要)
	PromoteWebsite        int8   `json:"promoteWebsite"`        //推广网址(1.需要 2.不需要)
	OtherMethod           int8   `json:"otherMethod"`           //其他方式(1.需要 2.不需要)
	IsMustChineseNickname int8   `json:"isMustChineseNickname"` //中文昵称是否必填(1.必须填2.非必填)
	IsMustEnglishNickname int8   `json:"isMustEnglishNickname"` //英文昵称是否必填(1.必须填2.非必填)
	IsMustEmail           int8   `json:"isMustEmail"`           //邮箱是否必填(1.必须填2.非必填)
	IsMustIdentity        int8   `json:"isMustIdentity"`        //证件是否必填(1.必须填2.非必填)
	IsMustQq              int8   `json:"isMustQq"`              //qq是否必填(1.必须填2.非必填)
	IsMustPhone           int8   `json:"isMustPhone"`           //手机号是否必填(1.必须填2.非必填)
	IsMustPromoteWebsite  int8   `json:"isMustPromoteWebsite"`  //推广网址是否必填(1.必须填2.非必填)
	IsMustMethod          int8   `json:"isMustMethod"`          //其他方式是否必填(1.必须填2.非必填)
}

//代理注册页面数据
type AgencyForm struct {
	Name   string `json:"name"`   //
	Notice string `json:"notice"` //
	Title  string `json:"title"`  //
}
