angular.module('app.CopyTemplate').controller('AddRegisterCopyCtrl', function($scope,copyTemplateService,$state,httpSvc,APP_CONFIG,popupSvc,$rootScope){

    $scope.registerJson = APP_CONFIG.option.option_register_type;
    $scope.typeId = '1';


    //提交获取富文本内容
    $scope.submit = function () {
        var summernote = $('.summernote').code();
        console.log(summernote);
        var postData = {
            title:$scope.title,
            type:$scope.typeId,
            content:summernote
        };
        copyTemplateService.addRegister(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }

});