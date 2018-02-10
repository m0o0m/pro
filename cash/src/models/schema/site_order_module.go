package schema

import "global"

//站点模块排序表
type SiteOrderModule struct {
	Id          int64  `xorm:"'id' PK autoincr"`
	SiteId      string `xorm:"site_id"`       // 站点id
	SiteIndexId string `xorm:"site_index_id"` // 站点前台id
	VideoModule string `xorm:"video_module"`  //视讯模块
	FcModule    string `xorm:"fc_module"`     //彩票模块
	DzModule    string `xorm:"dz_module"`     //电子模块
	SpModule    string `xorm:"sp_module"`     //体育模块
	Module      string `xorm:"module"`        //四大模块，电子，视讯，彩票，体育排序
}

func (*SiteOrderModule) TableName() string {
	return global.TablePrefix + "site_order_module"
}
