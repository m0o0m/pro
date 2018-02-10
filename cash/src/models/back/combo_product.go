package back

//套餐商品返回
type ComboProduct struct {
	PlatformId int64   `json:"platform_id"` //交易平台id
	ProductId  int64   `json:"product_id"`  //商品id
	Platform   string  `json:"platform"`    //交易平台名
	Proportion float64 `json:"proportion"`  //商品占成比
}
