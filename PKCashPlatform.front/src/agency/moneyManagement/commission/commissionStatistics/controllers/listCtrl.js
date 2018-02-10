angular.module('app.CommissionStatistics').controller('listCtrl',
    function ($scope, popupSvc, commissionService, $rootScope, APP_CONFIG) {
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            commissionService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.arr;
            });
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


    });

