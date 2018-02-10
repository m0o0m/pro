package back

//套餐商品返回
type ComboPro struct {
	Id         int64  `json:"id"`                   //套餐id
	ComboName  string `json:"comboName"`            //套餐名称
	SiteCount  int64  `xorm:"num" json:"siteCount"` //站点数量
	CreateTime int64  `json:"createTime"`           //创建时间
	Status     int8   `json:"status"`               //套餐状态
}

//套餐返回
type Combo struct {
	Id         int64  `json:"id"`         //套餐id
	ComboName  string `json:"comboName"`  //套餐名称
	SiteCount  int64  `json:"siteCount"`  //站点数量
	CreateTime int64  `json:"createTime"` //创建时间
	Status     int8   `json:"status"`     //套餐状态
}

//套餐详情返回
type ComboInfo struct {
	Id        int64  `json:"id"`        //套餐id
	ComboName string `json:"comboName"` //套餐名称
}

//套餐下拉框
type ComboDrop struct {
	Id        int64  `xorm:"id" json:"id"`                 //套餐id
	ComboName string `xorm:"combo_name" json:"combo_name"` //套餐名称
}
