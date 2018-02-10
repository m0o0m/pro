angular.module('app.websiteManagement', [ 'ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.websiteManagement').config(function ($stateProvider) {

    $stateProvider
        .state('app.websiteManagement', {
            abstract: true,
            data: {
                title: 'WebsiteManagement'
            }
        })
        .state('app.websiteManagement.WebsiteInformation', {
            url: '/websiteManagement/WebsiteInformation',
            data: {
                // title: '网站信息'
                title: 'WebsiteInformation'
            },
            views: {
                "content@app": {
                    controller: 'WebsiteInformationCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/WebsiteInformation.html"
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
        .state('app.websiteManagement.videoManagement', {
            url: '/websiteManagement/videoManagement',
            data: {
                // title: '视讯管理'
                title: 'videoManagement'
            },
            views: {
                "content@app": {
                    controller: 'videoManagementCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/videoManagement.html"
                }
            }
        })
        .state('app.websiteManagement.electronicManagement', {
            url: '/websiteManagement/electronicManagement',
            data: {
                // title: '电子管理'
                title: 'electronicManagement'
            },
            views: {
                "content@app": {
                    controller: 'electronicManagementCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/electronicManagement.html"
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
        .state('app.websiteManagement.sportManagement', {
            url: '/websiteManagement/sportManagement',
            data: {
                // title: '体育管理'
                title: 'sportManagement'
            },
            views: {
                "content@app": {
                    controller: 'sportManagementCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/sportManagement.html"
                }
            }
        })
        .state('app.websiteManagement.lotteryManagrment', {
            url: '/websiteManagement/lotteryManagrment',
            data: {
                // title: '彩票管理'
                title: 'lotteryManagrment'
            },
            views: {
                "content@app": {
                    controller: 'lotteryManagrmentCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/lotteryManagrment.html"
                }
            }
        })
        .state('app.websiteManagement.lotterySort', {
            url: '/websiteManagement/lotterySort',
            data: {
                //123 title: '彩票排序'
                title: 'lotteryManagrment'
            },
            views: {
                "content@app": {
                    controller: 'lotterySortCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/lotterySort.html"
                }
            }
        })
        .state('app.websiteManagement.modularManagement', {
            url: '/websiteManagement/modularManagement',
            data: {
                // title: '模块管理'
                title: 'modularManagement'
            },
            views: {
                "content@app": {
                    controller: 'modularManagementCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/modularManagement.html"
                }
            }
        })
        .state('app.websiteManagement.PreferencesSettings', {
            url: '/websiteManagement/PreferencesSettings',
            data: {
                // title: '自助申请优惠设置'
                title: 'PreferencesSettings'
            },
            views: {
                "content@app": {
                    controller: 'PreferencesSettingsCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/PreferencesSettings.html"
                }
            }
        })
        .state('app.websiteManagement.cacheManagement', {
            url: '/websiteManagement/cacheManagement',
            data: {
                // title: '缓存管理'
                title: 'cacheManagement'
            },
            views: {
                "content@app": {
                    controller: 'cacheManagementCtrl',
                    templateUrl: "views/webInformation/websiteManagement/views/cacheManagement.html"
                }
            }
        });
});