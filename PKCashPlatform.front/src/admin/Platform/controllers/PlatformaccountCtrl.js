angular.module('app.Platform').controller('PlatformAccountCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
    PlatformService.getRoleDrop().then(function (response) {
        $scope.selectJson = response.data;
    });
    $scope.json = APP_CONFIG.option;
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };

    var GetAllEmployee = function () {
        console.log($scope.type);
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            status:$scope.idinfo,
            role_id:$scope.type,
            accunt:$scope.account
        };
        PlatformService.getGetAdmin(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        })
    };

    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    //点击搜索
    $scope.search = function () {
        GetAllEmployee();
    };
    //新增账号
    $scope.adds = function () {
        var postData = {
            role_id:$scope.add.type*1,
            account:$scope.add.account,
            password:$scope.add.password,
            confirm_password:$scope.add.confirm_password,
            status:$scope.add.status
        };
        PlatformService.getPostAdmin(postData).then(function (response) {
           console.log(response);
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    //修改状态
    $scope.disables=function (item) {
        var sure = function () {
            var postData = {
                id:item.id
            };
            PlatformService.getAdminStatus(postData).then(function (response) {
                if (response.data.data === null) {
                    if(item.status==1){
                        item.status = 2;
                    }else{
                        item.status = 1;
                    }
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
    };
    //获取单个详情
    $scope.getID = function (item) {
        $scope.modifyes = item;
        var num = item.role_id;
        var numbers = $("#numbers").find("option");
        for (var j = 1; j < numbers.length; j++) {
            if ($(numbers[j]).val() == num) {
                $(numbers[j]).attr("selected", "selected");
            }
        }
    };
    //修改后提交
    $scope.submited = function () {
        console.log($scope.role_id);
        var postData = {
            role_id:$scope.role_id*1,
            account:$scope.modifyes.account,
            password:$scope.modifyes.password,
            confirm_password:$scope.modifyes.confirm_password,
            status:$scope.modifyes.status,
            id:$scope.modifyes.id
        };
        PlatformService.getPutAdmin(postData).then(function (response) {
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    // 删除
    $scope.del=function (id) {
        var del = function () {
            var postData = {
                id:id
            };
            PlatformService.getAdminDel(postData).then(function (response) {
                if(response.data.data===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
    };
});