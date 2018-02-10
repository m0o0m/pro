angular.module('app.analysisExit').controller('AnalysisExitCtrl',
    function($scope,APP_CONFIG,$LocalStorage,AnalysisExitService){

    $scope.sitId = function (site_index_id) {
        AnalysisExitService.getSiteSelect(site_index_id).then(function (response) {
            $scope.sharedJson = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        console.log($scope.siteId);
        $scope.sitId(user.site_index_id);
    }

    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        if($scope.strat_time===null){
            $scope.strat_time="";
        }
        if($scope.end_time===null){
            $scope.end_time="";
        }
        if($scope.type===null){
            $scope.type="";
        }
        if($scope.type_account===null){
            $scope.type_account="";
        }
        if($scope.account===null){
            $scope.account="";
        }
        if($scope.order_by===null){
            $scope.order_by="";
        }
        var postData = {
            site_index_id:$scope.site_index_id,
            start_time:$scope.start_time,
            end_time:$scope.end_time,
            type:$scope.type,
            type_account:$scope.type_account,
            account:$scope.account,
            order_by:$scope.order_by
        };
        AnalysisExitService.getAnalysisExit(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
            $scope.total = response.total;
            $scope.subtotal = response.subtotal;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
});


