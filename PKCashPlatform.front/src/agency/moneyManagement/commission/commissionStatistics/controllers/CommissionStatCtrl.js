angular.module('app.CommissionStatistics').controller('CommissionStatCtrl',
function($scope,$state,httpSvc,siteService){
    //获取站点
    siteService.getSite().then(function (response) {
        $scope.siteJson = response.data;
    });
    //httpSvc.get("/agent/first/drop").then(function (response) {
    //    $scope.siteJson=response.data;
    //});
    //httpSvc.get("/retirement/list").then(function (response) {
    //    $scope.siteJson=response.data;
    //});

});
