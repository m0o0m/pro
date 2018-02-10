
angular.module('app.site').controller('maintenanceItemCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        

    var GetAllEmployee = function () {
       siteService.mainTainProject().then(function (response) {
           $scope.paginationConf.totalItems = response.data.meta[0].count;
           $scope.list = response.data.data;
       });
        
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 10
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.del=function(){
        popupSvc.smallBox("success","删除成功");
    }

    $scope.setting=function(){
        $scope.is_show=true;
    }
});

