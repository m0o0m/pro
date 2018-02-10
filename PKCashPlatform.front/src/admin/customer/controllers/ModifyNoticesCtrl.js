angular.module('app.customer').controller('ModifyNoticeAdCtrl',
    function ($compile, $scope, popupSvc, configurationsettingService, $rootScope, APP_CONFIG, $stateParams) {

        var GetAllEmployee = function () {
            configurationsettingService.getContent({
                id: $stateParams.id
            }).then(function (response) {
                var html = response.data.content;
                var content = $compile(html)($scope);
                $('.summernote').code(content);
            });
        };
        GetAllEmployee();

        //提交获取富文本内容
        $scope.submit = function () {
            var content=$('.summernote').code();
            configurationsettingService.editContent({
                id: $stateParams.id,
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

