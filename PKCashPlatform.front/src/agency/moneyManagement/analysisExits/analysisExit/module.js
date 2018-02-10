angular.module('app.analysisExit', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.AnalysisExit', {
            abstract: true,
            data: {
                title: 'AnalysisExit'
            }
        })

        .state('app.AnalysisExit.AnalysisExit', {
            url: '/AnalysisExit/AnalysisExit',
            data: {
                title: 'AnalysisExit'
            },
            views: {
                "content@app": {
                    controller: 'AnalysisExitCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/AnalysisExit/views/AnalysisExit.html"
                }
            }
        })
        .state('app.AnalysisExit.MemberInquiries', {
            url: '/AnalysisExit/MemberInquiries',
            data: {
                title: 'MemberInquiries'
            },
            views: {
                "content@app": {
                    controller: 'MemberInquiriesCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/AnalysisExit/views/MemberInquiries.html"
                }
            }
        })
        .state('app.AnalysisExit.PurchaseAnalysis', {
            url: '/AnalysisExit/PurchaseAnalysis',
            data: {
                title: 'PurchaseAnalysis'
            },
            views: {
                "content@app": {
                    controller: 'PurchaseAnalysisCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/AnalysisExit/views/PurchaseAnalysis.html"
                }
            }
        })
        .state('app.AnalysisExit.PreferentialAnalysis', {
            url: '/AnalysisExit/PreferentialAnalysis',
            data: {
                title: 'PreferentialAnalysis'
            },
            views: {
                "content@app": {
                    controller: 'PreferentialAnalysisCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/AnalysisExit/views/PreferentialAnalysis.html"
                }
            }
        })
        .state('app.AnalysisExit.ValidList', {
            url: '/AnalysisExit/ValidList',
            data: {
                title: 'ValidList'
            },
            views: {
                "content@app": {
                    controller: 'ValidListCtrl',
                    templateUrl: "views/moneyManagement/analysisExits/AnalysisExit/views/ValidList.html"
                }
            }
        })
});
