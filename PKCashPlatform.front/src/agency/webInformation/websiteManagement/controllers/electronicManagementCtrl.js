angular.module('app.websiteManagement').controller('electronicManagementCtrl',
    function ($scope, popupSvc, electronicManagementService, $rootScope, APP_CONFIG, $LocalStorage) {
        $scope.user = JSON.parse($LocalStorage.getItem("user"));
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
        electronicManagementService.getTheme().then(function (response) {
            $scope.theme = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            electronicManagementService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.themeModify = function () {
            electronicManagementService.modifyTheme($scope.theme).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.initialization = function () {
            var sure = function () {
                electronicManagementService.initialization({
                    site_index_id: $scope.user.site_index_id
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });

            }
            popupSvc.smartMessageBox($rootScope.getWord("是否初始化原来版本?"), sure);
        };

        $scope.open = function (item) {
            $scope.order = item.order;
            $scope.new_order = item.order;
            $scope.id = item.id;
            $scope.name = item.name;
        };
        $scope.modify = function () {
            electronicManagementService.modifyOrder({
                order: $scope.order,
                id: $scope.id,
                new_order: $scope.new_order
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

    });