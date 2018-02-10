angular.module('app.CashManagement', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.CashManagement').config(function ($stateProvider) {

    $stateProvider
        .state('app.CashManagement', {
            abstract: true,
            data: {
                title: 'CashManagement'
            }
        })
        .state('app.CashManagement.CashManagement', {
            url: '/CashManagement/CashManagement',
            data: {
                title: 'CashManagement'
            },
            views: {
                "content@app": {
                    controller: 'CashManagementCtrl as CashManagement',
                    templateUrl: "views/moneyManagement/accessManagement/CashManagement/views/CashManagement.html"
                }
            }
        })
        .state('app.CashManagement.Monitor', {
            url: '/CashManagement/Monitor',
            data: {
                title: 'monitor'
            },
            views: {
                "content@app": {
                    controller: 'MonitorCtrl as Monitor',
                    templateUrl: "views/moneyManagement/accessManagement/CashManagement/views/Monitor.html"
                }
            }
        })
});