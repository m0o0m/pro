package input

//平台账号添加
type AdminAdd struct {
	Account         string `json:"account" valid:"Required;MinSize(4);MaxSize(12);ErrorCode(30009)"`  //账号
	Password        string `json:"password" valid:"Required;MinSize(6);MaxSize(12);ErrorCode(30010)"` //密码
	ConfirmPassword string `json:"confirmPassword" valid:"Required;ErrorCode(30011)"`                 //确认密码
	Status          int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"`               //状态
	RoleId          int64  `json:"roleId" valid:"Required;Min(1);ErrorCode(30077)"`                   //角色id
	LoginIp         string `json:"loginIp"`                                                           //登录ip限制（多个用逗号隔开）
}

//平台账号修改
type AdminEdit struct {
	Id              int64  `json:"id" valid:"Required;ErrorCode(30041)"`
	Password        string `json:"password" valid:"MinSize(6);MaxSize(12);ErrorCode(30010)"` //密码
	ConfirmPassword string `json:"confirmPassword" valid:"ErrorCode(30011)"`                 //确认密码
	Status          int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"`      //状态
	RoleId          int64  `json:"roleId" valid:"Required;Min(1);ErrorCode(30077)"`          //角色id
}

//平台账号修改(新)
type AdminEditNew struct {
	Id              int64  `json:"id" valid:"Required;ErrorCode(30041)"`
	Password        string `json:"password" valid:"Nullable(6,12);ErrorCode(30010)"`    //密码
	ConfirmPassword string `json:"confirmPassword" valid:"ErrorCode(30011)"`            //确认密码
	Status          int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"` //状态
	RoleId          int64  `json:"roleId" valid:"Required;Min(1);ErrorCode(30077)"`     //角色id
	LoginIp         string `json:"loginIp"`                                             //登录ip限制（多个用逗号隔开）
}

type AdminId struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(30041)"` //账号id
}

//帐号状态修改
type AdminStatus struct {
	Id     int64 `json:"id" valid:"Required;ErrorCode(30041)"`       //账号id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //状态
}

//账号列表筛选字段
type AdminList struct {
	Status       int8   `query:"status" valid:"Range(0,2);ErrorCode(30050)"`       //状态
	RoleId       int64  `query:"roleId" valid:"Min(0);ErrorCode(30077)"`           //角色id
	Account      string `query:"account" valid:"MaxSize(20);ErrorCode(50010)"`     //账号
	OnlineStatus int8   `query:"onlineStatus" valid:"Range(0,2);ErrorCode(50083)"` //在线状态
}

//修改密码
type UpdatePassword struct {
	Id             int64  `json:"id" valid:"Required"`              //账号id
	BeforePassword string `json:"before_password" valid:"Required"` //之前的密码
	NewPassword    string `json:"new_password" valid:"Required"`    //新密码
	RepeatPassword string `json:"repeat_password" valid:"Required"` //确认新密码
}

//平台管理员登录
type AdminLogin struct {
	Account  string `json:"account" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"`  //账号
	Password string `json:"password" valid:"MinSize(4);MaxSize(12);ErrorCode(20000)"` //密码
	Code     string `json:"code" valid:"MinSize(4);MaxSize(4);ErrorCode(20021)"`      //验证码
}

type InitPassword struct {
	Id int64 `json:"id" valid:"Required;ErrorCode(30041)"` //账号id
}
