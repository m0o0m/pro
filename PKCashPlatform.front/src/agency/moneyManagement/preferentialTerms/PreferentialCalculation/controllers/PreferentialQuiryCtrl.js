angular.module('app.Precalcula').controller('PreferentialQuiryCtrl',
    function ($scope, popupSvc, siteService, preferentialQuiryService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });

        var GetAllEmployee = function () {
            $scope.year = new Date().getFullYear();
            if ($scope.yy == undefined) {
                $scope.yy = "";
            }
            if ($scope.mm == undefined) {
                $scope.mm = "";
            }
            if ($scope.site_index_id == undefined) {
                $scope.site_index_id = "";
            }
            var postData = {
                year: $scope.yy,
                month: $scope.mm,
                site_index_id: $scope.site_index_id,
            };
            preferentialQuiryService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
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

        //筛选展开
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
