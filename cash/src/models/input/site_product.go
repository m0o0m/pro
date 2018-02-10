package input

//站点模块管理（站点商品查询）
type SiteProductList struct {
	SiteId      string `query:"site_id" valid:"Required;MaxSize(4);ErrorCode(60105)"`       //站点id
	SiteIndexId string `query:"site_index_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id

}

//站点模块管理（站点商品剔除）
type SiteProductEdit struct {
	SiteId      string  `json:"site_id" valid:"Required;MaxSize(4);ErrorCode(60105)"`       //站点id
	SiteIndexId string  `json:"site_index_id" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	ProductId   []int64 `json:"product_id"`                                                 //商品id
}
