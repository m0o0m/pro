angular.module('app.copyEditor').controller('discountEditorCtrl',
    function ($compile, $scope, popupSvc, siteService, attachmentService, $rootScope, APP_CONFIG,CopyEditorService) {

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

        //点击选择把图片放到富文本

        $scope.arr = [];
        $scope.select = function (img) {
            var content = $('.summernote').code();
            $scope.arr.push(img);
            var html = "";
            for (var i = 0; i < $scope.arr.length; i++) {
                html += "<img src='" + $scope.arr[i] + "'>";
            }
            var template = angular.element(html);
            var mobile = $compile(template)($scope);
            console.log(mobile);
            $('.summernote').code(mobile);

        }

        //点击查看跳转图片
        $scope.see_1 = function ($index) {
            var addCopy = document.getElementsByClassName('addCopy')[$index].placeholder;
            console.log(addCopy);
            $scope.seeurl = addCopy;
        }
        //提交获取富文本内容
        $scope.submit = function () {
            // console.log($('.summernote'));
            console.log($('.summernote').code());
            $('.summernote').code();

        }

    });
