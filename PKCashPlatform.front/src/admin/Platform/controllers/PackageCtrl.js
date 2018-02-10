angular.module('app.Platform').controller('PackageCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state) {
    //获取下拉框
    $scope.json = APP_CONFIG.option;
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
            status:$scope.type,
            combo_name:$scope.combo_name
        };
        PlatformService.getComboGet(postData).then(function (response) {
            console.log(response);
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
    //新增套餐
    $scope.adds = function () {
        var postData = {
            combo_name:$scope.add.combo_name
        };
        PlatformService.getComboPost(postData).then(function (response) {
            if (response.data.data === null) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
                GetAllEmployee();
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    };

    //修改状态
    $scope.disables=function (item) {
        var able = function () {
            var postData = {
                id:item.id
            };
            PlatformService.getComboStatus(postData).then(function (response) {
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
    $scope.getID = function (ides) {
        var postData = {
            id:ides
        };
        PlatformService.getComboInfo(postData).then(function (response) {
            $scope.modifyes = response.data.data;
        })
    };
    //修改后提交
    $scope.submiteds = function () {
        var data = {
            id:$scope.modifyes.id,
            combo_name:$scope.modifyes.combo_name
        };
        PlatformService.getComboPut(data).then(function (response) {
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
            PlatformService.getComboDel(postData).then(function (response) {
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
//点击跳转配置
    $scope.allocation  = function (idese) {
        $state.go('app.Platform.allocation',{
            ids:idese
        })
    }
});