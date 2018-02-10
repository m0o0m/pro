angular.module('app.ReportForm').controller('SummaryCtrl',
    function ($scope, popupSvc, commonService, financeSummaryService, $rootScope, APP_CONFIG) {
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
        var GetAllEmployee = function () {
            var postData = {
                site: $scope.site,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                section: $scope.section
            };
            financeSummaryService.getData(postData).then(function (response) {
                $scope.income = response.data.income;
                $scope.expenditure = response.data.expenditure;
                $scope.total = response.data.total;
                $scope.actual = response.data.actual;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };

    });
