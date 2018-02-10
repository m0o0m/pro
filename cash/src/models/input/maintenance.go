package input

import "models/back"

//
type Maintenance struct {
	Id          int64  `json:"id"`           //
	VType       string `json:"v_type"`       //唯一性标示
	ProductName string `json:"product_name"` //种类名称
	State       int64  `json:"state"`        //暂时不用 状态
	Content     string `json:"content"`      //维护内容
	Sids        string `json:"sids"`         //站点 0全部站点
	Rid         int64  `json:"rid"`          //来源id
}

type MaintenanceList struct {
	SiteId      string          `json:"site_id"`       //站点id
	SiteIndexId string          `json:"site_index_id"` //站点前台id
	Wtype       string          `json:"wtype"`         //维护分类 全站 all 单站 one
	Type        string          `json:"type"`          //维护类型 彩票 体育 电子 视讯
	IsState     int64           `json:"is_state"`      //
	Level       int64           `json:"level"`         //优先级
	Content     string          `json:"content"`       //维护内容
	Module      []back.InfoList `json:"module"`        //维护项目列表
}

//获取视讯电子列表
type InfoList struct {
	SiteId      string `query:"site_id"`       //站点id
	SiteIndexId string `query:"site_index_id"` //站点前台id
	Wtype       string `query:"wtype"`         //维护分类 全站 all 单站 one
	Type        string `query:"type"`          //维护类型 彩票 体育 电子 视讯
}

//总后台-站点管理-修改全站维护
type SiteMaintenance struct {
	Id      int64  `json:"id" valid:"Required;ErrorCode(60207)"`
	Content string `json:"content" valid:"Required;ErrorCode(30260)"` //维护内容
	SiteIdS string `json:"siteIdS" valid:"Required;ErrorCode(60207)"` //站点 0全部站点(以站点id_前台站点id的形式存储，多个用逗号隔开)
}

//总后台-站点管理-站点是否被选中
type SiteIsSelect struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(50013)"`
}
