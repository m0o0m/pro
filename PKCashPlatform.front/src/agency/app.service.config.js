appConfig.apiUrls = {
    APP_HOST: "http://localhost:3000",
    //HOST: "http://api.nothingness.info:3002/mock/39/agency",
    HOST: "http://api.nothingness.info:3002/mock/46", //请求地址
    CAPTCHA: "/captcha", //验证码接口
    LOGIN: "/login", //登录接口
    LANG: { //语言下拉
        "index": "/langs",
        "cn": "/lang/cn",
        "us": "/lang/us",
        "zh": "/lang/zh",
    },
    ACTIVITY: {
        "index": "/activitys",
        "msgs": "/activity/msgs",
        "notify": "/activity/notify",
    },
    //用户中心
    FIRST_DROPF: "/first/drop",
    SECOND_DROPF: "/second/drop",
    THIRD_DROPF: "/third/dropf", //会员站点下来框
    THIRD_AGENT_DROPF: "/third/dropf", //会员站点下来框

    //账号管理接口
    MEMBER: "/member", //会员列表
    MEMBER_TYPE: "/member/type", //会员查询类型筛选
    MEMBER_STATUS: "/member/status",
    MEMBER_INFO: "/member/info", //会员状态
    MEMBER_DETAIL: "/member/detail",
    MEMBER_LEVEL: "/member/level", //修改.用户层级

    MEMBER_LEVEL_REGRESS:"/member/level/regress",    //回归层级(level/regress)
    MEMBER_LEVEL_DROP:"/member/level/drop" ,      //会员层级下拉框(drop)
    MEMBER_LEVEL_MOVE:"/member/level/move",//移动会员分层(move)
    MEMBER_LEVEL_SELFREBATE:"/member/level/selfrebate", //开启自助返水(selfrebate)
    MEMBER_LEVEL_PAYSET:"/member/level/payset",     //获取,修改层级支付设定(payset)
    MEMBER_LEVEL_LIST:"/member/level/list",      //站点会员层级列表(list)
    MEMBER_LEVEL_LOCK:"/member/level/lock",      //锁定会员层级(lock)
    MEMBER_LEVEL_MEBER:"/member/level/memberlist", //会员详情列表(memberlist)
    MEMBER_LEVEL_INFO:"/member/level/info" ,      //获取会员层级信息(info)
    MEMBER_REGISTER:"/member/register/setting",  //获取,配置站点会员注册信息(setting)


    //子帐号 Sub-account
    SUB_ACCOUNT_LIST : "/agent/sub/list",  //查看子账号列表(list)
    SUB_ACCOUNT : "/agent/sub",  //添加，修改 删除子账号
    SUB_ACCOUNT_INFO : "/agent/sub/info",  //获取子账号详情
    SUB_ACCOUNT_STATUS : "/agent/sub/status",  //查看子账号列表
    SUB_ACCOUNT_PERMISSION : "/agent/sub/permission",  //子账号权限列表,修改子账号权限
    SUB_ACCOUNT_TOKEN:  "/agent/subtoken" , //服务器端生成key
    SUB_ACCOUNT_TOKEN_INFO:  "/agent/subtoken/info",    //查看口令验证设置信息
    SUB_ACCOUNT_TOKEN_STATUS:  "/agent/subtoken/status",    //查看口令验证设置信息

    //银行卡管理
    DISPENSING_BANK:  "/member/bank",  //修改 删除 取会员出款银行集合(bank
    DISPENSING_BANK_INFO:  "/member/bank/info",  //获取会员银行详情

    //代理管理 代理 dealer
    DEALER_LIST : "/agent/third",  //代理列表
    DEALER_DISCOUNT : "/agent/discount",  //代理设置的会员优惠
    DEALER_ADD : "/third_agency",  //增加代理 修改详情 
    DEALER_BASEINFO : "/third_agency/info",  //查看代理基本资料
    DEALER_STATUS : "/third_agency/status",  //启用/禁用代理(status)
    DEALER_DOMAIN : "/agent/domain",  //查看 修改 添加 删除推广域名
    DEALER_INFO : "/agent/third/info",  //查看 修改代理详情
    DEALER_APPLICANT : "/agent/register",  //查看 删除 审核代理申请列表
    DEALER_APPLICANT_SET : "/agent/register/setting",  //查看 修改代理注册申请设定
    DEALER_APPLICANT_ONE : "/agent/register/one",  //查询一条代理申请(one)

    //总代
    DISTRIBUTOR_LIST : "/agent/second"  , //总代列表(agent/second)
    DISTRIBUTOR : "/second_agency", //增加,修改总代
    DISTRIBUTOR_INFO : "/second_agency/info", //查看总代详情(info)
    DISTRIBUTOR_STATUS : "/second_agency/status", //启用/禁用总代(status)
    DISTRIBUTOR_DISCOUNT : "/agent/second/discount", //会员优惠(discount)

    //股东
    SHAREHOLDER_LIST:"/agent/first",//股东列表(agent/first)
    SHAREHOLDER:"/first_agency",  //新增 修改股东
    SHAREHOLDER_DISCOUNT:"/agent/first/discount", //获取,修改优惠设定
    SHAREHOLDER_STATUS:"/first_agency/status", //启用／禁用股东(status)
    SHAREHOLDER_INFO:"/first_agency/info", //获取股东详情(info)

    SYSTEM_SEARCH : "/agent/search"  ,         //体系查询(agent/search)

    //系统设置--公告管理
    SYSTERM_NOTICE: "/systerm/notice/systermNotice", //系统公告
    SYSTERM_INFORMATION: "/systerm/notice/information", //公告信息

    //系统设置--会员消息
    MEMBER_NEWS:"/systerm/memberNews",                              //会员消息
    DELETE_NEWS:"/systerm/memberNews",                              //删除信息
    PREFERENCESNEWS:"/systerm/preferencesNews",
    MEMBER_SYSTEM:"/memberNews/type/drop",                          //会员体系下拉框
    SYSTEM_MEMBER_NEWS:"/systerm/memberNews",                       //发布新消息
    SYSTEM_PREFERENCES_NEWS:"/systerm/preferencesNews",             //注册优惠消息模板

    //系统设置--最新消息
    SYSTEM_MOTICE_LIST:"/system/notice",                            //公告管理列表
    ADD_SYSTEM_MOTICE:"/systerm/notice",                            //添加最新消息
    ADD_TYPE_DROP:"/add/type/drop",                                 //类型下拉框
    GET_NOTICE_NEWS:"/notice/news",                                 //获取消息信息
    POST_NOTICE_NEWS:"/notice/news",                                //提交消息信息
    DEL_SYSTEM_NOTICE:"/systerm/notice",                            //删除信息

    //系统设置--日志管理
    SYSTEM_MEMBER_LOGIN_LIST:"/system/memberLogin",                 //会员登录列表
    SYSTEM_ADMIN_LOGIN:"/system/adminLogin",                        //管理员列表
    SYSTEM_LOG_LIST:"/system/log",                                  //操作日志
    SYSTEM_AUDIT:"/system/audit",                                   //自动稽核


    //资金管理
    OUT_MONEY:"/out_money",
    REFUSE_MONEY:"/refuse/out_money",                                //拒绝出款
    CANCLE_COMPANY:"/cancle/company_income",                         //取消一条公司入款
    CONFIRM_MOENY:"/confirm/out_money"  ,                            //确认出款
    MONITOR_DEPOSIT:"/Monitor/Deposit",                              //公司入款-监控列表(公司入款)
    MONITOR_ONLINE:"/Monitor/online",                                //公司入款-监控列表(线上入款)
    MONITOR_MOENY: "/Monitor/Money",                                 //公司入款-监控列表(出款管理)
    MEMVER_BASEINFO:"/member/baseinfo",                              //人工存款-账号搜索
    MANUAL_ACCESS:"/manualAccess",                                   //人工存款-存入
    MANUAL_ACCESS_BATCH:"/manualAccess/batch",                       //人工存款-批量存款
    MEMBER_DROP:"/member_drop",                                      //获取层级
    MANUAL_WITHDRAWAL:"/manualWithdrawal",                           //人工取款-存入
    COMPANY_INCOME:"/companyIncome",                                 //公司入款-列表
    DEPOSIT_LIST:"/deposit/list",                                    //线上公司入款
    THIRDPAID_LIST:"/thirdPaidList",                                 //获取商户
    MANUAL_ACCESS_RECORD:"/manualAccess/record",                     //出入账目记录
    MANUAL_ACCESS_COLLECT:"/manualAccess/collect",                   //出入款账目汇总
    TYPE_LIST:"/type/list",                                          //获取转入转出项目
    MEMBER_BALANCE:"/member/balance",                                //获取转出项目金额
    QUOTA_SUMBIT:"/Quota/Submit",                                    //额度转换提交
    BALANCE_CONVERSION:"/balanceConversion",                         //额度转换记录列表
    AUDIT_AUDITLOG:"/audit/auditlog",                                //稽核日志查询
    ACTUAL_PURCHASE:"/actual/purchase",                              //获取实际购买
    ACTUAL_MEMBER_AUDITNOW:"/audit/memberauditnow",                  //稽核日志
    BANK_OUTBANK:"/bank/outbank",                                    //出款银行剔除
    OUTBANK_STATUS:"/bank/outbank/status",                           //更改出款银行状态
    PAYMENT_LIST:"/payment/list",                                    //入款银行设定列表
    ADD_PAYMENT:"/add/payment",                                      //添加入款银行设定
    PAYMENT:"/payment",                                              //获取单个入款银行设定详情
    PAYMENT_PUT:"/payment/put",                                      //入款银行设定(修改)
    PAYMENT_DEPOSIT:"/payment/deposit",                              //存款记录
    PAYMENT_STATUS:"/payment/status",                                //入款银行设定(状态)
    PAYMENT_DELETE:"/payment/delete",                                //入款银行设定(删除)
    BANK_INCOME:"/bank/income",                                      //入款银行剔除
    BANK_INCOME_STATUS:"/bank/income/status",                        //入款银行剔除(状态)
    THIRD_PAID_LIST: "/PaidList",                                    //第三方下拉
    BANK_THIRD:"/bank/third",                                        //第三方银行剔除列表
    THIRD_STATUS: "/third/status",                                   //第三方银行剔除(状态)
    ONLINE_SETUP:"/onlineSetup",                                     //线上支付设定(列表)
    NEWONLINE_SETUP:"/newOnlineSetup",                               //线上支付设定(新增)
    PAID_TYPE:"/paidType",                                          //获取类型
    ONLINE_SETUP_SINGLE :"/onlineSetup/single",                      //线上支付设定详情
    NEWONLINE_STUP_MODIFY:"/newOnlineSetup/modify",                  //线上支付设定(修改)
    STOP_ONLINESETUP:"/stop/onlineSetup",                            //线上支付设定(修改状态)
    ONLINE_SETUP_DEL: "/onlineSetup/del",                           ////线上支付设定(删除)
    PAYSET_LIST:"/payset/list",                                      //支付参数设定(公司自定设置)
    CURRENCY:"/currency",                                            //获取币别
    PAYSET_ADD:"/payset/add",                                       //支付参数设定(添加公司设定)
    PAYSET_PUBLIC:"/payset/public",                                 //支付参数设定(币别设定)
    PAYSET_DETAIL:"/payset/detail",                                 //获取单个公司设定详情
    PAYSET_MODIFY:"/payset/modify",                                 //修改单个公司自定设定
    PAYSET_DELETE: "/payset/delete",                                //支付参数设定(删除公司设定)
    DENOMINATION:"/denomination",                                   //币别删除
    PAYSET_ONE:"/payset/public/one",                                //获取详情公司设定
    PAYSETES:"/payset",                                             //修改后提交
    PAYSET_PUBLIC_ONE:"/payset/public/ones",                         //查看币别
    STATISTICS:"/statistics",                                       //优惠统计
    DEPOSIT_DISCOUNY:"/Deposit/Discount",                            //存入
    RETREAT_WATER_SELF_SEARCH:"/retreat/water/self/search",         //自助返水列表
    RETREAT_DETAIL:"/retreat/water/self/detai",                     //自助返水明细
    RETREAT_WATER_EDIT:"/retreat/water/edit",                       //优惠查询(明细-冲销)
    COMBO_PLATFORM:"/combo_platform",                               //获取商品
    RETREAT_WATER_SET_ADD:"/retreat/water/set/add",                 //添加设定
    RETREAT_WATER_SET_DETAIL:"/retreat/water/set/detail",           //获取返点优惠查询详情
    RETREAT_WATER_SET_EDIT:"/retreat/water/set/edit",               //修改后提交
    OVERRIDE_ADDONE:"/override/addone",                             //代理退佣设定(新增)
    OVERRIDE_DETAIL:"/override/detail",                             //代理退佣设定(详情)
    OVERRIDE_MODIFY:"/override/modify",                             //代理退佣设定(修改)
    //资金管理--会员分析接口
    USER_INCOME_SEARCH: "/user/income_search",                      //出入款分析
    USER_BUY_SEARCH: "/user/buy_search",                            //购买分析
    USER_BUY_TYPE_SELECT: "/user/buy/type_select",                  //购买分析--类型商品下拉框
    USER_BUY_AGENCY_SELECT: "/user/buy/agency_select",              //购买分析--代理账号
    USER_RETURN_SEARCH: "/user/return_search",                      //退水分析
    USER_SEARCH: "/user/search",                                    //有效会员列表
    USER_MONEY_SEARCH: "/user/money_search",                        //现金系统
    USER_MONEY_DELETE: "/user/money_delete",                        //现金系统--删除
    USER_BALANCE_SEARCH: "/user/balance_search",                    //会员查询
    //报表管理
    REPORTQUERY:"/Reportquery",                                     //报表查询
    SHAREHOLDERS_STATEMENT:"/Shareholders/statement",               //获取股东报表
    GENERAL_GENERATION_REPORT:"/General/generation/Report",         //获取总代
    PROXY_REPORT:"/proxy/report",                                   //获取代理
    MEMBERSHIP_REPORT:"/membership/report",                         //获取会员


    //资金管理--会员返佣接口
    REBATE_LIST:"/rebate/list",                                      //返佣查询
    REBATE_DETAILS:"/rebate/details",                                //返佣查询--明细
    REBATE_WRITEOFF:"/rebate/writeoff",                              //返佣查询--明细--冲销
    REBATE_SET_GET_ALL:"/rebateSet/getAll",                          //返佣优惠设定
    REBATE_SET_ADD_ALL:"/rebateSet/addAll",                          //返佣优惠设定-新增详情
    REBATE_SET_SUBMIT:"/rebateSet/submit",                           //返佣优惠设定-新增/修改-提交
    REBATE_SET_GET_ONE:"/rebateSet/getOne",                          //返佣优惠设定-修改详情
    REBATE_SET_DEL:"/rebateSet/del",                                 //返佣优惠设定--删除
    USER_REBATE_SEARCH:"/userRebate/search",                         //会员返佣-搜索
    USER_REBATE_DEPOSIT:"/userRebate/deposit",                       //会员返佣-搜索页面-存入
    SPREAD_INFO:"/spread/info",                                      //会员推广查询
    SPREAD_NUM_INFO:"/spread/numInfo",                               //会员推广查询--推荐会员数
    SPREAD_LIST:"/spread/list",                                      //会员推广设定
    SPREAD_EDIT:"/spread/edit",                                      //会员推广设定--修改
    SPREAD_ADD:"/spread/add",                                        //会员推广设定--添加

    //资金管理--额度统计
    QUOTA_LIST:"/quota/list",                                       //额度统计
    QUOTA_RECHARGE:"/quota/recharge",                               //充值记录
    SITE_SINGLE_RECORD:"/siteSingleRecord",                         //掉单列表
    SUB_SINGLE_RECORD:"/subSingleRecord",                           //掉单申请--提交
    RECORD_ORDERNUM:"/record/orderNum",                             //额度充值--第三方--订单号
    BANK_ORDERNUM:"/bank/ordernum",                                 //额度充值--银行卡--订单号
    THREE_BANK:"/three/bank",                                       //额度充值--第三方--支付银行
    RECORD_CARD_BANK:"/record/card/bank",                           //额度充值--银行卡--收款银行
    THREE_SUB:"/three/sub",                                         //额度充值--第三方--提交
    BANK_SUB:"/bank/sub",                                           //额度充值--银行卡--提交
    QUOTA_RECORD:"/quota/record",                                   //额度记录


    //网站资讯管理--站点资料编辑
    COLOR_DROP: "/color/drop",                                       //会员中心颜色下拉
    SITE_WEBSITE: "/site/website",                                   //网站信息
    VIDEO_TYPE_DROP: "/video/type/drop",                             //类型下拉
    VIDEO_STYLE_DROP: "/video/style/drop",                           //风格下拉
    SITE_VIDEO: "/site/video",                                       //视讯管理
    SITE_VIDEO_USE: "/site/video/use",                               //视讯管理--使用
    SITE_VIDEO_BACK: "/site/video/back",                             //视讯管理--还原老版本
    SITE_SPORTS: "/site/sports",                                     //体育管理
    SITE_ELECTRONICS_THEME: "/site/electronics/theme",                         //电子管理--获取主题配置信息
    SITE_ELECTRONICS: "/site/electronics",                                     //电子管理
    SITE_ELECTRONICS_INITIALIZATION: "/site/electronics/initialization",       //电子管理--初始化
    SITE_LOTTERY: "/site/lottery",                                          //彩票管理--初始化
    SITE_LOTTERY_ENCLOSURE: "/site/lottery/enclosure",                      //彩票管理--获取附件信息
    SITE_LOTTERY_ENCLOSURE_SELECT: "/site/lottery/enclosure/select",        //彩票管理--选择附件信息
    SITE_LOTTERY_HALL_SOURCE_DROP: "/site/lottery/hall/source/drop",        //彩票大厅--来源下拉
    SITE_LOTTERY_HALL: "/site/lottery/hall",                                //彩票大厅
    SITE_MODULE: "/site/module",                                            //模块管理
    SITE_APPLICATION: "/site/application",                                  //模块管理
    SITE_CACHE_PAGE_DROP: "/site/cache/page/drop",                          //缓存管理--界面下拉
    SITE_CACHE: "/site/cache",                                              //缓存管理

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
    GRAPHIC_NOTICE_CONTENT: "/graphic/notice/content",                      //公告弹框--广告编辑
    GRAPHIC_ADVERTISEMENT: "/graphic/advertisement",                        //广告管理
    GRAPHIC_ADVERTISEMENT_SORT: "/graphic/advertisement/sort",              //广告管理--排序
    GRAPHIC_ADVERTISEMENT_STATUS: "/graphic/advertisement/status",          //广告管理--启用禁用
    GRAPHIC_ADVERTISEMENT_DETAIL: "/graphic/adver/detail",          //广告管理--详情

    //网站资讯管理--案件编辑
    PENDING_CASE:"/caseEditor/pendingCase",                                 //待审案件
    PENDING_CASE_SEND:"/caseEditor/pendingCase/sendAudit",                  //待审案件--发送审核
    PENDING_CASE_DEL:"/caseEditor/pendingCase/delete",                      //待审案件--删除
    AUDIT_CASE:"/caseEditor/auditCase",                                     //待审中案件
    AUDIT_CASE_DEL:"/caseEditor/auditCase/delete",                          //待审中案件--删除
    THROUGH_CASE:"/caseEditor/throughCase",                                 //通过案件
    REVOKE_CASE:"/caseEditor/revokeCase",                                   //撤销案件

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

    //资金管理--退佣统计
    REBATE_SEARCH:"/rebate/search",                                     //退佣统计列表
    OVERRIDE_LIST:"/override/list",                                     //代理退佣设定列表
    OVERRIDE_DELETE:"/override/delete",                                 //代理退佣设定删除
    PERIODS:"/periods",                                                 //期数管理
    PERIODS_DEL:"/periods/del",                                         //期数管理--删除
    PERIODS_COMMISSION:"/periods/commission",                           //期数管理--退佣冲销
    PERIODS_ADD:"/periods/add",                                         //期数管理--新增
    PERIODS_MODIFY:"/periods/modify",                                   //期数管理--修改
    POUNDAGE_LISTSET:"/poundage/listset",                               //手续费设定
    POUNDAGE_DEL:"/poundage/del",                                       //手续费设定--删除
    POUNDAGE_GETLIST:"/poundage/getlist",                               //手续费设定--详情
    POUNDAGE_UPDATE:"/poundage/update",                                 //手续费设定--修改
    RETREAT_WATER_SET_LIST:"/retreat/water/set/list",                   //返点优惠设定
    RETURN_LEVEL:"/return/level",                                       //返点层级下拉
    RETREAT_WATER_SET_DEL:"/retreat/water/set/del",                     //返点优惠设定--删除
    PREFERENTIAL_INQUIRIES:"/preferential/inquiries",                   //优惠查询
    RETREAT_WATER_DETAIL:"/retreat/water/detail",                       //优惠查询--明细

};


