angular.module('app.copyEditor').controller('DepositCopyCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService,CopyEditorService,$state){

    var GetAllEmployee = function () {
        var postData = {
            site_index_id: $scope.site_index_id
        }
        CopyEditorService.getDepositCopy(postData).then(function (response) {
            console.log(response)
            $scope.list = response.data.list;
        });
    }
  GetAllEmployee();

    $scope.search = function () {
        GetAllEmployee();
    }
    $scope.modify = function (id,title,code,status,color) {
        $scope.ids = id;
        $scope.tit = title;
        $scope.code = code;
        $scope.color = color;
        $scope.status = status;
        $scope.code = code;
        var status_1 = document.getElementsByName('status');
        for(var i = 0;i < 2;i++) {
            if (status_1[i].value == status) {
                status_1[i].checked = 'checked';
            }
        }
    }
    $scope.sub = function () {
        $scope.status_1 = $("input[name='status']:checked").val();
        var postData = {
            id: $scope.ids,
            title: $scope.tit,
            code: $scope.code,
            status: $scope.status_1,
            color: $scope.color
        }
        CopyEditorService.getDepositCopySub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    }

    $scope.mContent = function (id) {
        $state.go('app.CopyEditor.DepositEditor',{
            id:id
        })
    }
    $scope.module = function (id) {
        $state.go('app.CopyEditor.Mould',{
            id:id
        })
    }

    $scope.keep = function (id) {
        var del = function () {
            var postData = {
                id: id
            };
            CopyEditorService.getDepositCopyKeep(postData).then(function (response) {
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
});