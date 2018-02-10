angular.module('app.customer').controller('LogCtrl', function(httpSvc,popupSvc,$scope,APP_CONFIG,customerService,$rootScope,$LocalStorage){

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
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id:$scope.site_index_id,
            account:$scope.account,
            start_time:$scope.startTime,
            end_time:$scope.endTime
        };
        customerService.getAuditRecord(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 10
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    var GetAllEmploye = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id:$scope.site_index_ID,
            account:$scope.accounts,
            start_time:$scope.startT,
            end_time:$scope.endT
        };
        customerService.getAuditLog(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.listL = response.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 10
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmploye);
    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.searchL = function () {
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

