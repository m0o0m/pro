angular.module('app.websiteManagement').controller('WebsiteInformationCtrl',
    function ($scope, popupSvc, websiteInformationService, $rootScope) {

        websiteInformationService.getColorSelect().then(function (response) {
            $scope.colorDrop = response.data;
        });
        websiteInformationService.getInformation().then(function (response) {
            $scope.information=angular.copy(response.data);
        });
        $scope.modify = function () {
            websiteInformationService.modifyInformation($scope.information).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
    });