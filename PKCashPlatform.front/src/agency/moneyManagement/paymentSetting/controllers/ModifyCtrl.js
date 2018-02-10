/**
 * Created by apple on 17/11/28.
 */
angular.module('app.PaymentSetting').controller('ModifyCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$state,$stateParams){
        $scope.option_outbank = APP_CONFIG.option.option_outbank
    //获取站点
        paymentSettingService.getDropSelect().then(function (res) {
            $scope.siteJson = res.data;
        });
        //获取层级
        paymentSettingService.getLevel().then(function (res) {
            console.log(res);
            $scope.drp = res.data;
        })
        $scope.ids = $stateParams.Id;
        //获取单个详情
        paymentSettingService.paymentDetail($scope.ids).then(function (res) {
            console.log(res);
            $scope.Froms = res.one_bank_pay_set;
                $scope.level_id = res.level_id;
                console.log($scope.level_id);
                console.log($scope.Froms.id);
        });

        $scope.isSelected = function (id) {
            return $.inArray(id, $scope.level_id)!=-1;
        };

        $scope.se = function () {
            $scope.arr = [];
            $('input[class="disable"]:checked').each(function(){
            $scope.arr.push($(this).val());
         });
        var formData = new FormData($("#formid_a")[0]);
        formData.append("level", $scope.arr);
        paymentSettingService.paymentPut(formData).then(function (data) {
            if(data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
                    $state.go('app.PaymentSetting.DepositBank');
                }else {
                    popupSvc.smallBox("fail",data.msg);
                    };
        });




    };


});

