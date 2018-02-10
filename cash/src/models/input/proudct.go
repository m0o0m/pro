package input

//添加商品
type Product struct {
	ProductName string `json:"productName" valid:"MinSize(1);MaxSize(30);ErrorCode(10080)"` //商品名
	TypeId      int64  `json:"typeId" valid:"Required;Min(1);ErrorCode(10084)"`             //商品类型id
	Api         string `json:"api" valid:"MinSize(1);MaxSize(255);ErrorCode(10087)"`        //商品域名
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(10088)"`                  //状态
	PlatformId  int64  `json:"platformId" valid:"Required"`                                 //平台id
	VType       string `json:"vType" valid:"Required"`                                      //游戏类型
}

//获取单个商品信息
type ProductInfo struct {
	ProductId int64 `query:"productId" json:"productId" valid:"Required;Min(1);ErrorCode(10091)"`
}

//修改商品信息
type ProductEdit struct {
	ProductId   int64  `json:"productId" valid:"Required;Min(1);ErrorCode(10091)"`
	ProductName string `json:"productName" valid:"MinSize(1);MaxSize(30);ErrorCode(10080)"` //商品名
	TypeId      int64  `json:"typeId" valid:"Required;Min(1);ErrorCode(10084)"`             //商品类型id
	Api         string `json:"api" valid:"MinSize(1);MaxSize(30);ErrorCode(10087)"`         //商品域名
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(10088)"`                  //状态
	PlatformId  int64  `json:"platformId" valid:"Required"`                                 //平台id
	VType       string `json:"vType" valid:"Required"`                                      //游戏类型
}

//启用或禁用商品
type ProductStatus struct {
	ProductId int64 `json:"productId" valid:"Required;Min(1);ErrorCode(10091)"`
}

//获取商品列表
type ProductList struct {
	ProductId   int64  `query:"productId"`   //商品id
	Status      int8   `query:"status"`      //状态
	Title       string `query:"title"`       //分类名称
	ProductName string `query:"productName"` //商品名称
}

//商品分类
type ProductTypeId struct {
	TypeId int64 `query:"typeId" valid:"Required;ErrorCode(30180)"` //商品类型
}
