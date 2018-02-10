package schema

import "global"

//站点维护表
type SiteModule struct {
	Id          int64  `xorm:"'id' PK autoincr"`
	VType       string `xorm:"'v_type' PK notnull"`      //唯一性标示ag  og bbin  mg lebo ct具体商品维护  indexid前台维护  admin后台维护
	ProductName string `xorm:"product_name"`             //种类名称
	State       int64  `xorm:"'state' default(1)"`       //1启用2停用
	Content     string `xorm:"content"`                  //维护内容
	SiteIds     string `xorm:"'site_id_s' default('0')"` //站点0全部站点
	FType       int64  `xorm:"'f_type' PK notnull"`      //来源id 1全部 2pc 3wap 4app
}

func (*SiteModule) TableName() string {
	return global.TablePrefix + "site_module"
}
