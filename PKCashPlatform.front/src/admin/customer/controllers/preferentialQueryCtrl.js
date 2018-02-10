angular.module('app.customer').controller('preferentialQueryCtrl',
    function ($scope, popupSvc, commonService, customerPreferentialQueryService, $rootScope, APP_CONFIG) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                id: $scope.id
            };

            customerPreferentialQueryService.getBullet(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });

        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search=function(){
            GetAllEmployee()
        };

    });