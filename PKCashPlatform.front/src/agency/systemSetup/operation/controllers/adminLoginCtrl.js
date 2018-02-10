angular.module('app.memberMessage').controller('adminLoginCtrl',
function(httpSvc,$scope,APP_CONFIG,operationService){
    //获取站点
    $scope.siteId = function (site_index_id) {
        operationService.getDropSelect(site_index_id).then(function (response) {
            $scope.sharedJson = response.data;
        });
    };
    $scope.siteId();

    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };

    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site: $scope.site_index_id,
            account: $scope.account,
            ip: $scope.ip,
            start_time: $scope.start_time,
            end_time: $scope.end_time
        };
        operationService.setSystemAdminLogin(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.search = function () {
        GetAllEmployee();
    }
});