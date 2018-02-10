angular.module('app.ReportForm').controller('quotaRecordCtrl',
    function ($scope, popupSvc, commonService, financeQuotaRecordService, $rootScope, APP_CONFIG) {
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

        financeQuotaRecordService.getTransactionType().then(function (response) {
            $scope.json1 = response.data;
        });
        financeQuotaRecordService.getVideoType().then(function (response) {
            $scope.json2 = response.data;
        });
        financeQuotaRecordService.getTransactionCategory().then(function (response) {
            $scope.json3 = response.data;
        });

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                site: $scope.site,
                account: $scope.account,
                vedioType: $scope.vedioType,
                transaction: $scope.transaction,
                transactionType: $scope.transactionType
            };
            financeQuotaRecordService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };

    });

