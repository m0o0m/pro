package input

//查询一条公共支付设定参数(get)
type OnePublicPaySet struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId"`
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
}

//增加一条公司设置的支付参数设定(post)
type PaymentSetAdd struct {
	SiteId      string `json:"siteId"`
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	Title       string `json:"title" valid:"Required;MaxSize(20);ErrorCode(50110)"`
}

//公司自设定支付列表
type PaymentSetList struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId"`
}

//公司支付设定设置
type PaymentSetUp struct {
	SiteId                      string  `json:"siteId"`
	SiteIndexId                 string  `json:"siteIndexId"`
	Id                          int64   `json:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	IsFree                      int8    `json:"isFree" valid:"Required;Range(0,2);ErrorCode(50014)"`                            // 出款是否免手续费
	FreeNum                     int     `json:"freeNum" valid:"Required;Min(0);ErrorCode(50016)"`                               // 出款免费次数
	OutCharge                   float64 `json:"outCharge" valid:"Required;MinFloat64(0.00);ErrorCode(50017)"`                   // 出款手续费金额
	OnceQuotaChangeLimmit       float64 `json:"onceQuotaChangeLimmit" valid:"Required;MinFloat64(0.00);ErrorCode(50015)"`       // 单次额度转换下限
	OnlineIsDepositDiscount     int8    `json:"onlineIsDepositDiscount" valid:"Required;Range(1,2);ErrorCode(50018)"`           // 线上入款存款优惠 1首次 2每次
	OnlineIsDeposit             int8    `json:"onlineIsDeposit" valid:"Required;Range(1,2);ErrorCode(50112)"`                   // 可放弃线上入款存款优惠 1是 2否
	OnlineDiscountStandard      float64 `json:"onlineDiscountStandard" valid:"Required;MinFloat64(0.00);ErrorCode(50019)"`      // 线上入款存款优惠标准
	OnlineDiscountPercent       float64 `json:"onlineDiscountPercent" valid:"Required;MinFloat64(0.00);ErrorCode(50020)"`       // 线上入款存款优惠百分比
	OnlineDepositMax            float64 `json:"onlineDepositMax" valid:"Required;MinFloat64(0.00);ErrorCode(50021)"`            // 线上入款单次最高存款金額
	OnlineDepositMin            float64 `json:"onlineDepositMin" valid:"Required;MinFloat64(0.00);ErrorCode(50022)"`            // 线上入款单次最低存款金額
	OnlineDiscountUp            float64 `json:"onlineDiscountUp" valid:"Required;MinFloat64(0.00);ErrorCode(50023)"`            // 线上入款優惠上限金額
	OnlineOtherDiscountStandard float64 `json:"onlineOtherDiscountStandard" valid:"Required;MinFloat64(0.00);ErrorCode(50024)"` //线上入款其他存款优惠标准
	OnlineOtherDiscountPercent  float64 `json:"onlineOtherDiscountPercent" valid:"Required;MinFloat64(0.00);ErrorCode(50025)"`  // 线上入款其他存款优惠百分比
	OnlineOtherDiscountUp       float64 `json:"onlineOtherDiscountUp" valid:"Required;MinFloat64(0.00);ErrorCode(50026)"`       //线上入款其他優惠上限
	OnlineOtherDiscountUpDay    float64 `json:"onlineOtherDiscountUpDay" valid:"Required;MinFloat64(0.00);ErrorCode(50027)"`    // 线上入款其他優惠24小時內最高上限
	OnlineIsMultipleAudit       int8    `json:"onlineIsMultipleAudit" valid:"Required;Min(1);ErrorCode(50028)"`                 // 是否开启线上入款综合额度稽核
	OnlineMultipleAuditTimes    int     `json:"onlineMultipleAuditTimes" valid:"Required;Min(0);ErrorCode(50029)"`              // 线上入款综合额度稽核倍数
	OnlineIsNormalAudit         int8    `json:"onlineIsNormalAudit" valid:"Required;Range(1,2);ErrorCode(50030)"`               // 是否开启线上入款常态额度稽核
	OnlineNormalAuditPercent    float64 `json:"onlineNormalAuditPercent" valid:"Required;ErrorCode(50031)"`                     // 线上入款常态稽核百分比
	LineIsDepositDiscount       int8    `json:"lineIsDepositDiscount" valid:"Required;Range(1,2);ErrorCode(50032)"`             // 公司入款存款优惠 1首次 2每次
	LineIsDeposit               int8    `json:"lineIsDeposit" valid:"Required;Range(1,2);ErrorCode(50033)"`                     // 可放弃公司入款存款优惠 1是 2否
	LineDiscountStandard        float64 `json:"lineDiscountStandard" valid:"Required;MinFloat64(0.00);ErrorCode(50034)"`        // 公司入款存款优惠标准
	LineDiscountPercent         float64 `json:"lineDiscountPercent" valid:"Required;MinFloat64(0.00);ErrorCode(50035)"`         // 公司入款存款优惠百分比
	LineDepositMax              float64 `json:"lineDepositMax" valid:"Required;MinFloat64(0.00);ErrorCode(50036)"`              // 公司入款单次最高存款金額
	LineDepositMin              float64 `json:"lineDepositMin" valid:"Required;MinFloat64(0.00);ErrorCode(50037)"`              // 公司入款单次最低存款金額
	LineDiscountUp              float64 `json:"lineDiscountUp" valid:"Required;MinFloat64(0.00);ErrorCode(50038)"`              // 公司入款優惠上限金額
	LineOtherDiscountStandard   float64 `json:"lineOtherDiscountStandard" valid:"Required;MinFloat64(0.00);ErrorCode(50039)"`   // 公司入款其他存款优惠标准
	LineOtherDiscountPercent    float64 `json:"lineOtherDiscountPercent" valid:"Required;MinFloat64(0.00);ErrorCode(50040)"`    // 公司入款其他存款优惠百分比
	LineOtherDiscountUp         float64 `json:"lineOtherDiscountUp" valid:"Required;MinFloat64(0.00);ErrorCode(50041)"`         // 公司入款其他優惠上限
	LineOtherDiscountUpDay      float64 `json:"lineOtherDiscountUpDay" valid:"Required;MinFloat64(0.00);ErrorCode(50042)"`      // 公司入款其他優惠24小時內最高上限
	LineIsMultipleAudit         int8    `json:"lineIsMultipleAudit" valid:"Required;Range(1,2);ErrorCode(50043)"`               // 是否开启公司入款综合额度稽核
	LineMultipleAuditTimes      int     `json:"lineMultipleAuditTimes" valid:"Required;Min(1);ErrorCode(50044)"`                // 公司入款综合额度稽核倍数
	LineIsNormalAudit           int8    `json:"lineIsNormalAudit" valid:"Required;Range(1,2);ErrorCode(50045)"`                 // 是否开启公司入款常态额度稽核
	LineNormalAuditPercent      float64 `json:"lineNormalAuditPercent" valid:"Required;MinFloat64(0.00);ErrorCode(50046)"`      // 公司入款常态稽核百分比
	AuditRelaxQuota             float64 `json:"auditRelaxQuota" valid:"Required;MinFloat64(0.00);ErrorCode(50047)"`             //常态稽核放宽额度
	AuditAdminRate              float64 `json:"auditAdminRate" valid:"Required;MinFloat64(0.00);ErrorCode(50048)"`              // 常态稽核行政费率
}

//修改公司支付设定名称
type PaymentChangeName struct {
	SiteId      string `json:"site_id"`
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"`
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(50013)"`
	Title       string `json:"title" valid:"Required;MaxSize(20);ErrorCode(50110)"`
}

//删除公司支付设定
type PaymentDelette struct {
	SiteId      string `json:"siteId"`
	SiteIndexId string `json:"siteIndexId"`
	Id          int64  `json:"id" valid:"Required;Min(1);ErrorCode(50013)"`
}

//查询一条公司支付设定
type PaymentSetOne struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId"`
	Id          int64  `query:"id" valid:"Required;Min(1);ErrorCode(50013)"`
}
