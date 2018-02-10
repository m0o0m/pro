angular.module('app.PaymentSetting').controller('OnlinePaymentCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state){
        //下拉框初始化
        $scope.option_online =  APP_CONFIG.option.option_online;

        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });
        //获取第三方
        paymentSettingService.thirdPaidList().then(function (res) {
            $scope.third = res.data;
        });


    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site,
            third_id: $scope.third_ids,
            status:$scope.typeed
        }
        paymentSettingService.onlineSetup(postData).then(function (response) {
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
    $scope.disables=function (item) {
        if (item.status === 2 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        };
        var sure = function () {
            paymentSettingService.stopOnlineSstup(item.id).then(function (response) {
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
    // 删除
    $scope.del=function (id) {
        console.log(id);
        var sure = function () {
            paymentSettingService.onlineSetupDel(id).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    GetAllEmployee();
                }else {
                    popupSvc.smallBox("fail",response.msg);
                };
            });
            // httpSvc.del("/online_setup/del",{
            //     id:id
            // }).then(function (response) {
            //     if(response===null){
            //         popupSvc.smallBox("success","删除成功");
            //         GetAllEmployee();
            //
            //     }else {
            //         popupSvc.smallBox("fail",response.msg);
            //     }
            // });
        }
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
    };

    $scope.set = function () {
        $state.go('app.PaymentSetting.NewSettings')
    };
    $scope.mos = function () {
        $state.go('app.PaymentSetting.DepositRecord')
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

