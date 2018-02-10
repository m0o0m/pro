
angular.module('app.PaymentSetting').controller('PayParameterCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state){
     //获取币别
        paymentSettingService.currency().then(function (response) {
            $scope.da = response.data;
        })

     //获取站点
        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });

        var GetAllEmployee = function () {
        var postData = {
            site:$scope.site,
            account: $scope.accounts,
        }
        paymentSettingService.paysetList(postData).then(function (response) {
            console.log(response);
                 $scope.paginationConf.totalItems = response.mate.count;
                 $scope.list = response.data;
                // console.log($scope.list);
        });


    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    // 更改状态
    $scope.disables=function (status,id) {
        console.log(id);
        var sure = function () {
            httpSvc.put("/product/status",{
                product_id:id
            }).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success","修改状态成功");
                    GetAllEmployee();

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        }
        popupSvc.smartMessageBox("确定更改状态?",sure);
    };
    $scope.Det = function () {
        $state.go('app.PaymentSetting.DetailsSetting')

    };

    //点击添加
    $scope.addAcount = function () {
        var postData = {
            site_index_id:$scope.site,
            id:$scope.formData.parent_id*1,
            title:$scope.name
        }
        paymentSettingService.paysetAdd(postData).then(function (data) {
            if(data==null){
                popupSvc.smallBox("success",$rootScope.getWord("success"));
                GetAllEmployee();
                // $state.go('app.PaymentSetting.DepositBank')
            }else {
                popupSvc.smallBox("fail",data.msg);
            }
        });

    };
    //获取单个详情
        $scope.datali = function (id) {
            paymentSettingService.paysetDetail(id).then(function (data) {
               $scope.currency= data.data[0];
               console.log($scope.currency.id*1);
            }) ;
        };
       //点击修改后提交
        $scope.submitd = function () {
            var postData = {
                site_index_id:$scope.currency.site_index_id,
                id:$scope.currency.id,
                name:$scope.currency.title
            }
            paymentSettingService.paysetModify(postData).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));

                }else {
                    popupSvc.smallBox("fail",response.msg);
                };
            });
        };
        // 更改状态
        $scope.disables=function (id) {
            var sure = function () {
                paymentSettingService.paysetDel(id).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    };
                });
            }
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
        };



    //点击到币别
    $scope.bibie = function () {

        var GetAllEmployees = function () {
            var postData = {
                site:$scope.site,
                account: $scope.accounts,
            }
            paymentSettingService.paysetPublic(postData).then(function (response) {
                    $scope.paginationConfs.totalItems = response.meta.count;
                    $scope.list1 = response.data;
            });


        }
        $scope.paginationConfs = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };

        $scope.$watch('paginationConfs.currentPage + paginationConfs.itemsPerPage', GetAllEmployees);

        // 删除币别
        $scope.disables1=function (id) {
            var sure = function () {
                paymentSettingService.denomination(id).then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    };
                });
            }
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
        };


    };



});

