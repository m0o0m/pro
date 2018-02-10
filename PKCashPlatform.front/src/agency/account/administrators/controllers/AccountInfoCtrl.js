angular.module('app.administrators').controller('AccountInfoCtrl',
    function (AccountService, popupSvc, $scope, $rootScope, $stateParams, $state) {
        //获取当前id 数据
        $scope.ids = $stateParams.ids;
        AccountService.getMembersDetail($scope.ids).then(function (response) {
            console.log(response);
            $scope.detail = response;
            var myDate = new Date(response.birthday * 1000);
            var year = myDate.getFullYear().toString();
            var month = myDate.getMonth() + 1;
            var day = myDate.getDate();
            if (month < 10) {
                month = '0' + month;
            }
            if (day < 10) {
                day = '0' + day;
            }
            $scope.detail.birthday = year + '-' + month + '-' + day;
        });
        //修改后提交
        $scope.submit = function () {
            $scope.detail.id = $scope.ids;
            AccountService.putMembersDetail($scope.detail).then(function (response) {
                console.log(response);
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    $state.go('app.administrators.accounts');
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
    });