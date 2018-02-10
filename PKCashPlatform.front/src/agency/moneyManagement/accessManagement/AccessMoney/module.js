angular.module('app.AccessMoney', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.AccessMoney').config(function ($stateProvider) {

    $stateProvider
        .state('app.AccessMoney', {
            abstract: true,
            data: {
                // title: '存款与取款'
                title: 'ExitManagement'
            }
        })
        .state('app.AccessMoney.AccessMoney', {
            url: '/AccessMoney/AccessMoney',
            data: {
                // title: '人工存款'
                title: 'DepositsWithdrawals'
            },
            views: {
                "content@app": {
                    controller: 'AccessMoneyCtrl as AccessMoney',
                    templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/AccessMoney.html"
                }
            }
        })
        .state('app.AccessMoney.Deposit', {
            url: '/AccessMoney/Deposit',
            data: {
                // title: '人工存款/批量操作'
                title: 'bulkDeposit'
            },
            views: {
                "content@app": {
                    controller: 'DepositCtrl',
                    templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/Deposit.html"
                }
            }
        })
        .state('app.AccessMoney.History', {
        url: '/AccessMoney/History',
        data: {
            // title: '出入款记录'
            title: 'HistoryQuery'
        },
        views: {
            "content@app": {
                controller: 'HistoryCtrl as History',
                templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/History.html"
            }
        }
    })
        .state('app.AccessMoney.Quota', {
            url: '/AccessMoney/Quota',
            data: {
                // title: '额度转换'
                title: 'QuotaConversion'
            },
            views: {
                "content@app": {
                    controller: 'QuotaCtrl as Quota',
                    templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/Quota.html"
                }
            }
        })
        .state('app.AccessMoney.QuotaRecord', {
            url: '/AccessMoney/QuotaRecord',
            data: {
                // title: '额度转换记录'
                title: 'MarginReplacement'
            },
            views: {
                "content@app": {
                    controller: 'QuotaRecordCtrl as QuotaRecord',
                    templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/QuotaRecord.html"
                }
            }
        })
        .state('app.AccessMoney.ArtificialWithdrawal', {
            url: '/AccessMoney/ArtificialWithdrawal',
            data: {
                // title: '人工取款'
                title: 'ArtificialWithdrawal'
            },
            views: {
                "content@app": {
                    controller: 'ArtificialWithdrawalCtrl as Artificial',
                    templateUrl: "views/moneyManagement/accessManagement/AccessMoney/views/ArtificialWithdrawal.html"
                }
            }
        })
        // .state('app.AccessMoney.AccessMoney', {
        //     url: '/AccessMoney/AccessMoney',
        //     data: {
        //         title: '存款与取款'
        //     },
        //     views: {
        //         "content@app": {
        //             controller: 'AccessMoneyCtrl as AccessMoney',
        //             templateUrl: "app/views/AccessMoney/views/AccessMoney.html"
        //         }
        //     }
        // })
});