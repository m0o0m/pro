package input

//返点优惠商品列表
type ListRetreatWaterProduct struct {
	SetId     int64   `json:"set_id"`                                                        //返佣设定id
	ProductId int64   `json:"product_id"`                                                    //商品分类id
	Rate      float64 `json:"rate" valid:"Match(/(?!0\.00)(\d+\.\d{2}$)/);ErrorCode(30161)"` //返点比例
}
