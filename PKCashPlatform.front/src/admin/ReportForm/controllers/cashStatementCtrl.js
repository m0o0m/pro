angular.module('app.ReportForm').controller('cashStatementCtrl',
    function ($scope, popupSvc, commonService, financeCashService, $rootScope, APP_CONFIG) {
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
        //获取方式下拉
        financeCashService.getWay().then(function (response) {
            $scope.wayOption = response.data;
        });

        $scope.refreshOption=APP_CONFIG.option.refresh_time1;
        $scope.sourceOption=APP_CONFIG.option.option_source;
        $scope.amountOption=APP_CONFIG.option.option_amount;

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                status: $scope.status,
                key: $scope.key
            };
            financeCashService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };

        $scope.selectAll = function ($event) {
            var target = $event.target;
            var isCheck = $(target).is(':checked');
            var list = $("tbody").find('input[type="checkbox"]');
            for (var i = 0; i < list.length; i++) {
                if (isCheck) {
                    $(list[i]).prop("checked", true);
                } else {
                    $(list[i]).prop("checked", false);
                }
            }
        };

    });

