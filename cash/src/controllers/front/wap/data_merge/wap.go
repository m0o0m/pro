package data_merge

import "models/function"

//data_merge
const (
	WAP_INDEX          = "wap/index"
	WAP_FOOTER         = "wap/footer"
	WAP_MEM_FOOTER     = "wap/module/home/account/footer"
	WAP_APPLY_PRO      = "wap/module/activity/index"
	WAP_LOGIN          = "wap/module/login/login"
	WAP_REG            = "wap/module/login/enrolment"
	WAP_CONVERT        = "wap/module/home/convert/index"
	WAP_FAST           = "wap/module/home/deposit/fast"
	WAP_BANK           = "wap/module/home/deposit/bank"
	WAP_FINISHED       = "wap/module/home/deposit/finished"
	WAP_FINISHED2      = "wap/module/home/deposit/finished2"
	WAP_CARRY          = "wap/module/home/deposit/carry"
	WAP_WITHDRAW       = "wap/module/home/withdrawal/index"
	WAP_ACCOUNT        = "wap/module/home/index"
	WAP_MESCENTER      = "wap/module/activity/bulletin"
	WAP_RECORD         = "wap/module/home/account/recording"
	WAP_StatisticsThis = "wap/module/home/Statistics/StatisticsThis"
	WAP_StatisticsLast = "wap/module/home/Statistics/Statistics"
	WAP_INFO           = "wap/module/home/account/index"
	WAP_EGAME          = "wap/egame"
	MODIFY_PAS         = "wap/module/home/account/modifyPas"
	MODIFY_INFO        = "wap/module/home/account/modifyInfo"
	BANK_CARD          = "wap/module/home/account/BankCard"
	BANK_CARD_ADD      = "wap/module/home/account/BankCardAdd"
	WAP_RETURNWATER    = "wap/module/activity/returnWater"
	WAP_DRAW_WRTIE     = "wap/module/home/withdrawal/carry"
	WAP_APPLYSELF      = "wap/module/home/account/apply"
	WAP_PAY_CALLBACK   = "wap/module/home/deposit/payCallback"
)

var (
	noticeBean                    = new(function.NoticeBean)                 //公告
	noteGameBean                  = new(function.NoteGameBean)               //电子游戏管理
	siteOperateBean               = new(function.SiteOperateBean)            //网站运营
	webAdv                        = new(function.WebAdvBean)                 //弹窗广告
	siteDomainBean                = new(function.SiteDomainBean)             //站点域名配置
	siteDown                      = new(function.SiteDownBean)               //视讯下载链接表
	webInfoBean                   = new(function.WebInfoBean)                //站点信息
	siteIWordBean                 = new(function.SiteIwordBean)              //站点文案
	productBean                   = new(function.ProductBean)                //商品信息
	siteIwordBean                 = new(function.SiteIwordBean)              //文案管理
	baseInfoBean                  = new(function.BaseInfoBean)               //会员个人资料
	memMessageBean                = new(function.MemberMessageBean)          //会员个人资料
	sitePromotionConfig           = new(function.SitePromotionConfigBean)    //自助优惠申请配置
	webLogo                       = new(function.WebLogoBean)                //站点logo图片管理
	siteInfoBean                  = new(function.SiteInfoBean)               //站点信息
	webFloatBean                  = new(function.WebFloatBean)               //左右浮动图
	paidTypeBean                  = new(function.PaidTypeBean)               //支付类型
	betreport                     = new(function.BetReportBean)              //报表
	memberRegister                = new(function.MemberRegisterSettingBean)  //注册
	memberCompanyIncomeBean       = new(function.MemberCompanyIncomeBean)    //会员公司入款
	bankCardBean                  = new(function.BankCardBean)               //银行L
	siteProductBean               = new(function.SiteProductBean)            //商品列表
	member_level_bean             = new(function.MemberLevelBean)            //会员层级
	member_retreat_water_selfbean = new(function.MemberRetreatWaterSelfBean) //反水打码
	MemberBankBean                = new(function.MemberBankBean)             //会员银行卡列表
	MemberSelfInfoBean            = new(function.MemberSelfInfoBean)         //会员个人资料操作
	PoundageBean                  = new(function.PoundageBean)               //出款手续费设定
	drawMoney                     = new(function.DrawMoneyBean)              //取款管理
	SiteInfoBean                  = new(function.SiteInfoBean)               //获取站点联系方式
	memberBean                    = new(function.MemberBean)                 //查询会员信息
)
