angular.module('app.customer').controller('AdvertisingcontentCtrl',
    function ($compile, $scope, popupSvc, configurationsettingService, $rootScope, APP_CONFIG) {

        //提交获取富文本内容
        $scope.submit = function () {
            var content = $('.summernote').code();
            configurationsettingService.add({
                content: content
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }

    });
