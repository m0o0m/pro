package back

//平台账号列表
type Admin struct {
	Id           int64  `json:"id"`
	Status       int8   `json:"status"`       //状态  1：启用  2：禁用
	RoleName     string `json:"roleName"`     //角色名称
	Account      string `json:"account"`      //账号
	RoleId       int64  `json:"roleId"`       //角色id
	CreateTime   int64  `json:"createTime"`   //新增时间
	LoginIp      string `json:"loginIp"`      //登录ip限制（没有则没限制）
	OnlineStatus int8   `json:"onlineStatus"` //在线状态 1在线 2离线
}

//平台账号详情
type AdminInfo struct {
	Id         int64  `json:"id"`
	Status     int8   `json:"status"`     //状态  1：启用  2：禁用
	RoleId     int64  `json:"roleId"`     //角色id
	RoleName   string `json:"roleName"`   //角色名称
	Account    string `json:"account"`    //账号
	CreateTime int64  `json:"createTime"` //新增时间
	LoginIp    string `json:"loginIp"`    //登录ip限制（多个用逗号隔开）

}
