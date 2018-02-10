package front

import (
	"controllers/front"
	"controllers/front/page"
	"github.com/labstack/echo"
	"router"
)

//页面模板
func PageTemplate(c *echo.Echo) {
	//主动缓存,因为这个功能是代理后台的,但是代理后台和页面前台是两个程序,内存不共享,所以以下两个接口使用的代理后台的中间件
	cacheCtrl := new(front.CacheController)
	c.POST("/site/genPageCache", cacheCtrl.GenPageCacheBySite, router.GetRedisToken) //整站生成页面缓存
	c.POST("/site/delPageCache", cacheCtrl.DelPageCacheBySite, router.GetRedisToken) //整站清除页面缓存

	e := c.Group("", router.PageCheck)

	homeCtrl := new(page.HomeController)
	e.GET("/", homeCtrl.NIndex, router.WhoIndex) //这个加控制器代理检测,区分移动端和pc端的跳转
	e.GET("/index", homeCtrl.NIndex)             //首页
	e.GET("/index.html", homeCtrl.NIndex)        //首页
	e.GET("/notice/data", homeCtrl.NoticeData)   //点击弹出公告
	e.GET("/quick/pay", homeCtrl.QuickPay)       //快速充值中心
	e.GET("/applypro", homeCtrl.Applypro)        //优惠活动大厅
	//e.GET("/red/log", homeCtrl.RedPacketLog)   //红包

	e.GET("/maintain", homeCtrl.Maintain) //新版本维护功能

	loginInfo := new(page.LoginInfoController)
	e.GET("/login/info", loginInfo.LoginInfo) //条款页

	wapview := new(page.WapviewController) //手机下注
	e.GET("/wapview", wapview.Wapview)     //手机下注
	e.GET("/download", homeCtrl.Download)  //下载专区

	e.GET("/sports", homeCtrl.Sports)                          //体育
	e.GET("/lottery", homeCtrl.Lottery)                        //彩票
	e.GET("/livetop", homeCtrl.LiveTop)                        //视讯
	e.GET("/egame", homeCtrl.EGame, router.SingleProductCheck) //电子
	e.GET("/youhui", homeCtrl.YouHui)                          //优惠
	//e.GET("/lotterypk", homeCtrl.LotteryPk)                    //PK彩票大厅
	e.GET("/detect", homeCtrl.Detect)                 //线路检测
	e.GET("/zhuce", homeCtrl.Register)                //会员注册
	e.GET("/daili/shenqing", homeCtrl.AgencyRegister) //代理注册
	e.GET("/iword", homeCtrl.IWord)                   //文案展示
	//加载or清除缓存的暂时写这里,后面加到代理后台和总后台;页面接口请写到上面👆
	//cacheCtrl := new(page.CacheController)
	e.GET("/cache/add/theme", cacheCtrl.CacheTheme) //缓存皮肤路径

	video_rule := new(page.VideoRuleController)
	e.GET("/rule", video_rule.VideoRule) //视讯规则页面

	isoview := new(page.IsoviewController)
	e.GET("/isoview", isoview.Isoview) //查看ios安装信任教程

	video_play := new(page.VideoPlay)
	e.GET("/video/login", video_play.PostVideoPlay) //跳转游戏地址

	//以下均要登陆
	f := e.Group("", router.PcMemberCheck)
	//会员中心页面
	f.GET("/member/account", homeCtrl.MemberAccount)       //会员中心首页--我的账户
	f.GET("/member/bank", homeCtrl.MemberAccount)          //会员中心--存款
	f.GET("/member/bank/third", homeCtrl.MemberAccount)    //会员中心--存款--第三方支付
	f.GET("/member/bank/complete", homeCtrl.MemberAccount) //会员中心--存款--支付完成
	f.GET("/member/withdraw", homeCtrl.MemberAccount)      //会员中心--取款
	f.GET("/member/convert", homeCtrl.MemberAccount)       //会员中心--额度转换
	f.GET("/member/record", homeCtrl.MemberAccount)        //会员中心--交易记录
	f.GET("/member/report", homeCtrl.MemberAccount)        //会员中心--报表统计
	f.GET("/member/spread", homeCtrl.MemberAccount)        //会员中心--我要推广
	f.GET("/member/mescenter", homeCtrl.MemberAccount)     //会员中心--消息中心
	f.GET("/member/draw/write", homeCtrl.MemberAccount)    //会员中心--出款写入
	f.GET("/pay/callback", homeCtrl.PayCallback)
}
