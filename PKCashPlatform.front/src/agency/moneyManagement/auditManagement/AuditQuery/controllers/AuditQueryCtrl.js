angular.module('app.AuditQuery').controller('AuditQueryCtrl',
    function(AccessMoneyService,httpSvc, popupSvc, $stateParams, $scope, $rootScope, APP_CONFIG){
        // //获取站点下拉框
        AccessMoneyService.getDropSelect().then(function (response) {
            $scope.siteJson = response.data;
            console.log($scope.site);
        });

        var GetAllEmployee = function () {
        var postData = {
            site:$scope.site,
            account: $scope.accounts,
            };
            AccessMoneyService.auditLog(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
                console.log($scope.list);
                $scope.v1 =response.Effective_betting;
                $scope.v2 = response.Preferential_audit;
                $scope.v3 = response.Total
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
    };

    //获取实际购买
    $scope.Actual = function (id) {
        AccessMoneyService.actualPurchase(id).then(function (res) {
            console.log(res.data.data);
            $scope.Purchase = res.data.data;
        });
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