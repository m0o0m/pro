angular.module('app.quotaStatistics').controller('QuotaRecordsCtrl', function($scope,APP_CONFIG,$state,QuotaStatisticsService){
    QuotaStatisticsService.getTypeList().then(function (response) {
        console.log(response);
        $scope.video = response.data.data;
    });
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {

        if($scope.start_time == undefined){
            $scope.start_time = "";
        }
        if($scope.end_time == undefined){
            $scope.end_time = "";
        }
        if($scope.admin_name == undefined){
            $scope.admin_name = "";
        }
        if($scope.vd_type == undefined){
            $scope.vd_type = "";
        }
        if($scope.do_type == undefined){
            $scope.do_type = "";
        }
        if($scope.cash_type == undefined){
            $scope.cash_type = "";
        }
        var postData = {
            admin_name: $scope.adminName,
            vd_type: $scope.vdType,
            do_type: $scope.doType,
            cash_type: $scope.cashType,
            start_time: $scope.startTime,
            end_time: $scope.endTime
        }
        QuotaStatisticsService.getQuotaRecord(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data.quota_record;
            $scope.total = response.data.data;
        })
    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.search = function () {
        GetAllEmployee();
    }
    //筛选展开
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
});
