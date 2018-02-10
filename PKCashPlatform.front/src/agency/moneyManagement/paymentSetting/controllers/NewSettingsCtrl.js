angular.module('app.PaymentSetting').controller('NewSettingsCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state){
    //设备
     $scope.shebei = APP_CONFIG.option.shebei;

    //站点
    paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });
    //层级
    paymentSettingService.getLevel().then(function (res) {
            console.log(res);
            $scope.drp = res.data;
        });
    //获取类型
    paymentSettingService.paidType().then(function (res) {
            console.log(res);
            $scope.paid_types = res.data;
    })  ;


    //点击添加
    $scope.se = function () {
        var arr = [];
        $('input[name="radio-inline2"]:checked').each(function(){
            arr.push($(this).val());
        });
        console.log(arr);
        var postData = {
            level:arr,
            paid_domain:$scope.paid_domain,
            site_index_id:$scope.site,
            backaddress:$scope.backaddress,
            merchat_id:$scope.merchat_id,
            private_key:$scope.private_key,
            public_key:$scope.public_key,
            paid_limit:$scope.paid_limit,
            paid_platform:$scope.paid_platform,
            suitable_equipmentL:$scope.suitable_equipment,
            paid_type:$scope.paid_type,
            paid_code:$scope.paid_code,
            status:$scope.status
        };
        paymentSettingService.newOnlineSetup(postData).then(function (res) {
            if(res==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
                $state.go('app.PaymentSetting.OnlinePayment');
            }else {
                popupSvc.smallBox("fail",res.msg);
            }
        })

        // httpSvc.post('/new_online_setup',{
        //
        // }).then(function (res) {
        //     console.log(res);
        //
        // });
    }
});

