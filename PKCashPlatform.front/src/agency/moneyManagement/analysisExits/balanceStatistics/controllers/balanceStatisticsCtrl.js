angular.module('app.balanceStatistics').controller('balanceStatisticsCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,BalanceStatisticsService){
    $scope.sitId = function (site_index_id) {
        BalanceStatisticsService.getAgencySelect().then(function (response) {
            $scope.agency_select = response.data.data;
        });
    }
    var user=JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin=user.site_index_id==='';
    if($scope.isSuperAdmin){
        //获取站点
        $scope.sitId();
    }else{
        $scope.site_index_id=user.site_index_id;
    }
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        var checkbox = document.getElementById('test');
        if(checkbox.checked){
            $scope.check=1;
        }else{
            $scope.check=0;
        }
        if($scope.start_time ===null){
            $scope.start_time="";
        }
        if($scope.end_time===null){
            $scope.end_time="";
        }
        if($scope.agency===null){
            $scope.agency="";
        }
        if($scope.account===null){
            $scope.account="";
        }

        var postData = {
            start_time:$scope.start_time,
            end_time:$scope.end_time,
            agency:$scope.agency,
            account:$scope.account,
            is_check:$scope.check
        };
        BalanceStatisticsService.getBalance(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
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


