angular.module('app.ReportForm', [
    'ui.router'
]).config(function ($stateProvider) {
    $stateProvider
        .state('app.ReportForm', {
            abstract: true,
            data: {
                title: 'ReportForm'
            }
        })

        .state('app.ReportForm.ReportForm', {
            url: '/ReportForm/ReportForm',
            data: {
                title: 'ReportForm'
            },
            views: {
                "content@app": {
                    controller: 'ReportFormCtrl',
                    templateUrl: "views/ReportForm/views/ReportForm.html"
                }
            }
        })
        .state('app.ReportForm.ReportDetails', {
            url: '/ReportForm/ReportDetails',
            data: {
                title: 'ReportDetails'
            },
            views: {
                "content@app": {
                    controller: 'ReDetailCtrl',
                    templateUrl: "views/ReportForm/views/ReportDetails.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })
        .state('app.ReportForm.Memberdetails', {
            url: '/ReportForm/Memberdetails',
            data: {
                title: 'Memberdetails'
            },
            views: {
                "content@app": {
                    controller: 'MemberdetailsCtrl',
                    templateUrl: "views/ReportForm/views/Memberdetails.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })
        .state('app.ReportForm.hierarchical', {
            url: '/ReportForm/hierarchical',
            data: {
                title: '层级管理'
            },
            views: {
                "content@app": {
                    controller: 'hierarchicalCtrl',
                    templateUrl: "views/ReportForm/views/hierarchical.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })
        .state('app.ReportForm.hierarchicalSetting', {
            url: '/ReportForm/hierarchicalSetting',
            data: {
                title: '层级设定'
            },
            views: {
                "content@app": {
                    controller: 'hierarchicalSetCtrl',
                    templateUrl: "views/ReportForm/views/hierarchicalSetting.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })
        .state('app.ReportForm.Summary', {
            url: '/ReportForm/Summary',
            data: {
                title: '出入款项目汇总'
            },
            views: {
                "content@app": {
                    controller: 'SummaryCtrl',
                    templateUrl: "views/ReportForm/views/Summary.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })
        .state('app.ReportForm.packageManagement', {
            url: '/ReportForm/packageManagement',
            data: {
                title: '套餐管理'
            },
            views: {
                "content@app": {
                    controller: 'packageManagementCtrl',
                    templateUrl: "views/ReportForm/views/packageManagement.html"
                }
            }
        })
        .state('app.ReportForm.cashStatement', {
            url: '/ReportForm/cashStatement',
            data: {
                title: '现金报表'
            },
            views: {
                "content@app": {
                    controller: 'cashStatementCtrl',
                    templateUrl: "views/ReportForm/views/cashStatement.html"
                }
            }
        })
        .state('app.ReportForm.referential', {
            url: '/ReportForm/referential',
            data: {
                title: '优惠统计'
            },
            views: {
                "content@app": {
                    controller: 'referentialCtrl',
                    templateUrl: "views/ReportForm/views/referential.html"
                }
            }
        })


        .state('app.ReportForm.Arrears', {
            url: '/ReportForm/Arrears',
            data: {
                title: '催款查询'
            },
            views: {
                "content@app": {
                    controller: 'ArrearsCtrl',
                    templateUrl: "views/ReportForm/views/Arrears.html"
                }
            }
        })
        .state('app.ReportForm.QuotaNum', {
            url: '/ReportForm/QuotaNum',
            data: {
                title: '额度统计'
            },
            views: {
                "content@app": {
                    controller: 'QuotaNumCtrl',
                    templateUrl: "views/ReportForm/views/QuotaNum.html"
                }
            }
        })
        .state('app.ReportForm.quotaRecord', {
            url: '/ReportForm/quotaRecord',
            data: {
                title: '额度记录'
            },
            views: {
                "content@app": {
                    controller: 'quotaRecordCtrl',
                    templateUrl: "views/ReportForm/views/quotaRecord.html"
                }
            }
        })
        .state('app.ReportForm.rechargeRecord', {
            url: '/ReportForm/rechargeRecord',
            data: {
                title: '充值记录'
            },
            views: {
                "content@app": {
                    controller: 'rechargeRecordCtrl',
                    templateUrl: "views/ReportForm/views/rechargeRecord.html"
                }
            }
        })
        .state('app.ReportForm.moneyManagement', {
            url: '/ReportForm/moneyManagement',
            data: {
                title: '入款管理'
            },
            views: {
                "content@app": {
                    controller: 'moneyManagementCtrl',
                    templateUrl: "views/ReportForm/views/moneyManagement.html"
                }
            }
        })
        .state('app.ReportForm.IncomeNum', {
            url: '/ReportForm/IncomeNum',
            data: {
                title: '入款统计'
            },
            views: {
                "content@app": {
                    controller: 'IncomeNumCtrl',
                    templateUrl: "views/ReportForm/views/IncomeNum.html"
                }
            }
        })
        .state('app.ReportForm.DataCenter', {
            url: '/ReportForm/DataCenter',
            data: {
                title: '数据中心'
            },
            views: {
                "content@app": {
                    controller: 'DataCenterCtrl',
                    templateUrl: "views/ReportForm/views/DataCenter.html"
                }
            }
        })
        .state('app.ReportForm.reportStatistics', {
            url: '/ReportForm/reportStatistics',
            data: {
                title: '站点统计数据'
            },
            views: {
                "content@app": {
                    controller: 'reportStatisticsCtrl',
                    templateUrl: "views/ReportForm/views/reportStatistics.html"
                }
            }
        })
        .state('app.ReportForm.reportQuery', {
            url: '/ReportForm/reportQuery',
            data: {
                title: '报表查询'
            },
            views: {
                "content@app": {
                    controller: 'reportQueryCtrl',
                    templateUrl: "views/ReportForm/views/ReportQuery.html"
                }
            }
        })
        .state('app.ReportForm.billQuery', {
            url: '/ReportForm/billQuery',
            data: {
                title: '账单查询'
            },
            views: {
                "content@app": {
                    controller: 'billQueryCtrl',
                    templateUrl: "views/ReportForm/views/billQuery.html"
                }
            }
        })
        .state('app.ReportForm.report', {
            url: '/ReportForm/report',
            data: {
                title: '报表统计'
            },
            views: {
                "content@app": {
                    controller: 'reportCtrl',
                    templateUrl: "views/ReportForm/views/report.html"
                }
            }
        })


});
