angular.module('app.quotaStatistics').controller('DropListCtrl', function($scope,APP_CONFIG,$state,QuotaStatisticsService){
    var GetAllEmployee = function () {
        //获取转入转出项目
        QuotaStatisticsService.getTypeList().then(function(response){
            console.log(response);
            $scope.listData = response.data.data;
        });
        $scope.json = APP_CONFIG.option;
        if($scope.start_time == undefined){
            $scope.start_time = "";
        }
        if($scope.end_time == undefined){
            $scope.end_time = "";
        }
        if($scope.username == undefined){
            $scope.username = "";
        }
        if($scope.type == undefined){
            $scope.type = "";
        }
        if($scope.ctype == undefined){
            $scope.vtype = "";
        }
        if($scope.type == undefined){
            $scope.type = "";
        }
        var postData = {
            type: $scope.type,
            username: $scope.username,
            ctype: $scope.ctype,
            vtype: $scope.vtype,
            start_time: $scope.start_time,
            end_time: $scope.end_time
        };
        QuotaStatisticsService.getDropList(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data.site_single_record_back;
            $scope.total = response.data.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.search = function () {
        GetAllEmployee();
    };

    $scope.xin = function () {
        $state.go('app.QuotaStatistics.SeparateApplication')
    };
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
