angular.module('app.websiteManagement').controller('modularManagementCtrl',
    function ($scope, popupSvc, modularManagementService, $rootScope, APP_CONFIG) {
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
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
        modularManagementService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                site: $scope.site,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            modularManagementService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });

        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };


        $scope.open = function (item) {
            $scope.order = item.order;
            $scope.new_order = item.order;
            $scope.id = item.id;
            $scope.name = item.name;
        };
        $scope.modify = function () {
            modularManagementService.modify({
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