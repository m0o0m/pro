angular.module('app.PaymentSetting').controller('DepositBankCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state){
        $scope.option_status =  APP_CONFIG.option.option_status;
        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });


    var GetAllEmployee = function () {
        var postData = {
            site:$scope.site,
            account: $scope.accounts,
            status:$scope.status
        };
        paymentSettingService.paymentList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
                console.log($scope.list);
        });


    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    //点击搜索
    $scope.search = function () {
        GetAllEmployee();
    }


    // 更改状态
    $scope.disables=function (status,id,item) {
        console.log(id);
        if (item.status === 2 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        };
        var sure = function () {
            paymentSettingService.paymentStatus(id).then(function (response) {
                    if(response===null){
                        item.status = status;
                        popupSvc.smallBox("success",$rootScope.getWord("success"));

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    };
            });

        }
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
    };

    $scope.set = function () {
        $state.go('app.PaymentSetting.DepositRecord')
    };
    $scope.modi = function () {
        $state.go('app.PaymentSetting.ModifySettings');
    };

    //删除
    $scope.del = function (id) {
        console.log(id);
        var sure = function () {
            paymentSettingService.paymentDelete(id*1).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    GetAllEmployee();
                }else {
                    popupSvc.smallBox("fail",response.msg);
                };
            });
        }
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
    };


    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };
});
