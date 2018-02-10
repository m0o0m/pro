angular.module('app.PaymentSetting').controller('ThirdBankEliminCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope){

        //获取站点
        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
            console.log($scope.siteJson);
            $scope.site = $scope.siteJson[0].site_index_id;
        })
        //获取第三方
        paymentSettingService.thirdPaidList().then(function (res) {
            $scope.third = res.data;
            $scope.third_ids = $scope.third[0].payId
        });



    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site,
            paid_tpye: $scope.third_ids,
        }
        paymentSettingService.thirdBank(postData).then(function (response) {
                 $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
                console.log($scope.list);
        })


    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    // 更改状态
    $scope.disables=function (item) {
        // console.log(id);
        var status = 1;
        if (item.status === 2 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        }
        var sure = function () {
            paymentSettingService.thirdStatus(item.id,status).then(function (response) {
                if(response===null){
                    item.status = status;
                    popupSvc.smallBox("success",$rootScope.getWord("success"));

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
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
