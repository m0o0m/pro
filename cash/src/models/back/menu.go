package back

//角色、菜单返回
type MenuBack struct {
	RoleId int64 `xorm:"role_id" json:"roleId"` //角色id
	MenuId int64 `xorm:"menu_id" json:"menuId"` //菜单id'
}

//菜单列表返回
type MenuListBack struct {
	Id          int64  `xorm:"id" json:"id"`
	MenuName    string `xorm:"menu_name" json:"menuName"`       //菜单名称
	Route       string `xorm:"route" json:"route"`              //菜单路由
	LanguageKey string `xorm:"language_key" json:"languageKey"` // 前端国际化标识
	Type        string `xorm:"type" json:"type"`                // 菜单类型(agency代理,admin平台)
	ParentId    int64  `xorm:"parent_id" json:"parentId"`       //父级id
	Sort        int64  `xorm:"sort" json:"sort"`                //菜单排序
	Icon        string `xorm:"icon" json:"icon"`                //菜单icon
	Status      int8   `xorm:"status" json:"status"`            //状态
	Level       int8   `xorn:"level" json:"level"`              //菜单等级
	CreateTime  int64  `xorm:"create_time" json:"createTime"`   //创建时间
}

//菜单返回（下拉框）
type MenuIdNameBack struct {
	Id       int64  `xorm:"id" json:"id"`
	MenuName string `xorm:"menu_name" json:"menuName"` //菜单名称
}

//菜单列表返回
type Trees struct {
	Id          int64  `xorm:"id" json:"id"`
	MenuName    string `xorm:"menu_name" json:"menu_name"`       //菜单名称
	LanguageKey string `xorm:"language_key" json:"language_key"` // 前端国际化标识
	Type        string `xorm:"type" json:"type"`                 // 菜单类型(agency代理,admin平台)
	Route       string `xorm:"route" json:"route"`               //菜单路由
	Sort        int64  `xorm:"sort" json:"sort"`                 //菜单排序
	Icon        string `xorm:"icon" json:"icon"`                 //菜单icon
	Status      int8   `xorm:"status" json:"status"`             //状态
	Level       int8   `xorm:"level" json:"level"`               //菜单等级
	IsMenu      int8   `xorm:"-" json:"is_menu"`
	Children    []Trees
}

type MenuId struct {
	MenuId int64 `json:"menuId"` //菜单id
}
