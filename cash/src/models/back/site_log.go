package back

//操作日志
type SiteDoLog struct {
	Id             int64  `json:"id"`
	OperateTime    int64  `json:"operateTime"`    //日期
	OperateAccount string `json:"operateAccount"` //操作者账号
	OperateInfo    string `json:"operateInfo"`    //操作详情（记录）
	Ip             string `json:"ip"`             //ip
}

//登录日志
type SiteLoginLog struct {
	Id          int64  `json:"id"`
	Domains     string `json:"domains"`     //登录域名
	Device      int8   `json:"device"`      //登录设备
	Account     string `json:"account"`     //登录账号
	LoginResult int8   `json:"loginResult"` //登录是否成功(1.成功2.失败)
	LoginTime   int64  `json:"loginTime"`   //登入时间
	LoginIp     string `json:"loginIp"`     //登录ip
}
