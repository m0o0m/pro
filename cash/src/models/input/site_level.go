package input

//添加会员层级
type SiteLevelAdd struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	LevelId     string `json:"levelId" valid:"Required;ErrorCode(30196)"`                //层级id
	Description string `json:"description" valid:"Required;MaxSize(50)ErrorCode(30197)"` //层级描述
}

//会员层级详情
type SiteLevelInfo struct {
	SiteId      string `query:"siteId" json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`           //站点id
	SiteIndexId string `query:"siteIndexId" json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	LevelId     string `query:"levelId" json:"levelId" valid:"Required;ErrorCode(30196)"`                    //层级id
}

//站点会员层级列表
type SiteLevelList struct {
	SiteId string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
}

//修改层级
type SiteLevelEdit struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	LevelId     string `json:"levelId" valid:"Required;ErrorCode(30196)"`                //层级id
	Description string `json:"description" valid:"Required;MaxSize(50)ErrorCode(30197)"` //层级描述
}

//添加站点层级
type AddSiteLevel struct {
	Lid       int64                  `json:"lid"`       //层级编号
	LevelName string                 `json:"levelName"` //层级名字
	Talk      string                 `json:"talk"`      //描述
	Remark    string                 `json:"remark"`    //备注
	Platforms []SiteLevelProductRate `json:"platforms"` //站点对应各平台占成比
}

//修改站点层级
type EditSiteLevel struct {
	Id        int64                  `json:"id"`        //层级id
	LevelName string                 `json:"levelName"` //层级名字
	Talk      string                 `json:"talk"`      //描述
	Remark    string                 `json:"remark"`    //备注
	Platforms []SiteLevelProductRate `json:"platforms"` //站点对应各平台占成比
}

//删除站点层级
type DelSiteLevel struct {
	Id int64 `json:"id" query:"id"` //层级id
}

//移动站点层级
type MoveSiteLevel struct {
	OldId  int64  `json:"oldId"`  //老层级id
	NewId  int64  `json:"newId"`  //新层级id
	SiteId string `json:"siteId"` //站点id
}

//站点层级平台占成比
type SiteLevelProductRate struct {
	PlatformId int64   `json:"platformId"` //平台id
	Proportion float64 `json:"proportion"` //占成比
}

//获取层级设定数据
type ListSiteLevel struct {
	Id        int64                  `json:"id"`         //层级id
	Lid       int64                  `json:"lid"`        //层级编号
	LevelName string                 `json:"level_name"` //层级名字
	SiteLevel string                 `json:"site_level"` //包含站点
	Talk      string                 `json:"talk"`       //描述
	Remark    string                 `json:"remark"`     //备注
	Platforms []SiteLevelProductRate `json:"platforms"`  //站点对应各平台占成比
}

//站点列表(以站点搜索)
type SiteList struct {
	SiteId string `json:"site_id" query:"siteId"` //站点id
}

//站点列表下拉框
type SiteListDrop struct {
	SiteId string `json:"site_id" query:"siteId" valid:"Required;MaxSize(4);ErrorCode(30041)"` //站点id
}
