package input

type ComboProduct struct {
	ProductId  int64   `json:"productId" valid:"Required;Min(1);ErrorCode(10091)"`                  //商品id
	PlatformId int64   `json:"platformId" valid:"Required;Min(1);ErrorCode(30161)"`                 //交易平台id
	Proportion float64 `json:"proportion" valid:"Match(/(?!0\.00)(\d+\.\d{2}$)/);ErrorCode(30120)"` //商品占成比(小数点后保留两位)
}

//配置套餐商品
type ComboProducts struct {
	ComboId int64          `json:"comboId" valid:"Required;Min(1);ErrorCode(30104)"`
	Params  []ComboProduct `json:"params"`
}

//查看套餐商品
type ComboProductId struct {
	ComboId int64 `query:"comboId" valid:"Min(1);ErrorCode(30104)"`
}

type ProductTypeName struct {
	ComboId int64  `query:"comboId" valid:"Required;Min(1);ErrorCode(30104)"` //套餐id
	Name    string `query:"name"`                                             //商品类型名称
}
