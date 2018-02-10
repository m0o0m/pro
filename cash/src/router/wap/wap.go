//[app端路由接口]
package wap

import (
	"controllers/wap/activitycenter"
	"controllers/wap/conversion"
	"controllers/wap/drawmoney"
	"controllers/wap/login"
	"controllers/wap/memberinfo"
	"controllers/wap/messagecenter"
	"controllers/wap/records"
	"controllers/wap/report"
	"controllers/wap/selfhelp"
	"controllers/wap/sitelogo"
	"controllers/wap/wdeposit"
	"router"

	"github.com/labstack/echo"
)

func WapRouter(c *echo.Echo) {
	c.GET("/code", router.AppVerCode) //验证码
	c.GET("/codeimage", router.GetImage)
	wapLogin := new(login.WapLoginController)
	c.POST("/login", wapLogin.MemberLogin)       //会员登录
	c.POST("/register", wapLogin.MemberRegister) //注册
	c.GET("/setUp", wapLogin.GetMemberRegister)  //获取会员注册设定
	//首页logo，轮播图，公告
	l_p_n := new(sitelogo.InfoLogoController)
	c.GET("/homepage/info", l_p_n.HomePageLogo) //首页logo,公告，轮播图

	e := c.Group("", router.AppCheck)
	e.PUT("/logout", wapLogin.MemberLogout) //退出登录

	//交易记录管理
	record := new(records.WapRecordController)
	e.GET("/wap/bet/record", record.WapBetRecordList)                //投注记录列表
	e.GET("/wap/member/cash/record", record.WapMemberCashRecordList) //现金流水列表

	//额度转换ReportStatistics
	balanceConversion := new(conversion.WapBalanceConversionController)
	e.GET("/wap/balance", balanceConversion.WapBalance)                           //会员中心--会员余额刷新
	e.POST("/wap/balance/conversion", balanceConversion.WapBalanceConversion)     //额度转换-余额转换	（TODO:视讯接口是模拟的）
	e.POST("/platform/balance/refresh", balanceConversion.PlatformBalanceRefresh) //额度转换-单个平台余额刷新（TODO:视讯接口是模拟的）
	e.GET("/wap/platform/balance", balanceConversion.WapGetPlatformBalance)       //额度转换-获取各平台余额&&一键刷新

	//会员消息中心
	message := new(messagecenter.MemberMessage)
	e.GET("/member/person/message/list", message.MessageList) //会员个人消息列表
	e.GET("/member/notice/message/list", message.NoticeList)  //会员游戏公告列表
	e.GET("/member/person/message/info", message.MessageInfo) //个人消息详情
	e.GET("/member/notice/message/info", message.NoticeInfo)  //游戏公告详情
	e.DELETE("/member/person/message", message.MessageDel)    //删除个人消息
	e.DELETE("/member/notice/message", message.NoticeDel)     //删除会员游戏公告消息
	//会员活动中心
	activity := new(activitycenter.MemberActivity)
	c.GET("/member/activity/list", activity.WapActivityList) //会员活动列表
	e.GET("/member/activity/info", activity.WapActivityInfo) //单个会员活动详情

	//会员资料
	memberInfo := new(memberinfo.MemberInfoController)
	e.GET("/member/info/means", memberInfo.MemberSelfInfo)          //会员个人基本资料
	e.PUT("/member/info/email", memberInfo.EmailAddOrChange)        //修改/添加邮箱
	e.PUT("/member/info/birth", memberInfo.BirthAddOrChange)        //修改/添加生日
	e.PUT("/member/info/mobile", memberInfo.PhoneAddOrChange)       //修改/添加手机号
	e.GET("/member/info/code", memberInfo.PhoneCode)                //手机验证码
	e.GET("/member/info/homepage", memberInfo.MemberHomePage)       //会员中心主页
	e.PUT("/member/info/password", memberInfo.MemberPasswordChange) //会员修改密码

	//会员银行卡
	memberBankCard := new(memberinfo.MemberBankController)
	e.GET("/member/bank/card", memberBankCard.MemberBankList)        //会员银行卡
	e.POST("/member/card/add", memberBankCard.MemberBankAdd)         //添加银行卡
	e.PUT("/member/card/unbind", memberBankCard.MemberBankUnBind)    //银行卡解绑
	e.PUT("/member/card/bind", memberBankCard.MemberBankBind)        //银行卡绑定
	e.DELETE("/member/card/delete", memberBankCard.MemberBankDelete) //银行卡删除
	e.GET("/member/card/drop", memberBankCard.MemberBankDrop)        //银行下拉框
	e.GET("/member/card/one", memberBankCard.MemberBankCardOneInfo)  //会员银行卡详情

	//报表统计管理
	reports := new(report.ReportController)
	e.GET("/report/statistics", reports.ReportStatistics) //报表统计

	self := new(selfhelp.SelfHelpController)
	e.GET("/wap/member/rebate", self.WapMemberRebate) //推广返佣比例列表
	e.GET("/wap/leader/boards", self.Leaderboards)    //排行榜

	//取款管理
	drawMoney := new(drawmoney.WapDrawMoneyController)
	e.GET("/wap/member/info", drawMoney.GetMemberInfo)    //获取会员信息
	e.PUT("/wap/draw/money", drawMoney.Withdrawal)        //取款
	e.GET("/wap/draw/progress", drawMoney.DrawalProgress) //取款进度

	selfhelps := new(selfhelp.SelfHelpController)        //自助反水
	e.GET("/rewater/list", selfhelps.GetSelfHelpList)    //刚开始进来的返水列表和会员今日已经反水
	e.GET("/oneClick/rewater", selfhelps.OneClickSee)    //一键查看所有的反水额度以及有效投注
	e.GET("/single/rewater", selfhelps.GetSingleReWater) //查看单个的反水额度
	e.PUT("/pickup/all", selfhelps.AllReWaterPickUp)     //一键领取所有的

	//线上存款
	deposit := new(wdeposit.WapDeposit)
	e.POST("/online/wechat/deposit", deposit.WechatDeposit) //线上存款-微信 todo 未完成
	e.POST("/online/bank/deposit", deposit.BankDeposit)     //线上存款-网银 todo 未完成
	e.POST("/online/card/deposit", deposit.CardDeposit)     //线上存款-网银 todo 未完成
	//公司存款
	e.POST("/company/wechat/deposit", deposit.WechatCompanyDeposit) //公司存款-微信 todo 未完成
	e.POST("/company/bank/deposit", deposit.BankCompanyDeposit)     //公司存款-网银 todo 未完成

	wapDeposit := new(wdeposit.WapDepositController)
	e.GET("/company/getPaySetData", wapDeposit.GetPaySetData) //获取会员对应层级对应配置的支付设定
}
