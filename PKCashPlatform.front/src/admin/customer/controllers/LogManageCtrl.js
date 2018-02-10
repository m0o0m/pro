angular.module('app.customer').controller('LogManageCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,APP_CONFIG,customerService){
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };

    //获取站点
    $scope.siteId = function () {
        customerService.getSite().then(function (response) {
            $scope.siteJson  = response.data;
        });
    };
    $scope.siteId();

    //登录日志
    $scope.accountJson = APP_CONFIG.option.option_account_id;
    console.log($scope.accountJson);
    $scope.account_id = '1';
    $scope.loginJson = APP_CONFIG.option.option_login_id;
    $scope.login_id = '1';
    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site_index_id,
            strat_time:$scope.startTime,
            end_time:$scope.endTime,
            site_id:$scope.site_id,
            account_id:$scope.account_id,
            login_id:$scope.login_id,
            text:$scope.text,
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        };
        customerService.getLoginLog(postData).then(function (response) {
            $scope.commodity = response.data;
            $scope.paginationConf.totalItems = response.meta.count;
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

});

