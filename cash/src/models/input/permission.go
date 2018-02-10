package input

//添加权限
type PermissionAdd struct {
	PermissionName string `json:"permissionName" valid:"Required;MaxSize(45);ErrorCode(30085)"` //权限名称
	Module         string `json:"module" valid:"Required;MaxSize(45);ErrorCode(30092)"`         //所属模块
	Route          string `json:"route" valid:"Required;MaxSize(45);ErrorCode(30086)"`          //权限路由
	Method         string `json:"method" valid:"Required;MaxSize(45);ErrorCode(30087)"`         //路由请求方式
	Status         int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"`          //权限状态
	Type           int8   `json:"type" valid:"Required;Range(1,2);ErrorCode(50113)"`
}

//权限id
type PermissionId struct {
	Id int64 `query:"id"  valid:"Required;ErrorCode(30088)"` //权限id
}

//权限状态修改
type ColumnStatus struct {
	Id     int64 `json:"id"  valid:"Required;ErrorCode(30088)"`      //权限id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //状态
}

//权限删除（新）
type ColumnDelete struct {
	Id int64 `json:"id"  valid:"Required;ErrorCode(30088)"` //权限id
}

//修改权限
type PermissionUpdate struct {
	Id             int64  `json:"id"  valid:"Required;ErrorCode(30088)"`                        //权限id
	PermissionName string `json:"permissionName" valid:"Required;MaxSize(45);ErrorCode(30085)"` //权限名称
	Module         string `json:"module" valid:"Required;MaxSize(45);ErrorCode(30092)"`         //所属模块
	Route          string `json:"route" valid:"Required;MaxSize(45);ErrorCode(30086)"`          //权限路由
	Method         string `json:"method" valid:"Required;MaxSize(45);ErrorCode(30087)"`         //路由请求方式
}

//权限列表
type PermissionList struct {
	Type int8 `query:"type" valid:"Required;Range(1,2);ErrorCode(50114)"` //1、agency，2、admin
}
