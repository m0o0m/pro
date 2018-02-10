angular.module('app.caseEditor', ['ui.router', 'datatables', 'datatables.bootstrap']);

angular.module('app.caseEditor').config(function ($stateProvider) {

    $stateProvider
        .state('app.CaseEditor', {
            abstract: true,
            data: {
                // title: '案件编辑'
                title: 'CaseEditor'
            }
        })
        .state('app.CaseEditor.PendingCase', {
            url: '/CaseEditor/PendingCase',
            data: {
                // title: '待审案件'
                title: 'ThePendingCase'
            },
            views: {
                "content@app": {
                    controller: 'PendingCaseCtrl',
                    templateUrl: "views/webInformation/caseEditor/views/PendingCase.html"
                }
            }
        })
        .state('app.CaseEditor.AuditCase', {
            url: '/CaseEditor/AuditCase',
            data: {
                // title: '待审中案件'
                title: 'InThePendingCase'
            },
            views: {
                "content@app": {
                    controller: 'AuditCaseCtrl',
                    templateUrl: "views/webInformation/caseEditor/views/AuditCase.html"
                }
            }
        })
        .state('app.CaseEditor.ThroughCase', {
            url: '/CaseEditor/ThroughCase',
            data: {
                // title: '通过案件'
                title: 'ThroughCase'
            },
            views: {
                "content@app": {
                    controller: 'ThroughCaseCtrl',
                    templateUrl: "views/webInformation/caseEditor/views/ThroughCase.html"
                }
            }
        })
        .state('app.CaseEditor.RevokeCase', {
            url: '/CaseEditor/RevokeCase',
            data: {
                // title: '撤销案件'
                title: 'QuashCase'
            },
            views: {
                "content@app": {
                    controller: 'RevokeCaseCtrl',
                    templateUrl: "views/webInformation/caseEditor/views/RevokeCase.html"
                }
            }
        })

});