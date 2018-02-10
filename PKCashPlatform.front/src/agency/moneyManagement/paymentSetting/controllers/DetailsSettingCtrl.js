
angular.module('app.PaymentSetting').controller('DetailsSettingCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$stateParams,$state){
        //获取id
        $scope.id = $stateParams.id;
        //获取详情信息
        paymentSettingService.paysetOne($scope.id).then(function (res) {
            $scope.data = res.data;
            console.log($scope.data.is_free);

        });

        //修改后提交
        $scope.submits = function () {
            var postData = $scope.data;
            paymentSettingService.paysets(postData).then(function (data) {
                if(data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                    $state.go('app.PaymentSetting.PayParameter');
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };




});

