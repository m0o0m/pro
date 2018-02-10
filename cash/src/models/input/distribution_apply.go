package input

//代理注册设定(添加或修改)
type AgentSetDo struct {
	SiteId                string //站点id
	SiteIndexId           string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`           //站点前台id
	RegisterProxy         int8   `json:"registerProxy"  valid:"Range(1,2);ErrorCode(30017)"`        //是否启用代理注册(1.启用 2.不启用)
	ChineseNickname       int8   `json:"chineseNickname" valid:"Range(1,2);ErrorCode(30018)"`       //是否需要中文昵称(1. 需要 2.不需要)
	EnglishNickname       int8   `json:"englishNickname" valid:"Range(1,2);ErrorCode(30019)"`       //是否需要英文昵称(1.需要 2.不需要)
	NeedCard              int8   `json:"needCard" valid:"Range(1,2);ErrorCode(30020)"`              //是否需要证件号(1.需要 2.不需要)
	NeedEmail             int8   `json:"needEmail" valid:"Range(1,2);ErrorCode(30021)"`             //是否需要邮箱(1.需要 2.不需要)
	NeedQq                int8   `json:"needQq" valid:"Range(1,2);ErrorCode(30022)"`                //是否需要qq(1.需要 2.不需要)
	NeedPhone             int8   `json:"needPhone" valid:"Range(1,2);ErrorCode(30023)"`             //是否需要手机号(1.需要 2.不需要)
	PromoteWebsite        int8   `json:"promoteWebsite" valid:"Range(1,2);ErrorCode(30024)"`        //推广网址(1.需要 2.不需要)
	OtherMethod           int8   `json:"otherMethod" valid:"Range(1,2);ErrorCode(30025)"`           //其他方式(1.需要 2.不需要)
	IsMustChineseNickname int8   `json:"isMustChineseNickname" valid:"Range(1,2);ErrorCode(30026)"` //中文昵称是否必填(1.必须填2.非必填)
	IsMustEnglishNickname int8   `json:"isMustEnglishNickname" valid:"Range(1,2);ErrorCode(30032)"` //英文昵称是否必填(1.必须填2.非必填)
	IsMustEmail           int8   `json:"isMustEmail" valid:"Range(1,2);ErrorCode(30027)"`           //邮箱是否必填(1.必须填2.非必填)
	IsMustIdentity        int8   `json:"isMustIdentity" valid:"Range(1,2);ErrorCode(30028)"`        //证件是否必填(1.必须填2.非必填)
	IsMustQq              int8   `json:"isMustQq" valid:"Range(1,2);ErrorCode(30033)"`              //qq是否必填(1.必须填2.非必填)
	IsMustPhone           int8   `json:"isMustPhone" valid:"Range(1,2);ErrorCode(30029)"`           //手机号是否必填(1.必须填2.非必填)
	IsMustPromoteWebsite  int8   `json:"isMustPromoteWebsite" valid:"Range(1,2);ErrorCode(30030)"`  //推广网址是否必填(1.必须填2.非必填)
	IsMustMethod          int8   `json:"isMustMethod" valid:"Range(1,2);ErrorCode(30031)"`          //其他方式是否必填(1.必须填2.非必填)
}

//代理注册设定(查看)
type AgentSet struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//删除代理申请
type AgentRegState struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`                                           //站点前台id
	RegisterId  int64  `json:"registerId" query:"register_id" form:"register_id" valid:"Required;Min(1);ErrorCode(30002)"` //代理申请id
}

//审核代理申请
type AgentRegEdit struct {
	RegisterId      int64  `json:"registerId" valid:"Required;Min(1);ErrorCode(30002)"` //代理申请id
	SiteId          string //站点id
	SiteIndexId     string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`   //站点前台id
	ParentId        int64  `json:"parentId" valid:"Required;Min(1);ErrorCode(30008)"` //上级id
	Username        string `json:"username" valid:"Required;ErrorCode(10027)"`
	Account         string `json:"account" valid:"Required;MinSize(4);MaxSize(12);ErrorCode(30009)"` //账号
	Password        string `json:"password" valid:"MinSize(4);MaxSize(12);ErrorCode(30010)"`         //密码
	ConfirmPassword string `json:"confirmPassword" valid:"MaxSize(12);ErrorCode(30011)"`             //确认密码
}

//代理申请列表
type AgencyIndex struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Status      int8   `query:"status" valid:"Range(0,2);ErrorCode(30007)"`      //状态
	Key         string `query:"key"`                                             //下拉框搜索条件
	Value       string `query:"value"`                                           //下拉框搜索值
}

//获取一条代理注册申请设定
type OneAgencyReg struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `query:"id" valid:"Required;ErrorCode(50013)"`
}
