angular.module('app.copyEditor').controller('HomeCopyCtrl', function($scope,APP_CONFIG,$LocalStorage,CopyEditorService,popupSvc,$rootScope,$state){
    $scope.sitId = function (site_index_id) {
        CopyEditorService.getSiteSelect(site_index_id).then(function (response) {
            $scope.sharedJson = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        $scope.sitId(user.site_index_id);
    }

    var GetAllEmployee = function () {
        var postData = {
            site_index_id: $scope.site_index_id
        };
        CopyEditorService.getHomeCopy(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.list;
        })
    };
   GetAllEmployee();
    $scope.modify = function (id,title,is_use,is_login,status,order,type,code) {
        $scope.ids = id;
        $scope.title = title;
        $scope.form = is_use;
        $scope.login = is_login;
        $scope.status = status;
        $scope.num = order;
        $scope.code = code;
        var login = document.getElementsByName('login');
        var form = document.getElementsByName('form');
        var status_1 = document.getElementsByName('status');
        for(var i = 0;i < 2;i++) {
            if (login[i].value == is_login) {
                login[i].checked = 'checked';
            }
            if (form[i].value == is_use) {
                form[i].checked = 'checked';
            }
            if (status_1[i].value == status) {
                status_1[i].checked = 'checked';
            }
        }
    };

    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.storage = function (id) {
        var del = function () {
            var postData = {
                id: id
            };
            CopyEditorService.getHomeCopyKeep(postData).then(function (response) {
                if (response.data.data===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);

    };

    $scope.sub = function () {
        $scope.form_1 = $("input[name='form']:checked").val();
        $scope.login_1 = $("input[name='login']:checked").val();
        $scope.status_1 = $("input[name='status']:checked").val();
        var postData = {
            id: $scope.ids,
            title: $scope.title,
            orderBuy: $scope.num,
            code: $scope.code,
            isUse: $scope.form_1,
            islogin: $scope.login_1,
            status: $scope.status_1
        };
        CopyEditorService.getHomeCopySub(postData).then(function (response) {
            console.log(response);
            if(response.data.data==null){
                GetAllEmployee();
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    $scope.modifyContent = function (id) {
        $state.go("app.CopyEditor.HomeEditor",{
            id:id
        })
    }
});