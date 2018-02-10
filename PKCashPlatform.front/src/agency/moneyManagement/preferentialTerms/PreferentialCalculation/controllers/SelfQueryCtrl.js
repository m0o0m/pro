angular.module('app.Precalcula').controller('SelfQueryCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state,httpSvc){

        //获取站点
        siteService.getSite().then(function (response) {
            $scope.names = response.data;
                console.log(response);
                $scope.names = response.data;
                $scope.site_length = $scope.names.length;
                console.log($scope.site_length);
                $scope.siteid = $scope.site_index_id;
                console.log($scope.siteid);
                if ($scope.site_index_id == undefined) {
                    $scope.site_index_id = "";
                }
        });

    var GetAllEmployee = function () {

        if($scope.start_time == undefined){
            $scope.start_time = "";
        }
        if($scope.end_time == undefined){
            $scope.end_time = "";
        }
        if($scope.site_index_id == undefined){
            $scope.site_index_id = "";
        }
        if($scope.order_num == undefined){
            $scope.order_num = "";
        }
        if($scope.account == undefined){
            $scope.account = "";
        }
        var postData = {
            site_index_id: $scope.site_index_id,
            order_num: $scope.order_num,
            account: $scope.account,
            start_time: $scope.start_time,
            end_time: $scope.end_time
        };
        precalculaService.retreatWaterSelfSearch(postData).then(function (response) {
                console.log(response);
                $scope.list = response.data;
                $scope.paginationConf.totalItems = response.meta.count;
        });

    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.search = function () {
        GetAllEmployee();
    }
    //明细
    // $scope.detail = function (id,start,end) {
    //     console.log(start);
    //     console.log(end);
    //     $state.go('app.Precalcula.waterdetails',{
    //         id:id,
    //         start:start,
    //         end:end
    //     })
    // }

    //筛选展开
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
