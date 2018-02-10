
angular.module('app.Platform').controller('RoleCtrl',
    function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
        //新增
        $scope.addAcount = function () {
            PlatformService.getRoleAdd($scope.formData).then(function (response) {
                if (response.datd.data===null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            });
        };

        // 停用
        $scope.disable=function (id,status) {
            var sure = function () {
                var postData = {
                    id: id,
                    status: status
                };
                PlatformService.getRoleStatus(postData).then(function (response) {
                    if (response.data.data === null) {
                        GetAllEmployee();
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), sure);
        };
        // 删除
        $scope.delete = function (id) {
            var del = function () {
                var postData = {
                    id: id
                };
                PlatformService.getRoleDel(postData).then(function (response) {
                    if (response.data.data === null) {
                        GetAllEmployee();
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), del);
        };

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
            };
            PlatformService.getRole(postData).then(function (response) {
                console.log(response);
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.data;
            })
        };

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    });
