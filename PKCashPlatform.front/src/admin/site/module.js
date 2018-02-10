angular.module('app.site', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.site')
    .config(function ($stateProvider) {
    $stateProvider
        .state('app.site', {
            abstract: true,
            data: {
                // title : '站点'
                title : 'site'
            }
        })
        .state('app.site.siteManager', {
            url: '/site/siteManager',
            data: {
                title : '站点管理'
            },
            views: {
                "content@app": {
                    controller : 'siteManagerCtrl',
                    templateUrl: "views/site/views/siteManager.html"
                }
            }
        })
        .state('app.site.siteCommand', {
            url: '/site/siteCommand',
            data: {
                title : '站点口令设置'
            },
            views: {
                "content@app": {
                    controller : 'siteCommandCtrl',
                    templateUrl: "views/site/views/siteCommand.html"
                }
            }
        })
        .state('app.site.informationAudit', {
            url: '/site/informationAudit',
            data: {
                title : '资讯审核'
            },
            views: {
                "content@app": {
                    controller : 'informationAuditCtrl',
                    templateUrl: "views/site/views/informationAudit.html"
                }
            }
        })
        .state('app.site.informationHistory', {
            url: '/site/informationHistory',
            data: {
                title : '历史记录'
            },
            views: {
                "content@app": {
                    controller : 'informationHistoryCtrl',
                    templateUrl: "views/site/views/informationHistory.html"
                }
            }
        })
        .state('app.site.siteList', {
            url: '/site/siteList',
            data: {
                // title : '站点列表'
                title : 'Site List'
            },
            views: {
                "content@app": {
                    controller : 'SiteCtrl as site',
                    templateUrl: "views/site/views/sitelist.html"
                }
            }
        })
        .state('app.site.addSite', {
            url:'/site/addSite',
            data:{
                //123 title:"添加站点"
                title:"添加站点"
            },
            views:{
                "content@app": {
                    controller: 'AddsiteCtrl as addsite',
                    templateUrl: "views/site/views/addSite.html"
                }
            }
        })
        .state('app.site.addSite1', {
            url:'/site/addSite1',
            data:{
                //123 title:"添加站点"
                title:"添加站点"
            },
            views:{
                "content@app": {
                    controller: 'AddSite1Ctrl as addsite1',
                    templateUrl: "views/site/views/addSite1.html"
                }
            }
        })
        .state('app.site.seeSite', {
            url:'/site/seeSite',
            data:{
                //123 title:"查看站点"
                title:"查看站点"
            },
            views:{
                "content@app": {
                    controller: 'SeeSiteCtrl as seeSite',
                    templateUrl: "views/site/views/seeSite.html"
                }
            }
        })
        .state('app.site.modifySite', {
            url:'/site/modifySite',
            data:{
                // title:"修改站点信息"
                title:"Modify site information"
            },
            params:{id:null},
            views:{
                "content@app": {
                    controller: 'modifySiteCtrl as modifySite',
                    templateUrl: "views/site/views/modifySite.html"
                }
            }
        })
        .state('app.site.maintenance', {
            url:'/site/maintenance',
            data:{
                //123 title:"查看站点"
                title:"维护管理"
            },
            views:{
                "content@app": {
                    controller: 'maintenanceCtrl',
                    templateUrl: "views/site/views/maintenance.html"
                }
            }
        })
        .state('app.site.maintenanceItem', {
            url:'/site/maintenanceItem',
            data:{
                //123 title:"查看站点"
                title:"维护管理"
            },
            views:{
                "content@app": {
                    controller: 'maintenanceItemCtrl',
                    templateUrl: "views/site/views/maintenanceItem.html"
                }
            }
        })

        .state('app.site.multistation', {
            url:'/site/multistation',
            data:{
                //123 title:"查看站点"
                title:"multistation"
            },
            views:{
                "content@app": {
                    controller: 'multistationCtrl',
                    templateUrl: "views/site/views/multistation.html"
                }
            }
        })
        .state('app.site.url', {
            url:'/site/url',
            data:{
                //123 title:"查看站点"
                title:"下载链接地址"
            },
            views:{
                "content@app": {
                    controller: 'urlCtrl',
                    templateUrl: "views/site/views/url.html"
                }
            }
        })
        .state('app.site.IPSwitch', {
            url:'/site/IPSwitch',
            data:{
                // title:"修改站点信息"
                title:"IPSwitch"
            },
            params:{id:null},
            views:{
                "content@app": {
                    controller: 'IPSwitchCtrl',
                    templateUrl: "views/site/views/IPSwitch.html"
                }
            }
        })
        .state('app.site.Proxynavigation', {
            url:'/site/Proxynavigation',
            data:{
                // title:"修改站点信息"
                title:"Proxynavigation"
            },
            params:{id:null},
            views:{
                "content@app": {
                    controller: 'ProxynavigationCtrl',
                    templateUrl: "views/site/views/Proxynavigation.html"
                }
            }
        })
        .state('app.site.JSversion', {
            url:'/site/JSversion',
            data:{
                // title:"修改站点信息"
                title:"JSversion"
            },
            params:{id:null},
            views:{
                "content@app": {
                    controller: 'JSversionCtrl',
                    templateUrl: "views/site/views/JSversion.html"
                }
            }
        })
        .state('app.site.Whitelist', {
            url:'/site/Whitelist',
            data:{
                // title:"ip白名单"
                title:"Whitelist"
            },
            views:{
                "content@app": {
                    controller: 'WhitelistCtrl',
                    templateUrl: "views/site/views/Whitelist.html"
                }
            }
        })
        .state('app.site.Announcement', {
            url:'/site/Announcement',
            data:{
                // title:"公告管理"
                title:"Announcement"
            },
            views:{
                "content@app": {
                    controller: 'announceCtrl',
                    templateUrl: "views/site/views/Announcement.html"
                }
            }
        })
        .state('app.site.videoModule', {
            url:'/site/videoModule/:id',
            data:{
                // title:"视讯模块管理"
                title:"videoModule"
            },
            views:{
                "content@app": {
                    controller: 'videoModuleCtrl',
                    templateUrl: "views/site/views/videoModule.html"
                }
            }
        })
        .state('app.site.negative', {
            url:'/site/negative/:id',
            data:{
                // title:"视讯模块管理"
                title:"negative"
            },
            views:{
                "content@app": {
                    controller: 'negativeCtrl',
                    templateUrl: "views/site/views/negative.html"
                }
            }
        })
        .state('app.site.reception', {
            url:'/site/reception/:id',
            data:{
                // title:"视讯模块管理"
                title:"reception"
            },
            views:{
                "content@app": {
                    controller: 'receptionCtrl',
                    templateUrl: "views/site/views/reception.html"
                }
            }
        })
        .state('app.site.data', {
            url:'/site/data/:id',
            data:{
                // title:"视讯模块管理"
                title:"data"
            },
            views:{
                "content@app": {
                    controller: 'dataCtrl',
                    templateUrl: "views/site/views/data.html"
                }
            }
        })
        .state('app.site.dataLevel', {
            url:'/site/dataLevel/:id',
            data:{
                // title:"视讯模块管理"
                title:"dataLevel"
            },
            views:{
                "content@app": {
                    controller: 'dataLevelCtrl',
                    templateUrl: "views/site/views/dataLevel.html"
                }
            }
        })
        .state('app.site.ProxyData', {
            url:'/site/ProxyData/:id',
            data:{
                // title:"代理数据"
                title:"ProxyData"
            },
            views:{
                "content@app": {
                    controller: 'ProxyDataCtrl',
                    templateUrl: "views/site/views/ProxyData.html"
                }
            }
        })
        .state('app.site.admin', {
            url:'/site/admin/:id',
            data:{
                // title:"代理数据"
                title:"admin"
            },
            views:{
                "content@app": {
                    controller: 'adminCtrl',
                    templateUrl: "views/site/views/admin.html"
                }
            }
        })
        .state('app.site.Maintenanceset', {
            url:'/site/Maintenanceset/:id',
            data:{
                // title:"代理数据"
                title:"Maintenanceset"
            },
            views:{
                "content@app": {
                    controller: 'MaintenancesetCtrl',
                    templateUrl: "views/site/views/Maintenanceset.html"
                }
            }
        })


    });
