package back

import "models/schema"

//返回单个商品信息
type ProductInfo struct {
	Id          int64  `xorm:"id" json:"id"`                    //
	ProductName string `xorm:"product_name" json:"productName"` //商品名
	Title       string `xorm:"title" json:"title"`              //类型名称
	TypeId      int64  `xorm:"type_id" json:"typeId"`           //商品类型id
	Api         string `xorm:"api" json:"api"`                  //商品域名
	Status      int8   `xorm:"status" json:"status"`            //状态 1正常  2停用
}

//返回商品列表信息
type ProductList struct {
	Id          int64  `xorm:"id" json:"id"`                    //
	ProductName string `xorm:"product_name" json:"productName"` //商品名
	Title       string `xorm:"title" json:"title"`              //类型名称
	Api         string `xorm:"api" json:"api"`                  //商品域名
	CreateTime  string `xorm:"create_time" json:"createTime"`   //创建时间
	Status      int8   `xorm:"status" json:"status"`            //状态 1正常  2停用
	TypeId      int64  `xorm:"type_id" json:"typeId"`           //商品类型id
}

//商品分类和商品返回
type AllProductClassifyListBack struct {
	Id       int64  `xorm:"id" json:"id"`       //商品分类id
	Title    string `xorm:"title" json:"title"` //商品分类名
	Children []PlatformlistBack
}

//商品返回结构体
type PlatformlistBack struct {
	Id         int64   `xorm:"id" json:"id"`                   //交易平台id
	TypeId     int64   `xorm:"type_id" json:"typeId"`          //商品类型id
	ProductId  int64   `xorm:"product_id" json:"productId"`    //商品类型id
	Platform   string  `xorm:"platform" json:"platform"`       //交易平台名称
	Proportion float64 `xorm:"-" json:"proportion"`            //平台占成比
	IsProduct  int8    `xorm:"-" json:"isProduct"`             //是否在套餐中设置
	PlatformId int64   `xorm:"platform_id" json:"platform_id"` //平台id
}

//返回商品列表(转出/转入项目)
type TypeList struct {
	Id       int64  `xorm:"id" json:"id"`
	Platform string `xorm:"platform" json:"product_name"`
}

//返回商品列表(报表统计)
type ProductlistRep struct {
	Id          int64  `xorm:"id" json:"id"`
	ProductName string `xorm:"product_name" json:"productName"`
	VType       string `xorm:"v_type" json:"vType"`
	PlatformId  int64  `xorm:"platform_id" json:"platform_id"`
}

//返回商品剔除
type SiteProductDel struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` // 站点
	ProductId   int64  `xorm:"product_id" json:"product_id"`       // 商品id
	Count       int64  `xorm:"count" json:"count"`                 // 商品数量
}

type ProductName struct {
	Id          int64  `json:"id"`
	ProductName string `json:"product_name"` //商品名
}

//商品分类和商品返回
type WapProductList struct {
	Id       int64  `xorm:"id" json:"id"` //商品分类id
	Title    string `xorm:"title"`        //商品分类名
	Children []schema.Product
}

//返回商品剔除
type ProductDel struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` // 站点
	ProductId   int64  `xorm:"product_id" json:"product_id"`       // 商品id
}

//商品下拉框
type ProductListDropBack struct {
	Id          int64  `xorm:"id" json:"id"`                    //主键id
	ProductName string `xorm:"product_name" json:"productName"` //商品名 30
	VType       string `xorm:"v_type" json:"vType"`             //游戏类型
}
