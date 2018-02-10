/**
 * Created by apple on 17/12/13.
 */

angular.module('app.PaymentSetting').controller('seedetailCtrl',
    function(httpSvc,popupSvc,$scope,paymentSettingService,APP_CONFIG,$rootScope,$stateParams,$state){
        //获取id
        $scope.id = $stateParams.id;
        //获取详情信息
        paymentSettingService.paysetPublicOne($scope.id).then(function (res) {
            $scope.data = res.data;
            console.log($scope.data.is_free);

        });





    });

