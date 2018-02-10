angular.module('app.ReportForm').controller('reportQueryCtrl',
    function ($scope, popupSvc, commonService, financeReportService, $rootScope, APP_CONFIG, $state) {
        $scope.datepickerOptions = {
            changeMonth: true,
            changeYear: true
        };
        commonService.getSite().then(function (response) {
            $scope.siteList = response.data;
        });
        //获取模块
        financeReportService.getGame().then(function (response) {
            $scope.gameList = response.data;
        });

        $scope.detail=function(){
            var module=[];
            var list=$(".check");
            for(var i=0; i<list.length; i++){
                if($(list[i]).prop("checked")){
                    var id=$(list[i]).parent().find(".module-id")[0].innerHTML;
                    module.push(id);
                }
            };
            console.log(module);

            $state.go('app.ReportForm.report',{
                ogStartTime:$scope.ogStartTime,
                ogEndTime:$scope.ogEndTime,
                otherStartTime:$scope.otherStartTime,
                otherEndTime:$scope.otherEndTime,
                module: module
            });
        }

    });

