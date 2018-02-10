angular.module('app.CommissionStatistics').controller('FeeSettingCtrl',
    function ($scope, popupSvc, siteService, feeSettingService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id
            };

            feeSettingService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.list;
            });

        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };
        $scope.del = function (id) {
            var fn=function() {
                feeSettingService.del({
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
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation")+"？", fn);
        };


    });

