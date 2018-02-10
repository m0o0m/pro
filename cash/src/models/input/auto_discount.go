package input

//DiscountList 自助优惠申请列表查询条件struct
type DiscountList struct {
	SiteId       string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId  string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Status       int8   `query:"status" valid:"Range(0,3);ErrorCode(60036)"`      //状态
	ApplyAccount string `query:"applyAccount"`                                    //账号
	StartDate    string `query:"startDate"`                                       //开始时间
	EndDate      string `query:"endDate"`                                         //结束时间
}

//SelfDiscountSwitch 自助优惠开关列表
type SelfDiscountSwitch struct {
	SiteId  string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"` //站点id
	IndexId string `query:"siteIndexId"`
}

//AutoDiscountStatus 自主优惠设定修改
type AutoDiscountStatus struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Status      int8   `json:"status" valid:"Range(1,2);ErrorCode(30050)"`               //状态
}

//AutoDiscountStatus 自主优惠申请提交
type SelfHelpApllyAdd struct {
	//SiteId      string  `json:"site_id" valid:"MaxSize(4);ErrorCode(60105)"`       //站点id
	//SiteIndexId string  `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Account     string  `json:"account"`     //申请者账号
	Applyreason string  `json:"applyreason"` //申请原因
	ApplyMoney  float64 `json:"apply_money"` //申请金额
	Code        string  `json:"code"`        //验证码
	//MemberId    int64   `json:"member_id"`                                         //会员id
	ProId      int64  `json:"pro_id"`      //id
	ProTitle   string `json:"pro_title"`   //标题
	ProContent string `json:"pro_content"` //内容
}

//AutoDiscountStatus wap优惠申请提交
type ApllySubmit struct {
	Account     string  `json:"account"`     //申请者账号
	Applyreason string  `json:"applyreason"` //申请原因
	ApplyMoney  float64 `json:"apply_money"` //申请金额
	ProId       int64   `json:"pro_id"`      //id
	ProTitle    string  `json:"pro_title"`   //标题
	ProContent  string  `json:"pro_content"` //内容
}

//AutoDiscountInfo 获取自助优惠详情
type AutoDiscountInfo struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(50013)"` //自助优惠申请id
}

//拒绝一条优惠申请
type RefuseApplyFor struct {
	Id     int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`         //优惠申请主键id
	Reason string `json:"reason" valid:"Required;MinSize(1);ErrorCode(60116)"` //拒绝原因
}

//通过一条优惠申请
type ThroughApplyFor struct {
	Id            int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`
	ThroughRemark string `json:"through_remark" valid:"Required;MinSize(1);ErrorCode()"` //通过备注
}
