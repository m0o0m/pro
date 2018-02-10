angular.module('app.analysisExit').controller('MemberInquiriesCtrl', function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG,$state){
    var GetAllEmployee = function () {
        var postData = {
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        }

        httpSvc.get("/report.json", postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data[0].audit.length;
            $scope.bank = response.data[0].audit;
            console.log($scope.bank);
        })
    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 20
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

});