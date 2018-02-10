angular.module('app.ReportForm').controller('MemberdetailsCtrl',
    function(httpSvc,popupSvc,$scope,financeReportService,APP_CONFIG,$stateParams){
        $scope.id = $stateParams.id;
        financeReportService.membershipReport($scope.id).then(function (res) {
            $scope.list2 = res.data;
        });

    });
