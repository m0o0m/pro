'use strict';

appConfig.apiUrls = {
    APP_HOST: "http://localhost:3000",
    //HOST: "http://localhost:9797",
    HOST: "http://api.nothingness.info:3002/mock/53",
    CAPTCHA: "/captcha",
    LOGIN: "/login",
    LANG: {                                                     //语言下拉
        "index": "/langs",
        "cn": "/lang/cn",
        "us": "/lang/us",
        "zh": "/lang/zh"
    },
    ACTIVITY: {
        "index": "/activitys",
        "msgs": "/activity/msgs",
        "notify": "/activity/notify"
    },
    //公告接口
    THIRD_DROPF: "/third/dropf", //站点下拉框
    MEMBER_STATUS:"",
    //维护管理
    MAINTENANCE:"/maintenance",                                    //维护管理列表
    MAIN_TAIN:"/Maintain",                                         //维护-全网维护/维护
    MAINTAIN_PROJECT:"/maintain/project",                          //维护项目管理列表
    COPYWRITING:"/copywriting",                                    //文案管理列表
    COPYWRITING_HISTORY:"/copywriting/history",                    //文案历史
    COPYWRITING_DETAIL:"/copywriting/detali",                      //文案详情
    COPYWRITING_STATUS:"/copywriting/status",                      //文案管理(通过/拒绝)
    COPYWRITING_DEL:"/copywriting/del",                            //删除文案
    SITEPASSWORD:"/sitepassword",                                 // 站点口令设置列表
    SITEPASSWORD_DETALI:"/sitePassword/detali",                   //获取单个口令
    SITEPASSWORD_MODIFY: "/sitePassword/modify",                  //修改单个口令
    LINKDOWNLOAD:"/linkDownload",                                 //下载链接地址列表
    ADD_VIDEO:"/addvideo",                                        //添加地址
    DOWNLOADLINKS_MODIFY:"/downloadlinks/modify",                 //获取单个下载链接
    DOWNLOAD_MODIFY:"/download/modify",                           //修改单个下载链接
    DOWNLOAD_STATUS:"/download/status",                           //修改单个下载链接状态
    JS_WEP:"/js/wep",                                             //js版本号-wep
    JS_TABLE:"/js/Table",                                         //js版本号-总表
    JS_PC:"/js/pc",                                               //js版本号-PC
    JS_DETAIL:"/js/detail",                                       //JS获取详情
    JS_del:"/js/del",                                             //删除PC/WEP
    TABLE_MODIFY:"/table/modify",                                 //总表编辑
    GENERATE_WEP:"/generate/wep",                                 //生成wep版本号
    GENERATE_PC:"/generate/pc",                                  //生成PC版本号
    IP_SWITCHING:"/IPswitching",                                 //IP开关操作
    IP_WHITE_LIST:"/IPWhiteList",                                //IP白名单
    IP_SWITCH_ADD:"/ipSwitch/add",                               //添加ip开关操作
    IP_SWITCH_MODIFY:"/ipSwitch/modify",                         //修改ip开关操作
    SITE_MANAGEMENT:"/site/list",                           //站点管理
    MODULE_MANAGEMENT:"/moduleManagement",                      //模块管理
    NEGATIVE:"/negative",                                       //获取负数列表
    MULTISTATION:"/multistation",                               //获取多站点

    //网站资讯管理--图文编辑
    GRAPHIC_LOGO: "/graphic/logo",                                          //LOGO管理
    GRAPHIC_LOGO_ENCLOSURE: "/graphic/logo/enclosure",                      //LOGO管理--附件管理
    GRAPHIC_LOGO_ENCLOSURE_SELECT: "/graphic/logo/enclosure/select",        //LOGO管理--附件选择
    GRAPHIC_FLOAT: "/graphic/float",                                        //浮动图片
    ENCLOSURE: "/enclosure",                                                //附件管理
    FLOAT_IMG: "/floatImg",                                                 //浮动图片管理
    FLOAT_IMG_ENCLOSURE: "/floatImg/enclosure",                             //浮动图片管理--附件管理
    FLOAT_IMG_ENCLOSURE_SELECT: "/floatImg/enclosure/select",               //浮动图片管理--附件选择
    GRAPHIC_SWIPER: "/graphic/swiper",                                      //轮播图片
    GRAPHIC_SWIPER_STATUS: "/graphic/swiper/status",                        //轮播图片--启用禁用
    GRAPHIC_SWIPER_STORAGE: "/graphic/swiper/storage",                      //轮播图片--存储案件
    GRAPHIC_SWIPER_ENCLOSURE: "/graphic/swiper/enclosure",                  //轮播图片--附件
    GRAPHIC_SWIPER_ENCLOSURE_SELECT: "/graphic/swiper/enclosure/select",    //轮播图片--附件选择
    GRAPHIC_NOTICE: "/graphic/notice",                                      //公告弹框
    GRAPHIC_NOTICE_AD: "/graphic/notice/ad",                                //公告弹框--弹框广告设置
    GRAPHIC_NOTICE_CONTENT: "/graphic/notice/content",                       //公告弹框--广告编辑
    //网站资讯--文案编辑
    DISCOUNT_DEL:"/discount/del",                                           // 优惠活动--删除	DELETE/discount/del
    LINE_DETECTION_MODIFY_SUB:"/lineDetection/modifySub",                   // 线路检测--编辑--提交	POST/lineDetection/modifySub
    LINE_DETECTION_ADD_SUB:"/lineDetection/addSub",                         // 线路检测--新增--提交	POST/lineDetection/addSub
    LINE_DETECTION:"/lineDetection",                                        // 线路检测	GET/lineDetection
    LINE_DETECTION_DEL:"/lineDetection/del",                                // 线路检测--删除	DEL/lineDetection/del
    WAP_DISCOUNT_CONTENT_DEL:"/wapDiscount/content/del",                    // WAP优惠活动--优惠内容--删除	DELETE/wapDiscount/content/del
    DEPOSIT_COPY:"/depositCopy",                                            // 存款文案	GET/depositCopy
    DEPOSIT_COPY_SUB:"/depositCopy/sub",                                    // 存款文案--编辑--提交	POST/depositCopy/sub
    DEPOSIT_EDITOR:"/depositEditor",                                        // 存款文案--编辑内容	GET/depositEditor
    DEPOSIT_EDITOR_SUB:"/depositEditor/sub",                                // 存款文案--编辑内容--提交	POST/depositEditor/sub
    DEPOSIT_COPY_KEEP:"/depositCopy/keep",                                  // 存款文案--储存案件	POST/depositCopy/keep
    DEPOSIT_COPY_MODULE:"/depositCopy/module",                              // 存款文案--模板选择	GET/depositCopy/module
    DEPOSIT_COPY_MODULE_CHOICE:"/depositCopy/moduleChoice",                 // 存款文案--模板选择--模板选择	POST/depositCopy/moduleChoice
    DEPOSIT_COPY_MODULE_KEEP:"/depositCopy/moduleKeep",                     // 存款文案--模板选择--储存案件	POST/depositCopy/moduleKeep
    REGISTER_COPY:"/registerCopy",                                          // 注册文案	GET/registerCopy
    DISCOUNT_SUB:"/discount/sub",                                           // 优惠活动--编辑--提交	POST/discount/sub
    REGISTER_EDITOR:"/registerEditor",                                      // 注册文案--编辑内容	GET/registerEditor
    REGISTER_EDITOR_SUB:"/registerEditor/sub",                              // 注册文案--编辑内容--提交	POST/registerEditor/sub
    REGISTER_COPY_KEEP:"/registerCopy/keep",                                // 注册文案--储存案件	POST/registerCopy/keep
    REGISTER_COPY_MODULE:"/registerCopy/module",                            // 注册文案--模板选择	GET/registerCopy/module
    REGISTER_COPY_MODULE_CHOICE:"/registerCopy/moduleChoice",               // 注册文案--模板选择--模板选择	POST/registerCopy/moduleChoice
    REGISTER_COPY_MODULE_KEEP:"/registerCopy/moduleKeep",                   // 注册文案--模板选择--储存案件	POST/registerCopy/moduleKeep
    DISCOUNT:"/discount",                                                   // 优惠活动	GET/discount
    REGISTER_COPY_SUB:"/registerCopy/sub",                                  // 注册文案--编辑--提交	POST/registerCopy/sub
    DISCOUNT_CONTENT:"/discount/content",                                   // 优惠活动--优惠内容	GET/discount/content
    DISCOUNT_MODIFY_CONTENT:"/discount/modifyContent",                      // 优惠活动--编辑内容	GET/discount/modifyContent
    DISCOUNT_MODIFY_SUB:"/discount/modifySub",                              // 优惠活动--编辑内容--提交	POST/discount/modifySub
    DISCOUNT_UPDATE:"/discount/update",                                     // 优惠活动--上传--提交	POST/discount/update
    WAP_DISCOUNT_C_M_SUB:"/wapDiscount/content/modifySub",                  // WAP优惠活动--优惠内容--编辑内容--提交	POST/wapDiscount/content/modifySub
    DISCOUNT_KEEP:"/discount/keep",                                         // 优惠活动--储存案件	POST/discount/keep
    DISCOUNT_ADD_SUB:"/discount/addSub",                                    // 优惠活动--新增--提交	POST/discount/addSub
    DISCOUNT_WIDTH:"/discount/width",                                       // 优惠活动--优惠宽度编辑	GET/discount/width
    DISCOUNT_WIDTH_SUB:"/discount/widthSub",                                // 优惠活动--优惠宽度编辑--提交	POST/discount/widthSub
    DISCOUNT_C_ADD_SUB:"/discount/content/addSub",                          // 优惠活动--优惠内容--添加优惠活动--提交	POST/discount/content/addSub
    DISCOUNT_C_M_SUBMIT:"/discount/content/modifySubmit",                   // 优惠活动--优惠内容--编辑--提交	POST/discount/content/modifySubmit
    WAP_DISCOUNT_C_M_SUBMIT:"/wapDiscount/content/modifySubmit",            // WAP优惠活动--优惠内容--编辑--提交	POST/discount/content/modifySubmit
    DISCOUNT_C_WIDTH_SUB:"/discount/content/widthSub",                      // 优惠活动--优惠内容--优惠宽度编辑--提交	POST/discount/content/widthSub
    DISCOUNT_C_UPDATE:"/discount/content/update",                           // 优惠活动--优惠内容--上传--提交	POST/discount/content/update
    DISCOUNT_C_M_CONTENT:"/discount/content/modifyContent",                 // 优惠活动--优惠内容--编辑内容	GET/discount/content/modifyContent
    DISCOUNT_C_M_SUB:"/discount/content/modifySub",                         // 优惠活动--优惠内容--编辑内容--提交	POST/discount/content/modifySub
    DISCOUNT_C_DEL:"/discount/content/del",                                 // 优惠活动--优惠内容--删除	DELETE/discount/content/del
    DISCOUNT_C_KEEP:"/discount/content/keep",                               // 优惠活动--优惠内容--储存案件	POST/discount/content/keep
    WAP_DISCOUNT_C_KEEP:"/wapDiscount/content/keep",                        // WAP优惠活动--优惠内容--储存案件	POST/discount/content/keep
    WAP_DISCOUNT:"/wapDiscount",                                            // WAP优惠活动	GET/wapDiscount
    WAP_DISCOUNT_ADD_SUB:"/wapDiscount/addSub",                             // WAP优惠活动--新增--提交	POST/wapDiscount/addSub
    WAP_DISCOUNT_SUB:"/wapDiscount/sub",                                    // WAP优惠活动--编辑--提交	POST/wapDiscount/sub
    WAP_DISCOUNT_C:"/wapDiscount/content",                                  // WAP优惠活动--优惠内容	GET/wapDiscount/content
    WAP_DISCOUNT_M_C:"/wapDiscount/modifyContent",                          // WAP优惠活动--编辑内容	GET/wapDiscount/modifyContent
    WAP_DISCOUNT_M_SUB:"/wapDiscount/modifySub",                            // WAP优惠活动--编辑内容--提交	POST/wapDiscount/modifySub
    WAP_DISCOUNT_UPDATE:"/wapDiscount/update",                              // WAP优惠活动--上传--提交	POST/wapDiscount/update
    WAP_DISCOUNT_DEL:"/wapDiscount/del",                                    // WAP优惠活动--删除	DELETE/wapDiscount/del
    WAP_DISCOUNT_C_ADD_SUB:"/wapDiscount/content/addSub",                   // WAP优惠活动--优惠内容--添加优惠活动--提交	POST/wapDiscount/content/addSub
    WAP_DISCOUNT_C_UPDATE:"/wapDiscount/content/update",                    // WAP优惠活动--优惠内容--上传--提交	POST/wapDiscount/content/update
    WAP_DISCOUNT_C_M_C:"/wapDiscount/content/modifyContent",                // WAP优惠活动--优惠内容--编辑内容	GET/wapDiscount/content/modifyContent
    HOME_COPY:"/homeCopy",                                                  // 首页文案	GET/homeCopy
    HOME_COPY_KEEP:"/homeCopy/keep",                                        // 首页文案--储存案件	POST/homeCopy/keep
    HOME_COPY_SUB:"/homeCopy/sub",                                          // 首页文案--编辑--提交	PUT/homeCopy/sub
    HOME_EDITOR:"/homeEditor",                                              // 首页文案--编辑内容	GET/homeEditor
    HOME_EDITOR_SUB:"/homeEditor/sub",                                      // 首页文案--编辑内容--提交	POST/homeEditor/sub
    HIERrchicaldata:"/hierarchicalData",                                    //数据-层级数据
    HIERrchicaldataModify:"/HierarchicalData/modify",                       //数据-层级数据-修改
    HIERrchicaldataDel:"/HierarchicalData/del",                              //数据-层级数据-删除
    HIERrchicaldataAdd:"/HierarchicalData/add",                              //数据-层级数据-添加
    PROXYDATA:"/ProxyData",                                                 //数据-代理数据
    PROXYDATA_MODIFY:"/Proxydata/modify",                                   //代理编辑
    ADDLOWERLEVEL:"/AddLowerLevel",                                         //添加代理下级
    AGENT_DEL:"/agent/del",                                                 //数据-代理数据-删除
    VIDEOCONFIFURATION:"/Videoconfiguration",                               //数据-视讯配置
    DATA_ADMIN:"/data/admin",                                               //数据-管理员
    DATA_ADMIN_MODIFY:"/data/admin/modify",                                 //数据-管理员（修改）
    DATA_ADMIN_ADD:"/data/admin/add",                                       //数据-管理员（添加）
    DATA_ADMIN_DEL:"/data/admin/del",                                       //数据-管理员(删除)
    MAINTENANCESETTINGS:"/Maintenancesettings",                             //全站维护设置
    SAVESETTINGS:"/Savesettings",                                           //保存设置
    SITE_ADD:"/site/add",                                                   //站点管理(添加站点)
    QUOTAOPERATION:"/Quotaoperation",                                       //站点管理(修改额度操作)
    MODULAR:"/Modular",                                                     //模块管理
    GOONLINE:"/goOnline",                                                   //站点管理(上线)
    NEGATIVE_ADD:"/negative/add",                                           //站点管理（负数新增）
    MULTISTATION_ADD:"/multistation/add",                                   //站点管理（多站点添加）
    MULTISTATION_ADDAGENT:"/multistation/addAgent",                         //站点管理（多站点添加代理）
    ONETOCH_AGENT:"/oneToch/agent",                                         //站点管理（一键生成三级代理添加）
    MULTISTATION_MODIFY:"/multistation/modify",                             //站点管理（多站点编辑）
    SITE_MODIFY:"/site/modify",                                             //站点管理（修改站点）
    SYSTEM_MOTICE_LIST:"/system/notice",                            //公告管理列表
    ADD_TYPE_DROP:"/add/type/drop",                                 //类型下拉框
    ADD_SYSTEM_MOTICE:"/systerm/notice",                            //添加最新消息
    DEL_SYSTEM_NOTICE:"/systerm/notice",                            //删除信息
    GET_NOTICE_NEWS:"/notice/news",                                 //获取消息信息
    POST_NOTICE_NEWS:"/notice/news",                                //提交消息信息
    GRAPHIC_ADVERTISEMENT: "/graphic/advertisement",                        //广告管理
    GRAPHIC_ADVERTISEMENT_SORT: "/graphic/advertisement/sort",              //广告管理--排序
    GRAPHIC_ADVERTISEMENT_STATUS: "/graphic/advertisement/status",          //广告管理--启用禁用
    GRAPHIC_ADVERTISEMENT_DETAIL: "/graphic/adver/detail",          //广告管理--详情
    //文案模板管理
    COPYTEMPLATE:"/copyTemplate/registerCopy",                     //注册文案模板
    COPY_STATUS:"/copyTemplate/registerCopy/status",               //状态修改
    COPYTEMPLATE_VIDEO:"/copyTemplate/videoCopy",                  //视讯文案列表
    PUT_VIDEOCOPY:"/copyTemplate/videoCopy/put",                   //视讯修改
    POST_VIDEOCOPY:"/copyTemplate/videoCopy/put",                  //视讯添加
    ADD_REGISTER:"/addRegister",                                   //注册文案-添加
    GET_ADD_REGISTER:"/addRegister",                               //注册文案-编辑获取
    PUT_ADD_REGISTER:"/addRegister",                               //注册文案-编辑提交



    //客户后台管理
    CUSTOMER_VIDEO:"/customer/video",                               //视讯账号管理
    CUSTOMER_VIDEO_TYPE:"/customer/video/type",                     //视讯账号管理--视讯类别
    CUSTOMER_EXCEPTIONMEMBER:"/customer/exceptionMember",           //异常会员查询
    CUSTOMER_EXCEPTIONMEMBER_HANDLE:"/customer/exceptionMember/handle",//异常会员--处理
    CUSTOMER_COMMONBULLET:"/customer/commonbullet",                 //公告弹框管理--公告弹框
    CUSTOMER_ANIMATION:"/customer/animation",                       //公告弹框管理--h5动画
    CUSTOMER_CONFIGURATIONSETTING:"/customer/configurationsetting", //站内广告
    CUSTOMER_NOTICEAD_SETTING:"/customer/noticeAd/setting",         //弹框广告设置
    CUSTOMER_NOTICEAD_CONTENT:"/customer/noticeAd/content",         //广告详情
    CUSTOMER_PREFERENTIALQUERY:"/customer/preferentialQuery",        //优惠查询
    CUSTOMER_PREFERENTIAL_LIST:"/customer/preferential/list",        //优惠查询--列表
    CUSTOMER_APPLICATIONINQUIRY:"/customer/applicationInquiry",      //自助优惠申请查询
    CUSTOMER_APPLICATIONSWITCH:"/customer/applicationSwitch",        //自助优惠开关
    PROXY_LIST:"/proxy/list",                                       //代理管理列表
    CHILDERN_SITE:"/children/site",                                 //子账号站点下拉
    USER_LIST:"/customer/userManagement",                           //用户列表
    USER_STATUS:"/user/status",                                     //修改状态
    USER_INFORMENT:"/user/informent",                               //会员详细资料
    LOGIN_LOG:"/siteLoginLog",                                      //日志管理--登录日志列表
    SITE_LOG:"/siteDoLog",                                          //日志管理--操作日志
    AUTOAUDIT:"/autoAudit",                                         //日志管理--自动稽核
    HIERARCHICAL_LIST:"/hierarchicalManag",                         //层级管理列表
    CHILD_LIST:"/child/list",                                       //子账号列表

    //财务报表
    FINANCE_HIERARCHICAL:"/finance/hierarchical",                  //层级设定
    FINANCE_HIERARCHICALMANAGER:"/finance/hierarchicalManager",    //层级管理
    FINANCE_CASH:"/finance/cash",                                  //现金报表
    FINANCE_CASH_TYPE:"/finance/cash/type",                        //现金报表--方式下拉
    FINANCE_DATACENTER:"/finance/dataCenter",                      //数据中心
    FINANCE_DATACENTER_TYPE:"/finance/dataCenter/type",            //数据中心--类型下拉
    FINANCE_REFERENTIAL:"/finance/referential",                    //优惠统计
    FINANCE_INCOME:"/finance/income",                              //入款统计
    FINANCE_INCOME_WAY:"/finance/income/way",                      //入款统计--入款方式下拉
    FINANCE_SUMMARY:"/finance/summary",                            //出入款账目汇总
    FINANCE_ARREARS:"/finance/arrears",                            //催款查询
    FINANCE_ARREARS_PRESS:"/finance/arrears/press",                //催款查询--催款
    FINANCE_REPORT_MODULE:"/finance/report/module",                //报表辅助查询--获取模块
    REPORT:"/report",                                              //报表统计
    BILLQUERY:"/billQuery",                                        //账单查询
    BILLQUERY_ISSUED:"/billQuery/issued",                          //账单查询--下发
    FINANCE_REPORTSTATISTICS:"/finance/reportStatistics",          //报表统计--站点数据统计
    FINANCE_REPORTSTATISTICS_TYPE:"/finance/reportStatistics/type",  //报表统计--类型下拉
    FINANCE_REPORT_GAME:"/finance/report/game",                     //报表查询--游戏
    FINANCE_QUOTANUM:"/finance/quotaNum",                           //额度统计
    FINANCE_QUOTARECORD:"/finance/quotaRecord",                     //额度记录
    TRANSACTION_TYPE:"/transaction/type",                           //额度记录--交易类型
    VIDEO_TYPE:"/video/type",                                       //额度记录--视讯类型
    TRANSACTION_CATEGORY:"/transaction/category",                   //额度记录--交易类别
    FINANCE_RECHARGERECORD:"/finance/rechargeRecord",               //充值记录
    FINANCE_THIRD:"/finance/third",                                 //入款管理--第三方管理
    FINANCE_BANK:"/finance/bank",                                   //入款管理--银行卡管理
    FINANCE_HIERARCHY_DROP:"/finance/hierarchy/drop",               //入款管理--层级下拉
    //财务报表--报表管理
    REPORTQUERY:"/reportquery",                                     //报表查询
    SHAREHOLDERS_STATEMENT:"/shareholders/statement",               //获取股东报表
    GENERAL_GENERATION_REPORT:"/general/generation/report",         //获取总代
    PROXY_REPORT:"/proxy/report",                                   //获取代理
    MEMBERSHIP_REPORT:"/membership/report",                         //获取会员

    // 管理员
    ROLE_DROP:"/role/drop",                                         // 角色下拉框
    GET_ADMIN:"/admin",                                             // 平台管理	GET/admin
    PUT_ADMIN:"/admin",                                             // 平台管理--修改	PUT/admin
    ADMIN_STATUS:"/admin",                                          // 平台管理--修改状态	PUT/admin/status
    DEL_ADMIN:"/admin",                                             // 平台管理--删除	DELETE/admin
    POST_ADMIN:"/admin",                                            // 平台管理--添加	POST/admin
    ADMIN_INFO:"/admin/info",                                       // 平台管理--获取详细	GET/admin/info
    COMBO_DROP:"/combo/drop",                                       // 开户人管理--添加--下一步--套餐
    HOLDER_LIST:"/holder/list",                                     // 开户人管理	GET /holder/list
    HOLDER_ADD:"/holder/add",                                       // 开户人管理--添加	POST /holder/add
    HOLDER:"/holder",                                               // 开户人管理--修改	GET /holder
    HOLDER_UPDATA:"/holder/updata",                                 // 开户人管理--修改--提交	PUT /holder/updata
    HOLDER_DISABLE:"/holder/disable",                               // 开户人管理--修改状态	PUT /holder/disable
    HOLDER_DEL:"/holder/delete",                                    // 开户人管理--删除	DELETE /holder/delete
    GET_ROLE:"/role",                                               // 角色管理	GET/role
    ROLE_STATUS:"/role/status",                                     // 角色管理--修改状态	PUT /role/status
    DEL_ROLE:"/role",                                               // 角色管理--删除	DELETE/role
    POST_ROLE:"/role",                                              // 角色管理--添加	POST/role
    ROLE_PERMISSION_POST:"/role/permission",                        // 角色管理--权限配置	POST/role/permission
    ROLE_PERMISSION_GET:"/role/permission",                         // 角色管理--权限配置--修改	GET /role/permission
    ROLE_MENU_GET:"/role/menu",                                     // 角色管理--菜单	GET/role/menu
    ROLE_MENU_POST:"/role/menu",                                    // 角色管理--菜单--修改	POST/role/menu
    MENU_ADMIN:"/menu_admin/list",                                  // 菜单管理(平台)admin	GET/menu_admin/list
    MENU_DEL:"/menu/delete",                                        // 菜单管理--删除	DELETE/menu/delete
    MENU_DROP:"/menu/drop",                                         // 菜单管理--修改--根据id取一级二级菜单	GET/menu/drop
    MENU_PUT:"/menu/put",                                           // 菜单管理--修改提交	PUT/menu/put
    MENU_ADD:"/menu/add",                                           // 菜单管理--添加	POST/menu/add
    MENU_STATUS:"/menu/status",                                     // 菜单管理--状态修改	PUT/menu/status
    MENU_DETAIL:"/menu/detail",                                     // 菜单管理--详情	GET/menu/detail
    MENU_AGENCY:"/menu_agency/list",                                // 菜单管理(代理)agency	GET/menu_agency/list
    PRODUCT_LIST:"/product/list",                                   // 商品管理	GET/product/list
    PRODUCT_DEL:"/product",                                         // 商品管理--删除	DELETE/product
    PRODUCT_STATUS:"/product/status",                               // 商品管理--修改状态	PUT/product/status
    PRODUCT_PUT:"/product",                                         // 商品管理--修改提交	PUT/product
    PRODUCT_GET:"/product",                                         // 商品管理--获取商品详情	GET/product
    PRODUCT_POST:"/product",                                        // 商品管理--添加	POST/product
    PRODUCT_TYPE_INFO:"/product/type/infos",                        // 商品管理--商品类型	GET/product/type/infos
    PRODUCT_TYPE_DEL:"/product/type",                              // 商品管理--类型管理--删除	DELETE/product/type
    PRODUCT_TYPE_STATUS:"/product/type/status",                     // 商品管理--类型管理--修改状态	PUT/product/type/status
    PRODUCT_TYPE_PUT:"/product/type",                               // 商品管理--类型管理--修改提交	PUT/product/type
    PRODUCT_TYPE_GET:"/product/type",                               // 商品管理--类型管理--获取商品详情	GET/product/type
    PRODUCT_TYPE_POST:"/product/type",                              // 商品管理--类型管理--添加	POST/product/type
    PRODUCT_TYPES_INFO:"/product/types/infos",                      // 商品管理--类型管理	GET/product/types/infos
    PERMISSION_POST:"/permission",                                  // 功能管理--添加	POST/permission
    PERMISSION_DEL:"/permission",                                   // 功能管理--删除	DELETE/permission
    PERMISSION_STATUS:"/permission/status",                         // 功能管理--修改状态	PUT/permission/status
    PERMISSION_PUT:"/permission",                                   // 功能管理--修改提交	PUT/permission
    PERMISSION_GET:"/permission",                                   // 功能管理	GET/permission
    PERMISSION_INFO:"/permission/info",                             // 功能管理--详情	GET/permission/info
    PRODUCT_TYPE:"/product_type",                                   // 套餐管理--配置--搜索	GET/product_type
    COMBO_PLATFORM_POST:"/combo_platform",                          // 套餐管理--配置--提交	POST/combo_platform
    COMBO_PLATFORM_GET:"/combo_platform",                           // 套餐管理--获取配置详情	GET/combo_platform
    COMBO_DEL:"/combo",                                             // 套餐管理--删除	DELETE/combo
    COMBO_PUT:"/combo",                                             // 套餐管理--修改提交	PUT/combo
    COMBO_INFO:"/combo/info",                                       // 套餐管理--详情	GET/combo/info
    COMBO_STATUS:"/combo/status",                                   // 套餐管理--修改状态	PUT/combo/status
    COMBO_POST:"/combo",                                            // 套餐管理--新增	POST/combo
    COMBO_GET:"/combo",                                             // 套餐管理	GET/combo
    SITE_CLOUMN_PRIVATE:"/siteCloumn/private",                      // 站点栏目--私有	POST/siteCloumn/private
    SITE_CLOUMN_POST:"/siteCloumn",                                 // 站点栏目--添加	POST/siteCloumn
    SITE_CLOUMN_SYNCHRO:"/siteCloumn/synchro",                      // 站点栏目--栏目同步	POST/siteCloumn/synchro
    SITE_CLOUMN_DEL:"/siteCloumn",                                  // 站点栏目--删除	DELETE/siteCloumn
    SITE_CLOUMN_PUT:"/siteCloumn",                                  // 站点栏目--修改提交	PUT/siteCloumn
    SITE_CLOUMN_GET:"/siteCloumn",                                  // 站点栏目	GET/siteCloumn
    LOG:"/log",                                                     // 日志管理	GET/log
    OPERATION:"/operation",                                         // 操作记录	GET/operation
    BANKIN_OUT:"/bankin_out/out",                                   // 出入款管理--出款管理	GET/bankin_out/out
    BANKIN_IN:"/bankin_out/in",                                     // 出入款管理--入款管理	GET/bankin_out/in
    CHILD_STATUS:"/child/status",                                   // 子账号管理--状态修改	PUT/child/status
    CHILD_PUT:"/child",                                             // 子账号管理--修改提交	PUT/child
    AUDIT_LOG:"/audit/log",                                         // 稽核日志--稽核日志	GET/audit/log
    AUDIT_RECORD:"/audit/record",                                   // 稽核日志--稽核列表	GET/audit/record
    NOTICE_KEY:"/customer/noticeKey" ,                               // 公告密钥管理	GET/customer/noticeKey

    FILE_MANAGER:"/filemanager"
};


