angular.module('app.GraphicEditor').controller('AttachmentCtrl',
    function ($scope, popupSvc, attachmentService, $rootScope, APP_CONFIG) {
        var GetAllEmployee = function () {
            attachmentService.getList({

            }).then(function (response) {
                $scope.enclosure = response.data;
            });

        };
        GetAllEmployee();
        $scope.modifyTitle = function (item) {
            attachmentService.modify({
                id: item.id,
                title: item.title
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (item) {
            attachmentService.del({
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


    });