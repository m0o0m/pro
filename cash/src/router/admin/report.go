package admin

import (
	"controllers/admin/import_member"
	"controllers/admin/rbac"
	"controllers/admin/report"
	"github.com/labstack/echo"
	"router"
)

//财务报表
func ReportRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.AdminRedisCheck, router.AdminAccessLog) //开启操作日志记录
	//e := c.Group("", router.AdminRedisCheck)
	//额度统计
	quota := new(report.QuotaCountController)
	e.GET("/finance/quotaNum", quota.QuotaList)                  //额度统计
	e.GET("/finance/quotaRecord", quota.QuotaRecordList)         //额度记录
	e.GET("/video/type", quota.GetPlatform)                      //视讯下拉框
	e.GET("/finance/rechargeRecord", quota.RechargeRecord)       //充值记录
	e.POST("/quota/recharge", quota.RechargeRecordUpdate)        //修改充值记录
	e.POST("/quota/payAddUpdate", quota.PostQuotaSetAddOrUpdate) //添加或者修改第三方或者银行卡
	e.GET("/quota/payList", quota.GetQuotaSetList)               //查询第三方或者银行卡
	e.GET("/quota/thirdType", quota.ThirdTypeDrop)               //支付类型下拉框

	//报表统计
	siteReport := new(report.ReportController)
	e.GET("/report/getTerm", siteReport.GetReportTerm)           //表查询页，显示所需查询条件
	e.GET("/reportquery", siteReport.GetSiteReportList)          //站点报表查询
	e.POST("/report", siteReport.GetSiteReportAccount)           //报表统计数据查询
	e.POST("/report/getExport", siteReport.GetSiteReportExport)  //账单导出
	e.GET("/billQuery", siteReport.GetSiteReportBill)            //账单查询
	e.PUT("/biiiQuery/batch", siteReport.GetSiteReportBillBatch) //账单批量下发
	e.POST("/billQuery", siteReport.PostSiteBillAdd)             //账单添加
	e.PUT("/billQuery", siteReport.PutSiteBillUpdate)            //账单修改
	e.DELETE("/billQuery", siteReport.DelSiteBill)               //账单删除

	//数据中心
	dataCenter := new(report.DataCenterController)
	e.GET("/finance/dataCenter", dataCenter.GetDataCenterList) //统计数据查询

	//现金报表
	cashRecord := new(report.CashRecordController)
	e.GET("/finance/cash", cashRecord.GetCashRecordList) //现金记录查询
	e.DELETE("/finance/cash", cashRecord.DelCashRecord)  //批量取消（硬删除）
	e.PUT("/finance/cash", cashRecord.PutCashRecord)     //批量删除（软删除）

	//催单管理
	pressmoney := new(report.PressMoneyController)
	e.GET("/reportform/arrears", pressmoney.GetPressMoneyList)         //催单列表
	e.GET("/reportform/arrearOne", pressmoney.GetPressMoneyOne)        //单条催单详情
	e.POST("/reportform/arrearsAdd", pressmoney.PostPressMoneyAdd)     //添加催单记录
	e.PUT("/reportform/arrearsUpdate", pressmoney.PutPressMoneyUpdate) //修改催单记录
	e.PUT("/reportform/arrearsStatus", pressmoney.PutPressMoneyStatus) //修改催单状态
	e.GET("/reportform/arrearsSite", pressmoney.PressMoneySiteDrop)    //站点id下拉框

	//站点层级
	siteLevel := new(report.SiteLevelController)
	e.POST("/finance/hierarchical", siteLevel.PostSiteLevelAdd)   //站点层级增加
	e.PUT("/finance/hierarchical", siteLevel.PutLevelInfoUpdate)  //站点层级修改
	e.DELETE("/finance/hierarchical", siteLevel.DelSiteLevelInfo) //站点层级删除
	e.GET("/hierarchicalData", siteLevel.GetSiteLevelList)        //站点层级列表
	e.GET("/hierarchical", siteLevel.GetSiteLevelInfo)            //获取单个层级详情
	e.PUT("/init/hierarchical", siteLevel.PutSiteLevelAll)        //初始化站点层级设置
	e.PUT("/move/hierarchical", siteLevel.PutSiteLevelUpdate)     //站点层级移动
	e.GET("/finance/hierarchical", siteLevel.GetSiteList)         //站点层级列表(以站点搜索)
	e.GET("/finance/hierarchy/drop", siteLevel.GetSiteListDrop)   //站点层级下拉框

	//套餐管理
	sitePackage := new(report.PackageController)
	combo := new(rbac.ComboController)
	e.GET("/combo", combo.Index)                     //查看套餐列表
	e.GET("/combo/info", combo.EditCombo)            //获取套餐详情
	e.PUT("/combo", combo.ComboEdit)                 //修改套餐
	e.GET("/comboPlatform", combo.ComboProduct)      //获取套餐平台
	e.POST("/comboPlatform", combo.ComboProductEdit) //设置套餐平台
	e.GET("/productType", combo.GetProductType)      //根据商品类型名称模糊搜索交易平台名
	e.GET("/combo/drop", sitePackage.ComboDrop)      //套餐下拉框
	e.PUT("/combo/status", combo.Status)             //修改套餐状态
	e.DELETE("/combo", combo.DeleteCombo)            //删除套餐
	e.POST("/combo", combo.ComboAdd)                 //添加套餐

	//优惠统计管理
	discountCount := new(report.DiscountCountController)
	e.GET("/finance/referential", discountCount.GetDiscountCountList)     //优惠统计数据查询
	e.POST("/discount/timingCount", discountCount.PutDiscountCountSwitch) //定时统计开关
	e.POST("/discount/recount", discountCount.PostDiscountCountUpdate)    //重新统计

	//入款统计管理
	depositCountController := new(report.DepositCountController)
	e.POST("/cashCount/timingCount", depositCountController.PutDepositCountSwitch) //定时统计开关
	e.POST("/cashCount/recount", depositCountController.PostDepositCountUpdate)    //重新统计
	e.GET("/finance/income", depositCountController.GetDepositCountList)           //入款统计数据查询

	//出入账目汇总
	cashCountController := new(report.CashCountController)
	e.GET("/finance/summary", cashCountController.GetCashCount)       //汇总查询
	e.POST("/cashCount/export", cashCountController.PutCashCountData) //导出excel

	//导入会员
	imu := new(import_member.UploadController)
	e.GET("/level/drop", imu.GetLevelDrop)                //层级下拉框
	e.GET("/first/agency/drop", imu.GetFirstAgencyDrop)   //股东下拉框
	e.GET("/second/agency/drop", imu.GetSecondAgencyDrop) //总代下拉框
	e.GET("/third/agency/drop", imu.GetThirdAgencyDrop)   //代理下拉框
	e.POST("/uploadMember", imu.Upload)                   //导入会员
}
