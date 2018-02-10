angular.module('app.CopyTemplate').controller('ModifyRegisterCopyCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$stateParams,copyTemplateService){
    $scope.registerJson = APP_CONFIG.option.option_register_type;
    $scope.id = $stateParams.id;
    var postData = {
        id: $scope.id
    };
    copyTemplateService.getAddRegister(postData).then(function (response) {
        console.log(response);
        $scope.title = response.data.data.title;
        $scope.type = response.data.data.type;
        $scope.typeId = $scope.type.toString();
        console.log($scope.typeId);
        $scope.list = response.data.data.content;
        $('.summernote').code($scope.list);
    });

    //提交获取富文本内容
    $scope.submit = function () {
        var summernote = $('.summernote').code();
        console.log(summernote);
        var postData = {
            title:$scope.title,
            type:$scope.typeId,
            content:summernote
        };
        copyTemplateService.putAddRegister(postData).then(function (response) {
            console.log(response);
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }

});