package input

//
type Site struct {
	SiteId      string `query:"siteId"`                                        //站点Id
	SiteIndexId string `query:"siteIndexId" valid:"Required;ErrorCode(10050)"` //站点Id
}

type From struct {
	ClientType int8 `json:"client_type" valid:"Required;Range(1,2);ErrorCode(70030)"` //1pc,2wap
}

//电子视讯彩票体育页面参数
type VType struct {
	VType string `query:"type"` //站点Id
}
