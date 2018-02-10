angular.module('app.copyEditor').controller('RegisterEditorCtrl', function($scope, popupSvc, siteService, attachmentService, $rootScope, APP_CONFIG,CopyEditorService,$stateParams) {
    $scope.id = $stateParams.id;
    $scope.id = $stateParams.id;
    var postData = {
        id: $scope.id
    }
    CopyEditorService.getRegisterEditor(postData).then(function (response) {
        console.log(response);
        $scope.list = response.data.content;
        console.log($scope.list);
        $('.summernote').code($scope.list);
    })
    $scope.submit = function () {
        var summernote = $('.summernote').code();
        console.log(summernote);
        var postData = {
            id:$scope.id,
            content:summernote
        }
        CopyEditorService.getDepositEditorSub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }

});