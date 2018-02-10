angular.module('app.AccessMoney').controller('HistoryCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope,$interval){
        
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
    //获取站点下拉框
    AccessMoneyService.getDropSelect().then(function (response) {
        $scope.siteJson = response.data;
        console.log($scope.site);
    });

    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            start_time: $scope.start_time,
            end_time: $scope.end_time,
            account: $scope.account,
            site_index_id: $scope.site_index_id
        };
          AccessMoneyService.manualAccessRecord(postData).then(function (response) {
                  $scope.paginationConf.totalItems = response.data[0].total_count;
                  $scope.data = response.data[0];
                  $scope.list = response.data[0].member_cash_record;

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
});