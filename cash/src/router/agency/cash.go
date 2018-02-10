package agency

import (
	"controllers/agency/cash"
	"controllers/agency/rebate"
	"github.com/labstack/echo"
	"router"
)

func CashRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	//e := c.Group("")
	//银行卡管理
	bank_card := new(cash.BankCardController)
	e.GET("/bank/income", bank_card.BankIncome)              //入款银行
	e.GET("/bank/outbank", bank_card.BankOut)                //出款银行
	e.GET("/bank/third", bank_card.BankThird)                //三方银行
	e.PUT("/bank/income/status", bank_card.BankIncomeStatus) //入款银行开启、剔除
	e.PUT("/bank/outbank/status", bank_card.BankOutStatus)   //出款银行开启、剔除
	e.PUT("/third/status", bank_card.BankThirdStatus)        //三方银行开启、剔除
	e.GET("/agency/bankDrop", bank_card.BankAgencyOutByDrop) //站点出款下拉框
	//人工存取款
	manualAccess := new(cash.CashController)
	e.POST("/manualAccess", manualAccess.ManualAccessDo)             //添加一条人工存款
	e.POST("/manualAccess/batch", manualAccess.ManualAccessBatchDo)  //添加多条人工存款
	e.POST("/manualWithdrawal", manualAccess.ManualWithdrawal)       //添加人工取款
	e.POST("/Quota/Submit", manualAccess.BalanceConversionDo)        //添加额度转换（调用视讯接口是模拟的,写在global/video中）
	e.GET("/balanceConversion", manualAccess.BalanceConversion)      //额度转换列表
	e.GET("/manualAccess/record", manualAccess.ManualAccess)         //存取款历史记录列表
	e.GET("/manualAccess/collect", manualAccess.ManualAccessCollect) //账目汇总
	//2018-1-29   todo   修改err打印
	e.GET("/member/baseinfo", manualAccess.MemberInfo)  //获取会员账号，余额，姓名
	e.GET("/member/balance", manualAccess.GetMoney)     //获取会员转出项目余额
	e.GET("/member/getbalance", manualAccess.GetBlance) //获取会员余额统计数据
	e.GET("/type/list", manualAccess.Index)             //平台列表(转出/转入项目)
	//入款银行设定
	payment := new(cash.PaymentController)
	e.GET("/payment/list", payment.BankIncomeList)        //入款银行设定列表
	e.POST("/add/payment", payment.BankIncomeDo)          //添加入款银行设定
	e.PUT("/payment/put", payment.BankIncomeEdit)         //修改入款银行设定
	e.GET("/payment", payment.BankIncome)                 //查询一条入款银行设定
	e.DELETE("/payment/delete", payment.BankIncomeDelete) //删除一条入款银行设定
	e.PUT("/payment/status", payment.BankIncomeStatus)    //开启、禁用
	e.GET("/payment/deposit", payment.DepositRecord)      //存款记录
	e.GET("/payment/level", payment.ApplicationLevel)     //适用层级
	//代理退佣设定
	override := new(cash.OverrideController)
	e.GET("/override/list", override.OverrideList)         //退佣列表
	e.GET("/override/getone", override.OverGetOne)         //获取单条数据
	e.GET("/override/delete", override.DeleteOver)         //删除单条数据
	e.POST("/override/update", override.UpdateOver)        //修改单条数据
	e.POST("/override/addone", override.OverRideAdd)       //新增单条数据
	e.POST("/override/update_money", override.UpdataMoney) //修改有效会员投注金额

	//会员推广设定
	spread := &rebate.SpreadController{}
	e.POST("/spread/add", spread.SpreadDo)       //添加会员推广设定
	e.PUT("/spread/edit", spread.SpreadDoSubmit) //修改会员推广设定
	e.GET("/spread/list", spread.SpreadList)     //查询会员推广设定
	e.GET("/spread/info", spread.SpreadInfo)     //查询会员推广信息

	//返点优惠设定
	member_retreat_water_set := new(cash.MemberRetreatWaterSetController)
	e.GET("/retreat/water/set/list", member_retreat_water_set.ListMemberRetreatWaterSet)     //返点优惠设定列表
	e.POST("/retreat/water/set/add", member_retreat_water_set.AddMemberRetreatWaterSet)      //添加返点设定
	e.PUT("/retreat/water/set/edit", member_retreat_water_set.EditMemberRetreatWaterSet)     //修改返点设定
	e.DELETE("/retreat/water/set/del", member_retreat_water_set.DelMemberRetreatWaterSet)    //删除返点设定
	e.GET("/retreat/water/set/detail", member_retreat_water_set.DetailMemberRetreatWaterSet) //获取单个详情
	e.GET("/preferential/inquiries", member_retreat_water_set.SearchMemberRetreatWaterSet)   //优惠查询列表
	//优惠查询
	member_retreat_water := new(cash.MemberRetreatWaterController)
	e.GET("/retreat/water/detail", member_retreat_water.DetailMemberRetreatWater) //优惠查询明细
	e.PUT("/retreat/water/edit", member_retreat_water.EditMemberRetreatWater)     //优惠查询冲销
	e.GET("/retreat/water/search", member_retreat_water.SearchMemberRetreatWater) //优惠查询列表

	//自助返水查询
	member_retreat_water_self := new(cash.MemberRetreatWaterSelfController)
	e.GET("/retreat/water/self/search", member_retreat_water_self.SearchMemberRetreatWaterSelf) //自助返水查询列表
	e.GET("/retreat/water/self/detail", member_retreat_water_self.DetailMemberRetreatWaterSelf) //自助返水查询明细
	//优惠统计
	bet_report_account := new(cash.BetReportAccountController)
	e.POST("/bet/report/account/count", bet_report_account.CountBetReportAccount) //优惠统计
	e.POST("/bet/report/account/store", bet_report_account.StoreBetReportAccount) //优惠统计-存入

	//会员返佣设定
	rebateSet := &rebate.RebateSetController{}
	e.POST("/rebateSet/add", rebateSet.AddOrUpdate) //添加或更新
	e.PUT("/rebateSet/del", rebateSet.Del)          //删除
	e.GET("/rebateSet/getOne", rebateSet.GetOne)    //查询详情
	e.GET("/rebateSet/getAll", rebateSet.GetAll)    //查询列表

	//会员返佣查询
	reBate := &rebate.RebateController{}
	e.GET("/rebate/count", reBate.Count)        //会员返佣统计
	e.POST("/rebate/save", reBate.Commit)       //会员返佣存入
	e.GET("/rebate/list", reBate.List)          //查询返佣列表
	e.GET("/rebate/details", reBate.Details)    //返佣详情
	e.POST("/rebate/writeoff", reBate.Writeoff) //返佣冲销

	//支付设定
	payment_set := new(cash.PaymentSetController)
	e.GET("/payset/public", payment_set.PublicPaySetList)         //币别列表
	e.GET("/currency", payment_set.PublicPaySetListCurrency)      //公共币种列表
	e.GET("/payset/public/ones", payment_set.OnePublicPaymentSet) //查询一条币别
	e.POST("/payset/add", payment_set.PaymentSetAdd)              //添加支付设定
	e.GET("/payset/list", payment_set.PaymentSetList)             //支付设定列表
	e.PUT("/payset", payment_set.PaymentSetUp)                    //支付设定设置
	e.PUT("/payset/modify", payment_set.PaymentSetChange)         //支付设定修改名称
	e.DELETE("/payset/delete", payment_set.PaymentSetDelete)      //删除支付设定
	e.GET("/payset/detail", payment_set.PaymentSetOne)            //查询一条支付设定

	//线上入款
	cashs := new(cash.CashController)
	e.GET("/thirdPaidList", cashs.ThirdPaidList)               //入款商户下拉
	e.PUT("/Monitor/oblin/confirm", cashs.OnlineIncomeDo)      //确定一条线上入款
	e.PUT("/Monitor/online/cancel", cashs.OnlineIncomeChannel) //取消一条线上入款
	e.GET("/deposit/list", cashs.OnlineIncomeList)             //线上入款列表

	//公司入款
	e.PUT("/Deposit/Cancel", cashs.CompanyIncomeDo)      //确定一条公司入款
	e.PUT("/Monitor/cancel", cashs.CompanyIncomeChannel) //取消一条公司入款 不再提醒
	e.GET("/companyIncome", cashs.CompanyIncome)         //公司入款列表
	e.GET("/getAgency", cashs.GetAgency)                 //代理下拉框
	e.GET("/getSetAgency", cashs.GetSetAgency)           //收款账号下拉框

	//出款管理
	e.GET("/outMoney", cashs.ListCash)                //出款管理列表
	e.GET("/memberDrop", cashs.MemberDrop)            //会员层级下拉框
	e.PUT("/prepareOut", cashs.ManageCashReady)       //预备出款
	e.PUT("/confirm/outMoney", cashs.ManageCashDo)    //确定出款
	e.PUT("/modify/cancel", cashs.ManageCashChannel)  //取消出款
	e.PUT("/refuse/outMoney", cashs.ManageCashRefuse) //拒绝出款
	e.GET("/outRemark", cashs.OutRemark)              //拒绝，取消出款原因
	//线上入款设定
	//todo 如果需要将请求第三方放在model里面开事务，需要将controllers的移动到models,第三方的错误没有搬到本地
	incomeSet := new(cash.OnlinePaidSetupController)
	e.DELETE("/onlineSetup/del", incomeSet.DelThisPaidSetup)                   //删除该支付设定
	e.GET("/third", incomeSet.GetThirdPayList)                                 //获取第三方支付平台[第三方接口]
	e.GET("/paidType", incomeSet.GetOnlinePaidType)                            //获取支付类型//[第三方接口]
	e.GET("/onlineSetup/depositeRecord", incomeSet.GetThisOnlineDepositRecord) //获取该支付设定的存款记录
	e.PUT("/stop/onlineSetup", incomeSet.StopThisOnlineSetup)                  //修改线上支付状态
	e.PUT("/newOnlineSetup/modify", incomeSet.ChangeOnlineSetup)               //修改该线上支付设定[第三方接口]
	e.POST("/newOnlineSetup", incomeSet.AddNewOnlinePaidSetup)                 //新增加线上支付设定[第三方接口]
	e.GET("/onlineSetup", incomeSet.OnlinePaidSetupList)                       //线上支付设定列表
	e.GET("/onlineSetup/single", incomeSet.GetPaidSetup)                       //获取某个支付设定详情[三方接口但是是从本地获取]
	e.GET("/paid_setup/bank", incomeSet.GetPaidBank)                           //某个支付类型下面的银行列表获取[第三方接口]

	//期数管理
	periods := new(cash.PeriodsController)
	e.GET("/periods/list", periods.PeriodsList)      //期数列表获取
	e.GET("/periods/getone", periods.PeriodsGetOne)  //获取单条期数数据
	e.GET("/periods/delete", periods.PeriodsDelete)  //删除单条期数数据
	e.GET("/periods/commission", periods.Commission) //退佣冲销
	e.POST("/periods/addone", periods.PeriodsAdd)    //新增一条期数数据
	e.POST("/periods/update", periods.PeriodsUpdate) //修改单条期数数据

	//退佣统计
	rebatecount := new(cash.RebateCountController)
	e.POST("/rebatecount/criteria", rebatecount.GetList)  //退佣统计条件
	e.POST("/rebatecount/display", rebatecount.CheckList) //退佣统计查询展示
	e.POST("/rebatecount/write", rebatecount.RebateFile)  //退佣存档

	//退佣查询
	retirement := new(cash.RetirementController)
	e.GET("/retirement/list", retirement.GetList)       //获得查询条件
	e.POST("/retirement/getlist", retirement.CheckList) //退佣查询

	//掉单申请
	siteSingleRecord := new(cash.SiteSingleRecord)
	e.GET("/siteSingleRecord", siteSingleRecord.Index)   //掉单列表
	e.POST("/subSingleRecord", siteSingleRecord.Add)     //添加掉单申请
	e.PUT("/site_single_record", siteSingleRecord.Check) //审核掉单申请
	//额度统计
	quota := new(cash.QuotaController)
	e.GET("/quota/list", quota.QuotaList)          //额度统计
	e.GET("/quota/record", quota.QuotaRecordList)  //额度记录
	e.GET("/videoSelect", quota.GetPlatform)       //视讯下拉框
	e.GET("/quota/recharge", quota.RechargeRecord) //充值记录

	//设定手续费
	poundage := new(cash.PoundageController)
	e.GET("/poundage/listset", poundage.Poundage)    //手续费设定
	e.GET("/poundage/getlist", poundage.GetList)     //获取单条手续费设定
	e.POST("/poundage/addset", poundage.PoundAdd)    //新增手续费设定
	e.POST("/poundage/update", poundage.PoundUpdate) //修改手续费设定

	//稽核管理
	audit := new(cash.AuditController)
	e.GET("/system/audit", audit.AuditLogList)        //稽核日志查询
	e.GET("/audit/memberauditnow", audit.GetNowAudit) //会员即时稽核
	//红包列表，添加
	redBag := new(cash.RedPacketSetController)
	e.GET("/redBag/list", redBag.List)           //红包列表
	e.GET("/redBag/one", redBag.ListInfo)        //红包详情
	e.POST("/redBag/one", redBag.Add)            //添加红包
	e.PUT("/redBag/one", redBag.RedBagChange)    //修改红包
	e.GET("/redBag/info", redBag.RedBagInfo)     //查看红包
	e.DELETE("/redBag/one", redBag.RedBagRemove) //红包终止
	//红包生成
	redGenerate := new(cash.RedPacketLogController)
	e.POST("/redBag/generate", redGenerate.GenerateRedPacket) //生成红包
	//红包样式
	redStyle := new(cash.RedPacketStyleController)
	e.GET("/redStyle/drop", redStyle.FindListDrop)           //红包样式下拉框
	e.GET("/redStyle/list", redStyle.FindList)               //红包样式列表
	e.GET("/redStyle", redStyle.GetDetails)                  //红包样式详情
	e.POST("/redStyle", redStyle.AddOrUpdate)                //红包样式添加或修改
	e.DELETE("/redStyle", redStyle.Del)                      //红包样式删除
	e.GET("/redStyle/picture", redStyle.FindListDropPicture) //红包样式图片

}
