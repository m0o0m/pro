angular.module('app.GraphicEditor').controller('UploadCtrl',
    function ($scope, popupSvc, attachmentService, $rootScope) {

        $scope.upload = function () {
            var fd = new FormData();
            var file =$("#file").get(0).files[0];
            fd.append('file', file);
            fd.append('fileName', $scope.name);
            attachmentService.upload(fd).then(function (response) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                //if (response === null) {
                //    popupSvc.smallBox("success", $rootScope.getWord("success"));
                //} else {
                //    popupSvc.smallBox("fail", response.msg);
                //}
            });
        }

    });