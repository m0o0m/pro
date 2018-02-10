package front

import (
	"controllers/front/wap"
	"controllers/front/wap/data_merge"
	"github.com/labstack/echo"
	"router"
)

//页面模板
func WapPageTemplate(c *echo.Echo) {
	e := c.Group("", router.WapPageCheck)

	//******************** 以下的是独立路由 *************************
	homeCtrl := new(wap.HomeController)
	//登陆
	e.GET("/m/login", homeCtrl.Login)
	e.GET("/m/mescenter", homeCtrl.MesCenter)        //消息中心
	e.GET("/m/egame", homeCtrl.EGame)                //电子页面
	e.GET("/m/index", homeCtrl.Index)                //首页
	e.GET("/m", homeCtrl.Index)                      //首页
	e.GET("/m/", homeCtrl.Index)                     //首页
	e.GET("/m/reg", homeCtrl.Register)               //会员注册
	e.GET("/m/discount", homeCtrl.Discount)          //优惠活动
	e.GET("/m/Withdrawal/write", homeCtrl.DrawWrite) //出款写入页面
	e.GET("/m/mescenter", homeCtrl.MesCenter)        //消息中心
	e.GET("/m/egame", homeCtrl.EGame)                //电子页面

	//进入游戏界面
	video_play := new(data_merge.VideoPlay)
	e.GET("/m/video/login", video_play.PostVideoPlay) //进入游戏平台

	//注册
	regiser := new(wap.SignController)
	e.POST("/m/register/reg", regiser.Register)            //会员注册
	wapAjaxReg := new(wap.AjaxRegisterController)          //注册判断
	e.GET("/wap/ajax/reg/repeat", wapAjaxReg.RegisterAjax) //判断是否重复
	wapAjax := new(wap.WapAjaxController)
	e.GET("/m/ajax/register/set", wapAjax.GetRegisterSet) //获取注册设定数据
	//会员中心
	accountInfo := new(wap.AccountController)

	//登陆后才能打开的页面
	f := e.Group("", router.WapMemberCheck)
	f.GET("/m/bank", homeCtrl.Bank)                     //存款
	f.GET("/m/pay/callback", homeCtrl.PayCallback)      //第三方支付回调
	f.GET("/m/withdraw", homeCtrl.Withdraw)             //取款
	f.GET("/m/account", accountInfo.Index)              //会员中心
	f.GET("/m/info", accountInfo.Info)                  //会员详情
	f.GET("/m/modifyPas", accountInfo.ModifyPas)        //修改密码
	f.GET("/m/modifyInfo", accountInfo.ModifyInfo)      // 添加会员出款银行
	f.GET("/m/bankCard", accountInfo.BankCard)          //会员出款银行卡列表
	f.GET("/m/bankCardAdd", accountInfo.BankCardAdd)    // 添加会员出款银行
	f.GET("/m/statisticsthis", homeCtrl.Statisticsthis) //本周报表
	f.GET("/m/statisticslast", homeCtrl.Statisticslast) //上周报表
	f.GET("/m/returnwater", homeCtrl.ReturnWater)       //自助反水
	f.GET("/m/draw/write", homeCtrl.DrawWrite)          //取款数据写入
	f.GET("/m/apply", homeCtrl.ApplySelf)               //自助优惠申请
	f.GET("/m/finished", homeCtrl.Finished)             //线上入款提交完成
	f.GET("/m/finished2", homeCtrl.Finished2)           //快捷支付入款提交完成
	f.GET("/m/carry", homeCtrl.Carry)                   //公司支付提交完成
	f.GET("/m/convert", homeCtrl.Convert)               //额度转换
	f.GET("/m/record", homeCtrl.RecordList)             //交易记录
	f.GET("/m/fast", homeCtrl.Fast)                     //快捷支付 todo:相关会员数据待补充，循环出来的样式待调整

}
