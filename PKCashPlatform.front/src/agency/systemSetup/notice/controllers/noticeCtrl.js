angular.module('app.notice').controller('noticeCtrl',
function($scope,noticeService){
    noticeService.setSystermInformation().then(function (response) {
        $scope.data = response.data.data;
    })
});

