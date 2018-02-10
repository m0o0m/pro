angular.module('app.GraphicEditor').controller('FloatGoCtrl',
    function ($scope, popupSvc, floatManagrmentService, $stateParams, $rootScope) {
        $scope.float_id = $stateParams.float_id;

        var GetAllEmployee = function () {
            floatManagrmentService.getList({
                id: $stateParams.id
            }).then(function (response) {
                $scope.list = response.data;
            });
        };
        GetAllEmployee();

        $scope.deleteImg = function (id) {
            var del = function () {
                floatManagrmentService.deleteImg({
                    id: id
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox("确定删除？", del);
        };
        $scope.open = function (item) {
            $scope.formData = item;
        };
        $scope.modifyImg = function () {
            floatManagrmentService.modifyImg($scope.formData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.addImg = function () {
            floatManagrmentService.addImg($scope.addData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.upload = function (id) {
            $scope.id = id;
            floatManagrmentService.getEnclosure({
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
            floatManagrmentService.modifyEnclosure({
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
            floatManagrmentService.deleteEnclosure({
                id: item.id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.submit = function () {
            floatManagrmentService.selectEnclosure({
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