appConfig.option = {
    option_state: [
        {name: "在线", value: "1"},
        {name: "全部", value: "-1"}
    ],
    option_status:[
        {name:"禁用",value:"2"},
        {name:"启用",value:"1"}
    ],
    option_online:[
        {name:"在线",value:"1"},
        {name:"离线",value:"2"}
    ],
    option_hidden:[
        {name:"全部",value:"-1"},
        {name:"显示",value:"0"},
        {name:"隐藏",value:"1"}
    ],
    option_sort:[
        {name:"新增日期",value:"create_time"},
        {name:"代理名称",value:"username"},
        {name:"代理账号",value:"account"}
    ],
    option_sort_1:[
        {name:"新增日期",value:"create_time"},
        {name:"股东名称",value:"username"},
        {name:"股东账号",value:"account"}
    ],
    option_sort_2:[
        {name:"新增日期",value:"create_time"},
        {name:"总代理名称",value:"username"},
        {name:"总代理账号",value:"account"}
    ],
    option_sortBig:[
        {name:"正序",value:"true"},
        {name:"倒序",value:"false"}
    ],
    option_types:[
        {name:"账号", value:1},
        {name:"注册ip",value:3},
        {name:"登录ip",value:4},
        {name:"手机",value:5},
        {name:"银行卡",value:6},
        {name:"邮箱",value:7},
        {name:"qq",value:8},
        {name:"微信",value:9}
    ],
    option_income_type:[
        {name:"公司入款",value:1},
        {name:"线上入款",value:2},
        {name:"人工入款",value:3},
        {name:"人工出款",value:4},
        {name:"线上出款",value:5}
    ],
    option_is_type:[
        {name:"会员账号", value:1},
        {name:"代理账号",value:2}
    ],
    option_is_order:[
        {name:"总金额", value:1},
        {name:"总笔数",value:2}
    ],
    option_type_m:[
        {name:"admin"},
        {name:"agency"}
    ],
    option_reg:[
        {name:"pc端",value:"1"},
        {name:"android端",value:"3"},
        {name:"ios端",value:"4"},
        {name:"wap端",value:"2"}
    ],
    option_type : [
        {name: "微信", value: "wechat"},
        {name: "电话号", value: "skype"},
        {name: "邮箱", value: "email"},
        {name: "qq", value: "qq"}
    ],
    option_type1 : [
        {name: "平台管理", "id": 2},
        {name: "代理平台", "id": 1}
    ],
    option_states:[
        {"status": "未处理", "id": "2"},
        {"status": "已通过", "id": "1"}
    ],
    option_order : [
        {name: "倒序", value: "true"},
        {name: "正序", value: "false"}
    ],
    option_level : [
        {name: "股东", value: "1"},
        {name: "总代理", value: "2"},
        {name: "代理", value: "3"},
        {name: "会员", value: "4"},
        {name: "全部", value: ""}
    ],
    option_vague : [
        {name: "是", value: "1"},
        {name: "否", value: "0"}
    ],
    option_commodity : [
        {name: "商品名称", value: "product_name"},
        {name: "商品id", value: "product_id"}
    ],
    option_method : [
        {name: "POST"},
        {name: "GET"},
        {name: "PUT"},
        {name: "HEAD"},
        {name: "DELETE"},
        {name: "OPTIONS"},
        {name: "TRACE"}
    ],
    option_system : [
        {name: "未反水", value: 1},
        {name: "返水", value: 2}
    ],
    option_detail : [
        {name: "已冲销", "id": 1},
        {name: "未冲销", "id": 2}
    ],
    option_site : [
        {name: "站点2", "id": 2},
        {name: "站点3", "id": 3},
        {name: "站点4", "id": 4},
        {name: "站点5", "id": 5},
        {name: "站点1", "id": 1}
    ],
    option_deposit:[
        {name: "人工存入", "id": 1},
        {name: "存款优惠", "id": 2},
        {name: "负数额度归零", "id": 3},
        {name: "取消出款", "id": 4},
        {name: "返点优惠", "id": 5},
        {name: "活动优惠", "id": 6},
        {name: "其他", "id": 7}
    ],
    option_deposit_type:[
        {name: "重复出款", "id": 8},
        {name: "公司入款误存", "id": 9},
        {name: "公司负数回冲", "id": 10},
        {name: "手动申请出款", "id": 11},
        {name: "扣除非法下注派彩", "id": 12},
        {name: "放弃存款优惠", "id": 13},
        {name: "其他", "id": 14}
    ],
    option_way:[
        {name: "额度转换", "id": 1},
        {name: "体育下注", "id": 2},
        {name: "彩票下注", "id": 3},
        {name: "彩票派彩", "id": 4},
        {name: "EG电子下注", "id": 5},
        {name: "EG彩票下注", "id": 6},
        {name: "EG彩票派彩", "id": 7}
    ],
    shebei:[
        {name: "pc", "id": 1},
        {name: "wap", "id": 2},
        {name:"全部","id":3 }
    ],
    money_status:[
        {name:"全部",id:0},
        {name:"未出款",id:5},
        {name:"已出款",id:1},
        {name:"已拒绝",id:4},
        {name:"已取消",id:3},
        {name:"预备出款",id:2}

    ],
    automatic:[
        {name:"是",id:1},
        {name:"否",id:2},
        {name:"请选择",id:0}
    ],
    select_by:[
        {name:"账号",id:1},
        {name:"操作者",id:2},
        {name:"请选择",id:0}
    ],
    refresh_time:[
        {name:"30s",id:"30000"},
        {name:"60s",id:"60000"},
        {name:"90s",id:"90000"},
        {name:"120s",id:"120000"},
        {name:"请选择",id:"0"}
    ],
    refresh_time1:[
        {name:"10s",id:"10000"},
        {name:"30s",id:"30000"},
        {name:"60s",id:"60000"},
    ],
    option_source:[
        {name:"PC","id":1},
        {name:"WAP","id":2},
        {name:"APP","id":3},
    ],
    option_amount:[
        {name:"10000","id":10000},
        {name:"20000","id":20000},
        {name:"50000","id":50000},
    ],
    option_time : [
        {name: "转时区", "id": 1},
        {name: "非转时区", "id": 2}
    ],
    option_money_type: [
        {name:"額度轉換","id":2},
        {name:"体育下注","id":21},
        {name:"体育派彩","id":22},
        {name:"彩票下注","id":23},
        {name:"冲销明细","id":24},
        {name:"彩票派彩","id":25},
        {name:"注单无效","id":26},
        {name:"注单取消","id":27},
        {name:"线上入款","id":28},
        {name:"公司入款","id":29},
        {name:"线上入款不含优惠","id":30},
        {name:"公司入款不含优惠","id":31},
        {name:"线上取款","id":32},
        {name:"公司/线上入款","id":33},
        {name:"优惠退水","id":34},
        {name:"自助返水","id":35},
        {name:"优惠活动","id":36},
        {name:"人工存入","id":37},
        {name:"人工取出","id":38},
        {name:"人工存款與取款","id":39},
        {name:"入款明细","id":40},
        {name:"出款明细","id":41},
        {name:"会员退佣","id":42},
        {name:"会员退佣冲销","id":43},
        {name:"额度掉单(取出)","id":44},
        {name:"额度掉单(存入)","id":45},
        {name:"会员推广返佣","id":46}
    ],
    option_refresh:[
        {name:"30秒","id":30},
        {name:"60秒","id":60},
        {name:"120秒","id":120},
        {name:"180秒","id":180}
    ],
    option_money_time:[
        {name:"北京","id":1},
        {name:"美东","id":2}
    ],
    option_deposit_item:[
        {name:"人工存入","id":1},
        {name:"存款优惠","id":2},
        {name:"负数额度归零","id":3},
        {name:"取消出款","id":4},
        {name:"返点优惠","id":5},
        {name:"活动优惠","id":6}
    ],
    option_handle:[
        {name:"未处理", "id": 3},
        {name:"已确认", "id": 1},
        {name:"已取消","id":2}
    ],
    option_query_criteria:[
        {name:"账号","id":1},
        {name:"订单号","id":2}
    ],
    option_discount:[
        {name:"有优惠","id":1},
        {name:"无优惠","id":2}
    ],
    option_month:[
        {name:"1月","id":"01"},
        {name:"2月","id":"02"},
        {name:"3月","id":"03"},
        {name:"4月","id":"04"},
        {name:"5月","id":"05"},
        {name:"6月","id":"06"},
        {name:"7月","id":"07"},
        {name:"8月","id":"08"},
        {name:"9月","id":"09"},
        {name:"10月","id":"10"},
        {name:"11月","id":"11"},
        {name:"12月","id":"12"}
    ],
    option_return:[
        {name:"有返佣","id":1}
    ],
    option_hierarchy:[
        {name: "股东", value: "1"},
        {name: "总代理", value: "2"},
        {name: "代理", value: "3"}
    ],
    option_extension:[
        {name: "账号", value: "1"},
        {name: "代理账号", value: "2"},
        {name: "推广id", value: "3"}
    ],
    option_account_type:[
        {name: "正式账号", value: "1"},
        {name: "试玩账号", value: "2"},
        {name: "带玩账号", value: "3"}
    ],
    option_account_id:[
        {name: "会员", value: "1"},
        {name: "管理员", value: "2"}
    ],
    option_login_id:[
        {name: "账号", value: "1"},
        {name: "IP", value: "2"}
    ],
    option_holder_order:[
        {name:"新增时间",value:"create_time"},
        {name:"股东数",value:"first_agency_count"},
        {name:"总代理数",value:"second_agency_count"},
        {name:"经销商数",value:"third_agency_count"},
        {name:"会员数",value:"member_count"},
        {name:"站点数",value:"site_index_id"}
        ],
    option_information_type:[
        {name:"公司入款",value:1},
        {name:"线上入款",value:2},
        {name:"联系我们",value:3},
        {name:"关于我们",value:4},
        {name:"代理联盟",value:5},
        {name:"存款帮助",value:6},
        {name:"常见问题",value:7},
        {name:"会员注册",value:8},
        {name:"代理注册",value:9},
        {name:"网站LOGO",value:10},
        {name:"首页轮播",value:11},
        {name:"左边浮动",value:12},
        {name:"右边浮动",value:13},
    ],
    option_onOff:[
        {name:"是",value:1},
        {name:"否",value:2},
    ],
    option_video_type:[
        {name:"lebo",value:1},
        {name:"bbin",value:2}
    ],
    option_register_type:[
        {name:"会员注册",id:1},
        {name:"代理注册",id:2},
        {name:"开户协议",id:3}
    ],
    option_name_type:[
        {name:"账号",id:1},
        {name:"名称",id:2}
    ],
    option_accounts_type:[
        {name:"所有账号",id:1},
        {name:"子账号",id:2},
        {name:"管理员",id:3}
    ],
    option_ip:[
        {name:"客户后台",value:1},
        {name:"代理后台",value:2},
        {name:"前台",value:3},
        {name:"wap",value:4}
    ],
    option_package:[
        {name:"套餐一",value:1},
        {name:"套餐二",value:2},
        {name:"套餐三",value:3},
    ],
    option_line:[
        {name:"线路一",value:1},
        {name:"线路二",value:2},
        {name:"线路三",value:3},
    ],
    option_log:[
        {name:"操作日志",value:1},
        {name:"登录日志",value:2}
    ]

};
