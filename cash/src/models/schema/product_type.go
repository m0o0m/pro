package schema

import "global"

//商品类型
type ProductType struct {
	Id         int64  `xorm:"'id' PK autoincr"`      //主键id
	Title      string `xorm:"title"`                 //类型名称
	Status     int8   `xorm:"status"`                //类型状态
	CreateTime int64  `xorm:"'create_time' created"` //添加时间
	DeleteTime int64  `xorm:"delete_time"`           //删除时间
}

func (*ProductType) TableName() string {
	return global.TablePrefix + "product_type"
}
