angular.module('app.customer').controller('configurationsettingCtrl',
    function ($scope, popupSvc, commonService, configurationsettingService, $rootScope, APP_CONFIG) {
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取弹框广告信息
        configurationsettingService.getAd().then(function (response) {
            $scope.adData = response.data;
        });
        //搜索
        $scope.search = function () {
            GetAllEmployee()
        };

        var GetAllEmployee = function () {
            configurationsettingService.getList({
                site: $scope.site
            }).then(function (response) {
                $scope.list = response.data;
            });
        };
        GetAllEmployee();


        $scope.modify = function (item) {
            $scope.data = item;
        };

        $scope.modifySubmit = function () {
            configurationsettingService.modify($scope.data).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (id) {
            configurationsettingService.del({
                id: id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.modifyAd = function () {
            configurationsettingService.mdifyAd($scope.adData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

    });