angular.module('app.ReportForm').controller('DataCenterCtrl',
    function ($scope, popupSvc, commonService, financeDataCenterService, $rootScope, APP_CONFIG) {
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
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取类型下拉
        financeDataCenterService.getWay().then(function (response) {
            $scope.typeOption = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                type: $scope.type,
                timeZone: $scope.timeZone,
                memberAccount: $scope.memberAccount,
                agentAccount: $scope.agentAccount
            };
            financeDataCenterService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.subtotal = response.data.subtotal;
                $scope.total = response.data.total;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };

    });

