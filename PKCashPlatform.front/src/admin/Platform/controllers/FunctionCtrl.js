angular.module('app.Platform').controller('FunctionCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService) {
    $scope.json = APP_CONFIG.option;
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
     var typeid = 2;
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            type:typeid
        };
        PlatformService.getPermissionGet(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    };
    //分页初始化
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    // 禁用or启用
    $scope.open = function (item) {
        var able = function () {
            var postData = {
                id:item.id*1
            };
            PlatformService.getPermissionStatus(postData).then(function (response) {
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
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), able);
    };
    //新增功能
    $scope.addF = function () {
        var postData = {
            permission_name:$scope.f_name,
            module:$scope.f_username,
            route:$scope.f_route,
            method:$scope.f_method,
            status:$scope.f_status*1,
            type:$scope.id*1
        };
        PlatformService.getPermissionPost(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    }
    //修改
    $scope.f_Id;
    $scope.modify = function (id) {
        $scope.f_Id = id;
        var postData = {
            id:id
        };
        PlatformService.getPermissionInfo(postData).then(function (response) {
            $scope.info = response.data.data;
        })
    };
    $scope.submit = function () {
        var data = {
            id:$scope.f_Id,
            permission_name:$scope.info.permission_name,
            module:$scope.info.module,
            route:$scope.info.route,
            method:$scope.info.method,
            status:$scope.info.status*1
        };
        PlatformService.getPermissionPut(data).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
    $scope.delete = function (id) {
        var del = function () {
            var postData = {
                id:id*1
            };
            PlatformService.getPermissionDel(postData).then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), del);
    };

    $scope.adminType = function () {
        typeid = 2 ;
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            type:typeid
        };
        PlatformService.getPermissionGet(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    };
    $scope.agencyType = function () {
        typeid = 1 ;
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            type:typeid
        };
        PlatformService.getPermissionGet(postData).then(function (response) {
            console.log(response);
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    }
});