package back

//代理下拉框
type GetAgency struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
}

//收款账号下拉
type GetSetAgency struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Title   string `json:"title"`
}

//代理管理
type AgencyManageListBack struct {
	Id          int64  `xorm:"id" json:"id"`                     //主键id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	Account     string `xorm:"account" json:"account"`           //登录账号
	IsLogin     int8   `xorm:"is_login" json:"isLogin"`          //是否在线
	Username    string `xorm:"username" json:"username"`         //代理名称
	Status      int8   `xorm:"status" json:"status"`             //状态 1正常2禁用
	CreateTime  int64  `xorm:"create_time" json:"createTime"`    //创建时间
	SpreadId    string `xorm:"spread_id" json:"spreadId"`        //推广id
	MemberNum   int64  `xorm:"member_count" json:"memberNum"`    //会员数量
}

//代理占成比
type AgencyOccupationRatioBack struct {
	AgencyId   int64   `xorm:"agency_id" json:"agency_id"`   // 代理id(agency表主键)
	Commission float64 `xorm:"commission" json:"commission"` // 比例,最大100%
	Title      string  `xorm:"title" json:"title"`           //类型名称
}

//站点管理-代理列表
type SiteAgencyList struct {
	Id          int64  `json:"id"`
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	Username    string `json:"userName"`    //代理名称
	Account     string `json:"account"`     //账号
	RoleName    string `json:"roleName"`    //角色名称
	RoleId      string `json:"roleId"`      //角色id
	Remark      string `json:"remark"`      //备注
	CreateTime  string `json:"createTime"`  //创建时间
}

//站点管理-代理详情
type SiteAgencyInfo struct {
	Id          int64  `json:"id"`
	SiteId      string `json:"site_id"`       //站点id
	SiteIndexId string `json:"site_index_id"` //站点前台id
	Username    string `json:"username"`      //代理名称
	ParentId    string `json:"parent_id"`     //父级账号
	Account     string `json:"account"`       //账号
	RoleName    string `json:"role_Name"`     //角色名称
	Remark      string `json:"remark"`        //备注
}

//站点管理-代理详情
type SiteAgencyInfoBack struct {
	Id            int64  `json:"id"`
	SiteId        string `json:"site_id"`        //站点id
	SiteIndexId   string `json:"site_index_id"`  //站点前台id
	Username      string `json:"username"`       //代理名称
	ParentAccount string `json:"parent_account"` //父级账号
	Account       string `json:"account"`        //账号
	RoleName      string `json:"role_name"`      //角色名称
	Remark        string `json:"remark"`         //备注
}

//账号
type SiteAgency struct {
	Id     int64 `json:"id"`
	RoleId int64 `json:"role_id"`
}

//收款人列表信息
type GetSetAccountInfo struct {
	Id           int64   `json:"id"`
	PaidTypeName string  `json:"paid_type_name"` //支付类型
	Title        string  `json:"title"`          //银行名称
	Account      string  `json:"account"`        //收款账号
	OpenBank     string  `json:"open_bank"`      //开户行
	Payee        string  `json:"payee"`          //收款人
	StopBalance  float64 `json:"stop_balance"`   //停用金额
	BankId       int64   `json:"bank_id"`        //银行id
	QrCode       string  `json:"qr_code"`        //二维码图片数据
	PayTypeId    int     `json:"pay_type_id"`    //支付类型id
}

//收款人列表信息
type GetPayeeInfo struct {
	Id           int64   `json:"id"`
	PaidTypeName string  `json:"paid_type_name"` //支付类型
	Title        string  `json:"title"`          //银行名称
	Account      string  `json:"account"`        //收款账号
	OpenBank     string  `json:"open_bank"`      //开户行
	Payee        string  `json:"payee"`          //收款人
	StopBalance  float64 `json:"stop_balance"`   //停用金额
	BankId       int64   `json:"bank_id"`        //银行id
	QrCode       string  `json:"qr_code"`        //二维码图片数据
	PayTypeId    int     `json:"pay_type_id"`    //支付类型id
	Sort         int64   `json:"sort"`           //排序
}
