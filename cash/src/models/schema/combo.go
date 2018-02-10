package schema

import "global"

//套餐表
type Combo struct {
	Id         int64  `xorm:"id PK autoincr"`
	ComboName  string `xorm:"combo_name"`            //套餐名称
	Status     int8   `xorm:"status"`                //状态 1正常2禁用
	CreateTime int64  `xorm:"'create_time' created"` //创建时间
	DeleteTime int64  `xorm:"delete_time"`           //软删除时间 0.表示未删除
}

func (*Combo) TableName() string {
	return global.TablePrefix + "combo"
}
