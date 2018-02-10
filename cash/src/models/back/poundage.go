package back

//手续费设定
type Poundage struct {
	SiteId              string  `xorm:"site_id" json:"siteId"`                            //操作站点id
	SiteIndexId         string  `xorm:"site_index_id" json:"siteIndexId"`                 //站点前台id
	IncomePoundageRatio float64 `xorm:"income_poundage_ratio" json:"incomePoundageRatio"` //入款手续费比例
	IncomePoundageUp    int     `xorm:"income_poundage_up" json:"incomePoundageUp"`       //入款手续费上限
	OutPoundageRatio    float64 `xorm:"out_poundage_ratio" json:"outPoundageRatio"`       //出款手续费比例
	OutPoundageUp       int     `xorm:"out_poundage_up" json:"outPoundageUp"`             //出款手续费上限
	IsDeliveryModel     int8    `xorm:"is_delivery_model" json:"isDeliveryModel"`         //是否开启交收模式
	UpdateTime          int64   `xorm:"update_time" json:"updateTime"`                    //修改时间
}
