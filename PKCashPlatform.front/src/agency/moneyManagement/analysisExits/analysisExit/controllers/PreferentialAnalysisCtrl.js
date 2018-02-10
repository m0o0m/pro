angular.module('app.analysisExit').controller('PreferentialAnalysisCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,AnalysisExitService){
    $scope.sitId = function (site_index_id) {
        AnalysisExitService.getSiteSelect(site_index_id).then(function (response) {
            $scope.names = response.data.data;
        });
    };
    var user=JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin=user.site_index_id==='';
    if($scope.isSuperAdmin){
        $scope.site_index_id=user.default_site;
        //获取站点
        $scope.sitId();
    }else{
        $scope.site_index_id=user.site_index_id;
    }

    var GetAllEmployee = function () {
        if($scope.start_time===null){
            $scope.start_time="";
        }
        if($scope.end_time===null){
            $scope.end_time="";
        }
        if($scope.account===null){
            $scope.account="";
        }
        var postData = {
            site_index_id:$scope.site_index_id,
            start_time:$scope.start_time,
            end_time:$scope.end_time,
            account:$scope.account
        };
        AnalysisExitService.getPreferentialAnalysis(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.list;
            $scope.name = response.product_name;
            $scope.length = $scope.name.length+3;
            console.log($scope.length);
            $scope.total=response.total;
            $scope.subtotal=response.subtotal;
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


