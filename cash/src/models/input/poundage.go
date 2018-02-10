package input

//手续费设定
type Poundage struct {
	SiteId              string  `json:"site_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       //操作站点id
	SiteIndexId         string  `json:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	IncomePoundageRatio float64 `json:"income_poundage_ratio" valid:"Required;Min(1);Max(4);ErrorCode(90301)"` //入款手续费比例
	IncomePoundageUp    int     `json:"income_poundage_up" valid:"Required;Min(1);Max(4);ErrorCode(90302)"`    //入款手续费上限
	OutPoundageRatio    float64 `json:"out_poundage_ratio" valid:"Required;Min(1);Max(4);ErrorCode(90303)"`    //出款手续费比例
	OutPoundageUp       int     `json:"out_poundage_up" valid:"Required;Min(1);Max(4);ErrorCode(90304)"`       //出款手续费上限
	IsDeliveryModel     int8    `json:"is_delivery_model" valid:"Required;Min(1);Max(4);ErrorCode(90305)"`     //是否开启交收模式
	UpdateTime          int64   `json:"update_time"`                                                           //修改时间
}
type PoundAdd struct {
	SiteId              string  `json:"site_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(60105)"`       //操作站点id
	SiteIndexId         string  `json:"site_index_id" valid:"Required;MinSize(1);MaxSize(4);ErrorCode(10050)"` //站点前台id
	IncomePoundageRatio float64 `json:"income_poundage_ratio" valid:"Required;Min(1);Max(4);ErrorCode(90301)"` //入款手续费比例
	IncomePoundageUp    int     `json:"income_poundage_up" valid:"Required;Min(1);Max(4);ErrorCode(90302)"`    //入款手续费上限
	OutPoundageRatio    float64 `json:"out_poundage_ratio" valid:"Required;Min(1);Max(4);ErrorCode(90303)"`    //出款手续费比例
	OutPoundageUp       int     `json:"out_poundage_up" valid:"Required;Min(1);Max(4);ErrorCode(90304)"`       //出款手续费上限
	IsDeliveryModel     int8    `json:"is_delivery_model" valid:"Required;Min(1);Max(4);ErrorCode(90305)"`     //是否开启交收模式
}
