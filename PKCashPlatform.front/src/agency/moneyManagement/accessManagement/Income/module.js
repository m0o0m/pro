angular.module('app.Income', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.Income').config(function ($stateProvider) {

    $stateProvider
        .state('app.Income', {
            abstract: true,
            data: {
                // title: '入款管理'
                title: 'MoneyManagement'
            }
        })
        .state('app.Income.Income', {
            url: '/Income/Income',
            data: {
                // title: '公司入款'
                title: 'companyInChina'
            },
            views: {
                "content@app": {
                    controller: 'IncomeCtrl as Income',
                    templateUrl: "views/moneyManagement/accessManagement/Income/views/Income.html"
                }
            }
        })
        .state('app.Income.Online', {
        url: '/Income/Online',
        data: {
            // title: '线上入款'
            title: 'ChinaOnline'
        },
        views: {
            "content@app": {
                controller: 'OnlineCtrl as Online',
                templateUrl: "views/moneyManagement/accessManagement/Income/views/Online.html"
            }
        }
    })

});