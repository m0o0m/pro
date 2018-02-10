angular.module('app.AuditQuery', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.AuditQuery').config(function ($stateProvider) {

    $stateProvider
        .state('app.AuditQuery', {
            abstract: true,
            data: {
                // title: '即时稽核查询'
                title: 'Immediate'
            }
        })

        .state('app.AuditQuery.AuditQuery', {
            url: '/AuditQuery/AuditQuery',
            data: {
                // title: '即时稽核查询'
                title: 'Immediate'
            },
            views: {
                "content@app": {
                    controller: 'AuditQueryCtrl as AuditQuery',
                    templateUrl: "views/moneyManagement/auditManagement/AuditQuery/views/AuditQuery.html"
                }
            }
        })

});