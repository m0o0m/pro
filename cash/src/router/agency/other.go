package agency

import (
	"controllers/agency/website"
	"github.com/labstack/echo"
	"global/langs"
	"router"
)

func OtherRouter(c *echo.Echo) {
	c.GET(apiPath+"/langs", router.Langs) //添加菜单
	c.GET(apiPath+"/lang/cn", func(ctx echo.Context) error {
		return ctx.JSONBlob(200, []byte(langs.CNLangsAgency))
	}) //中文
	c.GET(apiPath+"/lang/us", func(ctx echo.Context) error {
		return ctx.JSONBlob(200, []byte(langs.USLangsAgency))
	}) //英文
	c.GET(apiPath+"/activitys", router.Activitys)
	c.GET(apiPath+"/activity/msgs", router.Activity)
	c.GET(apiPath+"/activity/notify", router.Activity)
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	//自助优惠申请配置
	sitepromotion := new(website.WebInfoController)
	e.GET("/site/application", sitepromotion.GetSitePromotionConfig)        //查询自助优惠列表
	e.POST("/site/application", sitepromotion.PostSitePromotionConfigAdd)   //添加
	e.PUT("/site/application", sitepromotion.PutSitePromotionConfigStatus)  //修改
	e.PUT("/site/egame/colorconfig", sitepromotion.SetElectConfig)          //电子内页主题配置新增/修改
	e.GET("/site/egame/colorconfig", sitepromotion.GetElectConfig)          //电子内页主题配置列表获取
	e.POST("/site/egame/colorconfig", sitepromotion.GetElectInitialization) //初始化
	//附件管理
	sitethum := new(website.WebLogoController)
	e.GET("/site/lottery/enclosure", sitethum.GetSiteThumb)    //附件列表
	e.PUT("/site/lottery/enclosure", sitethum.PutSitethumb)    //附件修改
	e.DELETE("/site/lottery/enclosure", sitethum.DelSitethumb) //附件删除
	e.POST("/site/lottery/enclosure", sitethum.PostSitethumb)  //附件上传

	//站点资料编辑
	webInfo := new(website.WebInfoController)
	e.GET("/site/electronics", webInfo.GetDzOrderList)             //电子管理-查询列表
	e.PUT("/site/electronics", webInfo.PutDzOrderUpdate)           //电子管理-修改顺序
	e.GET("/site/sports", webInfo.GetSpOrderList)                  //体育管理-查询列表
	e.PUT("/site/sports", webInfo.PutSpOrderUpdate)                //体育大厅-修改顺序
	e.GET("/site/lottery", webInfo.GetFcOrderList)                 //彩票大厅-查询列表
	e.PUT("/site/lottery", webInfo.PutFcOrderUpdate)               //彩票管理-修改顺序
	e.POST("/site/lottery", webInfo.PutFcOrderUpdateReset)         //重置排序
	e.GET("/site/video", webInfo.GetVideoOrderList)                //视讯管理-查询列表
	e.PUT("/site/video", webInfo.PutVideoOrderUpdate)              //视讯管理-修改顺序
	e.GET("/video/type/drop", webInfo.GetVideoOrderTypeList)       //视讯管理-获取类型下拉框
	e.GET("/video/style/drop", webInfo.GetVideoOrderStyleList)     //视讯管理-获取风格下拉框
	e.POST("/site/video/use", webInfo.PostVideoOrderStyleUse)      //视讯管理-风格配置使用
	e.PUT("/site/video/back", webInfo.PutVideoOrderStyleUseUpdate) //视讯管理-风格还原默认
	e.GET("/site/website", webInfo.GetSiteInfo)                    //网站基本信息-获取pc&wap
	e.POST("/site/website", webInfo.PostSiteInfo)                  //网站基本信息-添加或修改pc
	//e.POST("/webinfo/site_info_wap", webInfo.PostSiteInfoWap)      //网站基本信息-添加或修改wap
	e.GET("/site/module", webInfo.GetModuleOrderList)   //模板管理-查询列表
	e.PUT("/site/module", webInfo.PutModuleOrderUpdate) //模板管理-修改顺序
	//主动缓存
	e.POST("/site/genPageCache", webInfo.GenPageCacheBySite)    //整站生成页面缓存
	e.GET("/site/genPageCache", webInfo.GenPageCacheBySiteDrop) //界面下拉框
	e.POST("/site/delPageCache", webInfo.DelPageCacheBySite)    //整站清除页面缓存

	//文案编辑
	webWord := new(website.WebwordController)

	e.GET("/homeCopy", webWord.IwordList)       //文案查询
	e.GET("/homeEditor", webWord.IwordInfor)    //文案内容
	e.PUT("/homeEditor/sub", webWord.IwordEidt) //文案修改

	e.GET("/discount", webWord.ActivityList)                 //优惠文案查询
	e.DELETE("/discount/content/del", webWord.ActivityDel)   //优惠删除
	e.GET("/discount/content", webWord.ActivityInfo)         //优惠文案查询单个
	e.POST("/discount/addSub", webWord.ActivityEdite)        //优惠文案标题修改/添加
	e.PUT("/discount/content", webWord.ActivityEditeContent) //优惠内容修改

	e.GET("/lineDetection", webWord.SiteDetectList)           //站点线路检测数据查询
	e.PUT("/lineDetection/modifySub", webWord.SiteDetectEdit) //站点线路检测数据修改
	e.DELETE("/lineDetection/del", webWord.SiteDetectDel)     //站点线路检测数据删除
	e.POST("/lineDetection/addSub", webWord.SiteDetectAdd)    //站点线路检测数据添加

}
