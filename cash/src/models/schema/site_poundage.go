package schema

import "global"

//代理手续费设定
type SitePoundage struct {
	SiteId              string  `xorm:"site_id"`               //操作站点id
	SiteIndexId         string  `xorm:"site_index_id"`         //站点前台id
	IncomePoundageRatio float64 `xorm:"income_poundage_ratio"` //入款手续费比例
	IncomePoundageUp    int     `xorm:"income_poundage_up"`    //入款手续费上限
	OutPoundageRatio    float64 `xorm:"out_poundage_ratio"`    //出款手续费比例
	OutPoundageUp       int     `xorm:"out_poundage_up"`       //出款手续费上限
	IsDeliveryModel     int8    `xorm:"is_delivery_model"`     //是否开启交收模式
	UpdateTime          int64   `xorm:"update_time"`           //修改时间
}

func (*SitePoundage) TableName() string {
	return global.TablePrefix + "site_poundage"
}
