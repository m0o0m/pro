angular.module('app.customer').controller('NoticeKeyCtrl', function(httpSvc,popupSvc,$scope,APP_CONFIG,customerService,$rootScope,$LocalStorage){

    $scope.sitId = function (site_index_id) {
        customerService.getSite(site_index_id).then(function (response) {
            $scope.sharedJson = response.data;
            console.log($scope.sharedJson);
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        console.log($scope.siteId);
        $scope.sitId(user.site_index_id);
    }
    $scope.json = APP_CONFIG.option;
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id:$scope.site_index_id,
            account:$scope.id,
            status:$scope.status
        };
        customerService.getNoticeKey(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
        })
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    //搜索
    $scope.search = function () {
        GetAllEmployee();
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
