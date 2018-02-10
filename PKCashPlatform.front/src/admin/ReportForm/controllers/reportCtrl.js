angular.module('app.ReportForm').controller('reportCtrl',
    function(httpSvc,popupSvc,$scope,financeReportService,APP_CONFIG,$rootScope,$stateParams,$state,$timeout){
        //获取参数
        var postData = {
            ogStartTime:$stateParams.ogStartTime,
            ogEndTime:$stateParams.ogEndTime,
            otherStartTime:$stateParams.otherStartTime,
            otherEndTime:$stateParams.otherEndTime,
            module: $stateParams.module
        };

        //获取详情列表
        financeReportService.getReport(postData).then(function (res) {
            $scope.data = res.data;
        });

        $scope.open=function(event){
            var target=event.target;
            var aa=$(target).parent();
            var bb=$(aa).next().find(".table");
            if($(bb).css('display')=="none"){
                $(bb).show();
            }else{
                $(bb).hide();
            }

        };

    });
