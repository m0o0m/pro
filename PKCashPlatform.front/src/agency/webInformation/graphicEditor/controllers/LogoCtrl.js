angular.module('app.GraphicEditor').controller('LogoCtrl',
    function ($scope, popupSvc, siteService, logoService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });

        var GetAllEmployee = function () {
            logoService.getList({
                site: $scope.site
            }).then(function (response) {
                $scope.list = response.data;
            });
        };
        GetAllEmployee();

        //搜索
        $scope.search = function () {
            GetAllEmployee();
        };

        //图片预览
        $scope.img = function (url) {
            $scope.img_src = url;
        };
        //修改LOGO
        $scope.modifyLogo = function (item) {
            $scope.modifyData = item;
        };
        $scope.modifyLogoSubmit = function (item) {
            logoService.modifyLogo({
                id: $scope.modifyData.id,
                copy: $scope.modifyData.copy,
                title: $scope.modifyData.title,
                status: $scope.modifyData.status
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.storage = function (id) {
            var del = function () {
                logoService.storage({
                    id: id
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmStorage")+"？", del);

        };

        $scope.upload = function (id) {
            $scope.id = id;
            logoService.getEnclosure({
                id: $scope.id
            }).then(function (response) {
                $scope.enclosure = response.data;
            });
        };

        $scope.select = function (item) {
            $scope.title = item.title;
            $scope.url = item.url;
            $scope.enclosure_id = item.id;
        };
        $scope.modifyTitle = function (item) {
            logoService.modifyEnclosure({
                id: item.id,
                title: item.title
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (item) {
            logoService.deleteEnclosure({
                id: item.id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    logoService.getEnclosure({
                        id: $scope.id
                    }).then(function (response) {
                        $scope.enclosure = response.data;
                    });
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.submit = function () {
            logoService.selectEnclosure({
                id: $scope.id,
                enclosure_id: $scope.enclosure_id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }


    });