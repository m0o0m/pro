angular.module('app.websiteManagement').controller('PreferencesSettingsCtrl',
    function ($scope, popupSvc, preferencesSettingsService, $rootScope, APP_CONFIG) {
        preferencesSettingsService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
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
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            preferencesSettingsService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });

        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


        $scope.InitializationSub = function (id) {
            var sure = function () {
                preferencesSettingsService.modify({
                    id: id
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("是否置为无效?"), sure);
        };

        $scope.submit = function () {
            preferencesSettingsService.add($scope.formData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
    });