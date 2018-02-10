angular.module('app.caseEditor').controller('AuditCaseCtrl',
    function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope){
        var GetAllEmployee = function () {
            CaseEditorService.getAuditCase().then(function (response) {
                console.log(response.data.list);
                $scope.list = response.data.list;
            })
        }
        GetAllEmployee();
        //删除
        $scope.delete=function (id) {
            var del = function () {
                var postData = {
                    id:id
                }
                CaseEditorService.getAuditCaseDel(postData).then(function (response) {
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