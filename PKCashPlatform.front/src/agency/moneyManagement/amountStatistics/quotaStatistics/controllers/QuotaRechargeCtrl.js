angular.module('app.quotaStatistics').controller('QuotaRechargeCtrl', function($scope,APP_CONFIG,$state,QuotaStatisticsService,popupSvc,$rootScope){
    QuotaStatisticsService.getRecordOrdernum().then(function (response) {
        console.log(response.data.order_num);
        $scope.threeOrder = response.data.order_num;
    });
    QuotaStatisticsService.getBankOrdernum().then(function (response) {
        console.log(response.data.order_num);
        $scope.order = response.data.order_num;
    });
    $scope.json = APP_CONFIG.option;
    QuotaStatisticsService.getThreeBank().then(function (response) {
        console.log(response);
        $scope.three = response.data;
    });
    QuotaStatisticsService.getRecordCardBank().then(function (response) {
        console.log(response);
        $scope.list = response.data;
    });
    $scope.bankSubmit = function () {
        var postData = {
            order_num: $scope.order,
            money: $scope.money*1,
            account: $scope.account,
            bank: $scope.bank*1
        }
        QuotaStatisticsService.getBankSub(postData).then(function (response) {
            console.log(response);
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    }
    $scope.submit = function () {
        var postData = {
            order_num: $scope.threeOrder,
            money: $scope.threeMoney*1,
            type: $scope.threeType*1,
            bank: $scope.threeBank*1
        }
        QuotaStatisticsService.getThreeSub(postData).then(function (response) {
            console.log(response);
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    }

});
