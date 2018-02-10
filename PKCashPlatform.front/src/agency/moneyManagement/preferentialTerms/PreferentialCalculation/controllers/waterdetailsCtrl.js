angular.module('app.Precalcula').controller('waterdetailsCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state,httpSvc,$stateParams){

    var GetAllEmployee = function () {

        precalculaService.retreatWaterDetail($stateParams.Id).then(function (response) {
            $scope.paginationConf.totalItems = response.list.length;
            $scope.list = response.list;
            $scope.arr = response.arr;
            $scope.toal = response.Total;
        });

    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


});