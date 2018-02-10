angular.module('app.websiteManagement').controller('videoManagementCtrl',
    function ($scope, popupSvc, videoManagementService, $rootScope, APP_CONFIG) {
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
        videoManagementService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        videoManagementService.getType().then(function (response) {
            $scope.typeDrop = response.data;
        });
        videoManagementService.getStyle().then(function (response) {
            $scope.styleDrop = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                site: $scope.site,
                type: $scope.type,
                style: $scope.style,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            videoManagementService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        }

        //使用
        $scope.disables = function () {
            var sure = function () {
                videoManagementService.use().then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation?"), sure);
        };
        //还原老版本
        $scope.reduction = function () {
            var sure = function () {
                videoManagementService.back().then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation?"), sure);
        };

        $scope.open = function (item) {
            $scope.order = item.order;
            $scope.new_order = item.order;
            $scope.id = item.id;
            $scope.name = item.name;
        };
        $scope.modify = function () {
            videoManagementService.modify({
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