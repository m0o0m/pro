
angular.module('app.site').controller('maintenanceCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){

    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site_index_id,
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        };

        siteService.mainteaceList(postData).then(function (response) {
            console.log(response);
            $scope.arr = response.data.arr;
            console.log(response.data.meta.count);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.list;
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


    $scope.setting=function(){
        $scope.is_show=true;
    };
    
    $scope.mainTain = function (id) {
        var postData = {
            id:id,
            site_index_id:$scope.site_index_id
        };
        siteService.mainTain(postData).then(function (response) {
            console.log(response);
            $scope.check = response.data.checkbox;
            console.log($scope.check);
        });
    };

    $scope.modelone = function () {
        $scope.mainTain();
    };

    $scope.modeloness = function () {
        $scope.mainTain();
    };


});

