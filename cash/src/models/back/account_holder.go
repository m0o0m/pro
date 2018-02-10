package back

//开户人列表
type AccountHolderListBack struct {
	Id           int64  `json:"id"`            //开户人id
	SiteId       string `json:"site_id"`       //站点id
	UserName     string `json:"user_name"`     //开户人名称
	Account      string `json:"account"`       //帐号
	SiteNumber   int64  `json:"site_number"`   //站点数量
	FirstNumber  int64  `json:"first_number"`  //股东数量
	SecondNumber int64  `json:"second_number"` //总代数量
	ThirdNumber  int64  `json:"third_number"`  //代理数量
	MemberNumber int64  `json:"member_number"` //会员数量
	CreateTime   int64  `json:"create_time"`   //新增时间
	IsLogin      int8   `json:"is_login"`      //是否在线
	Status       int8   `json:"status"`        //状态
}

//开户人信息
type AccountHolderInfo struct {
	Id          int64  `xorm:"'id' PK autoincr" json:"id"`       //主键id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	ParentId    int64  `xorm:"parent_id" json:"parentId"`        //所属上级
	Account     string `xorm:"account" json:"account"`           //登录账号
	Username    string `xorm:"username" json:"username"`         //代理名称
	Remark      string `xorm:"remark" json:"remark"`             //备注
	Status      int8   `xorm:"status" json:"status"`             //状态 1正常2禁用
}

//站点会员注册优惠
type HolderRegBack struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string  `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	AgencyId    int64   `xorm:"agency_id" json:"agency_id"`         //代理id
	Offer       float64 `xorm:"offer" json:"offer"`                 //加入会员赠送优惠金额
	AddMosaic   int64   `xorm:"add_mosaic" json:"add_mosaic"`       //优惠打码倍数
	IsIp        int8    `xorm:"is_ip" json:"is_ip"`                 //是否限制IP 1:是2:否
}

//开户人连表查询返回表
type HoldersBacks struct {
	Id                int64  `xorm:"id" json:"id"`                  //开户人id
	IsLogin           int8   `xorm:"is_login" json:"isLogin"`       //是否在线
	Username          string `xorm:"username" json:"userName"`      //开户人名称
	Account           string `xorm:"account" json:"account"`        //帐号
	CreateTime        int64  `xorm:"create_time" json:"createTime"` //新增时间
	Status            int8   `xorm:"status" json:"status"`          //状态
	ComboId           int64  `xorm:"combo_id" json:"comboId"`       //套餐id
	ComboName         string `xorm:"combo_name" json:"comboName"`   //套餐名称
	FirstAgencyCount  int64  `xorm:"A" json:"firstNumber"`          //股东个数
	SecondAgencyCount int64  `xorm:"B" json:"secondNumber"`         //总代个数
	ThirdAgencyCount  int64  `xorm:"C" json:"thirdNumber"`          //代理个数
	MemberCount       int64  `xorm:"D" json:"memberNumber"`         //会员个数
	SiteCount         int64  `xorm:"E" json:"siteNumber"`           //站点数
	Domain            string `xorm:"domain" json:"domain"`          //客户后台域名
}

//开户人在线人数以及总数
type OnlineNumberAndTotal struct {
	OnlineNumber int64 `json:"onlineNumber"` //在线人数
	TotalNumber  int64 `json:"totalNumber"`  //总人数
}

//会员在线人数、总数、启用数
type OnlineNumberAndTotalAndStatus struct {
	OnlineNumber int64 `json:"onlineNumber"` //在线人数
	TotalNumber  int64 `json:"totalNumber"`  //总人数
	StatusNumber int64 `json:"statusNumber"` //启用人数
}
