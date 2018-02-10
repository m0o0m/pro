angular.module('app.AccessMoney').controller('QuotaRecordCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG){
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                start_time:$scope.start_time,
                end_time:$scope.end_time,
                account:$scope.account
            };

            AccessMoneyService.balanceConversion(postData).then(function (response) {
                    console.log(response);
                        $scope.paginationConf.totalItems = response.meta.count;
                        $scope.datalist = response.data;
            });


        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        //点击查询
        $scope.search = function () {
            GetAllEmployee();
        };
    });