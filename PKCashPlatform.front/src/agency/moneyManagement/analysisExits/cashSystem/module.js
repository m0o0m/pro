angular.module('app.CashSystem', ['ui.router']);

angular.module('app.CashSystem').config(function ($stateProvider) {

    $stateProvider
        .state('app.CashSystem', {
            abstract: true,
            data: {
                title: 'CashSystem'
            }
        })

        .state('app.CashSystem.CashSystem', {
            url: '/CashSystem/CashSystem',
            data: {
                title: 'CashSystem'
            },
            views: {
                "content@app": {
                    controller: 'CashSystemCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/CashSystem/views/CashSystem.html"
                }
            }
        })



});