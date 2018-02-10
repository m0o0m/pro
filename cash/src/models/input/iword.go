package input

type IwordList struct {
	Id int64 `query:"id"` //类型代码
}
type TransferInfo struct {
	SiteId      string `query:"site_id"`       //站点id
	SiteIndexId string `query:"site_index_id"` //站点前台id
	PageName    string `query:"page_name"`     //页面标题
	TypeId      string `query:"type_id"`       //类型id
}
