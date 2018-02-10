/**
 * Created by apple on 17/9/5.
 */

angular.module('app.CommissionStatistics').controller('modifyCtrl', function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$stateParams){
    $scope.id = $stateParams.id;
    console.log($scope.id);
    var GetAllEmployee = function () {
        var postData = {
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        };

        httpSvc.get("/report.json", postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.length;
            $scope.commission = response.data[0].commission[$scope.id];
            console.log($scope.commission);
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 20
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


});

