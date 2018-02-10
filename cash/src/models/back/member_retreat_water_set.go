package back

//返回联表查询结果 返点优惠设定列表+优惠商品列表
type RetreatWaterSetListAndProduct struct {
	Id          int64   `json:"id"`
	ValidMoney  int64   `json:"valid_money"`  //有效总投注
	DiscountUp  int64   `json:"discount_up"`  //优惠上限
	ProductId   int64   `json:"product_id"`   //商品分类id
	ProductName string  `json:"product_name"` //商品名
	Rate        float64 `json:"rate"`         //比例
}

//返回返点优惠设定列表
type RetreatWaterSetList struct {
	ValidMoney int64                     `json:"validMoney"` //有效总投注
	DiscountUp int64                     `json:"discountUp"` //优惠上限
	Params     []RetreatWaterProductList `json:"params"`     //引用商品列表
}

//返点优惠商品列表
type RetreatWaterProductList struct {
	ProductId   int64   `json:"productId"`   //商品分类id
	ProductName string  `json:"productName"` //商品名
	Rate        float64 `json:"rate"`        //比例
}
