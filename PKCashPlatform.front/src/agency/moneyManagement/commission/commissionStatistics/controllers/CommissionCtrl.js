angular.module('app.CommissionStatistics').controller('CommissionCtrl',
    function ($scope, popupSvc, siteService, commissionService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id,
                period: $scope.period,
                isGet: $scope.isGet,
                member_min: $scope.member_min,
                member_max: $scope.member_max
            };
            commissionService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.arr;
            });
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search=function(){
            GetAllEmployee();
        }

    });

