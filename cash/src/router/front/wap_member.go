package front

import (
	"controllers/front/wap"
	"controllers/front/wap/data_merge"
	"github.com/labstack/echo"
	"router"
)

//页面模板
func WapMember(c *echo.Echo) {
	e := c.Group("", router.MemberCheck)

	loginDo := new(wap.LoginController)
	c.POST("/m/loginDo", loginDo.LoginDo)                   //登陆
	c.GET("/m/code", router.VerCode)                        //验证码
	e.PUT("/m/logout", loginDo.Logout)                      //登出
	e.GET("/m/ajaxLoginVerify", loginDo.GetAjaxLoginVerify) //登陆信息验证请求
	e.PUT("/m/editPassword", loginDo.EditPassword)          //修改密码
	//会员中心
	accountInfo := new(wap.AccountController)
	e.GET("/m/getInfo", accountInfo.GetInfo)                                       //会员个人基本资料
	e.GET("/m/detail", accountInfo.MemberDetail)                                   //查询会员详情
	e.PUT("/m/editInfo", accountInfo.PutInfo)                                      //修改会员资料
	e.GET("/m/back", accountInfo.Bank)                                             //会员出款银行列表
	e.GET("/m/backAddInfo", accountInfo.BankAddInfo)                               //银行列表
	e.POST("/m/bankAdd", accountInfo.BankAdd)                                      //添加会员出款银行信息
	e.POST("/m/memberCompanyIncome", accountInfo.AddCompanyIncome)                 //会员提交一条公司入款记录
	e.POST("/m/onlineIncome", accountInfo.AddOnlineIncome)                         //会员提交一条线上入款记录
	e.POST("/m/onlineIncome/card", accountInfo.CardDeposit)                        //会员提交一条线上存款-点卡
	e.GET("/m/getIncomeData", accountInfo.GetIncomeData)                           //获取存款数据
	c.POST("/m/getFastIncomeData", accountInfo.GetFastIncomeData)                  //获取快速充值中心存款数据--不用登录
	e.GET("/m/snatch", accountInfo.GetSnatch)                                      //抢红包
	e.GET("/m/company/getPaySetData", accountInfo.GetPaySetData)                   //获取会员对应层级对应配置的支付设定
	e.GET("/m/company/getPaySetDataByAccount", accountInfo.GetPaySetDataByAccount) //获取会员对应层级对应配置的支付设定-快捷支付
	e.GET("/m/applyData", accountInfo.ApplyData)                                   //优惠申请大厅数据
	e.POST("/m/applySubmit", accountInfo.ApplySubmit)                              //优惠申请提交

	report := new(data_merge.StatisticsAjax)
	e.GET("/m/getThisWeekReport", report.GetThisReportAjax) // 获取本周报表
	e.GET("/m/getLastWeekReport", report.GetLastReportAjax) // 获取上周报表
	returnwater := new(data_merge.ReturnWaterAjax)
	e.GET("/m/isself", returnwater.GetMemberIsSelf)              // 是否开启自助反水
	e.GET("/m/getmemberbet", returnwater.GetMemberGameBet)       //获取反水打码数据
	e.GET("/m/postreturnwater", returnwater.PostReturnWaterSelf) //存入反水打码数据
	e.GET("/m/member/bank", returnwater.GetMemberBankList)       //获取会员银行信息和银行列表
	e.GET("/m/member/getCard", returnwater.GetOneBank)           //获取单条会员出款卡号
	e.GET("/m/check/drawPass", returnwater.CheckDrawPass)        //检测支付密码
	wap_draw := new(data_merge.WapDrawBean)
	e.POST("/wap/draw/money/h", wap_draw.Withdrawal) //取款操作(使用)

	wapAjax := new(wap.WapAjaxController)
	e.GET("/wap/draw/write", wapAjax.DrawWriteData)                //写入出款数据
	c.GET("/m/getYouhui", wapAjax.GetActivity)                     //优惠数据
	c.GET("/m/redLog", wapAjax.RedPacketLog)                       //红包
	c.GET("/m/gameTitle", wapAjax.GetGameTitle)                    //获取电子标题
	c.GET("/m/gameData", wapAjax.GetGameData)                      //获取电子数据
	e.GET("/ajax/mes/list", wapAjax.GetMesList)                    //获取消息列表
	e.GET("/ajax/notice/list", wapAjax.GetNoticeList)              //获取公告列表
	e.GET("/ajax/record/info", wapAjax.GetRecordInfo)              //获取注单记录
	e.PUT("/ajax/mes/status", wapAjax.PutMesStatus)                //修改消息状态
	e.GET("/ajax/cash/record", wapAjax.GetCashList)                //获取交易记录
	e.GET("/ajax/wap/apply/config", wapAjax.AjaxGetApplyTitleList) //获取优惠申请活动列表
	e.GET("/ajax/wap/apply/list", wapAjax.AjaxGetApplyList)        //获取会员优惠活动申请列表

}
