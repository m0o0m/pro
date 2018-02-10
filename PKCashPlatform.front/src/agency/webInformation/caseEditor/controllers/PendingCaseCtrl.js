angular.module('app.caseEditor').controller('PendingCaseCtrl', function($scope,APP_CONFIG,$LocalStorage,CaseEditorService,popupSvc,$rootScope){
        var GetAllEmployee = function () {
            CaseEditorService.getPendingCase().then(function (response) {
               console.log(response.data.list);
                $scope.list = response.data.list;
            })
        }
        GetAllEmployee();

        //发送审核
        $scope.send = function (id) {
            var sendAudit = function () {
                var postData = {
                    id:id
                }
                CaseEditorService.getPendingCaseSend(postData).then(function (response) {
                    console.log(response);
                    if (response.data.data===null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sendAudit);

        };
        //删除
        $scope.delete=function (id) {
            var del = function () {
                var postData = {
                    id:id
                }
                CaseEditorService.getPendingCaseDel(postData).then(function (response) {
                    console.log(response);
                    if (response.data.data===null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));
                        GetAllEmployee();
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
        };

    });