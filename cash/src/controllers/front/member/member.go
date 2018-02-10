//[控制器] [前台] 会员个人中心
package member

import (
	"models/function"
	"models/thirdParty"
)

var (
	memberBalanceConversionBean         = new(function.MemberBalanceConversionBean)         //额度转换
	distributionApplyBeen               = new(function.DistributionApplyBeen)               //代理申请
	otherBean                           = new(function.OtherBean)                           //其他
	baseInfoBean                        = new(function.BaseInfoBean)                        //会员个人资料
	memberBean                          = new(function.MemberBean)                          //会员
	siteDomainBean                      = new(function.SiteDomainBean)                      //站点域名配置
	agencyBean                          = new(function.AgencyBean)                          //代理
	memberRegisterSettingBean           = new(function.MemberRegisterSettingBean)           //会员注册设定
	betRecordInfoBean                   = new(function.BetRecordInfoBean)                   //交易记录
	reportFormBean                      = new(function.ReportFormBean)                      //会员报表统计
	memberLevelBean                     = new(function.MemberLevelBean)                     //会员层级
	secondDistributionRegisterSetupBeen = new(function.SecondDistributionRegisterSetupBeen) //站点代理申请注册设定
	siteModuleBean                      = new(function.SiteModuleBean)                      //站点维护信息
	sitePaySetBean                      = new(function.SitePaySetBean)                      //站点支付设定
	selfHelpApplyforBean                = new(function.SelfHelpApplyforBean)                //自助优惠申请
	memMessageBean                      = new(function.MemberMessageBean)                   //会员个人资料
	redPacketSetBean                    = new(function.RedPacketSetBean)                    //红包
	betreport                           = new(function.BetReportBean)                       //报表
	productBean                         = new(function.ProductBean)                         //商品表
	rebateSetBean                       = new(function.MemberRebateSetBean)                 //返佣设定表
	member_level_bean                   = new(function.MemberLevelBean)                     //会员层级
	member_retreat_water_selfbean       = new(function.MemberRetreatWaterSelfBean)          //反水打码
	siteProductBean                     = new(function.SiteProductBean)                     //商品列表
	MemberSelfInfoBean                  = new(function.MemberSelfInfoBean)                  //会员消息
	BetRecordInfoBean                   = new(function.BetRecordInfoBean)                   //会员消息
	siteIWordBean                       = new(function.SiteIwordBean)                       //站点文案
	MemberCashRecordBean                = new(function.MemberCashRecordBean)                //出款记录
	noteGameBean                        = new(function.NoteGameBean)                        //电子游戏管理
	ManualAccessBean                    = new(function.ManualAccessBean)                    //人工存款
	drawMoney                           = new(function.DrawMoneyBean)                       //取款管理
	MemberBankBean                      = new(function.MemberBankBean)                      //会员银行卡列表
	NoticeBean                          = new(function.NoticeBean)                          //公告列表
	memberSpreadBean                    = new(function.MemberSpreadBean)                    //会员推广设定
	bankCardBean                        = new(function.BankCardBean)                        //银行
	paidTypeBean                        = new(function.PaidTypeBean)                        //支付类型
	memberCompanyIncomeBean             = new(function.MemberCompanyIncomeBean)             //公司入款
	onlineEntryRecordBean               = new(function.OnlineEntryRecordBean)               //公司入款记录
	onlinePaidSetupBean                 = new(function.OnlinePaidSetupBean)                 //支付设定
	onlineDepositBean                   = new(function.OnlineDeposit)                       //在线入款
	thirdAgencyBean                     = new(function.ThirdAgencyBean)                     //代理
	memberCashBean                      = new(function.MemberCashRecordBean)                //現金記錄
	registerStatusBean                  = new(function.RegisterStatusBean)                  //注册状态判断
	levelBean                           = new(function.SiteLevelBean)                       //站点层级
	platformBean                        = new(function.PlatformBean)                        //平台
	SitePromotionConfigBean             = new(function.SitePromotionConfigBean)             //自助申请优惠列表
	auditsBean                          = new(function.AuditsBean)                          //稽核记录
	payBean                             = new(thirdParty.PayBean)                           //第三方支付
	apiClientsBean                      = new(function.ApiClientsBean)                      //三方对接验证加密
	memberCashCountBean                 = new(function.MemberCashCountBean)                 //会员现金统计表
)
