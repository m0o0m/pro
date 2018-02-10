package back

//AutoDiscountBack 自助优惠返回struct
type AutoDiscountBack struct {
	Id             int64   `xorm:"id" json:"id"`
	Account        string  `xorm:"account" json:"account"`              //用户名
	Createtime     int64   `xorm:"createtime" json:"applyTime"`         //申请时间
	Updatetime     int64   `xorm:"updatetime" json:"auditTime"`         //审核时间
	PromotionTitle string  `xorm:"promotion_title" json:"activityType"` //活动类别
	Applyreason    string  `xorm:"applyreason" json:"applyReason"`      //申请理由
	ApplyMoney     float64 `xorm:"apply_money" json:"applyMoney"`       //申请金额
	GiveMoney      float64 `xorm:"give_money" json:"giveawayMoney"`     //赠送金额
	Status         int8    `xorm:"status" json:"status"`                //申请状态
}

//AutoDiscountSwitch 自助优惠开关列表返回struct
type AutoDiscountSwitch struct {
	SiteId            string `xorm:"id" json:"siteId"`                          //站点id
	IndexId           string `xorm:"index_id" json:"siteIndexId"`               //前台id
	SiteName          string `xorm:"site_name" json:"siteName"`                 //站点名称
	AutoDicountSwitch int8   `xorm:"self_help_switch" json:"autoDicountSwitch"` //自助优惠开关
}

//自助优惠申请详情
type MemberAutoApplypro struct {
	Id               int64   `json:"id"`
	PromotionTitle   string  `json:"promotionTitle"`   //活动标题
	PromotionContent string  `json:"promotionContent"` //活动内容
	Account          string  `json:"account"`          //用户名
	ApplyMoney       float64 `json:"applyMoney"`       //申请金额
	GiveMoney        float64 `json:"giveMoney"`        //入款金额
	Createtime       int64   `json:"createtime"`       //申请时间
	Updatetime       int64   `json:"updatetime"`       //审核时间
	HandlerName      string  `json:"handlerName"`      //操作者名称
	Status           int8    `json:"status"`           //1 待审核 2 审核通过 3审核不通过
	Applyreason      string  `json:"applyreason"`      //申请理由
	Denyreason       string  `json:"denyreason"`       //拒绝理由
	AgreeRemark      string  `json:"agreeRemark"`      //通过审核备注
	IsNormality      int8    `json:"isNormality"`      //1:不参加常态稽核,2:参加常态稽核
	IsComplex        int8    `json:"isComplex"`        //1:不参加综合打码稽核,2:参加综合打码稽核
	ComplexAudit     int64   `json:"complexAudit"`     //综合打码量
}
