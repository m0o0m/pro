angular.module('app.Platform').controller('OperationlogCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
    var GetAllEmployee = function () {
        if($scope.account==undefined){
            $scope.account = ""
        }
        if($scope.ip==undefined){
            $scope.ip = ""
        }
        if($scope.url==undefined){
            $scope.url = ""
        }
        if($scope.stratTime==undefined){
            $scope.stratTime = ""
        }
        if($scope.endTime==undefined){
            $scope.endTime = ""
        }
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            account: $scope.account,
            ip: $scope.ip,
            url: $scope.url,
            stratTime: $scope.stratTime,
            endTime: $scope.endTime
        };
        PlatformService.getOperation(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta[0].count;
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
