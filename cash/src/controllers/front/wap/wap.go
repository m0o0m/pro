package wap

import (
	"models/function"
	"models/thirdParty"
)

var (
	noticeBean                  = new(function.NoticeBean)                  //公告
	siteBean                    = new(function.SiteOperateBean)             //站点信息
	memberBean                  = new(function.MemberBean)                  //会员
	siteDomainBean              = new(function.SiteDomainBean)              //站点域名配置
	memberRegisterSettingBean   = new(function.MemberRegisterSettingBean)   //会员注册设定
	memberLevelBean             = new(function.MemberLevelBean)             //会员层级
	agencyBean                  = new(function.AgencyBean)                  //代理
	baseInfoBean                = new(function.BaseInfoBean)                //会员个人资料
	MemberSelfInfoBean          = new(function.MemberSelfInfoBean)          //会员消息
	BetRecordInfoBean           = new(function.BetRecordInfoBean)           //会员消息
	memMessageBean              = new(function.MemberMessageBean)           //会员个人信息
	bankCardBean                = new(function.BankCardBean)                //银行
	siteIWordBean               = new(function.SiteIwordBean)               //站点文案
	noteGameBean                = new(function.NoteGameBean)                //电子游戏管理
	memberCompanyIncomeBean     = new(function.MemberCompanyIncomeBean)     //会员公司入款
	thirdAgencyBean             = new(function.ThirdAgencyBean)             //代理
	MemberCashRecordBean        = new(function.MemberCashRecordBean)        //出款记录
	sitePaySetBean              = new(function.SitePaySetBean)              //站点支付设定
	onlineEntryRecordBean       = new(function.OnlineEntryRecordBean)       //线上入款纪录
	onlineDepositBean           = new(function.OnlineDeposit)               //线上存款所需的表结构
	onlinePaidSetupBean         = new(function.OnlinePaidSetupBean)         //第三方支付设定
	redPacketSetBean            = new(function.RedPacketSetBean)            //红包
	memberCashBean              = new(function.MemberCashRecordBean)        //現金記錄
	drawMoney                   = new(function.DrawMoneyBean)               //取款管理
	ManualAccessBean            = new(function.ManualAccessBean)            //人工存款
	sitePromotionConfig         = new(function.SitePromotionConfigBean)     //
	selfHelpApplyforBean        = new(function.SelfHelpApplyforBean)        //优惠自助申请
	SitePromotionConfigBean     = new(function.SitePromotionConfigBean)     //优惠自助申请
	payBean                     = new(thirdParty.PayBean)                   //第三方支付
	apiClientsBean              = new(function.ApiClientsBean)              //三方对接验证加密
	memberCashCountBean         = new(function.MemberCashCountBean)         //会员现金统计表
	memberBalanceConversionBean = new(function.MemberBalanceConversionBean) //视讯余额
	auditsBean                  = new(function.AuditsBean)                  //稽核
	paidTypeBean                = new(function.PaidTypeBean)                //支付类型
	MemberBankBean              = new(function.MemberBankBean)              //会员银行卡列表
	registerStatusBean          = new(function.RegisterStatusBean)          //注册验证
	//baseInfoBean	   		  = new(function.BaseInfoBean)		  		//会员个人资料
)
