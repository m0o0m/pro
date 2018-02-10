package back

//站点商品维护列表返回
type SiteProductBack struct {
	SiteId      string `json:"site_id"`       //站点id
	SiteIndexId string `json:"site_index_id"` //站点前台id
	ProductId   int64  `json:"product_id"`    //商品id
	ProductName string `json:"product_name"`  //商品名称
	TypeId      int64  `json:"type_id"`       //商品类型id
	Title       string `json:"title"`         //商品类型名称
	IsCheck     int8   `json:"is_check"`      //是否选中，2未勾选   0勾选
}

type SiteProduct struct {
	Id          int64  `json:"id"`           //商品id
	ProductName string `json:"product_name"` //商品名称
	TypeId      int64  `json:"type_id"`      //商品类型id
	Title       string `json:"title"`        //商品类型名称
}

//站点商品维护列表返回
type SiteProductDelBack struct {
	ProductId int64 `json:"product_id"` //商品id
}
