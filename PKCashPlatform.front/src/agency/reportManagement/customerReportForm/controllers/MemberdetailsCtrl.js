angular.module('app.ReportForm').controller('MemberdetailsCtrl',
    function(httpSvc,popupSvc,$scope,reportformService,APP_CONFIG,$stateParams){
    $scope.id = $stateParams.id;
    console.log($scope.id);
    reportformService.membershipReport($scope.id).then(function (res) {
        $scope.list2 = res.data;
    });










})
