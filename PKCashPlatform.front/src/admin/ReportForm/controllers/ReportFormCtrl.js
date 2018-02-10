angular.module('app.ReportForm').controller('ReportFormCtrl',
    function ($scope, popupSvc, commonService, financeReportService, $rootScope, APP_CONFIG, $state) {
        $scope.datepickerOptions = {
            changeMonth: true,
            changeYear: true
        };
        //获取模块
        financeReportService.getModule().then(function (response) {
            $scope.module = response.data;
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

            $state.go('app.ReportForm.ReportDetails',{
                ogStartTime:$scope.ogStartTime,
                ogEndTime:$scope.ogEndTime,
                otherStartTime:$scope.otherStartTime,
                otherEndTime:$scope.otherEndTime,
                module: module
            });
        }

    });

