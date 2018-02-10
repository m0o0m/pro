package back

//角色列表
type Role struct {
	Id         int64  `json:"id"`
	RoleName   string `json:"roleName"`   //中文名
	RoleMark   string `json:"roleMark"`   //英文名
	IsOperate  int8   `json:"isOperate"`  //是否可以删除和禁用
	Remark     string `json:"remark"`     //备注
	Status     int8   `json:"status"`     //状态
	CreateTime int64  `json:"createTime"` //创建时间
}

//角色权限列表
type RolePermission struct {
	RoleName       string `json:"roleName"`       //角色名
	Status         int8   `json:"status"`         //状态
	IsOperate      int8   `json:"isOperate"`      //是否可以操作禁用和删除
	PermissionId   int64  `json:"permissionId"`   //权限id
	PermissionName string `json:"permissionName"` //权限名
	Type           string `json:"type"`           //所属平台
}

//角色权限返回列表
type RolePermissionBack struct {
	RoleName    string        `json:"roleName"`  //角色名
	Status      int8          `json:"status"`    //状态
	IsOperate   int8          `json:"isOperate"` //是否可以操作禁用和删除
	Permissions []Permissions `json:"ps"`        //权限功能
}

//平台账号中的角色下拉框
type RoleList struct {
	Id       int64  `json:"id"`
	RoleName string `json:"roleName"` //中文名
}

//角色菜单返回列表
type RoleMenus struct {
	RoleName string  `json:"roleName"` //角色名称
	MenuList []Trees `json:"menuList"` //菜单列表
}

//子账号方法和路由
type RouteList struct {
	Route  string `xorm:"route"`  //路由
	Method string `xorm:"method"` //方法
}
