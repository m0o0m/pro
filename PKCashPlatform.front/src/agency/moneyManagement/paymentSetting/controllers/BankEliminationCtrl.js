angular.module('app.PaymentSetting').controller('BankEliminationCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope){

    paymentSettingService.getDropSelect().then(function (res) {
        $scope.siteJson = res.data;
    })


    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site1s,
            bank_name: $scope.bank_name,
        }
        paymentSettingService.bankOutbank(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;

        });


    };
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
        var status = 2;
        if (item.status === 2 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        }
        var sure = function () {
            paymentSettingService.outbankStatus(id).then(function (response) {
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