appConfig.option = {
    option_state: [{
        name: "在线",
        value: "1"
    },
        {
            name: "全部",
            value: "-1"
        }
    ],
    option_status: [{
        name: "禁用",
        value: "2"
    },
        {
            name: "启用",
            value: "1"
        }
    ],
    option_online: [{
        name: "在线",
        value: "1"
    },
        {
            name: "离线",
            value: "2"
        }
    ],
    option_hidden: [{
        name: "全部",
        value: "-1"
    },
        {
            name: "显示",
            value: "0"
        },
        {
            name: "隐藏",
            value: "1"
        }
    ],
    option_sort: [{
        name: "新增日期",
        value: "create_time"
    },
        {
            name: "代理名称",
            value: "username"
        },
        {
            name: "代理账号",
            value: "account"
        }
    ],
    option_sort_1: [{
        name: "新增日期",
        value: "create_time"
    },
        {
            name: "股东名称",
            value: "username"
        },
        {
            name: "股东账号",
            value: "account"
        }
    ],
    option_sort_2: [{
        name: "新增日期",
        value: "create_time"
    },
        {
            name: "总代理名称",
            value: "username"
        },
        {
            name: "总代理账号",
            value: "account"
        }
    ],
    option_sortBig: [{
        name: "正序",
        value: "true"
    },
        {
            name: "倒序",
            value: "false"
        }
    ],
    option_types: [{
        name: "账号",
        value: 1
    },
        {
            name: "注册ip",
            value: 3
        },
        {
            name: "登录ip",
            value: 4
        },
        {
            name: "手机",
            value: 5
        },
        {
            name: "银行卡",
            value: 6
        },
        {
            name: "邮箱",
            value: 7
        },
        {
            name: "qq",
            value: 8
        },
        {
            name: "微信",
            value: 9
        }
    ],
    option_income_type: [{
        name: "公司入款",
        value: 1
    },
        {
            name: "线上入款",
            value: 2
        },
        {
            name: "人工入款",
            value: 3
        },
        {
            name: "人工出款",
            value: 4
        },
        {
            name: "线上出款",
            value: 5
        }
    ],
    option_is_type: [{
        name: "会员账号",
        value: 1
    },
        {
            name: "代理账号",
            value: 2
        }
    ],
    option_is_order: [{
        name: "总金额",
        value: 1
    },
        {
            name: "总笔数",
            value: 2
        }
    ],
    option_type_m: [{
        name: "admin"
    },
        {
            name: "agency"
        }
    ],
    option_reg: [{
        name: "pc",
        value: "1"
    },
        {
            name: "android",
            value: "3"
        },
        {
            name: "ios",
            value: "4"
        },
        {
            name: "wap",
            value: "2"
        }
    ],
    option_type: [{
        name: "微信",
        value: "wechat"
    },
        {
            name: "电话号",
            value: "skype"
        },
        {
            name: "邮箱",
            value: "email"
        },
        {
            name: "qq",
            value: "qq"
        }
    ],
    option_type1: [{
        name: "平台管理",
        "id": 2
    },
        {
            name: "代理平台",
            "id": 1
        }
    ],
    option_states: [{
        "status": "未处理",
        "id": "2"
    },
        {
            "status": "已通过",
            "id": "1"
        }
    ],
    option_order: [{
        name: "倒序",
        value: "true"
    },
        {
            name: "正序",
            value: "false"
        }
    ],
    option_level: [{
        name: "股东",
        value: "1"
    },
        {
            name: "总代理",
            value: "2"
        },
        {
            name: "代理",
            value: "3"
        },
        {
            name: "会员",
            value: "4"
        },
        {
            name: "全部",
            value: ""
        }
    ],
    option_vague: [{
        name: "是",
        value: "1"
    },
        {
            name: "否",
            value: "0"
        }
    ],
    option_commodity: [{
        name: "商品名称",
        value: "product_name"
    },
        {
            name: "商品id",
            value: "product_id"
        }
    ],
    option_method: [{
        name: "POST"
    },
        {
            name: "GET"
        },
        {
            name: "PUT"
        },
        {
            name: "HEAD"
        },
        {
            name: "DELETE"
        },
        {
            name: "OPTIONS"
        },
        {
            name: "TRACE"
        }
    ],
    option_system: [{
        name: "未反水",
        value: 1
    },
        {
            name: "返水",
            value: 2
        }
    ],
    option_detail: [{
        name: "已冲销",
        id: 1
    },
        {
            name: "未冲销",
            id: 2
        }
    ],
    option_site: [{
        name: "站点2",
        id: 2
    },
        {
            name: "站点3",
            id: 3
        },
        {
            name: "站点4",
            id: 4
        },
        {
            name: "站点5",
            id: 5
        },
        {
            name: "站点1",
            id: 1
        }
    ],
    option_deposit: [{
        name: "人工存入",
        id: 1
    },
        {
            name: "存款优惠",
            id: 2
        },
        {
            name: "负数额度归零",
            id: 3
        },
        {
            name: "取消出款",
            id: 4
        },
        {
            name: "返点优惠",
            id: 5
        },
        {
            name: "活动优惠",
            id: 6
        },
        {
            name: "其他",
            id: 7
        }
    ],
    option_deposit_type: [{
        name: "重复出款",
        id: 8
    },
        {
            name: "公司入款误存",
            id: 9
        },
        {
            name: "公司负数回冲",
            id: 10
        },
        {
            name: "手动申请出款",
            id: 11
        },
        {
            name: "扣除非法下注派彩",
            id: 12
        },
        {
            name: "放弃存款优惠",
            id: 13
        },
        {
            name: "其他",
            id: 14
        }
    ],
    option_way: [{
        name: "额度转换",
        id: 1
    },
        {
            name: "体育下注",
            id: 2
        },
        {
            name: "彩票下注",
            id: 3
        },
        {
            name: "彩票派彩",
            id: 4
        },
        {
            name: "EG电子下注",
            id: 5
        },
        {
            name: "EG彩票下注",
            id: 6
        },
        {
            name: "EG彩票派彩",
            id: 7
        }
    ],
    shebei: [{
        name: "pc",
        id: 1
    },
        {
            name: "wap",
            id: 2
        },
        {
            name: "全部",
            id: 3
        }
    ],

    money_status: [{
        name: "全部",
        id: 0
    },
        {
            name: "未出款",
            id: 5
        },
        {
            name: "已出款",
            id: 1
        },
        {
            name: "已拒绝",
            id: 4
        },
        {
            name: "已取消",
            id: 3
        },
        {
            name: "预备出款",
            id: 2
        }
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
    option_time : [
        {name: "转时区", id: 1},
        {name: "非转时区", id: 2}
    ],
    option_money_type: [
        {name:"額度轉換",id:2},
        {name:"体育下注",id:21},
        {name:"体育派彩",id:22},
        {name:"彩票下注",id:23},
        {name:"冲销明细",id:24},
        {name:"彩票派彩",id:25},
        {name:"注单无效",id:26},
        {name:"注单取消",id:27},
        {name:"线上入款",id:28},
        {name:"公司入款",id:29},
        {name:"线上入款不含优惠",id:30},
        {name:"公司入款不含优惠",id:31},
        {name:"线上取款",id:32},
        {name:"公司/线上入款",id:33},
        {name:"优惠退水",id:34},
        {name:"自助返水",id:35},
        {name:"优惠活动",id:36},
        {name:"人工存入",id:37},
        {name:"人工取出",id:38},
        {name:"人工存款與取款",id:39},
        {name:"入款明细",id:40},
        {name:"出款明细",id:41},
        {name:"会员退佣",id:42},
        {name:"会员退佣冲销",id:43},
        {name:"额度掉单(取出)",id:44},
        {name:"额度掉单(存入)",id:45},
        {name:"会员推广返佣",id:46}
    ],
    option_refresh:[
        {name:"30秒",id:30},
        {name:"60秒",id:60},
        {name:"120秒",id:120},
        {name:"180秒",id:180}
    ],
    option_money_time:[
        {name:"北京",id:1},
        {name:"美东",id:2}
    ],
    option_deposit_item:[
        {name:"人工存入",id:1},
        {name:"存款优惠",id:2},
        {name:"负数额度归零",id:3},
        {name:"取消出款",id:4},
        {name:"返点优惠",id:5},
        {name:"活动优惠",id:6}
    ],
    option_handle:[
        {name:"未处理", id: 3},
        {name:"已确认", id: 1},
        {name:"已取消",id:2}
    ],
    option_query_criteria:[
        {name:"账号",id:1},
        {name:"订单号",id:2}
    ],
    option_discount:[
        {name:"有优惠",id:1},
        {name:"无优惠",id:2}
    ],
    option_month:[
        {name:"1月",id:"01"},
        {name:"2月",id:"02"},
        {name:"3月",id:"03"},
        {name:"4月",id:"04"},
        {name:"5月",id:"05"},
        {name:"6月",id:"06"},
        {name:"7月",id:"07"},
        {name:"8月",id:"08"},
        {name:"9月",id:"09"},
        {name:"10月",id:"10"},
        {name:"11月",id:"11"},
        {name:"12月",id:"12"}
    ],
    option_return:[
        {name:"有返佣",id:1}

    ],

    option_outbank:[
        {"name":"支付宝","id":2},
        {"name":"中国银行","id":1}
    ],
    option_droplist:[
        {name:"审核中",id:1},
        {name:"已通过",id:2},
        {name:"无效注单",id:3}
    ],
    option_record:[
        {name:"存入",id:1},
        {name:"取出",id:2},
        {name:"预借",id:3}
    ],
    option_record_type:[
        {name:"额度转换",id:1},
        {name:"额度加款",id:2},
        {name:"额度扣款",id:3},
        {name:"在线充值",id:4}
    ],
    option_time_zone:[
        {name:"非转时区",id:1},
        {name:"转时区",id:2}
    ],
    option_Preferential_member:[
        {name:"全部",id:1},
        {name:"会员",id:2}
    ],
    option_discount:[
        {name:"优惠内容",id:1},
        {name:"优惠分类",id:2}
    ]
};
