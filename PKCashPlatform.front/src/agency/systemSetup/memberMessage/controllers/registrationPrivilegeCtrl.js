angular.module('app.memberMessage').controller('registrationPrivilegeCtrl',
function(httpSvc,$scope,APP_CONFIG,popupSvc,memberMessageService,$rootScope){
    memberMessageService.getDropSelect().then(function (response) {
        $scope.siteJson=response.data;
    });
    $scope.add = function () {
        var postData = {
            site: $scope.site,
            content_1: $scope.content_1,
            content_2: $scope.content_2
        };
        memberMessageService.postPreferencesNews(postData).then(function (response) {
            if (response===null){
                popupSvc.smallBox("success",$rootScope.getWord("success"));
            }else {
                popupSvc.smallBox("fail",response.msg);
            }
        })
    }
});