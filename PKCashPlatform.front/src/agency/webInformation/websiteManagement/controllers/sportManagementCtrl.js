angular.module('app.websiteManagement').controller('sportManagementCtrl',
    function ($scope, popupSvc, sportManagementService, $rootScope, APP_CONFIG) {
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            sportManagementService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.open = function (item) {
            $scope.order = item.order;
            $scope.new_order = item.order;
            $scope.id = item.id;
            $scope.name = item.name;
        };
        $scope.modify = function () {
            sportManagementService.modify({
                order: $scope.order,
                id: $scope.id,
                new_order: $scope.new_order
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
    });