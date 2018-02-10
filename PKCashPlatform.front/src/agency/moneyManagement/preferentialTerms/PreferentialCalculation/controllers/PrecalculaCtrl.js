angular.module('app.Precalcula').controller('PrecalculaCtrl',
    function ($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取返点层级
        precalculaService.getLevel().then(function (response) {
            $scope.levelDrop = response.data;
        });

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                level: $scope.level
            };
            precalculaService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.arr;
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
        // 删除
        $scope.del = function (id) {
            var fn=function() {
                precalculaService.del({
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
            popupSvc.smartMessageBox($rootScope.getWord("confirmStorage")+"？", fn);
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
