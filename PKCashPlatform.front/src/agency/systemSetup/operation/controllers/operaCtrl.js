
angular.module('app.memberMessage').controller('operaCtrl',
function(httpSvc,$scope,APP_CONFIG,operationService){
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
            account: $scope.account,
            start_time: $scope.start_time,
            end_time: $scope.end_time
        };
        operationService.setSystemLog(postData).then(function (response) {
            console.log(response);
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

