angular.module('app.Platform').controller('AccountHolderCtrl', function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$state) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        
        $scope.desc=true;

        //套餐列表
        PlatformService.getComboDrop().then(function (response) {
            $scope.packageList=response.data;
        });
        $scope.json = APP_CONFIG.option;
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                status: $scope.status,
                is_login: $scope.is_login,
                account: $scope.account,
                pai_xu: $scope.order_by,
                shun_xu: $scope.desc
            };
            PlatformService.getHolderList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.data;
            })
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        // 停用
        $scope.disable=function (item) {
            var sure = function () {
                var postData = {
                    id: item.id,
                    status: item.status
                };
                PlatformService.getHolderDisable(postData).then(function (response) {
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
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), sure);
        };
        $scope.delete = function (id) {
            var del = function () {
                var postData = {
                    id: id
                };
                PlatformService.getHolderDel(postData).then(function (response) {
                    console.log(response);
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
        //点击搜索
        $scope.search = function () {
            GetAllEmployee();
        };
        //新增
        $scope.addAcount = function () {
            $scope.addForm.status=$scope.addForm.status*1;
            $scope.addForm.combo_id=$scope.addForm.combo_id*1;
            $scope.addForm.domain_up=$scope.addForm.domain_up*1;
            $scope.addForm.up_cose=$scope.addForm.up_cose*1;
            PlatformService.getHolderAdd($scope.addForm).then(function (response) {
                if (response.datd.data===null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        //修改
        $scope.modify=function (id) {
            $scope.id=id;
            var postData = {
                id: $scope.id
            };
            PlatformService.getHolder(postData).then(function (response) {
                $scope.modifyData = response.data.data;
                console.log($scope.modifyData);
            })
        };
        //提交修改
        $scope.modifySubmit=function () {
            var postData = {
                id: $scope.id,
                account: $scope.modifyData.account,
                password: $scope.modifyData.password,
                re_password: $scope.modifyData.re_password,
                operate_password: $scope.modifyData.operate_password,
                username: $scope.modifyData.username,
                status: $scope.modifyData.status*1
            };
            PlatformService.getHolderUpdata(postData).then(function (response) {
                $(".modal-backdrop").hide();
                $("#myModal2").hide();
                if(response.data.data===null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            });
        };
        $scope.order=function () {
            $scope.desc=!$scope.desc;
            GetAllEmployee();
        };
        $scope.orderChange=function(){
            GetAllEmployee();
        }
});