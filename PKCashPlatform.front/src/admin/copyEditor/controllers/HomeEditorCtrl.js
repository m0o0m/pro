angular.module('app.copyEditor').controller('HomeEditorCtrl', function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$stateParams,$state){
    $scope.id = $stateParams.id;
    console.log($scope.id);
    var postData = {
        id: $scope.id
    }
    CopyEditorService.getHomeEditor(postData).then(function (response) {
        console.log(response);
        $scope.list = response.data.content;
        console.log($scope.list);
        $('.summernote').code($scope.list);
    })
    $scope.sub = function () {
        var summernote = $('.summernote').code();
        console.log(summernote);
        var postData = {
            id:$scope.id,
            content:summernote
        }
        CopyEditorService.getHomeEditorSub(postData).then(function (response) {
            console.log(response);
            if(response.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                $state.go("app.CopyEditor.HomeCopy");
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };
});