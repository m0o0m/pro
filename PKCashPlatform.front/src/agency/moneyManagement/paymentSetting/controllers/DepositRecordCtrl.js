angular.module('app.PaymentSetting').controller('DepositRecordCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state,$stateParams){
    //获取的id
    $scope.id = $stateParams.Id;
    var GetAllEmployee = function () {
        var postData = {
            order_num:  $scope.order_num,
            set_id: $scope.id,
            start_time:$scope.startTime,
            end_time:$scope.endTime
        };
        paymentSettingService.paymentDeposit(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
                console.log($scope.list);
        });


    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);




});
