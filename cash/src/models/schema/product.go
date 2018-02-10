package schema

import "global"

//商品
type Product struct {
	Id          int64  `xorm:"'id' PK autoincr"`      //主键id
	ProductName string `xorm:"product_name"`          //商品名 30
	TypeId      int64  `xorm:"type_id"`               //商品类型id
	Status      int8   `xorm:"status"`                //商品状态
	PlatformId  int64  `xorm:"platform_id"`           //平台id
	VType       string `xorm:"v_type"`                //游戏类型
	Api         string `xorm:"api"`                   //商品域名
	CreateTime  int64  `xorm:"'create_time' created"` //添加时间
	DeleteTime  int64  `xorm:"delete_time"`           //删除时间
}

func (*Product) TableName() string {
	return global.TablePrefix + "product"
}
