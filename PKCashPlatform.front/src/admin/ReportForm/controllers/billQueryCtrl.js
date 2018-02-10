angular.module('app.ReportForm').controller('billQueryCtrl',
    function(httpSvc,popupSvc,$scope,financeReportService,APP_CONFIG,$rootScope,$stateParams,$state,$timeout){
        //获取参数
        var postData = {
            site:$scope.site,
            date:$scope.date,
            status:$scope.status
        };

        //获取详情列表
        financeReportService.getBill(postData).then(function (res) {
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

        $scope.modifyBill = function (id) {
            financeReportService.modifyBill({
                id: id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delBill = function (id) {
            financeReportService.delBill({
                id: id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.issuedBill = function (id) {
            financeReportService.issuedBill({
                id: id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };





    });
