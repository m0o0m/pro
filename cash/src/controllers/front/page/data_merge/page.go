package data_merge

import "models/function"

var (
	noticeBean              = new(function.NoticeBean)                          //公告
	noteGameBean            = new(function.NoteGameBean)                        //电子游戏管理
	siteOperateBean         = new(function.SiteOperateBean)                     //网站运营
	webAdv                  = new(function.WebAdvBean)                          //弹窗广告
	siteDomainBean          = new(function.SiteDomainBean)                      //站点域名配置
	siteDown                = new(function.SiteDownBean)                        //视讯下载链接表
	webInfoBean             = new(function.WebInfoBean)                         //站点信息
	siteIWordBean           = new(function.SiteIwordBean)                       //站点文案
	productBean             = new(function.ProductBean)                         //商品信息
	siteIwordBean           = new(function.SiteIwordBean)                       //文案管理
	baseInfoBean            = new(function.BaseInfoBean)                        //会员个人资料
	memMessageBean          = new(function.MemberMessageBean)                   //会员个人资料
	sitePromotionConfig     = new(function.SitePromotionConfigBean)             //自助优惠申请配置
	webLogo                 = new(function.WebLogoBean)                         //站点logo图片管理
	siteInfoBean            = new(function.SiteInfoBean)                        //站点信息
	webFloatBean            = new(function.WebFloatBean)                        //左右浮动图
	paidTypeBean            = new(function.PaidTypeBean)                        //支付类型
	siteOrderModule         = new(function.SiteOrderModule)                     //模块管理排序
	siteProductBean         = new(function.SiteProductBean)                     //商品列表
	siteInfoVideoUser       = new(function.SiteInfoVideoUser)                   //视讯模版选择
	pkGames                 = new(function.PkGames)                             //pk彩票彩种列表
	egGames                 = new(function.EgGames)                             //eg彩票彩种列表
	csGames                 = new(function.CsGames)                             //cs彩票彩种列表
	siteAgencyRegSet        = new(function.SecondDistributionRegisterSetupBeen) //代理注册
	memberRegister          = new(function.MemberRegisterSettingBean)           //会员注册
	memberCompanyIncomeBean = new(function.MemberCompanyIncomeBean)             //会员公司入款
	bankCardBean            = new(function.BankCardBean)                        //银行L
	drawMoney               = new(function.DrawMoneyBean)                       //取款管理
	AgencyBean              = new(function.AgencyBean)                          //代理查询
)

const (
	MAINTENANCE = "public/maintenance" //维护页面
	MAINTAIN    = "public/maintain"    //新版维护页面

	FC = "lottery" //彩票
	SP = "sport"   //体育
	DZ = "game"    //电子
	VD = "live"    //视讯
	YH = "youhui"  //视讯

	FOOTER_      = "footer_"
	HEADER_      = "header_"
	POP_ADV      = "pop_adv"
	N_INDEX      = "n_index"
	HEADER       = "header"
	FOOTER       = "footer"
	LIVE_TOP     = "livetop"
	WAPVIEW      = "wapview"
	LOGIN_INFO   = "login_info"
	DOWNLOAD     = "download"
	VIDEO_RULE   = "video_rule"
	ISOVIEW      = "isoview"
	NOTICE_DATA  = "notice_data"
	QUICK_PAY    = "quick_pay"
	APPLYPRO     = "applypro"
	PAY_CALLBACK = "pay_callback"

	DETECT      = "detect"
	ZHUCE       = "zhuce"
	ZHUCE_DAILI = "zhuce_daili"
	IWORD       = "iword"

	SPORTS    = "sports"
	LOTTERY   = "lottery"
	LOTTERYPK = "lotterypk"

	//会员中心页面
	MEMBERINDEX     = "new_member_main"
	MEMBERHEADER    = "member/memberheader"
	MEMBERACCOUNT   = "member/index"
	MEMBERTHIRD     = "member/deposit/Thirdparty"
	WITHDRAWALINDEX = "member/withdrawal/Index"
	MEMBERCOVERT    = "member/quota-conversion"
	MEMBERRECORD    = "member/trading-record"
	MEMBERREPORT    = "member/report"
	MEMBERSPREAD    = "member/promotion"
	MEMBERNEWS      = "member/news"
	MEMBERCOMPANY   = "member/deposit/OLBankingIndex"
	MEMBERCOMPLETE  = "member/deposit/complete"
	MEMBERDRAWWRITE = "member/withdrawal/Audit"
)
