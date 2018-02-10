angular.module('app.PaymentSetting').controller('ModifySettingsCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state){
        $scope.option_status =  APP_CONFIG.option.option_status;
        $scope.option_outbank = APP_CONFIG.option.option_outbank
        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });

    paymentSettingService.getLevel().then(function (res) {
        console.log(res);
        $scope.drp = res.data;
    })
        $scope.se = function () {
            $scope.arr = [];
            $('input[class="disable"]:checked').each(function(){
                $scope.arr.push($(this).val());
            });
                var formData = new FormData($("#formid_a")[0]);
                     formData.append("level", $scope.arr);
                     console.log(formData);


                 paymentSettingService.addPayment(formData).then(function (data) {
                     if(data==null){
                         popupSvc.smallBox("success",$rootScope.getWord('success'));
                            $state.go('app.PaymentSetting.DepositBank');
                     }else {
                            popupSvc.smallBox("fail",data.msg);
                     }

                 });




        };


});

