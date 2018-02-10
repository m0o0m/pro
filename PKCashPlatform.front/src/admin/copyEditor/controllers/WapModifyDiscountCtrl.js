angular.module('app.copyEditor').controller('WapModifyDiscountCtrl', function($compile,$scope, popupSvc, siteService, attachmentService, $rootScope, APP_CONFIG,CopyEditorService,$stateParams) {
    $scope.id = $stateParams.id;
    var postData = {
        id: $scope.id
    }
    CopyEditorService.getWapDiscount_M_C(postData).then(function (response) {
        console.log(response);
        $scope.list = response.data.content;
        console.log($scope.list);
        $('.summernote').code($scope.list);
    })
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
        var summernote = $('.summernote').code();
        console.log(summernote);
        var postData = {
            id:$scope.id,
            content:summernote
        }
        CopyEditorService.getWapDiscount_M_Sub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }

});
