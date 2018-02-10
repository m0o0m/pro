package back

//返回单个商品分类信息
type ProductType struct {
	Id     int64  `xorm:"id" json:"id"`
	Title  string `xorm:"title" json:"title"`
	Status int8   `xorm:"status" json:"status"`
}

//返回商品列表信息
type ProductTypeInfo struct {
	Id     int64  `xorm:"id" json:"id"`
	Title  string `xorm:"title" json:"title"`
	Status int8   `xorm:"status" json:"status"`
	Count  int64  `xorm:"- count" json:"count"`
}

//返回商品分类下拉框列表
type ProductTypeList struct {
	Id    int64  `xorm:"id" json:"id"`
	Title string `xorm:"title" json:"title"`
}

//返回商品分类列表
type ProductTypes struct {
	Id           int64  `xorm:"id" json:"id"`
	Title        string `xorm:"title" json:"title"`
	Status       int8   `xorm:"status" json:"status"`
	ProductCount int64  `json:"product_count"`
	TypeId       int64  `xorm:"type_id" json:"type_id"`
}
