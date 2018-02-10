package agency

import (
	"controllers/agency/website"
	"github.com/labstack/echo"
	"router"
)

func WebSiteRouter(c *echo.Echo) {
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	weblogo := new(website.WebLogoController)                           //站点logo图片管理
	e.GET("/web/logo/list", weblogo.GetWebLogoList)                     //站点logo图片列表
	e.PUT("/web/logo/update", weblogo.PutWebLogo)                       //站点logo管理修改(基本信息)
	e.PUT("/web/logo/way", weblogo.PutWebLogoWay)                       //站点logo管理修改(修改路径)
	e.POST("/web/logo/add", weblogo.PostWebLogo)                        //站点logo管理新增
	//站点左右浮动查询
	e.DELETE("/web/float/delete", weblogo.DeleteWebFloat)   //浮动图片删除
	e.GET("/web/float/status", weblogo.GetWebFloatStatus)   //浮动图片状态
	e.PUT("/web/float/status", weblogo.PutWebFloatStatus)   //站点浮动图片全禁用开启
	e.GET("/web/float/list", weblogo.GetWebFloatList)       //站点左右浮动列表
	e.PUT("/web/float/update", weblogo.PutWebFloatUpdate)   //站点左右浮动修改(基本信息)
	e.PUT("/web/float/picture", weblogo.PutWebFloatPicture) //站点左右浮动修改（图片地址）
	e.POST("/web/float/add", weblogo.PostWebFloat)          //站点左右浮动新增
	//站点轮播图管理
	e.GET("/web/flash/list", weblogo.GetWebFlashList)     //站点轮播图管理列表
	e.PUT("/web/flash/update", weblogo.PutWebFlashUpdate) //站点轮播图管理修改
	e.POST("/web/flash/add", weblogo.PostWebFlash)        //站点轮播图管理新增
	e.PUT("/web/flash/status", weblogo.PutWebFlash)       //站点轮播图管理状态更改
	//站点公告弹窗管理
	e.GET("/web/adv/list", weblogo.GetWebAdvList)         // 站点公告弹窗管理 列表
	e.PUT("/web/adv/update", weblogo.PutWebAdvUpdate)     // 站点公告弹窗管理 修改
	e.GET("/web/adv/detail", weblogo.GetWebAdvListDetail) //站点公告弹窗管理 详情
	e.POST("/web/adv/add", weblogo.PostWebAdv)            // 站点公告弹窗管理 新增
	e.GET("/web/adv/config", weblogo.PutConfigDetail)     //站点公告弹窗管理 配置详情
	e.PUT("/web/adv/config", weblogo.PutConfig)           // 站点公告弹窗管理 配置修改
	e.DELETE("/web/adv/delete", weblogo.DeleteAdv)        //站点公告弹窗管理 删除
	//站点广告管理
	e.GET("/web/pop/list", weblogo.GetWebPopList)         // 站点广告管理 列表
	e.GET("/web/pop/detail", weblogo.GetWebPopListDetail) //站点广告详情
	e.PUT("/web/pop/update", weblogo.PutWebPopUpdate)     // 站点广告管理 修改
	e.POST("/web/pop/add", weblogo.PostWebPop)            // 站点广告管理 新增
	e.PUT("/web/pop/status", weblogo.PutPopStatus)        //站点广告管理 状态修改

}
