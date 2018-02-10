angular.module('app.websiteManagement').controller('cacheManagementCtrl',
    function ($scope, popupSvc, cacheManagementService, $rootScope, APP_CONFIG) {
        cacheManagementService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        cacheManagementService.getPage().then(function (response) {
            $scope.pageDrop = response.data;
        });
        $scope.submit = function () {
            cacheManagementService.submit({
                site: $scope.site,
                page: $scope.page
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
    });