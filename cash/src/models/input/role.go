package input

//添加角色
type RoleAdd struct {
	RoleName string `json:"roleName" valid:"Required;MaxSize(45);ErrorCode(30070)"` //中文名
	Remark   string `json:"remark" valid:"MaxSize(255);ErrorCode(30072)"`           //备注
	Status   int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"`    //状态
}

//查看角色详情
type RoleId struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(30077)"` //角色id
}

//角色状态修改
type RoleStatus struct {
	Id     int64 `json:"id" valid:"Required;ErrorCode(30041)"`       //账号id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //状态
}

//角色列表
type PromissRoleId struct {
	Id    int64  `query:"id" valid:"Required;ErrorCode(30077)"` //角色id
	Type  int8   `query:"type" valid:"ErrorCode(50113)"`
	Typed string `query:"typed" valid:"ErrorCode(50113)"`
}

//修改角色
type RoleEdit struct {
	Id       int64  `json:"id" valid:"Required;Min(1);ErrorCode(30077)"`            //角色id
	RoleName string `json:"roleName" valid:"Required;MaxSize(45);ErrorCode(30070)"` //中文名
	Remark   string `json:"remark"`                                                 //备注
	Status   int8   `json:"status" valid:"Range(0,2);ErrorCode(30050)"`             //状态
}

//修改角色
type RoleEditNew struct {
	Id       int64  `json:"id" valid:"Required;Min(1);ErrorCode(30077)"`            //角色id
	RoleName string `json:"roleName" valid:"Required;MaxSize(45);ErrorCode(30070)"` //中文名
	Remark   string `json:"remark"`                                                 //备注
}

//角色权限
type RolePermission struct {
	RoleId       int64   `json:"roleId"  valid:"Required;ErrorCode(30077)"`  //角色id
	Status       int8    `json:"status"`                                     //角色状态
	RoleName     string  `json:"roleName" valid:"Required;ErrorCode(30116)"` //角色名称
	PermissionId []int64 `json:"permissionId" valid:"ErrorCode(30048)"`      //权限id
}

//设置角色菜单
type RoleMenu struct {
	RoleId int64   `json:"roleId" valid:"Required;ErrorCode(30077)"` //角色id
	MenuId []int64 `json:"menuId" valid:"Required;ErrorCode(50057)"` //菜单id
}
