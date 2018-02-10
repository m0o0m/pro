package back

//线路检测列表
type InfoDomain struct {
	Id          int64  `json:"id"`            //线路检测id
	SiteId      string `json:"site_id"`       //站点id
	SiteIndexId string `json:"site_index_id"` //站点前台id
	Domain      string `json:"domain"`        //域名
	Status      int8   `json:"status"`        //状态 1启用 2停用
}
