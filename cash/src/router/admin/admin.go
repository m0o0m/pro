//[路由] 总后台接口路由
package admin

import (
	"controllers/admin/customer"
	"controllers/admin/filemanager"
	"controllers/admin/login"
	"controllers/admin/note"
	"controllers/admin/product"
	"controllers/admin/rbac"
	"controllers/admin/site"
	"controllers/admin/templet"
	"controllers/agency/account"
	"global/langs"
	"router"

	"github.com/labstack/echo"
)

const apiPath = "/api"

func NewAdminRouter(c *echo.Echo) {
	adminLogin := new(login.LoginDoController)

	c.GET("/metrics", router.PromHandler)
	c.POST(apiPath+"/login", adminLogin.GetLoginDo) //平台管理员登录[不要进行校验]
	c.GET(apiPath+"/captcha", router.VerCode)       //验证码
	c.GET(apiPath+"/langs", router.Langs)           //语言选择
	c.GET(apiPath+"/lang/cn", func(ctx echo.Context) error {
		return ctx.JSONBlob(200, []byte(langs.CNLangsAdmin))
	}) //中文
	c.GET(apiPath+"/lang/us", func(ctx echo.Context) error {
		return ctx.JSONBlob(200, []byte(langs.USLangsAdmin))
	}) //英文
	c.GET(apiPath+"/activitys", router.Activitys)
	c.GET(apiPath+"/activity/msgs", router.Activity)
	c.GET(apiPath+"/activity/notify", router.Activity)
	member := new(customer.MemberController)
	e := c.Group(apiPath, router.AdminRedisCheck, router.AdminAccessLog) //开启操作日志记录
	e.GET("/logout", adminLogin.GetLoginOut)                             //退出登录
	//站点会员
	e.GET("/customer/userManagement", member.GetMemberList)  //站点会员查询
	e.PUT("/site/member/update", member.PutMemberInfoUpdate) //会员资料修改
	e.GET("/user/informent", member.MemberInfo)              //会员资料详情
	e.PUT("/user/status", member.PutMemberStatusUpdate)      //会员状态修改
	//站点层级
	level := new(customer.MemberLevelController)
	e.GET("/site/member/level/list", level.GetMemeberLevelList) //站点层级查询
	//站点公告
	notice := new(site.NoticeController)
	e.GET("/system/notice", notice.GetNoticeList)               //获取站点公告列表
	e.POST("/system/notice", notice.PostNoticeAdd)              //添加站点公告
	e.PUT("/system/notice/state", notice.PutNoticeStatusUpdate) //开启或关闭站点公告
	e.PUT("/system/notice", notice.PutNoticeUpdate)             //修改公告信息
	e.DELETE("/system/notice", notice.PutNoticeDel)             //批量删除公告信息
	//稽核
	auditLog := new(customer.AuditLogController)
	e.GET("/audit/log", auditLog.GetAuditLogList)       //稽核日志查询
	e.GET("/audit/record", auditLog.GetAuditRecordList) //稽核列表查询
	e.GET("/instant/audit", auditLog.GetAuditNowList)   //即时稽核查询
	//客户后台管理-日志管理
	siteLog := new(customer.SiteLogController)
	e.GET("/siteDoLog", siteLog.GetSiteDoLog)       //操作日志查询
	e.GET("/siteLoginLog", siteLog.GetSiteLoginLog) //登录日志查询
	e.GET("/autoAudit", siteLog.GetSiteAutoAudit)   //自动稽核
	//出入款
	bankinOut := new(customer.BankInOutController)
	e.GET("/bankinOut/in", bankinOut.GetBankInRecord)   //入款查询
	e.GET("/bankinOut/out", bankinOut.GetBankOutRecord) //出款查询

	//子帐号
	child := new(customer.ChildAccountController)
	e.GET("/child/list", child.GetChildAccountList)           //子帐号列表
	e.GET("/child/one", child.GetOneChildAccount)             //获取一条子帐号信息
	e.PUT("/child", child.PutChildAccountUpdate)              //修改子帐号信息
	e.PUT("/child/status", child.PutChildAccountStatusUpdate) //子帐号状态修改
	e.GET("/third/dropf", child.SiteSiteIndexIdByAll)         //子帐号、代理站点下拉框
	e.GET("/child/site", child.SiteIndexIdList)               //子站点下拉框

	//优惠查询
	discountSearch := new(customer.DiscountSearchController)
	e.GET("/customer/preferentialQuery", discountSearch.GetDiscountCountList) //优惠列表
	e.GET("/favourable/total", discountSearch.GetDiscountCount)               //优惠查询
	e.GET("/customer/preferential/list", discountSearch.GetDiscountList)      //查询明细

	//自助优惠申请
	autoDiscount := new(customer.AutoDisCountController)
	e.PUT("/customer/applicationSwitch", autoDiscount.PutAutoDiscountSet)   //自助优惠设定修改[开关]
	e.GET("/customer/applicationInquiry", autoDiscount.GetAutoDiscountList) //自助优惠申请列表
	e.GET("/customer/applicationInfo", autoDiscount.GetInfoMation)          //自助优惠申请详情
	//通过一条优惠申请
	e.PUT("/auto_discount/reject", autoDiscount.RefuseOneDiscountApply)   //拒绝一条优惠申请
	e.GET("/customer/applicationSwitch", autoDiscount.SelfDiscountSwitch) //自助优惠开关列表

	//公告弹窗管理
	noticePopup := new(customer.NoticePopupController)
	e.GET("/customer/animation", noticePopup.GetSiteH5Set)                //H5动画设置查询--列表
	e.PUT("/customer/animation", noticePopup.PutSiteH5Set)                //H5动画状态修改
	e.GET("/customer/noticeAd/setting", noticePopup.GetNoticePopupConfig) //获取单个站点公告弹窗配置
	e.PUT("/customer/noticeAd/setting", noticePopup.PutNoticePopupConfig) //站点公告弹窗配置修改
	e.GET("/customer/commonbullet", noticePopup.NoticePopupList)          //公告弹窗列表
	e.POST("/editor", noticePopup.Editor)                                 //富文本编辑内容(TODO:未测试)

	e.POST("/customer/configurationsetting", noticePopup.PostNoticePopupAdd)  //公告弹框设置--添加站内广告
	e.GET("/customer/configurationsetting", noticePopup.GetNoticePopupSet)    //公告弹窗设置--广告列表
	e.GET("/customer/noticeAd/content", noticePopup.GetNoticePopupSetInfo)    //公告弹窗设置--广告详情
	e.PUT("/customer/configurationsetting", noticePopup.PutNoticePopupSet)    //公告弹窗设置--广告修改
	e.DELETE("/customer/configurationsetting", noticePopup.DelNoticePopupSet) //公告弹窗设置--广告删除

	//代理管理
	agencyManage := new(customer.AgencyController)
	e.GET("/proxy/list", agencyManage.GetAgencyList)           //代理列表
	e.GET("/proxy/occupy", agencyManage.AgencyOccupationRatio) //代理占成

	//视讯账号管理
	videoMember := new(customer.VideoMemberController)

	e.GET("/customer/video/type", videoMember.GetVideoTypeList) //视讯类型下拉框
	e.GET("/customer/video", videoMember.GetVideoMemberList)    //视讯账号查询

	//会员层级
	siteset := new(site.SiteController)
	e.GET("/hierarchicalManag", siteset.GetMemberLevelList) //会员层级列表查询
	e.GET("/site/level/info", siteset.GetMemberLevelInfo)   //会员层级详情查询
	e.POST("/site/level", siteset.PostMemberLevelAdd)       //会员层级添加
	e.PUT("/site/level", siteset.PutMemberLevelUpdate)      //会员层级修改
	e.DELETE("/site/level", siteset.PutMemberLevelDel)      //会员层级删除

	//客户后台管理--层级管理
	memberLevel := new(account.MemberLevelController)
	e.GET("/level/list", memberLevel.Index) //站点会员层级列表

	//站点代理
	e.GET("/ProxyData", siteset.GetSiteAgentsList)         //站点代理数据列表查询
	e.GET("/site/agency/info", siteset.GetSiteAgentsInfo)  //站点代理数据查询
	e.POST("/AddLowerLevel", siteset.PostSiteAgentAdd)     //站点代理数据添加[添加下级]
	e.PUT("/Proxydata/modify", siteset.PutSiteAgentUpdate) //站点代理数据修改
	e.DELETE("/agent/del", siteset.DeleteSiteAgentDel)     //站点代理数据删除

	//模块管理
	/*e.GET("/site/product", siteset.GetSiteProductList)    //站点模块管理（站点商品查询）
	e.POST("/site/product", siteset.InsertSiteProductDel) //站点模块管理（站点商品剔除）*/

	e.GET("/holder/list", siteset.GetSiteAdminInfo)          //开户人列表
	e.POST("/holder/add", siteset.PostSiteAdminAdd)          //添加开户人
	e.PUT("/holder/update", siteset.PutSiteAdminUpdate)      //修改开户人
	e.DELETE("/holder/delete", siteset.PutSiteAdminDel)      //删除开户人
	e.GET("/holder", siteset.EditAccountHolderInfo)          //获取开户人详情
	e.PUT("/holder/disable", siteset.AccountHolderOpenClose) //修改开户人状态
	//站点管理-数据-管理员
	e.PUT("/init/agency/password", siteset.PutInitPassword) //初始化密码为123456
	e.POST("/add/holder", siteset.PostSiteAdmin)            //添加管理员

	//多站点
	e.POST("/oneToch/agent", siteset.PostSiteAgentsAdd)  //一键生成三级代理提交
	e.GET("/site/agency", siteset.ProductAgentsAccount)  //一键生成三级代理账号获取
	e.PUT("/Quotaoperation", siteset.SiteQuota)          //站点额度操作
	e.PUT("/goOnline", siteset.PutSiteOnlinTimeUpdate)   //站点上线时间
	e.POST("/site/addSite", siteset.PostSiteAdd)         // 站点增加
	e.GET("/site/info", siteset.SiteInfoByDomainAndInfo) //站点详情
	e.PUT("/site/edit", siteset.PutSiteUpdate)           //站点修改
	e.GET("/sitemanagement", siteset.GetSiteList)        //站点列表
	e.PUT("/site/put", siteset.PutSiteStatusUpdate)      //站点状态修改
	e.DELETE("/site/delete", siteset.SiteManageDelete)   //站点删除

	//站点管理-全站维护
	e.GET("/site/maintenance", siteset.GetSiteMaintenance)        //获取站点维护id,商品名称，维护原因
	e.GET("/site/select", siteset.SiteIsSelect)                   //获取站点是否选择维护
	e.PUT("/site/maintenance", siteset.PostSiteMaintenanceUpdate) //修改维护

	e.POST("/multistation/addAgent", siteset.PostSiteMore) //多站点-添加代理
	e.GET("/multistation", siteset.GetSiteMoreNews)        //多站点-列表
	e.POST("/multistation/add", siteset.PostSiteChildAdd)  //多站点-添加
	e.PUT("/site/domainEdit", siteset.SiteDomainEdit)      //多站点（站点修改）

	//多站点-前台文案
	e.GET("/iword/type", siteset.SiteCopyType)          //前台文案类型
	e.GET("/iword/list", siteset.SiteCopyList)          //前台首页文案列表
	e.POST("/site/iword/add", siteset.SiteCopyAdd)      //添加前台首页文案
	e.GET("/iword/info", siteset.SiteCopyListInfoOne)   //前台首页文案详情
	e.PUT("/site/iword/update", siteset.SiteCopyUpdate) //修改前台首页文案

	//多站点-前台轮播
	e.GET("/graphic/swiper", siteset.FlashList) //轮播查询
	e.POST("/site/flash/add", siteset.FlashAdd) //轮播添加

	//多站点-站点logo图片管理
	e.GET("/graphic/logo", siteset.LogoList)  //站点logo图片管理列表
	e.POST("/site/logo/add", siteset.LogoAdd) //站点logo图片添加

	//站点报表负数
	/*e.GET("/negative", siteset.GetSiteNegative)       //报表负数查询
	e.POST("/negative", siteset.PostSiteNegativeAdd)  //报表负数添加
	e.PUT("/negative", siteset.PutSiteNegativeUpdate) //报表负数修改*/

	//视讯模板管理
	videoTemp := new(templet.VideoTempController)
	e.GET("/copyTemplate/videoCopy", videoTemp.GetVideoTempList)       //视讯模板查询
	e.POST("/copyTemplate/videoCopy/put", videoTemp.PostVideoTempAdd)  //视讯模板添加
	e.PUT("/copyTemplate/videoCopy/put", videoTemp.PutVideoTempUpdate) //视讯模板修改

	//注册模板管理
	loginTemp := new(templet.LoginTempController)
	e.GET("/copyTemplate/registerCopy", loginTemp.GetLoginTempList)                //注册模板查询
	e.POST("/copyTemplate/registerCopy", loginTemp.PostLoginTempAdd)               //注册模板添加
	e.GET("/copyTemplate/registerCopy/detail", loginTemp.GetLoginTempListDetail)   //注册模板详情
	e.PUT("/copyTemplate/registerCopy", loginTemp.PutLoginTempUpdate)              //注册模板修改
	e.PUT("/copyTemplate/registerCopy/status", loginTemp.PutLoginTempStatusUpdate) //注册模板状态修改

	//视讯下载地址
	sitedown := new(site.DownloadController)
	e.GET("/linkDownload", sitedown.GetDownloadList)            //下载地址列表
	e.GET("/downloadlinks/modify", sitedown.GetSiteDownOne)     //单条下载地址详情
	e.POST("/addvideo", sitedown.PostDownloadAdd)               //添加下载地址
	e.PUT("/download/modify", sitedown.PutDownloadUpdate)       //修改下载地址
	e.PUT("/download/status", sitedown.PutDownloadStatusUpdate) //修改下载地址状态

	//异常会员
	abnormalMember := new(customer.AbnormalMemberController)
	e.GET("/customer/exceptionMember", abnormalMember.GetAbnormalMemberList)       //异常会员列表
	e.PUT("/customer/exceptionMember/handle", abnormalMember.PutAbnormalMemberSet) //异常会员处理

	//红包补数据
	redBag := new(customer.RedBagController)
	e.POST("/customer/editRedBag", redBag.PutRedBagSet) //红包补数据

	//ip开关
	ipSet := new(site.IpSetController)
	e.GET("/IPswitching", ipSet.GetIpSetList)       //ip控制列表查询
	e.POST("/ipSwitch/add", ipSet.PostIpSetAdd)     //ip区间添加
	e.PUT("/ipSwitch/modify", ipSet.PutIpSetUpdate) //Ip区间段修改

	//ip白名单
	e.GET("/IPWhiteList", ipSet.GetIpWhiteList) //Ip白名单列表查询
	e.POST("/white/add", ipSet.PostIpwhiteAdd)  //Ip白名单添加
	e.PUT("/white/put", ipSet.PutIpwhiteUpdate) //Ip白名单修改

	//客户后台栏目管理
	adminBackstage := new(site.AdminBackstargController)
	e.GET("/admin_menu/list", adminBackstage.GetAdminBackstargList)     //客户后台栏目列表查询
	e.POST("/admin_menu/add", adminBackstage.PostAdminBackstargAdd)     //客户后台栏目添加
	e.PUT("/admin_menu/put", adminBackstage.PutAdminBackstargUpdate)    //客户后台栏目修改
	e.DELETE("/admin_menu/delete", adminBackstage.PutAdminBackstargDel) //客户后台栏目删除
	e.PUT("/admin_menu/status", adminBackstage.MenuOpenAndClose)        //客户后台栏目开启禁用
	e.GET("/admin_menu/drop", adminBackstage.GetMoreMenu)               //菜单下拉框

	//代理后台栏目管理
	agencyBackstage := new(site.AgencyBackstargController)
	e.GET("/agent_menu/list", agencyBackstage.GetAgencyBackstargList)     //代理后台栏目列表查询
	e.POST("/agent_menu/add", agencyBackstage.PostAgencyBackstargAdd)     //代理后台栏目添加
	e.PUT("/agent_menu/put", agencyBackstage.PutAgencyBackstargUpdate)    //代理后台栏目修改
	e.DELETE("/agent_menu/delete", agencyBackstage.PutAgencyBackstargDel) //代理后台栏目删除
	e.PUT("/agent_menu/status", agencyBackstage.MenuOpenAndCloseAgency)   //代理后台栏目开启禁用
	e.GET("/agent_menu/drop", agencyBackstage.GetMoreMenuAgency)          //代理菜单下拉框

	//银行列表管理
	bankCard := new(site.BankCardController)
	e.GET("/bank/list", bankCard.GetBankCardList)                    //银行列表查询
	e.POST("/bank/add", bankCard.PostBankCardAdd)                    //银行卡添加
	e.PUT("/bank/put", bankCard.PutBankCardUpdate)                   //银行卡修改
	e.PUT("/bank/status", bankCard.PutBankCardStatusUpdate)          //银行状态修改
	e.DELETE("/bank/delete", bankCard.PutBankCardDel)                //银行卡删除
	e.POST("/bank_card/bank_card_redis", bankCard.PostBankCardRedis) //入款查询
	//站点口令验证
	sitePassword := new(site.SitePasswordController)
	e.GET("/sitepassword", sitePassword.GetSitePassList)              //站点口令列表查询
	e.POST("/sitePassword/modify", sitePassword.PutSitePassUpdate)    //站点口令修改
	e.PUT("/sitePassword/batch/del", sitePassword.PutBatchDelChanges) //站点口令批量停用启用

	//三方相关操作
	onlineCard := new(site.OnlineCardController)
	e.GET("/onlineCard/list", onlineCard.GetOnlineCardList)           //三方列表
	e.PUT("/onlineCard/update", onlineCard.PutOnlineCardStatusUpdate) //同步存储数据库和redis
	e.PUT("/onlineCard/status", onlineCard.PutOnlineCardStatusUpdate) //三方状态启用或禁用
	e.PUT("/onlineCard/info", onlineCard.UpdateOnlineCard)            //网银信息修改
	e.DELETE("/onlineCard", onlineCard.PutOnlineCardDel)              //网银删除
	e.POST("/onlineCard/add", onlineCard.PostOnlineCardAdd)           //添加网银
	//js版本控制
	sitejs := new(site.CdnJsController)
	e.GET("/js/Table", sitejs.GetCdnJsList)       //js版本列表
	e.PUT("/table/modify", sitejs.PutCdnJsUpdate) //js版本修改

	//[平台]所有模块管理
	notemodule := new(note.GameModuleController)
	e.GET("/moduleManagement", notemodule.GetGameTypeList)           //所有游戏种类查询
	e.POST("/note/game/add", notemodule.PostGameTypeAdd)             //添加游戏种类
	e.PUT("/note/game/update", notemodule.PutGameTypeUpdate)         //修改游戏类型
	e.DELETE("/delete/game/delete", notemodule.DelGameType)          //删除游戏类型
	e.GET("/note/platform/list", notemodule.GetGamePlatformList)     //查询所有平台列表
	e.POST("/note/platform/add", notemodule.PostGamePlatformAdd)     //添加平台
	e.PUT("/note/platform/update", notemodule.PutGamePlatformUpdate) //修改平台
	e.DELETE("/delete/platform/delete", notemodule.DelGamePlatform)  //删除平台
	e.GET("/note/product/list", notemodule.GetGameProductList)       //所有游戏查询
	e.POST("/note/product/add", notemodule.PostGameProductAdd)       //添加游戏
	e.PUT("/note/product/update", notemodule.PutGameProductUpdate)   //修改游戏
	e.DELETE("/delete/product/delete", notemodule.DelGameProduct)    //删除游戏

	//视讯白名单控制
	notebet := new(note.NoteBetController)
	e.GET("/note/game/white", notebet.GetGameWhite)             //视讯ip白名单列表
	e.POST("/note/game/whiteAdd", notebet.PostGameWhiteAdd)     //视讯ip白名单添加
	e.PUT("/note/game/whiteUpdate", notebet.PutGameWhiteUpdate) //视讯ip白名单修改
	e.DELETE("/note/game/whiteDel", notebet.DelGameWhite)       //视讯ip白名单删除(硬删除)

	//操作日志
	oR := new(rbac.OperateRecordController)
	e.GET("/operation", oR.GetOperaterecord) //总后台操作记录
	logc := new(rbac.LogController)
	e.GET("/log", logc.GetLogList) //总后台登录日志

	//权限
	column := new(rbac.ColumnController)
	e.GET("/permission", column.GetColumnList)               //权限列表
	e.POST("/permission", column.PostColumnAdd)              //权限添加
	e.GET("/permission/info", column.EditPermission)         //权限详情
	e.PUT("/permission", column.PutColumnUpdate)             //权限修改
	e.DELETE("/permission", column.DelColumn)                //权限删除
	e.PUT("/permission/status", column.PermissionEditStauts) //修改权限状态
	//管理员
	adminRbac := new(rbac.RbacAdminController)
	e.GET("/admin", adminRbac.GetRbacAdminList)                //管理员列表
	e.POST("/admin", adminRbac.PostRbacAdminAdd)               //管理员添加
	e.PUT("/admin", adminRbac.PutRbacAdminUpdate)              //管理员修改
	e.GET("/admin/info", adminRbac.EditAccount)                //管理员详情
	e.PUT("/admin/status", adminRbac.PutRbacAdminStatusUpdate) //管理员状态修改
	e.DELETE("/admin", adminRbac.DeleteAccountDelete)          //管理员删除
	e.PUT("/init/password", adminRbac.PutInitPassword)         //初始化密码为123456
	//权限组管理
	group := new(rbac.RbacGroupController)
	e.GET("/role", group.GetRbacGroupList)             //权限组列表(角色列表)
	e.POST("/role", group.PostRbacGroupAdd)            //权限组添加
	e.GET("/group/one", group.EditRole)                //权限组详情
	e.PUT("/group/put", group.PutRbacGroupUpdate)      //权限组修改
	e.PUT("/role/status", group.RoleEditStauts)        //修改权限组状态
	e.DELETE("/role", group.RoleState)                 //删除权限组
	e.GET("/role/permission", group.RolePermission)    //角色权限列表
	e.POST("/role/permission", group.RolePermissionDo) //设置角色权限
	e.POST("/role/menu", group.RoleMenuDo)             //设置角色菜单
	e.GET("/role/drop", group.RoleList)                //平台账号中的角色下拉框
	e.GET("/role/menu", group.RoleMenu)                //角色菜单列表
	//菜单
	menu := new(rbac.SiteColumnController)
	e.GET("/menu/role", menu.Admin)                    //根据登录人角色获取菜单
	e.GET("/menu_admin/list", menu.GetAdminColumnList) //站点栏目列表(admin)
	e.GET("/menu_agency/list", menu.GetSiteColumnList) //站点栏目列表(agency)
	e.POST("/menu/add", menu.PostSiteColumnAdd)        //平台栏目添加
	e.PUT("/menu/put", menu.PutSiteColumnUpdate)       //平台栏目修改
	e.DELETE("/menu/delete", menu.DelSiteColumn)       //平台栏目删除
	e.PUT("/menu/status", menu.MenuOpenAndClose)       //菜单开启，禁用
	e.GET("/menu/detail", menu.EditMenu)               //菜单详情
	e.GET("/menu/drop", menu.MoreMenuDrop)             //菜单下拉框

	// 维护管理
	maintain := new(site.MaintainController)
	e.GET("/maintain/index", maintain.Index)   // 维护信息 查询
	e.PUT("/maintain/create", maintain.Create) // 维护信息 添加/修改
	e.GET("/maintain/close", maintain.Close)   // 维护信息 关闭

	//维护管理  游戏种类，后台的维护
	maintenance := new(site.MaintenanceController)
	e.GET("/maintenance", maintenance.GetMaintenanceInfo)        //维护页面查询
	e.PUT("/tenance/update", maintenance.PutMaintenNoticeUpdate) //维护信息修改
	e.GET("/tenance/infolist", maintenance.GetInfoList)          //获取视讯 电子下级菜单

	//商品分类
	productType := new(product.ProductController)
	e.POST("/product/type", productType.ProductTypeAdd)          //添加商品分类
	e.GET("/product/type", productType.EditProductType)          //获取商品分类详情
	e.PUT("/product/type", productType.ProductTypeEdit)          //更新商品分类详情
	e.PUT("/product/type/status", productType.ProductTypeStatus) //启用或者禁用商品分类
	e.DELETE("/product/type", productType.ProductTypeDel)        //删除商品分类
	e.GET("/product/type/list", productType.ProductTypeList)     //商品分类下拉框
	e.GET("/product/type/infos", productType.ProductTypeIndex)   //商品分类列表

	//商品管理
	e.POST("/product", productType.ProductAdd)          //添加商品
	e.GET("/product", productType.EditProduct)          //获取商品信息
	e.PUT("/product", productType.ProductEdit)          //更新商品信息
	e.PUT("/product/status", productType.Status)        //更新商品状态
	e.DELETE("/product", productType.DeleteProduct)     //删除商品
	e.GET("/product/list", productType.Index)           //商品列表
	e.GET("/product/drop", productType.ProductListDrop) //商品下拉框

	//mongodb文件系统
	fm := new(filemanager.FileManager)
	e.Any("/filemanager", fm.Handle)

	//站点管理-广告管理
	ac := new(site.AdvertController)
	e.POST("/advert", ac.PostAdvertAdd)               //添加广告
	e.GET("/graphic/adver/detail", ac.GetAdvertInfo)  //获取广告信息
	e.PUT("/advert", ac.PutAdvertUpdate)              //编辑广告
	e.PUT("/advert/status", ac.PutAdvertStatusUpdate) //修改广告状态
	e.PUT("/advert/sort", ac.PutAdvertSortUpdate)     //修改广告排序
	e.GET("/advert/list", ac.GetAdvertList)           //广告列表
}
