angular.module('app.CommissionStatistics').controller('ModificationFeeCtrl',
    function ($scope, popupSvc, siteService, feeSettingService, $rootScope, APP_CONFIG, $stateParams) {
        feeSettingService.getInfo({
            id: $stateParams.id
        }).then(function (response) {
            $scope.income_poundage_ratio=response.data.income_poundage_ratio;
            $scope.income_poundage_up=response.data.income_poundage_up;
            $scope.out_poundage_ratio=response.data.out_poundage_ratio;
            $scope.out_poundage_up=response.data.out_poundage_up;
            $scope.is_delivery_model=response.data.is_delivery_model;
        });

        $scope.modify=function(){
            feeSettingService.modify({
                id: $stateParams.id,
                income_poundage_ratio: $scope.income_poundage_ratio,
                income_poundage_up: $scope.income_poundage_up,
                out_poundage_ratio: $scope.out_poundage_ratio,
                out_poundage_up: $scope.out_poundage_up,
                is_delivery_model: $scope.is_delivery_model
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }
    });

