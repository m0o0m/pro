package schema

import "global"

//公共支付设定表
type PublicPaySet struct {
	Id                          int64   `xorm:"'id' PK autoincr"`
	CreateTime                  int64   `xorm:"'create_time'created"`           // 创建时间
	Title                       string  `xorm:"title"`                          // 币别
	Code                        string  `xorm:"code"`                           // 代码
	IsFree                      int8    `xorm:"is_free"`                        // 出款是否免手续费
	FreeNum                     int     `xorm:"free_num"`                       // 出款免费次数
	OutCharge                   float64 `xorm:"out_charge"`                     // 出款手续费金额
	OnceQuotaChangeLimmit       float64 `xorm:"once_quota_change_limmit"`       // 单次额度转换下限
	OnlineIsDepositDiscount     int8    `xorm:"online_is_deposit_discount"`     // 线上入款存款优惠 1首次 2每次
	OnlineIsDeposit             int8    `xorm:"online_is_deposit"`              // 可放弃线上入款存款优惠 1是 2否
	OnlineDiscountStandard      float64 `xorm:"online_discount_standard"`       // 线上入款存款优惠标准
	OnlineDiscountPercent       float64 `xorm:"online_discount_percent"`        // 线上入款存款优惠百分比
	OnlineDepositMax            float64 `xorm:"online_deposit_max"`             // 线上入款单次最高存款金額
	OnlineDepositMin            float64 `xorm:"online_deposit_min"`             // 线上入款单次最低存款金額
	OnlineDiscountUp            float64 `xorm:"online_discount_up"`             // 线上入款優惠上限金額
	OnlineOtherDiscountStandard float64 `xorm:"online_other_discount_standard"` //线上入款其他存款优惠标准
	OnlineOtherDiscountPercent  float64 `xorm:"online_other_discount_percent"`  // 线上入款其他存款优惠百分比
	OnlineOtherDiscountUp       float64 `xorm:"online_other_discount_up"`       //线上入款其他優惠上限
	OnlineOtherDiscountUpDay    float64 `xorm:"online_other_discount_up_day"`   // 线上入款其他優惠24小時內最高上限
	OnlineIsMultipleAudit       int8    `xorm:"online_is_multiple_audit"`       // 是否开启线上入款综合额度稽核
	OnlineMultipleAuditTimes    int     `xorm:"online_multiple_audit_times"`    // 线上入款综合额度稽核倍数
	OnlineIsNormalAudit         int8    `xorm:"online_is_normal_audit"`         // 是否开启线上入款常态额度稽核
	OnlineNormalAuditPercent    float64 `xorm:"online_normal_audit_percent"`    // 线上入款常态稽核百分比
	LineIsDepositDiscount       int8    `xorm:"line_is_deposit_discount"`       // 公司入款存款优惠 1首次 2每次
	LineIsDeposit               int8    `xorm:"line_is_deposit"`                // 可放弃公司入款存款优惠 1是 2否
	LineDiscountStandard        float64 `xorm:"line_discount_standard"`         // 公司入款存款优惠标准
	LineDiscountPercent         float64 `xorm:"line_discount_percent"`          // 公司入款存款优惠百分比
	LineDepositMax              float64 `xorm:"line_deposit_max"`               // 公司入款单次最高存款金額
	LineDepositMin              float64 `xorm:"line_deposit_min"`               // 公司入款单次最低存款金額
	LineDiscountUp              float64 `xorm:"line_discount_up"`               // 公司入款優惠上限金額
	LineOtherDiscountStandard   float64 `xorm:"line_other_discount_standard"`   // 公司入款其他存款优惠标准
	LineOtherDiscountPercent    float64 `xorm:"line_other_discount_percent"`    // 公司入款其他存款优惠百分比
	LineOtherDiscountUp         float64 `xorm:"line_other_discount_up"`         // 公司入款其他優惠上限
	LineOtherDiscountUpDay      float64 `xorm:"line_other_discount_up_day"`     // 公司入款其他優惠24小時內最高上限
	LineIsMultipleAudit         int8    `xorm:"line_is_multiple_audit"`         // 是否开启公司入款综合额度稽核
	LineMultipleAuditTimes      int     `xorm:"line_multiple_audit_times"`      // 公司入款综合额度稽核倍数
	LineIsNormalAudit           int8    `xorm:"line_is_normal_audit"`           // 是否开启公司入款常态额度稽核
	LineNormalAuditPercent      float64 `xorm:"line_normal_audit_percent"`      // 公司入款常态稽核百分比
	AuditRelaxQuota             float64 `xorm:"audit_relax_quota"`              //常态稽核放宽额度
	AuditAdminRate              float64 `xorm:"audit_admin_rate"`               // 常态稽核行政费率
}

func (*PublicPaySet) TableName() string {
	return global.TablePrefix + "public_pay_set"
}
