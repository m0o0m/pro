angular.module('app.quotaStatistics').controller('RechargeRecordCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,QuotaStatisticsService){
    var GetAllEmployee = function () {
        if($scope.start_time == undefined){
            $scope.start_time = "";
        }
        if($scope.end_time == undefined){
            $scope.end_time = "";
        }
        if($scope.status == undefined){
            $scope.status = "";
        }
        if($scope.type == undefined){
            $scope.type = "";
        }
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            status:$scope.status,
            type:$scope.type,
            start_time:$scope.start_time,
            end_time:$scope.end_time
        }
        QuotaStatisticsService.getQuotaRecharge(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
            console.log($scope.list);
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