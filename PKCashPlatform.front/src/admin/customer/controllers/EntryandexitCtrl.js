angular.module('app.customer').controller('EntryandexitCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,customerService) {
    $scope.json = APP_CONFIG.option;
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
    $scope.startip = "1";
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id:$scope.site_index_id,
            account:$scope.account,
            order_num:$scope.orderNum,
            in_type:$scope.type,
            start_time:$scope.startTime,
            end_time:$scope.endTime,
            site:$scope.site
        };
        customerService.getBankIn(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.listIn = response.data;
        })

    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    var GetAllEmploye = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id:$scope.site_index_id_O,
            account:$scope.account_O,
            start_time:$scope.startTime_O,
            end_time:$scope.endTime_O,
            site:$scope.site_O
        };
        customerService.getBankOut(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.listOut = response.data;
        })

    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmploye);
    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.searchOut = function () {
        GetAllEmploye();
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