package schema

import "global"

//站点线路检测表
type InfoDomain struct {
	Id          int64  `xorm:"'id' PK autoincr"` //文案id
	SiteId      string `xorm:"site_id"`          //站点id
	SiteIndexId string `xorm:"site_index_id"`    //站点前台id
	Domain      string `xorm:"domain"`           //域名
	Status      int8   `xorm:"status"`           //状态 1启用 2停用

}

func (*InfoDomain) TableName() string {
	return global.TablePrefix + "info_domain"
}
