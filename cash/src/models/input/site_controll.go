//站点管理列表[admin]
package input

//站点列表
type SiteManageList struct {
	Id            int64  `query:"id"`            //开户人Id
	Status        int8   `query:"status"`        //状态
	SiteId        string `query:"siteId"`        //站点id
	MoreCondition int8   `query:"moreCondition"` //多条件
	MoreContent   string `query:"moreContent"`   //多条件内容
	SiteName      string `query:"siteName" `     //站点名称
	SiteDomain    string `query:"siteDomain" `   //主域名
	//ComboId         int64 `query:"comboId" valid:"Min(0);ErrorCode(30104)"`              //套餐id
}

//站点状态修改
type SiteManageStatus struct {
	Status      int8   `json:"status" valid:"Range(0,4);ErrorCode(30050)"`               //状态
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//站点域名编辑
type SiteDomainEdit struct {
	Id          int64   `json:"id"  valid:"Min(0);ErrorCode(30041)"`                               //域名配置id
	SiteId      string  `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`               //站点id
	SiteIndexId string  `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`          //站点前台id
	SiteName    string  `json:"siteName" valid:"Required;MinSize(1);MaxSize(30);ErrorCode(60009)"` //站点名称
	Domain      string  `json:"domain" valid:"Required;ErrorCode(60112)"`                          //主域名
	IsDownApp   int8    `json:"isDownApp" valid:"Required;Range(1,2);ErrorCode(60114)"`            //是否可以下载app
	UpCose      float64 `json:"upCose"`                                                            //收费金额
	DomainUp    int     `json:"domainUp"`                                                          //域名上限
}

//站点详情
type SiteDomainInfo struct {
	Site      string `query:"site" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndex string `query:"siteIndex" valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台id
}
