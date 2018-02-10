package agency

import (
	"controllers/agency/website"
	"github.com/labstack/echo"

	"router"
)

func IwordRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	webWord := new(website.WebwordController)
	e.GET("/iword/deposit/list", webWord.IwordList)                       //存款文案查询
	e.GET("/iword/deposit/detail", webWord.IwordInfor)                    //存款文案获取
	e.PUT("/iword/deposit", webWord.IwordEidt)                            //存款文案修改
	e.GET("/webword/activity_list", webWord.ActivityList)                 //优惠文案查询
	e.GET("/webword/activity_info", webWord.ActivityInfo)                 //优惠文案查询
	e.POST("/webword/activity_eidt", webWord.ActivityEdite)               //优惠文案标题修改/添加
	e.PUT("/webword/activity_eidt_content", webWord.ActivityEditeContent) //优惠内容修改
	e.DELETE("/webword/activity_del", webWord.ActivityDel)                //优惠删除
	e.GET("/webword/site_detect", webWord.SiteDetectList)                 //站点线路检测数据查询
	e.PUT("/webword/site_detect_edit", webWord.SiteDetectEdit)            //站点线路检测数据修改
	e.DELETE("/webword/site_detect_del", webWord.SiteDetectDel)           //站点线路检测数据删除
	e.POST("/webword/site_detect_add", webWord.SiteDetectAdd)             //站点线路检测数据添加
}
