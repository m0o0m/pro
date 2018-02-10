package back

//站点管理-层级列表
type SiteMemberLevelList struct {
	LevelId          string  `json:"levelId"`          //层级
	SiteIndexId      string  `json:"siteIndexId"`      //前台id
	Description      string  `json:"description"`      //描述
	DepositNumber    int64   `json:"depositNumber"`    //存款次数
	DepositTotal     float64 `json:"depositTotal"`     //存款总额
	WithdrawalNumber int64   `json:"withdrawalNumber"` //提款次数
}

//站点管理-层级详情
type SiteMemberLevelInfo struct {
	SiteId      string `json:"site_id"`       //站点id
	LevelId     string `json:"level_id"`      //层级
	SiteIndexId string `json:"site_index_id"` //站点前台id
	Description string `json:"description"`   //描述
}

//层级查询数据
type AllLevelBySite struct {
	ListSiteLevel []ListSiteLevel            `json:"listSiteLevel"`
	PingTai       []SiteLevelProductPlatform `json:"pingTai"`
}

//获取层级设定数据
type ListSiteLevel struct {
	Id        int64                  `json:"id"`        //层级id
	Lid       int64                  `json:"lid"`       //层级编号
	LevelName string                 `json:"levelName"` //层级名字
	SiteLevel string                 `json:"siteLevel"` //包含站点
	Talk      string                 `json:"talk"`      //描述
	Remark    string                 `json:"remark"`    //备注
	Platforms []SiteLevelProductRate `json:"platforms"` //站点对应各平台占成比
}

//获取层级设定数据all
type ListSiteLevelAll struct {
	Id         int64   `json:"id"`          //层级id
	Lid        int64   `json:"lid"`         //层级编号
	LevelName  string  `json:"level_name"`  //层级名字
	SiteLevel  string  `json:"site_level"`  //包含站点
	Talk       string  `json:"talk"`        //描述
	Remark     string  `json:"remark"`      //备注
	PlatformId int64   `json:"platform_id"` //站点平台id
	Platform   string  `json:"platform"`    //平台名称
	Proportion float64 `json:"proportion"`  //占成比
}

//站点层级平台占成比
type SiteLevelProductRate struct {
	PlatformId int64   `json:"platformId"` //平台id
	Platform   string  `json:"platform"`   //平台名称
	Proportion float64 `json:"proportion"` //占成比
}

//单个站点层级详情
type DetailSiteLevel struct {
	Id        int64                  `json:"id"`         //层级id
	Lid       int64                  `json:"lid"`        //层级编号
	LevelName string                 `json:"level_name"` //层级名字
	Talk      string                 `json:"talk"`       //描述
	Remark    string                 `json:"remark"`     //备注
	Platforms []SiteLevelProductRate `json:"platforms"`  //站点对应各平台占成比
}

//站点层级列表(以站点搜索)
type ListSite struct {
	Id        int64                  `json:"id"`        //层级id
	Lid       int64                  `json:"lid"`       //层级编号
	LevelName string                 `json:"levelName"` //层级名字
	SiteId    string                 `json:"siteId"`    //包含站点
	Talk      string                 `json:"talk"`      //描述
	Remark    string                 `json:"remark"`    //备注
	Platforms []SiteLevelProductRate `json:"platforms"` //站点对应各平台占成比
}

//站点层级列表下拉框
type ListSiteDropBack struct {
	Id        int64  `json:"id"`        //层级id
	Lid       int64  `json:"lid"`       //层级编号
	LevelName string `json:"levelName"` //层级名字
}

//CENGJI
type AllLevelBySiteIdBack struct {
	ListSite []ListSite                 `json:"listSite"`
	PingTai  []SiteLevelProductPlatform `json:"pingTai"`
}

//站点层级平台
type SiteLevelProductPlatform struct {
	PlatformId int64  `json:"platformId"` //平台id
	Platform   string `json:"platform"`   //平台名称
}
