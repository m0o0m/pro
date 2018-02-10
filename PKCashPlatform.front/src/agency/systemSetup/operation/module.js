angular.module("app.operation",['ui.router', 'datatables', 'datatables.bootstrap']).config(function ($stateProvider) {
    $stateProvider
        .state('app.Operation',{
            abstract:true,
            data:{
                title:"LogManagement"
            }
        })
        .state('app.Operation.Operation',{
            url: '/Operation/Operation',
            data: {
                title: 'OperationLog'
            },
            views:{
                'content@app':{
                    controller:'operaCtrl',
                    templateUrl: 'views/systemSetup/operation/views/operation.html'
                }
            }
        })
        .state('app.Operation.memberLogin',{
            url: '/Operation/memberLogin',
            data: {
                title: 'MemberLogin'
            },
            views:{
                'content@app':{
                    controller:'memberLoginCtrl',
                    templateUrl: 'views/systemSetup/operation/views/memberLogin.html'
                }
            }
        })
        .state('app.Operation.adminLogin',{
            url: '/Operation/adminLogin',
            data: {
                title: 'AdministratorLogin'
            },
            views:{
                'content@app':{
                    controller:'adminLoginCtrl',
                    templateUrl: 'views/systemSetup/operation/views/adminLogin.html'
                }
            }
        })
        .state('app.Operation.automaticAudit',{
            url: '/Operation/automaticAudit',
            data: {
                title: 'AutomaticAudit'
            },
            views:{
                'content@app':{
                    controller:'automaticAuditCtrl',
                    templateUrl: 'views/systemSetup/operation/views/automaticAudit.html'
                }
            }
        })
});
