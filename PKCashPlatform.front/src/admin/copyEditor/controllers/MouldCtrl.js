angular.module('app.copyEditor').controller('MouldCtrl',  function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$stateParams){
    $scope.id = $stateParams.id;

    var GetAllEmployee = function () {
        var postData = {
            id: $scope.id
        }
        CopyEditorService.getDepositCopyModule(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data;

        })
    }
    GetAllEmployee();

    $scope.storage = function (id) {
        var del = function () {
            var postData = {
                id: id
            };
            CopyEditorService.getDepositCopyModule_k(postData).then(function (response) {
                if (response.data.data===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
    }
    $scope.choise = function (id) {
        var postData = {
            id: id
        };
        CopyEditorService.getDepositCopyModule_C(postData).then(function (response) {
            if (response.data.data===null){
                GetAllEmployee();
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }
});