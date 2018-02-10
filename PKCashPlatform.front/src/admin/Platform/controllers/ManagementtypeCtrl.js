angular.module('app.Platform').controller('ManagementtypeCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state) {
    $scope.toggleAdd = function () {
        if (!$scope.newTodow) {
            $scope.newTodow = {
                state: 'Important'
            };
        } else {
            $scope.newTodow = undefined;
        }
    };
    var GetAllEmployee = function () {
        var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            title:$scope.title
        };
        PlatformService.getProductTypeInfo(postData).then(function (response) {
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
    //点击搜索
    $scope.search = function () {
        GetAllEmployee();
    };
    //点击添加
    $scope.add = function () {
        var postData = {
            title:$scope.formData.title,
            status:$scope.formData.status
        };
        PlatformService.getProductTypePost(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };
    // 更改状态
    $scope.disables=function (item) {
        var able = function () {
            var postData = {
                id:item.id
            };
            PlatformService.getProductTypeStatus(postData).then(function (response) {
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
    //获取单个详情
    $scope.getID = function (ids) {
        var postData = {
            id:ids
        };
        PlatformService.getProductTypeGet(postData).then(function (response) {
            $scope.modifyes = response.data.data;
        })
    };
    //修改后提交
    $scope.submit = function () {
        var data = {
            id:$scope.modifyes.id,
            title:$scope.modifyes.title,
            status:$scope.modifyes.status
        };
        PlatformService.getProductTypePut(data).then(function (response) {
            if (response.data.data === null) {
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
            PlatformService.getProductTypeDel(postData).then(function (response) {
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

});