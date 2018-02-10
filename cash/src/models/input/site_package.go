package input

//获取套餐列表
type GetList struct {
	DeleteTime int64 `query:"deleteTime"` //删除时间
	Status     int8  `query:"status" `    //状态
}

//新增套餐
type PackAdd struct {
	ComboName string           `json:"comboName" valid:"Required;MinSize(1);MaxSize(20);ErrorCode(30114)"` //套餐名字
	List      []AddProductList `json:"list"`                                                               //新增套餐商品列表
}

//新增套餐商品列表
type AddProductList struct {
	ComboId    int64   `json:"comboId" valid:"Min(1);Required;ErrorCode(30104)"`     //套餐id
	ProductId  int64   `json:"productId"  valid:"Min(1);Required;ErrorCode(10091)"`  //商品id
	Proportion float64 `json:"proportion"  valid:"Min(1);Required;ErrorCode(30120)"` //占成比
	PlatformId int8    `json:"platformId"  valid:"Min(1);Required;ErrorCode(30161)"` //平台id
}

//修改套餐
type PackUpdate struct {
	Id        int64            `json:"id" valid:"Min(1);Required;ErrorCode(30104)"`                        //套餐id
	ComboName string           `json:"comboName" valid:"Required;MinSize(1);MaxSize(20);ErrorCode(30114)"` //套餐名字
	List      []AddProductList `json:"list"`                                                               //新增套餐商品列表
}

//获取套餐详情
type GetPackage struct {
	Id int64 `query:"id"` //套餐id
}
