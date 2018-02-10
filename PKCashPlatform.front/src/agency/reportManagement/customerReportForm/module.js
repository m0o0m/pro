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
                title: 'Report form'
            },
            views: {
                "content@app": {
                    controller: 'ReportFormCtrl',
                    templateUrl: "views/reportManagement/customerReportForm/views/ReportForm.html"
                }
            }
        })
        .state('app.ReportForm.ReportDetails', {
            url: '/ReportForm/ReportDetails',
            data: {
                title: 'Report details'
            },
            params:{
                site_index_id:null,
                time_zone:null,
                username:null,
                rtype:null,
                v_type:null,
                start_time:null,
                end_time:null
            },
            views: {
                "content@app": {
                    controller: 'ReDetailCtrl',
                    templateUrl: "views/reportManagement/customerReportForm/views/ReportDetails.html"
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
                title: 'Membership details'
            },
            params:{id:null},
            views: {
                "content@app": {
                    controller: 'MemberdetailsCtrl',
                    templateUrl: "views/reportManagement/customerReportForm/views/Memberdetails.html"
                }
            },
            resolve: {
                scripts: function(lazyScript){
                    return lazyScript.register('vendor.ui.js')
                }
            }
        })



});
