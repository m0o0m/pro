package agency

import (
	"controllers/agency/report"
	"github.com/labstack/echo"
	"router"
)

func ReportRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	centerController := new(report.CenterController)
	e.GET("/report/center", centerController.GetCenter) //数据中心

	e.GET("/report/search", centerController.RepSearch) //报表查询页面需要的数据

	e.POST("/report/list", centerController.ReportList) //报表查询

	e.GET("/report/click", centerController.ReportClick)              //点击金额加载数据
	e.GET("/report/bills", centerController.ReportBills)              //账单查询
	e.GET("/report/periods", centerController.Periods)                //报表期数列表
	e.GET("/report/press", centerController.PressMoneySiteInfo)       //缴款账单
	e.PUT("/report/press/edit", centerController.PutSiteUpdateStatus) //缴款提交
	e.PUT("/report/pre/payment", centerController.PrePayment)         //预缴款
}
