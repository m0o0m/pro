angular.module('app.analysisExit').controller('PurchaseAnalysisCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,AnalysisExitService){
    $scope.sitId = function (site_index_id) {
        AnalysisExitService.getSiteSelect(site_index_id).then(function (response) {
            $scope.names = response.data.data;
        });
    };
    var user=JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin=user.site_index_id==='';
    if($scope.isSuperAdmin){
        //获取站点
        $scope.sitId();
    }else{
        $scope.site_index_id=user.site_index_id;
    }
    AnalysisExitService.getTypeSelect().then(function (response) {
       console.log(response.data.data);
       $scope.type_select = response.data.data;
    });
    AnalysisExitService.getAgencySelect().then(function (response) {
        console.log(response);
        $scope.agency_select = response.data.data;

    });
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        console.log($scope.type);
        if($scope.start_time===null){
            $scope.start_time="";
        }
        if($scope.end_time===null){
            $scope.end_time="";
        }
        if($scope.agency===null){
            $scope.agency="";
        }
        if($scope.type===null){
            $scope.type="";
        }
        if($scope.commodity===null){
            $scope.commodity="";
        }
        if($scope.time===null){
            $scope.time="";
        }
        var postData = {
            site_index_id:$scope.site_index_id,
            start_time:$scope.start_time,
            end_time:$scope.end_time,
            agency:$scope.agency,
            type:$scope.type,
            commodity:$scope.commodity,
            time:$scope.time
        };
        AnalysisExitService.getPurchaseAnalysis(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
            $scope.total = response.total[0];
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


