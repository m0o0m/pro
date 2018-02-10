package input

//添加商品分类
type ProductType struct {
	Title  string `json:"title" valid:"MinSize(1);MaxSize(30);ErrorCode(10080)"`
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(10081)"`
}

//获取单个商品分类信息
type ProductTypeInfo struct {
	Id int64 `query:"id" valid:"Required;Min(1);ErrorCode(10084)"`
}

//修改商品分类信息
type ProductTypeEdit struct {
	Id     int64  `json:"id" valid:"Required;Min(1);ErrorCode(10084)"`
	Title  string `json:"title" valid:"MinSize(1);MaxSize(30);ErrorCode(10080)"`
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(10081)"`
}

//禁用/启用商品分类
type ProductTypeStatus struct {
	Id int64 `json:"id" valid:"Required;Min(1);ErrorCode(10084)"`
}

//获取商品分类列表
type ProductTypeList struct {
	SiteId string `query:"site_id"`
	Title  string `query:"title" valid:"MaxSize(30);ErrorCode(10080)"`
}

//根据商品分类获取商品
type ProductById struct {
	SiteId      string `query:"site_id"`
	SiteIndexId string `query:"site_index_id"`
	Id          int64  `query:"id"`
}

//游戏平台查询
type GamePlatform struct {
	Platform string `query:"platform"` //平台名称
	Id       int64  `query:"id"`       //平台id
}
type PlatformAdd struct {
	Platform string `json:"platform"` //平台名称
	Status   int8   `json:"status"`   //状态
}
type PlatformUpdate struct {
	Id       int64  `json:"id"`       //平台id
	Platform string `json:"platform"` //平台名称
	Status   int8   `json:"status"`   //状态
}
type PlatformDelete struct {
	Id       int64  `json:"id"`       //平台id
	Platform string `json:"platform"` //平台名称
}

//获取彩票下游戏列表
type GameList struct {
	GameType string `query:"game_type"` //彩票分类
}
