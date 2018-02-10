package back

//获取套餐列表
type GetList struct {
	ComboId     int64   `json:"comboId"`     //套餐id
	ProductId   int64   `json:"productId"`   //商品id
	Proportion  float64 `json:"proportion"`  //占成比
	ProductName string  `json:"productName"` //商品名
	ComboName   string  `json:"comboName"`   //套餐名字
}

//返回套餐列表
type ReturnPackList struct {
	ComboId   int64     `json:"comboId"`   //套餐id
	ComboName string    `json:"comboName"` //套餐名字
	List      []GetList `json:"list"`      //套餐列表
}
type PackAdd struct {
	Id        int64  `xorm:"id PK autoincr" json:"id"` //套餐id
	ComboName string `json:"comboName"`                //套餐名字
}

//套餐下拉框
type ComboDropBack struct {
	Id        int64  `xorm:"id" json:"id"`                //套餐id
	ComboName string `xorm:"combo_name" json:"comboName"` //套餐名字
}
