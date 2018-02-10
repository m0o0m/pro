angular.module("app.customer",['ui.router', 'datatables', 'datatables.bootstrap']).config(function ($stateProvider) {
    
    $stateProvider
        .state('app.customer',{
            abstract:true,
            data:{
                //123 title:"客户管理"
                title:"客户后台管理"
            }
        })
        .state('app.customer.customer',{
            url: '/customer/customer',
            data: {
                //123 title:"客户管理"
                title: '客户管理'
            },
            views:{
                'content@app':{
                    controller:'custCtrl as datatables',
                    templateUrl: 'views/customer/views/customer.html'
                }
            }
        })
        .state('app.customer.hierarchicalManag',{
            url: '/customer/hierarchicalManag',
            data: {

                title: '层级管理'
            },
            views:{
                'content@app':{
                    controller:'hierarchicalManagCtrl',
                    templateUrl: 'views/customer/views/hierarchicalManag.html'
                }
            }
        })
        .state('app.customer.AddHierarchy',{
            url: '/customer/AddHierarchy',
            data: {

                title: '添加层级'
            },
            views:{
                'content@app':{
                    controller:'AddHierarchyCtrl',
                    templateUrl: 'views/customer/views/AddHierarchy.html'
                }
            }
        })
        .state('app.customer.MembershipDetails',{
            url: '/customer/MembershipDetails',
            data: {

                title: '会员详情'
            },
            views:{
                'content@app':{
                    controller:'MembershipDetailsCtrl',
                    templateUrl: 'views/customer/views/MembershipDetails.html'
                }
            }
        })
        .state('app.customer.Entryandexit',{
            url: '/customer/Entryandexit',
            data: {

                title: '出入款管理'
            },
            views:{
                'content@app':{
                    controller:'EntryandexitCtrl',
                    templateUrl: 'views/customer/views/Entryandexit.html'
                }
            }
        })
        .state('app.customer.agent',{
            url: '/customer/agent',
            data: {

                title: '代理管理'
            },
            views:{
                'content@app':{
                    controller:'agentCtrl',
                    templateUrl: 'views/customer/views/agent.html'
                }
            }
        })
        .state('app.customer.Commonbullet',{
            url: '/customer/Commonbullet',
            data: {

                title: '公共弹框管理'
            },
            views:{
                'content@app':{
                    controller:'CommonbulletCtrl',
                    templateUrl: 'views/customer/views/Commonbullet.html'
                }
            }
        })
        .state('app.customer.configurationsetting', {
            url: '/customer/configurationsetting',
            data: {

                title: '配置设置'
            },
            views: {
                'content@app': {
                    controller: 'configurationsettingCtrl',
                    templateUrl: 'views/customer/views/configurationsetting.html'

                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ]);
                }
            }
        })
        .state('app.customer.NoticeKey',{
            url: '/customer/NoticeKey',
            data: {
                //123 title:"客户管理"
                title: '公告密钥管理'
            },
            views:{
                'content@app':{
                    controller:'NoticeKeyCtrl',
                    templateUrl: 'views/customer/views/NoticeKey.html'
                }
            }
        })
        .state('app.customer.Log',{
            url: '/customer/Log',
            data: {
                //123 title:"客户管理"
                title: '稽核日志'
            },
            views:{
                'content@app':{
                    controller:'LogCtrl',
                    templateUrl: 'views/customer/views/Log.html'
                }
            }
        })
        .state('app.customer.LogManage',{
            url: '/customer/LogManage',
            data: {
                //123 title:"客户管理"
                title: '登录日志'
            },
            views:{
                'content@app':{
                    controller:'LogManageCtrl',
                    templateUrl: 'views/customer/views/LogManage.html'
                }
            }
        })
        .state('app.customer.operationLog',{
            url: '/customer/operationLog',
            data: {
                title: '操作日志'
            },
            views:{
                'content@app':{
                    controller:'operationLogCtrl',
                    templateUrl: 'views/customer/views/operationLog.html'
                }
            }
        })
        .state('app.customer.automaticAudit',{
            url: '/customer/automaticAudit',
            data: {
                title: '自动稽核'
            },
            views:{
                'content@app':{
                    controller:'automaticAuditCtrl',
                    templateUrl: 'views/customer/views/automaticAudit.html'
                }
            }
        })
        .state('app.customer.VideoAccount',{
            url: '/customer/VideoAccount',
            data: {
                //123 title:"客户管理"
                title: '视讯账号管理'
            },
            views:{
                'content@app':{
                    controller:'VideoAccountCtrl',
                    templateUrl: 'views/customer/views/VideoAccount.html'

                }
            }
        })
        .state('app.customer.exceptionMember',{
            url: '/customer/exceptionMember',
            data: {
                title: '异常会员查询'
            },
            views:{
                'content@app':{
                    controller:'exceptionMemberCtrl',
                    templateUrl: 'views/customer/views/exceptionMember.html'

                }
            }
        })
        .state('app.customer.userManagement',{
            url: '/customer/userManagement',
            data: {
                title: '用户管理'
            },
            views:{
                'content@app':{
                    controller:'userManagementCtrl',
                    templateUrl: 'views/customer/views/userManagement.html'

                }
            }
        })
        .state('app.customer.preferentialQuery',{
            url: '/customer/preferentialQuery',
            data: {
                title: '优惠查询'
            },
            views:{
                'content@app':{
                    controller:'preferentialQueryCtrl',
                    templateUrl: 'views/customer/views/preferentialQuery.html'

                }
            }
        })
        .state('app.customer.preferentialDetail',{
            url: '/customer/preferentialDetail',
            data: {
                title: '优惠查询'
            },
            views:{
                'content@app':{
                    controller:'preferentialDetailCtrl',
                    templateUrl: 'views/customer/views/preferentialDetail.html'

                }
            }
        })
        .state('app.customer.childAccount',{
            url: '/customer/childAccount',
            data: {
                title: '子账号管理'
            },
            views:{
                'content@app':{
                    controller:'childAccountCtrl',
                    templateUrl: 'views/customer/views/childAccount.html'

                }
            }
        })
        .state('app.customer.applicationInquiry',{
            url: '/customer/applicationInquiry',
            data: {
                title: '自助优惠申请查询'
            },
            views:{
                'content@app':{
                    controller:'applicationInquiryCtrl',
                    templateUrl: 'views/customer/views/applicationInquiry.html'

                }
            }
        })
        .state('app.customer.applicationSwitch',{
            url: '/customer/applicationSwitch',
            data: {
                title: '自助优惠开关'
            },
            views:{
                'content@app':{
                    controller:'applicationSwitchCtrl',
                    templateUrl: 'views/customer/views/applicationSwitch.html'
                }
            }
        })

        .state('app.customer.Advertisingcontent',{
            url: '/customer/Advertisingcontent',
            data: {
                //123 title:"客户管理"
                title: '广告内容'
            },
            views:{
                'content@app':{
                    controller:'AdvertisingcontentCtrl',
                    templateUrl: 'views/customer/views/Advertisingcontent.html'

                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ])

                }
            }
        })
        .state('app.customer.ModifyNotice',{
            url: '/customer/ModifyNotice?id',
            data: {
                //123 title:"客户管理"
                title: '广告内容'
            },
            views:{
                'content@app':{
                    controller:'ModifyNoticeAdCtrl',
                    templateUrl: 'views/customer/views/ModifyNotice.html'

                }
            },
            resolve: {
                srcipts: function(lazyScript){
                    return lazyScript.register([
                        'vendor.ui.js'
                    ])

                }
            }
        })
});