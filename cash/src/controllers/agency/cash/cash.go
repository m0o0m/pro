//[控制器] [代理] 现金系统模块
package cash

import (
	"models/function"
)

var manualAccessBean = new(function.ManualAccessBean)                       //人工存取款
var memberBean = new(function.MemberBean)                                   //会员
var memberBalanceConversionBean = new(function.MemberBalanceConversionBean) //额度转换
var onlineEntryRecordBean = new(function.OnlineEntryRecordBean)             //线上入款
var memberCompanyIncomeBean = new(function.MemberCompanyIncomeBean)         //公司入款
var makeMoneyBean = new(function.MakeMoneyBean)                             //出款管理
var productBean = new(function.ProductBean)                                 //商品
var memberRetreatWaterSetBean = new(function.MemberRetreatWaterSetBean)     //返点优惠设定
var memberRetreatWaterBean = new(function.MemberRetreatWaterBean)           //优惠查询
var memberRetreatWaterSelfBean = new(function.MemberRetreatWaterSelfBean)   //自助返水查询
var betReportAccountBean = new(function.BetReportAccountBean)               //优惠统计
var overRideBean = new(function.OverRideBean)                               //代理退佣设定
var retirementBean = new(function.RetirementBean)                           //退佣查询
var periodsBean = new(function.PeriodsBean)                                 //期数管理
var siteSingleRecordBean = new(function.SiteSingleRecordBean)               //掉单
var poundageBean = new(function.PoundageBean)                               //手续费设定
var bankCardBean = new(function.BankCardBean)                               //银行
var onlinePaidSetupBean = new(function.OnlinePaidSetupBean)                 //第三方支付设定
var paidTypeBean = new(function.PaidTypeBean)                               //支付类型
var paymentBean = new(function.PaymentBean)                                 //
var memberLevelBean = new(function.MemberLevelBean)                         //会员层级
var onlineIncomeThirdBean = new(function.OnlineIncomeThirdBean)             //第三方支付平台
var paymentSetBean = new(function.PaymentSetBean)                           //支付设定
var quotaCountBean = new(function.QuotaCountBean)                           //额度统计
var auditsBean = new(function.AuditsBean)                                   //稽核日志
var rebateCountBean = new(function.RebateCountBean)                         //退佣统计

var (
	redPacketStyleBean = new(function.RedPacketStyleBean)
	redPacketSetBean   = new(function.RedPacketSetBean)
	redPacketLogBean   = new(function.RedPacketLogBean)
)
