angular.module('app.GraphicEditor').controller('SwiperCtrl',
    function ($scope, popupSvc, siteService, swiperService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //搜索
        $scope.search = function () {
            GetAllEmployee()
        };
        var GetAllEmployee = function () {
            swiperService.getList({
                site: $scope.site
            }).then(function (response) {
                $scope.list = response.data;
            });
        };
        GetAllEmployee();
        $scope.see = function (img) {
            $scope.img = img;
        };

        $scope.disable=function(id, status){
            var fn = function () {
                swiperService.disable({
                    id: id,
                    status: status === 1 ? 2 : 1
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            if (status === 1) {
                popupSvc.smartMessageBox($rootScope.getWord("confirmEnable")+"？", fn);
            } else {
                popupSvc.smartMessageBox($rootScope.getWord("confirmDisable")+"？", fn);
            }
        };

        $scope.storage = function (type) {
            var fn = function () {
                swiperService.storage({
                    type: type
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmStorage")+"？", fn);
        };

        $scope.modify = function (id) {
            $scope.id = id;
            swiperService.getEnclosure({
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
            swiperService.modifyEnclosure({
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
            swiperService.deleteEnclosure({
                id: item.id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    swiperService.getEnclosure({
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
            swiperService.selectEnclosure({
                id: $scope.id,
                enclosure_id: $scope.enclosure_id,
                link: $scope.link
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