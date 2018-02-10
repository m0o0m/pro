package back

//权限
type Permissions struct {
	Module     string       `json:"module"`     //所属模块
	Permission []Permission `json:"permission"` //权限
}

type Module struct {
	Module string `json:"module"`
}

//权限返回列表
type Permission struct {
	Id             int64  `json:"id"`             //主键id
	PermissionName string `json:"permissionName"` //权限名称
	Route          string `json:"route"`          //权限路由
	Module         string `json:"module"`         //所属模块
	Method         string `json:"method"`         //路由请求方式
	Status         int8   `json:"status"`         //权限状态
	Type           string `json:"type"`
	CreateTime     int64  `json:"createTime"`            //添加时间
	IsPermission   int8   `xorm:"-" json:"isPermission"` //是否拥有权限(1:是    0:否)
}

//权限列表
type Pmn struct {
	Id             int64  `json:"id"`             //主键id
	Module         string `json:"module"`         //所属模块
	PermissionName string `json:"permissionName"` //权限名称
	Route          string `json:"route"`          //权限路由
	Method         string `json:"method"`         //路由请求方式
	Status         int8   `json:"status"`         //权限状态
	Type           string `json:"type"`
	CreateTime     int64  `json:"createTime"`            //添加时间
	IsPermission   int8   `xorm:"-" json:"isPermission"` //是否拥有权限(1:是    2:否)
}

//权限列表返回
type PermissionList struct {
	Id             int64  `json:"id"`             //主键id
	Module         string `json:"module"`         //所属模块名称
	PermissionName string `json:"permissionName"` //权限名称
	Type           string `json:"type"`           //所属平台
	Route          string `json:"route"`          //权限路由
	Method         string `json:"method"`         //路由请求方式
	Status         int8   `json:"status"`         //权限状态
	CreateTime     int64  `json:"createTime"`     //添加时间
}

//站点
type SiteIndexBySiteBack struct {
	Id       string `xorm:"'id' PK"`   //主键id
	IndexId  string `xorm:"index_id"`  //前台id
	SiteName string `xorm:"site_name"` //站点名称
	Status   int8   `xorm:"status"`    //状态(1.正常2.关闭)
}
