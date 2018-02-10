package input

//增加推广域名
type AgencyThirdDomain struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`     //站点前台id
	AgencyId    int64  `json:"agencyId" valid:"Required;Min(1);ErrorCode(30117)"`   //代理id
	Domain      string `json:"domain" valid:"Required;MinSize(1);ErrorCode(30068)"` //推广域名（唯一判断）
}

//修改推广域名
type AgencyThirdDomainEdit struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(30118)"`
	AgencyId    int64  `json:"agencyId" valid:"Required;Min(1);ErrorCode(30117)"`   //代理id
	Domain      string `json:"domain" valid:"Required;MinSize(1);ErrorCode(30068)"` //推广域名（唯一判断）
}

//删除推广域名
type AgencyThirdDomainDel struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(30118)"`
}

//推广域名列表
type AgencyThirdDomainList struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	AgencyId    int64  `query:"agencyId" `
	Domain      string `query:"domain"`
}
