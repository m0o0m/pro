//子帐号管理[admin]
package back

//子帐号列表
type ChildAccountListBack struct {
	Id            int64  `xorm:"id" json:"id"`                         //主键id
	SiteId        string `xorm:"site_id" json:"siteId"`                //站点id
	SiteIndexId   string `xorm:"site_index_id" json:"siteIndexId"`     //站点前台id
	Username      string `xorm:"username" json:"username"`             //代理名称
	Account       string `xorm:"account" json:"account"`               //登录账号
	LastLoginIp   string `xorm:"last_login_ip" json:"lastLoginIp"`     //上次登录IP
	LastLoginTime int64  `xorm:"last_login_time" json:"lastLoginTime"` //上次登录时间
	LoginErrCount int64  `xorm:"login_err_count" json:"loginErrCount"` //登录错误次数
	CreateTime    int64  `xorm:"create_time" json:"createTime"`        //创建时间
	Status        int8   `xorm:"status" json:"status"`                 //状态 1正常2禁用
	IsLogin       int8   `xorm:"is_login" json:"isLogin"`              //是否在线
}

//子帐号信息
type OneChildInfo struct {
	SiteId   string `xorm:"site_id" json:"siteId"`    //子帐号站点id
	Account  string `xorm:"account" json:"account"`   //帐号
	Username string `xorm:"username" json:"username"` //用户名
}

//站点下拉框
type SiteSiteIndexBack struct {
	Id        string `xorm:"id" json:"id"`                //主键id
	IndexId   string `xorm:"index_id" json:"indexId"`     //前台id
	SiteName  string `xorm:"site_name" json:"siteName"`   //站点名称
	IsDefault int8   `xorm:"is_default" json:"isDefault"` //是否默认站点
}

//子站点下拉框
type SiteIndexId struct {
	IndexId string `json:"indexId"` //前台id
}
