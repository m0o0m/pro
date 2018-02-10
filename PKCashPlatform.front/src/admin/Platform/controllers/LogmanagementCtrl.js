angular.module('app.Platform').controller('LogmanagementCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        if($scope.account==undefined){
            $scope.account = ""
        }
        if($scope.ip==undefined){
            $scope.ip = ""
        }
        if($scope.logInfo==undefined){
            $scope.logInfo = ""
        }
        if($scope.type==undefined){
            $scope.type = ""
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
            logInfo: $scope.logInfo,
            type: $scope.type,
            stratTime: $scope.stratTime,
            endTime: $scope.endTime
        };
        PlatformService.getLog(postData).then(function (response) {
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
