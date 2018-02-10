package input

//添加或者修改会员层级信息
type MemberLevel struct {
	SiteId       string  `json:"siteId" `
	SiteIndexId  string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	LevelId      string  `json:"levelId" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)" `
	PaySetId     int64   `json:"paySetId" valid:"Range(1,2147483648);ErrorCode(10004)"`
	Description  string  `json:"description" valid:"MinSize(1);MaxSize(50);ErrorCode(10006)" `
	DepositNum   int64   `json:"depositNum" valid:"Range(0,99999999);ErrorCode(10009)"`
	DepositCount float64 `json:"depositCount" valid:"Match(/(?!0\.00)(\d+\.\d{2}$)/);ErrorCode(10010)"`
	StartTime    string  `json:"startTime"  valid:"Required;ErrorCode(10007)"`
	EndTime      string  `json:"endTime" valid:"Required;ErrorCode(10008)" `
	Remark       string  `json:"remark" valid:"MinSize(1);MaxSize(255)"`
}

//获取某个站点，站点前台,层级名称的信息
type LevelInfoGet struct {
	SiteId      string ` json:"siteId" query:"siteId" `
	SiteIndexId string ` json:"siteIndexId" query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	LevelId     string ` json:"levelId" query:"levelId" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)"`
}

//获取站点,站点前台的层级列表
type LevelIndex struct {
	SiteId      string ` query:"siteId" valid:"MaxSize(4);ErrorCode(10050)"`
	SiteIndexId string ` query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	Remark      string `query:"remark" `     //备注
	Description string `query:"description"` //层级描述
	LevelId     string `query:"levelId"`     //会员层级名称
}

//获取会员层级下拉
type MemberLevels struct {
	SiteId      string `query:"site_id"`
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"`
}

//更新站点，站点前台，层级名称的信息
type MemberLevelUpdate struct {
	SiteId       string  `json:"siteId" `
	SiteIndexId  string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	OldLevelId   string  `json:"oldLevelId" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)" `
	NewLevelId   string  `json:"newLevelId" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)"`
	Description  string  `json:"description" valid:"MinSize(1);MaxSize(50);ErrorCode(10006)" `
	DepositNum   int64   `json:"depositNum" valid:"Range(0,99999999);ErrorCode(10009)"`
	DepositCount float64 `json:"depositCount" valid:"Match(/(?!0\.00)(\d+\.\d{2}$)/);ErrorCode(10010)"`
	StartTime    string  `json:"startTime"  valid:"Required;ErrorCode(10007)"`
	EndTime      string  `json:"endTime" valid:"Required;ErrorCode(10008)" `
	Remark       string  `json:"remark" valid:"MinSize(1);MaxSize(255)"`
}

//修改站点,站点前台,层级名称的自助返水
type MemberLevelSelfRebate struct {
	SiteId       string `json:"siteId" `
	SiteIndexId  string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	LevelId      string `json:"levelId" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)"`
	IsSelfRebate int8   `json:"isSelfRebate" valid:"Range(1,2);ErrorCode(10054)"`
}

//获取层级支付设定
type MemberLevelPaySet struct {
	SiteId      string `query:"siteId"`                                                   //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErroeCode(10050)"` //站点前台id
	LevelId     string `query:"levelId"`                                                  //层级id
}

//修改层级支付设定
type MemberLevelPaySetUpdata struct {
	SiteId      string `json:"siteId"`      //站点Id
	SiteIndexId string `json:"siteIndexId"` //站点前台Id
	LevelId     string `json:"levelId"`     //会员层级名称
	PaySetId    int64  `json:"paySetId"`    //支付设置Id
}

//获取某个站点，站点前台,层级名称会员信息
type LevelMember struct {
	SiteId        string `query:"siteId"`
	SiteIndexId   string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	LevelId       string `query:"levelId" valid:"MaxSize(4);ErrorCode(10005)"`        //层级
	Account       string `query:"account" valid:"MaxSize(20);ErrorCode(20000)"`       //账号
	IsFuzzy       int8   `query:"isFuzzy" valid:"Range(0,1);ErrorCode(30247)"`        //是否模糊查询 1，是
	CreateTime    string `query:"createTime"`                                         //登录时间
	LastLoginIp   string `query:"lastLoginIp"`                                        //最后登录ip
	AccountList   string `query:"accountList"`                                        //批量查询（多个用逗号隔开）
	IsLockedLevel int8   `query:"isLockedLevel"  valid:"Range(0,2);ErrorCode(30248)"` //是否锁定1是  2否
}

//锁定会员层级
type LockMember struct {
	SiteId      string `  query:"siteId" `
	SiteIndexId string `  query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	MemberId    int64  `json:"memberId" valid:"Range(1,99999999);ErrorCode(10020)"`
	Lock        int8   `json:"lock" valid:"Range(1,2);ErrorCode(10055)"`
}

//移动层级
type MoveMemberLevel struct {
	SiteId      string `json:"siteId" `
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`
	MoveIn      string `json:"moveIn" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)"`
	MoveOut     string `json:"moveOut" valid:"MinSize(1);MaxSize(4);ErrorCode(10005)"`
}

type MemberLevelPaySetAndFirstDeposit struct {
	SiteId      string `query:"site_id"`       //站点id
	SiteIndexId string `query:"site_index_id"` //站点前台id
	LevelId     string `query:"level_id"`      //层级id
	Account     string `query:"account"`       //会员账号
}
