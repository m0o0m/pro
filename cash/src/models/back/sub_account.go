package back

//子账号返回列表
type SubAccount struct {
	SiteId        string `json:"siteId"`
	Id            int64  `json:"id"`
	IsLogin       int8   `json:"isLogin"`       //登录时存的key
	Username      string `json:"username"`      //代理名称
	Account       string `json:"account"`       //登录账号
	LastLoginIp   string `json:"lastLoginIp"`   //上次登录IP
	LastLoginTime int64  `json:"lastLoginTime"` //上次登录时间
	CreateTime    int64  `json:"createTime"`    //创建时间
	LoginErrCount int64  `json:"loginErrCount"` //登录错误次数
	Status        int8   `json:"status"`        //状态 1正常2禁用
}

//子账号返回详情
type SubAccountInfo struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`  //账号
	Username string `json:"username"` //名称
}

//子账号权限列表
type SubAgencyPermission struct {
	Account        string `json:"account"`         //账号
	PermissionId   int64  `json:"permission_id"`   //权限id
	Module         string `json:"module"`          //所属模块
	PermissionName string `json:"permission_Name"` //权限名
}

//子账号权限返回列表
type SubAgencyPermissionBack struct {
	Account     string        `json:"account"`        //账号
	Permissions []Permissions `json:"permissionList"` //权限功能
}

//子账号口令验证信息详情
type SubAccessTokenInfo struct {
	SiteId  int64  `xorm:"site_id" json:"site_id"`   //站点id
	Status  int    `xorm:"status" json:"status"`     //状态	1启用 2未启用
	PassKey string `xorm:"pass_key" json:"pass_key"` //密钥
}

//子账号
type SubAcc struct {
	Account string `json:"account"` //账号
}

//
type Key struct {
	Key string `json:"key"` //服务器端生成的密钥
}
