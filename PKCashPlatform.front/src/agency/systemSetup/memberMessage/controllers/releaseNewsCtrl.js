angular.module('app.memberMessage').controller('releaseNewsCtrl',
function(httpSvc,$scope,APP_CONFIG,memberMessageService,popupSvc,$rootScope){
    memberMessageService.getDropSelect().then(function (response) {
        $scope.siteJson=response.data;
    });
    memberMessageService.getSystemSelect().then(function (response) {
        $scope.typeDrop=response.data;
    });
    $scope.add = function () {
        var postData = {
            site: $scope.site,
            system: $scope.system,
            title: $scope.title,
            content: $scope.content
        };
        memberMessageService.postMemberNews(postData).then(function (response) {
            if (response===null){
                popupSvc.smallBox("success",$rootScope.getWord("success"));
            }else {
                popupSvc.smallBox("fail",response.msg);
            }
        })
    }
});