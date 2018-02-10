angular.module('app.Summary').controller('SummaryCtrl',
    function (httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope,$interval) {
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


        //获取入款商户
        AccessMoneyService.getAccountSelect().then(function (response) {
            $scope.agencyList = response.data.data;

        });

    $scope.getDate = function (AddDayCount) {
        var dd = new Date();
        dd.setDate(dd.getDate()+AddDayCount);//获取AddDayCount天后的日期
        var y = dd.getFullYear();
        var m = dd.getMonth()+1;//获取当前月份的日期
        m=m>=10?m:'0'+m;
        var d = dd.getDate();
        d=d>=10?d:'0'+d;
        $scope.date_time = y+"-"+m+"-"+d;
        GetAllEmployee();
    };

    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            site_index_id: $scope.site_index_id,
            start_time: $scope.start_time,
            end_time: $scope.end_time,
            account: $scope.account,
            agency_id: $scope.agency_id,
            date_time: $scope.date_time
        }
           AccessMoneyService.manualAccessCollect(postData).then(function (response) {
                   if(!response.code){
                       $scope.paginationConf.totalItems = response.data.total_count;
                       $scope.subtotal=response.data;
                       $scope.data = response.data;
                       console.log(response.data);
                   }else{
                       $scope.paginationConf.totalItems = 0;
                       $scope.subtotal = null;
                       $scope.list = null;
                   }
           });
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 10
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.search = function () {
        GetAllEmployee();
    };
});
