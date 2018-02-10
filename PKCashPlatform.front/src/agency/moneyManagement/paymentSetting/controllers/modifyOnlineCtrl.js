/**
 * Created by apple on 17/11/30.
 */
angular.module('app.PaymentSetting').controller('modifyOnlineCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$stateParams,$state){

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
        //获取详情
        paymentSettingService.onlineSetupSingle($stateParams.Id).then(function (res) {
                console.log(res);
                $scope.res = res;
                console.log($scope.res);
                $scope.s = $scope.res.level
                console.log($scope.res.paid_type*1);
        })
        $scope.isSelected = function (id) {
            return $.inArray(id, $scope.s)!=-1;
        };


    //点击修改后提交
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
        paymentSettingService.newOnlineSetupModify(postData).then(function (data) {
            if(data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
                $state.go('app.PaymentSetting.OnlinePayment');
            }else {
                popupSvc.smallBox("fail",data.msg);
            }
        });

    };
});

