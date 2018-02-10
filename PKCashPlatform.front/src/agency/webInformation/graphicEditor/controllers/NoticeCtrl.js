angular.module('app.GraphicEditor').controller('NoticesCtrl',
    function ($scope, popupSvc, siteService, announcementService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取弹框广告信息
        announcementService.getAd().then(function (response) {
            $scope.adData = response.data;
        });
        //搜索
        $scope.search = function () {
            GetAllEmployee()
        };

        var GetAllEmployee = function () {
            announcementService.getList({
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
            announcementService.modify($scope.data).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (id) {
            announcementService.del({
                id: id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.modifyAd=function(){
            announcementService.mdifyAd($scope.adData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

    });