angular.module('app.Platform', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.Platform', {
            abstract: true,
            data: {
                title: 'Platform'
            }
        })
        .state('app.Platform.Commodity', {
            url: '/Platform/Commodity',
            data: {
                title: 'Commodity'
            },
            views: {
                "content@app": {
                    controller: 'CommodityCtrl',
                    templateUrl: "views/Platform/views/Commodity.html"
                }
            }
        })

        .state('app.Platform.Platformaccount', {
            url: '/Platform/PlatformAccount',
            data: {
                title: 'PlatformAccount'
            },
            views: {
                "content@app": {
                    controller: 'PlatformAccountCtrl',
                    templateUrl: "views/Platform/views/PlatformAccount.html"
                }
            }
        })

        .state('app.Platform.accountHolder', {
            url: '/Platform/accountHolder',
            data: {
                // title: '开户人管理'
                title: 'A_10'
            },
            views: {
                "content@app": {
                    controller: 'AccountHolderCtrl',
                    templateUrl: "views/Platform/views/accountHolder.html"
                }
            }

        })
        .state('app.Platform.site', {
            url: '/Platform/site?id&site',
            data: {
                // title: '站点管理'
                title: 'Site management'
            },
            views: {
                "content@app": {
                    controller: 'SiteCtrl',
                    templateUrl: "views/Platform/views/site.html"
                }
            }

        })
        .state('app.Platform.domain', {
            url: '/Platform/domain?id&site',
            data: {
                //123 title: '域名管理'
                title: '域名管理'
            },
            views: {
                "content@app": {
                    controller: 'DomainCtrl',
                    templateUrl: "views/Platform/views/domain.html"
                }
            }

        })
        .state('app.Platform.role', {
            url: '/Platform/role',
            data: {
                // title: '角色管理'
                title: 'RoleManagement'
            },
            views: {
                "content@app": {
                    controller: 'RoleCtrl',
                    templateUrl: "views/Platform/views/role.html"
                }
            }

        })
        .state('app.Platform.permissionConfig', {
            url: '/Platform/permissionConfig?id&&role_mark',
            data: {
                //123 title: '权限配置'
                title: '权限配置'
            },
            views: {
                "content@app": {
                    controller: 'PermissionConfigCtrl',
                    templateUrl: "views/Platform/views/permissionConfig.html"
                }
            }
        })
        .state('app.Platform.Managementtype', {
            url: '/Platform/Managementtype',
            data: {
                //123 title: '类型管理'
                title: 'Typemanagement'
            },
            views: {
                "content@app": {
                    controller: 'ManagementtypeCtrl',
                    templateUrl: "views/Platform/views/Managementtype.html"
                }
            }
        })
        .state('app.Platform.Function', {
            url: '/Platform/Function',
            data: {
                //123 title: '功能管理'
                title: 'Functionmanagement'
            },
            views: {
                "content@app": {
                    controller: 'FunctionCtrl',
                    templateUrl: "views/Platform/views/Function.html"
                }
            }
        })

        // .state('app.app.Platform.menu', {
        //     url: '/Platform/menu',
        //     data: {
        //         title: 'SettlementRequest'
        //     },
        //     views: {
        //         "content@app": {
        //             controller: 'RequestCtrl',
        //             templateUrl: "app/views/settlement/views/SettlementRequest.html"
        //         },
        //         resolve: {
        //             srcipts: function(lazyScript){
        //                 return lazyScript.register([
        //                     'app/vendor.ui.js'
        //                 ])
        //
        //             }
        //         }
        //     }
        // })
        //

        .state('app.Platform.menu', {
            url: '/Platform/menu',
            data: {
                // title: '菜单管理'
                title: 'MenuManagement'
            },
            views: {
                "content@app": {
                    controller: 'menuCtrl',
                    templateUrl: "views/Platform/views/menu.html"
                },
                resolve: {
                    srcipts: function(lazyScript){
                        return lazyScript.register([
                            'build/vendor.ui.js'
                        ])

                    }
                }
            }
        })
        .state('app.Platform.Package', {
            url: '/Platform/Package',
            data: {
                // title: '套餐管理'
                title: 'PackageManagement'
            },
            views: {
                "content@app": {
                    controller: 'PackageCtrl',
                    templateUrl: "views/Platform/views/Package.html"
                }
            }

        })
        .state('app.Platform.allocation', {
            url: '/Platform/allocation?ids',
            data: {
                // title: '套餐管理'
                title: 'configuration'
            },
            params:{ids:null},
            views: {
                "content@app": {
                    controller: 'allocationCtrl',
                    templateUrl: "views/Platform/views/allocation.html"
                }
            }

        })
        .state('app.Platform.menuCarte', {
            url: '/Platform/menuCarte?idese',
            data: {
                // title: '菜单管理'
                title: 'menuCarte'
            },
            params:{idese:null},
            views: {
                "content@app": {
                    controller: 'menuCarteCtrl',
                    templateUrl: "views/Platform/views/menuCarte.html"
                },
                resolve: {
                    srcipts: function(lazyScript){
                        return lazyScript.register([
                            'build/vendor.ui.js'
                        ])

                    }
                }
            }
        })
        .state('app.Platform.Operationlog', {
            url: '/Platform/Operationlog',
            data: {
                title: '操作日志'
            },
            views: {
                "content@app": {
                    controller: 'OperationlogCtrl',
                    templateUrl: "views/Platform/views/Operationlog.html"
                }
            }
        })
        .state('app.Platform.Logmanagement', {
            url: '/Platform/Logmanagement',
            data: {
                title: '日志管理'
            },
            views: {
                "content@app": {
                    controller: 'LogmanagementCtrl',
                    templateUrl: "views/Platform/views/Logmanagement.html"
                }
            }
        })
        .state('app.Platform.Sitecolumn', {
            url: '/Platform/Sitecolumn',
            data: {
                title: '站点栏目'
            },
            views: {
                "content@app": {
                    controller: 'SitecolumnCtrl',
                    templateUrl: "views/Platform/views/Sitecolumn.html"
                }
            }
        })
});
