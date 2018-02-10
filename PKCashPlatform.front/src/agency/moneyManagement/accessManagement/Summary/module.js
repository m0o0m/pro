angular.module('app.Summary', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.Summary').config(function ($stateProvider) {

    $stateProvider
        .state('app.Summary', {
            abstract: true,
            data: {
                // title: '出入项目汇总'
                title: 'entryAndExit'
            }
        })

        .state('app.Summary.Summary', {
            url: '/Summary/Summary',
            data: {
                // title: '出入项目汇总'
                title: 'entryAndExit'
            },
            views: {
                "content@app": {
                    controller: 'SummaryCtrl as Summary',
                    templateUrl: "views/moneyManagement/accessManagement/Summary/views/Summary.html"
                }
            }
        })

});