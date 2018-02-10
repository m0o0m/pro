angular.module('app.AuditLog').controller('AuditLogCtrl',
    function(AccessMoneyService,httpSvc, popupSvc, $stateParams, $scope, $rootScope, APP_CONFIG){
        //获取站点下拉框
        AccessMoneyService.getDropSelect().then(function (response) {
            $scope.siteJson = response.data;
            console.log($scope.site);
        });
        

    var GetAllEmployee = function () {
        var postData = {
            site_id:$scope.site,
            site_index_id:$scope.site_index,
            account: $scope.accounts,
            start_time:$scope.startTime,
            end_time:$scope.endTime
        }
        AccessMoneyService.memberAuditnow(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta[0].count;
            $scope.list = response.data.data;
            console.log(response);
        });
    };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    //点击搜索
     $scope.search = function () {
         GetAllEmployee();
     }
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