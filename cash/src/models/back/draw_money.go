package back

//会员账号，余额，出款银行卡信息
type WapMemberInfo struct {
	Account string    `json:"account"` //会员账号
	Balance float64   `json:"balance"` //账号余额
	WapBank []WapBank `json:"bank"`    //银行信息
}

//银行信息
type WapBank struct {
	Id    int64  `json:"id"`    //银行id
	Card  string `json:"card"`  //银行卡号
	Title string `json:"title"` //银行名称
}

//出款稽核
type WapMemberAudit struct {
	BeginTime     int64   `json:"begin_time"`     //稽核开始时间
	NormalMoney   float64 `json:"normal_money"`   //常态稽核金额
	MultipleMoney float64 `json:"multiple_money"` //综合稽核金额
	AdminMoney    float64 `xorm:"admin_money"`    //扣除的行政费用
	DepositMoney  float64 `xorm:"deposit_money"`  //优惠金额
	RelaxMoney    int64   `xorm:"relax_money"`    //放宽额度
	Money         float64 `json:"money"`          //存款金额
}

//有效打码
type WapBetValid struct {
	BetValid float64 `json:"bet_valid"` //有效打码
}

//取款进度
type WapDrawalProgress struct {
	OutwardNum      float64 `json:"outwardNum"`      //出款金额
	FavourableMoney float64 `json:"favourableMoney"` //优惠金额
	ExpeneseMoney   float64 `json:"expeneseMoney"`   //行政费
	OutwardMoney    float64 `json:"outwardMoney"`    //实际出款金额
	OutStatus       int     `json:"outStatus"`       //出款状态5待审核1已出款2预备出款3取消出款4拒绝出款
	CreateTime      int64   `json:"createTime"`      //提交时间
	OutTime         int64   `json:"outTime"`         //出款时间
	BeginTime       int64   `json:"beginTime"`       //存款开始时间
	EndTime         int64   `json:"endTime"`         //存款结束时间(也是稽核完成时间)
	TradeNo         string  `json:"tradeNo"`         //订单号
	Balance         float64 `json:"balance"`         //投注金额
}

//单条出款银行信息
type WapOutBank struct {
	Id       int64  `json:"id"`        //银行id
	Card     string `json:"card"`      //银行卡号
	CardName string `json:"card_name"` //出款人姓名
}

//出款稽核整合数据
type AddOutData struct {
	BeginTime    int64   `json:"begin_time"`    //存款日期
	EndTime      int64   `json:"end_time"`      //存款日期
	Vdbet        float64 `json:"bet_valid"`     //有效打码
	Dzbet        float64 `json:"bet_valid"`     //有效打码
	Spbet        float64 `json:"bet_valid"`     //有效打码
	Fcbet        float64 `json:"bet_valid"`     //有效打码
	Allbet       float64 `json:"bet_valid"`     //有效打码
	IsNormal     int64   `json:"normal_status"` //通过常态稽核 1为通过，2为未通过，未通过扣取行政费用
	NormalStatus int64   `json:"normal_status"` //常态稽核是否达到
	DepositMoney float64 `json:"deposit_money"` //优惠稽核
	MuiStatus    int64   `json:"mui_status"`    //综合稽核是否达到
	AdminMoney   float64 `json:"admin_money"`   //扣除行政费用 综合稽核
	OutStatus    int64   `json:"out_status"`    //是否允许出款 1 允许 2不允许提出金额减去费用小于0）
	RelaxMoney   int64   `json:"relax_money"`   //放宽额度
	OutCharge    float64 `json:"out_charge"`    //实际出款金额
	Charge       float64 `json:"charge"`        //手续费
	Money        float64 `json:"money"`         //提出金额
	OutMoney     float64 `json:"out_money"`     //出款金额
	OrderNum     string  `json:"order_num"`     //订单号
	IsFirst      int8    `json:"is_first"`      //是否第一次出款
	ClientType   int64   `json:"client_type"`   //出款客户端
}

//出款稽核整合数据
type ShowOutData struct {
	CreateTime   int64          `json:"create_time"`   //存款日期
	BetValid     WapBetValid    `json:"bet_valid"`     //有效打码
	MemberAudit  WapMemberAudit `json:"member_audit"`  //出款稽核
	NormalStatus int64          `json:"normal_status"` //常态稽核是否达到
	MuiStatus    int64          `json:"mui_status"`    //综合稽核是否达到
	OutStatus    int64          `json:"out_status"`    //是否允许出款 0 不允许 1允许
	AdminMoney   float64        `json:"admin_money"`   //扣除行政费用 综合稽核
	DepositMoney float64        `json:"deposit_money"` //优惠稽核
	OutCharge    float64        `json:"out_charge"`    //实际出款金额
	Charge       float64        `json:"charge"`        //手续费
	OutMoney     float64        `json:"out_money"`     //出款金额
	OrderNum     string         `json:"order_num"`     //订单号
	IsFirst      int8           `json:"is_first"`      //是否第一次出款
	ClientType   int64
	//OrderNum     string         `json:"order_num"`     //订单号
}

//出款现金表存入数据

//出款表存入数据
type SaveMakeMoney struct {
	SiteId          string  `xorm:"site_id" json:"site_id"`                     //操作站点id
	SiteIndexId     string  `xorm:"site_index_id" json:"site_index_id"`         //站点前台id
	MemberId        int64   `xorm:"member_id" json:"member_id"`                 //会员id
	UserName        string  `xorm:"user_name" json:"user_name"`                 //会员账号
	LevelId         string  `xorm:"level_id" json:"level_id"`                   //会员所属层级
	AgencyId        int64   `xorm:"agency_id" json:"agency_id"`                 //会员所属代理id
	AgencyAccount   string  `xorm:"agency_account" json:"agency_account"`       //会员所属代理账号
	IsFirst         int8    `xorm:"is_first" json:"is_first"`                   //是否首次出款,0不是1是
	OutwardNum      float64 `xorm:"outward_num" json:"outward_num"`             //提出金额
	Charge          float64 `xorm:"charge" json:"charge"`                       //手续费
	FavourableMoney float64 `xorm:"favourable_money" json:"favourable_money"`   //优惠金额
	ExpeneseMoney   float64 `xorm:"expenese_money" json:"expenese_money"`       //行政费
	DoAgencyId      float64 `xorm:"do_agency_id" json:"do_agency_id"`           //操作人id
	DoAgencyAccount string  `xorm:"do_agency_account" json:"do_agency_account"` //操作人账号
	OutwardMoney    float64 `xorm:"outward_money" json:"outward_money"`         //实际出款金额
	Balance         float64 `xorm:"balance" json:"balance"`                     //账户余额
	FavourableOut   int     `xorm:"favourable_out" json:"favourable_out"`       //是否优惠扣除0不是1是
	OutStatus       int     `xorm:"out_status" json:"out_status"`               //出款状态1已出款2预备出款3取消出款4拒绝出款5待审核
	Remark          string  `xorm:"remark" json:"remark"`                       //会员备注
	CreateTime      int64   `xorm:"create_time" json:"create_time created"`     //提出时间
	DoUrl           string  `xorm:"do_url" json:"do_url"`                       //会员提交网址
	ClientType      int64   `xorm:"client_type" json:"client_type"`             //客户端类型1pc 2wap 3android 4ios
	OutTime         int64   `xorm:"out_time" json:"out_time"`                   //出款时间
	OrderNum        string  `xorm:"order_num"`                                  // 订单号
}
