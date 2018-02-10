angular.module('app.ReportForm').controller('hierarchicalCtrl',
    function ($scope, popupSvc, commonService, financeHierarchicalService, $rootScope, APP_CONFIG) {
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
        //获取层级下拉
        financeHierarchicalService.getDrop().then(function (response) {
            $scope.drop = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                siteId: $scope.siteId
            };
            financeHierarchicalService.getHierarchi(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.platforms;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee()
        };
        $scope.open = function (id) {
            $scope.id = id
        };
        $scope.submit = function () {
            financeHierarchicalService.modifyHierarchi({
                id: $scope.id,
                hierarchy: $scope.hierarchy
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }


    });
