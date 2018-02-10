package schema

import "global"

//菜单表
type Menu struct {
	Id          int64  `xorm:"id"`
	MenuName    string `xorm:"menu_name"`             //菜单名称
	Route       string `xorm:"route"`                 //菜单路由
	ParentId    int64  `xorm:"parent_id"`             //父级id
	Sort        int64  `xorm:"sort"`                  //菜单排序
	Icon        string `xorm:"icon"`                  //菜单icon
	CreateTime  int64  `xorm:"'create_time' created"` //创建时间
	Status      int8   `xorm:"status"`                //状态
	Level       int8   `xorm:"level"`                 //菜单等级
	LanguageKey string `xorm:"language_key"`          //前端国际化标识
	Type        string `xorm:"type"`                  //菜单类型(agency代理,admin平台)
	DeleteTime  int64  `xorm:"delete_time"`           //软删除时间 0.表示未删除
}

func (*Menu) TableName() string {
	return global.TablePrefix + "menu"
}
