angular.module('app.quotaStatistics', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.QuotaStatistics', {
            abstract: true,
            data: {
                title: 'QuotaStatistics'
            }
        })

        .state('app.QuotaStatistics.QuotaStatistics', {
            url: '/QuotaStatistics/QuotaStatistics',
            data: {
                title: 'QuotaStatistics'
            },
            views: {
                "content@app": {
                    controller: 'QuotaStatisticsCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/QuotaStatistics.html"
                }
            }
        })
        .state('app.QuotaStatistics.RechargeRecord', {
            url: '/QuotaStatistics/RechargeRecord',
            data: {
                title: 'RechargeRecord'
            },
            views: {
                "content@app": {
                    controller: 'RechargeRecordCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/RechargeRecord.html"
                }
            }
        })
        .state('app.QuotaStatistics.DropList', {
            url: '/QuotaStatistics/DropList',
            data: {
                title: 'DropList'
            },
            views: {
                "content@app": {
                    controller: 'DropListCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/DropList.html"
                }
            }
        })
        .state('app.QuotaStatistics.QuotaRecharge', {
            url: '/QuotaStatistics/QuotaRecharge',
            data: {
                title: 'QuotaRecharge'
            },
            views: {
                "content@app": {
                    controller: 'QuotaRechargeCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/QuotaRecharge.html"
                }
            }
        })
        .state('app.QuotaStatistics.QuotaRecord', {
            url: '/QuotaStatistics/QuotaRecord',
            data: {
                title: 'QuotaRecord'
            },
            views: {
                "content@app": {
                    controller: 'QuotaRecordsCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/QuotaRecord.html"
                }
            }
        })
        .state('app.QuotaStatistics.SeparateApplication', {
            url: '/QuotaStatistics/SeparateApplication',
            data: {
                title: 'SeparateApplication'
            },
            views: {
                "content@app": {
                    controller: 'SeparateApplicationCtrl',
                    templateUrl: "views/moneyManagement/amountStatistics/quotaStatistics/views/SeparateApplication.html"
                }
            }
        })
});