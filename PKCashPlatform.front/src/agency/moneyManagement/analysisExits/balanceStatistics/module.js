angular.module("app.balanceStatistics",['ui.router', 'datatables', 'datatables.bootstrap']).config(function ($stateProvider) {
    $stateProvider
        .state('app.balanceStatistics',{
            abstract:true,
            data:{
                title:"MemberInquiries"
            }
        })
        .state('app.balanceStatistics.balanceStatistics',{
            url: '/balanceStatistics/balanceStatistics',
            data: {
                title: 'MemberInquiries'
            },
            views:{
                'content@app':{
                    controller:'balanceStatisticsCtrl',
                    templateUrl: 'views/moneyManagement/analysisExits/balanceStatistics/views/balanceStatistics.html'
                }
            }
        })

});/**
 * Created by apple on 17/9/4.
 */
