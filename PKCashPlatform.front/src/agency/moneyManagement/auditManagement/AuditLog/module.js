angular.module('app.AuditLog', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.AuditLog').config(function ($stateProvider) {

    $stateProvider
        .state('app.AuditLog', {
            abstract: true,
            data: {
                // title: '稽核日志'
                title: 'auditLog'
            }
        })

        .state('app.AuditLog.AuditLog', {
            url: '/AuditLog/AuditLog',
            data: {
                // title: '稽核日志'
                title: 'auditLog'
            },
            views: {
                "content@app": {
                    controller: 'AuditLogCtrl as AuditLog',
                    templateUrl: "views/moneyManagement/auditManagement/AuditLog/views/AuditLog.html"
                }
            }
        })

});