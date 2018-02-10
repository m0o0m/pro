angular.module('app.GraphicEditor').controller('ModifyNoticesCtrl',
    function ($compile, $scope, popupSvc, siteService, attachmentService, announcementService, $rootScope, APP_CONFIG, $stateParams) {

        var GetAllEmployee = function () {
            announcementService.getContent({
                id: $stateParams.id
            }).then(function (response) {
                var html = response.data.content;
                var content = $compile(html)($scope);
                $('.summernote').code(content);
            });
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

        $scope.select = function (imgUrl) {
            var content = $('.summernote').code();
            var img = "<img src='"+imgUrl+"'>";
            content = "<span>"+content+img+"</span>";
            content = $compile(content)($scope);
            $('.summernote').code(content);
        };

        //提交获取富文本内容
        $scope.submit = function () {
            var content=$('.summernote').code();
            announcementService.editContent({
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

