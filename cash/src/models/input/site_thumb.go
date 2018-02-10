package input

//附件列表
type GetSiteThumb struct {
	SiteId      string `query:"siteId"`      //站点id
	SiteIndexId string `query:"siteIndexId"` //站点前台id
	State       int    `query:"status"`      //1可用 2禁用
}

type SiteThumbEdit struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(30041)"`                       //主键id
	SiteId      string `json:"siteId"`                                                   //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	FileName    string `json:"fileName" valid:"Required;MaxSize(50);ErrorCode(10001)"`   //文件名称
	State       int8   `json:"status"`                                                   //1可用 2禁用
}

type SiteThumbDelete struct {
	Id          int64  `json:"id"`          //主键id
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
}
