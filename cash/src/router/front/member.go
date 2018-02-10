package front

import (
	"controllers/front/member"
	"controllers/front/websocket"
	"controllers/wap/conversion"
	"github.com/labstack/echo"
	"router"
)

func Center(c *echo.Echo) {
	sign := new(member.SignController)
	cashCtrl := new(member.CacheController)
	e := c.Group("", router.MemberCheck)
	c.GET("/code", router.VerCode)                           //验证码
	c.POST("/login", sign.Login)                             //登录
	c.POST("/register", sign.Register)                       //注册
	c.POST("/refreshSiteModule", cashCtrl.RefreshSiteModule) //刷新站点维护信息缓存

	agencyReg := new(member.AgencyRegisterController)
	c.GET("/agency/register/set", agencyReg.Set)   //获取代理注册申请设定
	c.POST("/agency/register", agencyReg.Register) //提交代理注册申请

	bankInCtrl := new(websocket.BankIn)
	c.GET("/ws/deposit", bankInCtrl.DepositSuccess)   //存款推送服务
	c.GET("/ws/withdraw", bankInCtrl.WithdrawSuccess) //取款推送服务

	other := new(member.OtherController)
	c.GET("/product/name", other.GetProductName)                          //获取商品名
	e.GET("/pay/set/level", other.GetPaySetInfoById)                      //根据id获取站点层级支付设定
	e.GET("/member/level/payset", other.PaySet)                           //获取层级支付设定pay_set_id
	e.POST("/self/help/apply", other.SelfHelpApply)                       //自助优惠申请
	e.GET("/member/GetCompanyData", other.GetCompanyData)                 //公司入款数据
	e.GET("/member/GetIncomeData", other.GetIncomeData)                   //存款数据
	c.POST("/member/GetFastIncomeData", other.GetFastIncomeData)          //快速充值存款数据--不用登录
	e.POST("/member/memberCompanyIncome", other.AddCompanyIncome)         //会员提交一条公司入款记录
	e.POST("/member/memberOnlineBankIncome", other.AddOnlineBankIncome)   //会员提交一条第三方线上网银
	e.POST("/member/onlineIncome", other.AddOnlineIncome)                 //会员提交一条线上入款记录
	e.POST("/member/onlineIncome/card", other.CardDeposit)                //会员提交一条线上存款-点卡
	c.POST("/member/onlineIncomeCallback", other.AddOnlineIncomeCallback) //线上存款支付后回调接口
	c.POST("/member/onlineIncomeBank", other.GetOnlineIncomeBank)         //网银在线银行列表

	e.PUT("/member/password", sign.EditPassword)        //修改会员密码
	e.PUT("/logout", sign.Logout, router.PcMemberCheck) //登出

	bankPay := new(member.BankPayController)
	e.POST("/balance/conversions", bankPay.BalanceConversion) //额度转换
	e.GET("/balance", bankPay.Balance)                        //个人所有余额
	e.POST("/flyback", bankPay.Flyback)                       //一键回归

	payRecord := new(member.PayRecordController)
	e.GET("/product/types", payRecord.GetProductType)         //商品分类（交易记录二级菜单栏）
	e.GET("/product/types/product", payRecord.GetProductName) //商品分类下的商品(交易记录第三栏的类型)
	e.GET("/cathectic", payRecord.Cathectic)                  //投注记录(彩票，视讯，电子，捕鱼，体育的返回数据用GameResult返回的)
	e.GET("/member/cash/records", payRecord.MemberCashRecord) //现金流水 (//状态,流水项目没搞懂，没做返回)

	reportForm := new(member.ReportFormController)
	e.GET("/report/form", reportForm.Report) //报表统计

	//base_info
	memberInfo := new(member.BaseInfoController)
	e.GET("/base/info/info", memberInfo.Info) // 会员个人基本资料
	//e.GET("/base/info/record", memberInfo.PayRecord)       //今日交易记录
	e.GET("/base/info/record", memberInfo.GetPayRecord)    //今日交易记录
	e.GET("/base/info/list", memberInfo.Bank)              //获取个人出款银行列表
	e.GET("/base/info/bank", memberInfo.OneMemberBankInfo) //获取一条出款银行数据
	e.GET("/base/addInfo", memberInfo.BankAddInfo)         // 银行列表
	e.GET("/base/info/detail", memberInfo.MemberDetail)    //获取会员的详细信息
	e.PUT("/base/info/info", memberInfo.EditInfo)          //修改个人基本资料(密码)
	e.POST("/base/info/add", memberInfo.BankAdd)           //添加出款银行
	e.PUT("/base/info/put", memberInfo.BankEdit)           //修改出款银行
	e.DELETE("/base/info/delete", memberInfo.BankDelete)   //删除出款银行
	e.PUT("/base/info/phone", memberInfo.PhoneNumEdit)     //手机号修改
	e.PUT("/base/info/email", memberInfo.EmailNumEdit)     //邮箱修改
	e.PUT("/base/info/birth", memberInfo.BirthNumEdit)     //出生日期修改
	e.PUT("/base/info/means", memberInfo.MeansEdit)        //修改资料
	e.GET("/ajax/login/in", memberInfo.GetAjaxLoginStatus) //登陆验证
	e.GET("/base/rebate/all", memberInfo.GetMemberRebate)  //获取会员返佣详情

	//额度转换ReportStatistics
	balanceConversion := new(conversion.WapBalanceConversionController)
	//e.GET("/wap/balance", balanceConversion.WapBalance)                           //会员中心--会员余额刷新
	//e.POST("/wap/balance/conversion", balanceConversion.WapBalanceConversion)     //额度转换-余额转换	（TODO:视讯接口是模拟的）
	e.POST("/platform/balance/refresh", balanceConversion.PlatformBalanceRefresh) //额度转换-单个平台余额刷新（TODO:视讯接口是模拟的）
	//e.GET("/wap/platform/balance", balanceConversion.WapGetPlatformBalance)       //额度转换-获取各平台余额&&一键刷新
	//我要推广
	//红包
	redPacket := new(member.RedController)
	e.GET("/snatch", redPacket.GetSnatch)     // 抢红包
	c.GET("/red/log", redPacket.RedPacketLog) //红包

	//会员中心
	report := new(member.StatisticsAjax)
	e.GET("/member/reportthis", report.GetThisReportAjax) // 本周报表
	e.GET("/member/reportlast", report.GetLastReportAjax) // 上周报表

	returnwater := new(member.ReturnWaterAjax)
	e.GET("/member/isself", returnwater.GetMemberIsSelf)              // 是否开启自助反水
	e.GET("/member/getmemberbet", returnwater.GetMemberGameBet)       //获取反水打码数据
	e.GET("/member/postreturnwater", returnwater.PostReturnWaterSelf) //存入反水打码数据
	//  消息中心 交易记录ajax
	recordajax := new(member.RecordAjax)
	e.GET("/ajax/get/mes/list", recordajax.GetMesList)            //消息列表
	e.PUT("/mesAjax/mes/status", recordajax.PutMesStatus)         //修改消息状态
	e.GET("/ajax/cashRecord", recordajax.GetCashList)             //获取交易记录
	e.GET("/ajax/record/infoList", recordajax.GetRecordInfo)      //获取投注记录
	e.GET("/ajax/noticeList", recordajax.GetNoticeList)           //获取公告列表
	e.GET("/ajax/delete/mes", recordajax.DeleteMes)               //删除消息
	e.GET("/ajax/get/gameList", recordajax.GetGameList)           //根据id获取游戏列表
	e.GET("/ajax/get/drawList", recordajax.GetDrawData)           //获取会员出款数据
	e.GET("/ajax/check/drawPass", recordajax.CheckDrawPass)       //检测支付密码
	e.GET("/ajax/draw/write", recordajax.DrawWriteData)           //出款数据写入
	c.GET("/ajax/egame", recordajax.GetGameData)                  //电子数据加载
	e.GET("/ajax/get/oneCard", recordajax.GetOneBank)             //获取单条会员出款卡号
	e.GET("/ajax/get/balance", recordajax.AjaxGetBalance)         //获取单条会员金额
	e.GET("/ajax/get/game/list", recordajax.AjaxFcList)           //获取彩票下游戏列表
	e.GET("/ajax/apply/list", recordajax.AjaxGetApplyList)        //获取优惠申请列表
	e.GET("/ajax/apply/config", recordajax.AjaxGetApplyTitleList) //获取优惠申请活动列表
	e.GET("/ajax/draw/data", recordajax.AjaxGetDrawData)          //获取会员是否有未处理出款

	//ajax请求数据
	DrawBean := new(member.DrawBean)
	e.POST("/member/draw/data", DrawBean.Withdrawal) //出款页面ajax数据
	//ajax获取条件
	ajax_register := new(member.AjaxRegister)
	c.GET("/ajax/register", ajax_register.RegisterAjax)                     //获取注册判断条件
	c.GET("/ajax/pc/isReg", ajax_register.GetIsRegStatus)                   //检测是否开启注册状态
	c.POST("/member/result/memberCompanyIncome", other.CompanyIncomeResult) //模拟一条公司入款结果:测试推送的
	c.POST("/member/result/drawWrite", recordajax.DrawWriteResult)          //模拟一条公司出款结果:测试推送的
}
