angular.module('app.quotaStatistics').controller('QuotaStatisticsCtrl', function(httpSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,QuotaStatisticsService){
    console.log(1111);
    var GetAllEmployee = function () {
        console.log($scope.account);
        console.log($scope.start_time);
        console.log($scope.end_time);
        if($scope.account == undefined){
            $scope.account = "";
        }
        if($scope.start_time == undefined){
            $scope.start_time = "";
        }
        if($scope.end_time == undefined){
            $scope.end_time = "";
        }
        var postData = {
            account:$scope.account,
            start_time:$scope.start_time,
            end_time:$scope.end_time
        }
        QuotaStatisticsService.getQuotaList(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data.Data;
            $scope.total = response.data.data.Total;
        })
    }
    GetAllEmployee();

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
