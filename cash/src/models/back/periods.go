package back

//退佣列表
type Periods struct {
	Id          int64  `xorm:"id" json:"'id'"`                     //id
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	Title       string `xorm:"title" json:"title"`                 // 期数名称
	StartTime   int64  `xorm:"start_time" json:"start_time"`       // 开始时间
	EndTime     int64  `xorm:"end_time" json:"end_time"`           //结束时间
	Status      int    `xorm:"status" json:"status"`               // 退佣状态 0未退佣1已退佣
}
type SiteList struct {
	SiteId      string `xorm:"id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"index_id" json:"site_index_id"` //前台站点id
}
type ComData struct {
	Id          int64  `json:"id" valid:"Required;ErrorCode(60000)"` //期数id
	SiteId      string `json:"site_id"`                              //站点id
	SiteIndexId string `json:"site_index_id" `                       //前台站点id
	Status      int8   `json:"status"`                               //退佣状态
}
