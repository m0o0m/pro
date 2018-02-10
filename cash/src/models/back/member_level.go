package back

//返回单个层级信息
type MemberLevel struct {
	LevelId      string  `xorm:"level_id" json:"levelId"`
	Description  string  `xorm:"description" json:"description" `
	DepositNum   int64   `xorm:"deposit_num" json:"depositNum" `
	DepositCount float64 `xorm:"deposit_count" json:"depositCount" `
	StartTime    string  `xorm:"start_time" json:"startTime" `
	EndTime      string  `xorm:"end_time" json:"endTime" `
	Remark       string  `xorm:"remark" json:"remark" `
}

//返回层级列表
type MemberLevelList struct {
	SiteIndexId   string  `xorm:"site_index_id" json:"siteIndexId"`     //站点前台Id
	LevelId       string  `xorm:"level_id" json:"levelId"`              //会员层级名称
	Description   string  `xorm:"description" json:"description" `      //层级描述
	PaySetId      int64   `xorm:"pay_set_id" json:"paySetId"`           //支付设置Id
	IsSelfRebate  int8    `xorm:"is_self_rebate" json:"isSelfRebate"`   //是否开启自动返水功能。(1.开启2.未开启)
	IsDefault     int8    `xorm:"is_default" json:"isDefault"`          //是否为默认层级
	Count         int64   `xorm:"count" json:"count"`                   //会员数量
	DepositNum    int64   `xorm:"deposit_num" json:"depositNum" `       //取款次数
	DepositCount  float64 `xorm:"deposit_count" json:"depositCount" `   //取款总额
	StartTime     string  `xorm:"start_time" json:"startTime" `         //会员加入开始时间
	EndTime       string  `xorm:"end_time" json:"endTime" `             //会员加入结束时间
	DepositNumber int64   `xorm:"deposit_number" json:"depositNumber" ` //存款次数
	DepositTotal  float64 `xorm:"deposit_total" json:"depositTotal" `   //存款总额
	Remark        string  `xorm:"remark" json:"remark" `                //备注

}

//查询会员支付设定
type MemberLevelPaySetBack struct {
	SiteId      string `xorm:"'site_id'" json:"siteId"`            //站点Id
	SiteIndexId string `xorm:"'site_index_id'" json:"siteIndexId"` //站点前台Id
	LevelId     string `xorm:"'level_id'"  json:"levelId"`         //会员层级名称
	PaySetId    int64  `xorm:"pay_set_id" json:"paySetId" `        //支付设置Id
}

//返回层级名称下拉框列表
type MemberLevelDrop struct {
	LevelId     string `xorm:"level_id" json:"levelId"`
	Description string `xorm:"description" json:"description"`
}

//返回会员详情列表
type MemberInfoList struct {
	Id            int64   `xorm:"id" json:"id"`            //会员id
	LevelId       string  `xorm:"level_id" json:"levelId"` //层级id
	IsDefault     int8    `xorm:"is_default" json:"isDefault"`
	Account       string  `xorm:"account" json:"account"`   //账号
	Realname      string  `xorm:"realname" json:"realname"` //姓名
	IsLockedLevel int8    `xorm:"is_locked_level" json:"isLockedLevel"`
	LastLoginTime string  `xorm:"last_login_time" json:"lastLoginTime"`
	CreateTime    string  `xorm:"create_time" json:"createTime"`
	LoginIp       string  `xorm:"login_ip" json:"lastLoginIp"`
	LoginCount    int64   `xorm:"login_count" json:"loginCount"`
	Mobile        string  `xorm:"mobile" json:"mobile"`
	Email         string  `xorm:"email" json:"email"`
	Qq            string  `xorm:"qq" json:"qq"`
	DepositNum    float64 `xorm:"deposit_num" json:"depositNum"`
	DepositCount  float64 `xorm:"deposit_count" json:"depositCount"`
	DepositMax    float64 `xorm:"deposit_max" json:"depositMax"`
	DrawNum       int64   `xorm:"draw_num" json:"drawNum"`
	DrawCount     float64 `xorm:"draw_count" json:"drawCount" `
	DrawMax       float64 `xorm:"draw_max"  json:"drawMax"`
	SpreadMoney   float64 `xorm:"spread_money" json:"spreadMoney"`
}

//站点查询层级列表
type SiteLevelList struct {
	SiteId       string  `xorm:"site_id" json:"site_id"`
	SiteIndexId  string  `xorm:"site_index_id" json:"site_index_id"`
	LevelId      string  `xorm:"level_id" json:"level_id"`
	Description  string  `xorm:"description" json:"description" `
	Count        int64   `xorm:"count" json:"count"`
	DepositNum   int64   `xorm:"deposit_num" json:"deposit_num" `
	DepositCount float64 `xorm:"deposit_count" json:"deposit_count" `
	StartTime    string  `xorm:"start_time" json:"start_time" `
	EndTime      string  `xorm:"end_time" json:"end_time" `
	Remark       string  `xorm:"remark" json:"remark" `
}

//返回层级名称下拉框列表
type MemberLevelDrops struct {
	LevelId     string `xorm:"level_id" json:"levelId"`        //层级id
	Description string `xorm:"description" json:"description"` //层级描述
}

//入款商户下拉框列表
type ThirdPaidList struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	PaidTypeName string `json:"paidTypeName"`
}
