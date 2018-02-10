angular.module('app.auth').controller('modifyPwdCtrl',
    function ($scope,$rootScope,httpSvc,popupSvc) {
        $scope.modify = function () {
            httpSvc.post("/password",{
                old_password: $scope.old_password,
                new_password: $scope.new_password,
                reply_password: $scope.reply_password
            }).then(function (response) {
                if(response===null){
                    $rootScope.logout();
                    popupSvc.smallBox("success","修改成功，请重新登录！");
                }else{
                    popupSvc.smallBox("success", response.msg);
                }
            })
        }
    });
