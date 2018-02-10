package back

//返回稽核日志列表
type AuditLogList struct {
	Id          int64  `xorm:"id" json:"id"`                     //会员id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //前台id
	MemberId    int64  `xorm:"member_id" json:"memberId"`        //会员id
	Account     string `xorm:"account" json:"account"`           //账号
	Content     string `xorm:"content" json:"content"`           //日志详细内容
	UpdateDate  string `xorm:"update_date" json:"updateDate"`    //更新时间
	Status      int    `xorm:"type" json:"type"`                 //清除状态
}

type MemberAuditNow struct {
	Id            int64   `xorm:"id" json:"id"`                        //会员id
	SiteId        string  `xorm:"site_id" json:"siteId"`               //站点id
	SiteIndexId   string  `xorm:"site_index_id" json:"siteIndexId"`    //前台id
	MemberId      int64   `xorm:"member_id" json:"memberId"`           //会员id
	Account       string  `xorm:"account" json:"account"`              //账号
	BeginTime     int64   `xorm:"begin_time" json:"beginTime"`         //稽核开始时间
	EndTime       int64   `xorm:"end_time" json:"endTime"`             //稽核结束时间
	NormalMoney   float64 `xorm:"normal_money" json:"normalMoney"`     //常态稽核
	MultipleMoney float64 `xorm:"multiple_money" json:"multipleMoney"` //综合稽核
	Money         float64 `xorm:"money" json:"money"`                  //存款金额
	AdminMoney    float64 `xorm:"admin_money" json:"adminMoney"`       //扣除的行政费用
	DepositMoney  float64 `xorm:"deposit_money" json:"depositMoney"`   //优惠金额
	RelaxMoney    int64   `xorm:"relax_money" json:"relaxMoney"`       //放宽额度
	Status        int     `xorm:"status" json:"status"`                //状态
}

//查询稽核期间的打码情况(统计表)
type AuditBet struct {
	Id          int64   `xorm:"id" json:"id"`                     //会员id
	SiteId      string  `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"` //前台id
	Account     string  `xorm:"account" json:"account"`           //账号
	BetValid    float64 `xorm:"bet_valid" json:"betValid"`        //有效打码
	GameType    int     `xorm:"game_type" json:"gameType"`        //游戏类型
}

//查询稽核期间的打码情况(注单表)
type AuditBetRecord struct {
	SiteId      string  `xorm:"site_id" json:"site_id"`     //站点id
	SiteIndexId string  `xorm:"index_id" json:"index_id"`   //前台id
	Username    string  `xorm:"username" json:"username"`   //账号
	BetValid    float64 `xorm:"bet_valid" json:"bet_valid"` //有效打码
	GameType    int     `xorm:"game_type" json:"game_type"` //游戏类型
}

//即时稽核返回详情
type MemberAuditNowBack struct {
	BeginTime     int64   `xorm:"begin_time" json:"beginTime"`         //稽核开始时间
	EndTime       int64   `xorm:"end_time" json:"endTime"`             //稽核结束时间
	NormalMoney   float64 `xorm:"normal_money" json:"normalMoney"`     //常态稽核
	MultipleMoney float64 `xorm:"multiple_money" json:"multipleMoney"` //综合稽核
	Money         float64 `xorm:"money" json:"money"`                  //存款金额
	AdminMoney    float64 `xorm:"admin_money" json:"adminMoney"`       //扣除的行政费用
	DepositMoney  float64 `xorm:"deposit_money" json:"depositMoney"`   //优惠金额
	RelaxMoney    int64   `xorm:"relax_money" json:"relaxMoney"`       //放宽额度
	Vdbet         float64 `xorm:"vd_bet"  json:"vdBet"`                //视讯打码
	Dzbet         float64 `xorm:"dz_bet"  json:"dzBet"`                //电子打码
	Fcbet         float64 `xorm:"fc_bet"  json:"fcBet"`                //彩票打码
	Spbet         float64 `xorm:"sp_bet"  json:"spBet"`                //体育打码
	Allbet        float64 `xorm:"all_bet" json:"allBet"`               //总打码
	IsNormal      int     `xorm:"is_normal" json:"isNormal"`           //常态稽核是否通过  1通过  2不通过
	IsMultiple    int     `xorm:"is_multiple" json:"isMultiple"`       //综合稽核是否通过  1通过  2不通过
}

//取款各种稽核和费用返回详情
type MemberAuditOutBack struct {
	MemberAuditNowBack
	OutStatus  int64   `json:"outStatus"`  //是否允许出款 0 不允许 1允许
	OutCharge  float64 `json:"outCharge"`  //实际出款金额
	Charge     float64 `json:"charge"`     //手续费
	OutMoney   float64 `json:"outMoney"`   //出款金额
	OrderNum   string  `json:"orderNum"`   //订单号
	ClientType int64   `json:"clientType"` //客户端类型
}

//查询会员是否存在
type AuditMember1 struct {
	SiteId      string `json:"site_id"`       //站点id
	SiteIndexId string `json:"site_index_id"` //站点前台id
	Account     string `json:"account"`       //会员账号
}

//稽核日志列表(后台管理)
type AuditLogAdminList struct {
	Id        int64  `json:"id"`        //会员id
	SiteId    string `json:"siteId"`    //站点id
	Account   string `json:"account"`   //会员账号
	BeginTime int64  `json:"beginTime"` //稽核开始时间
	EndTime   int64  `json:"endTime"`   //稽核更改时间
}

//稽核日志列表(后台管理)
type AuditAdminList struct {
	Id         int64  `json:"id"`         //会员id
	SiteId     string `json:"siteId"`     //站点id
	Account    string `json:"account"`    //会员账号
	Content    string `json:"content"`    //内容
	UpdateDate int64  `json:"updateDate"` //稽核更改时间
}
