package input

//站点口令列表
type SitePassList struct {
	SiteId   string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"` //站点id
	PageSize int    `query:"pageSize"`
	Page     int    `query:"page"`
}

//站点口令修改
type SitePassUpdate struct {
	SiteId  string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id
	Status  int    `json:"status" valid:"Range(1,2);ErrorCode(60105)"`          //状态
	PassKey string `json:"pass_key" valid:"Required;ErrorCode(30213)"`          //密钥
	Account string `json:"account"`                                             //操作人账号
}

type BatchDelChanges struct {
	SiteId  []string `json:"siteId"`  //要修改的站点slice
	Account string   `json:"account"` //操作人
	Status  int8     `json:"status" valid:"Range(1,2);ErrorCode(60105)"`
}
