angular.module('app.Precalcula').controller('AddFeeSettingCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state){
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取商品
        precalculaService.comboPlatform().then(function (response) {
            $scope.commodity = response.Children;
            console.log($scope.commodity);
        });

        $scope.ckecked = function (event) {
            var parent = $(event.target).parent()[0];
            var input= $(parent).find(".inputsese")[0];
            if(event.target.checked){
                $(input).removeClass('ng-hide');
                $(input).show();
            }else {
                $(input).hide();
            }
        };
        $scope.sumbit = function () {

            var check_val = [];
            var test  = document.getElementsByClassName('test');
            for (var j=0;j<test.length;j++){
                if(test[j].checked){
                    var parent =$(test[j]).parent()[0];
                    var obj ={
                        platform_id:test[j].value*1,
                        proportion:$(parent).find('.inputse')[0].value*1
                    };
                    check_val.push(obj);
                };
            };
            var postData = {
                site_index_id:$scope.site_id,
                self_profit:$scope.self_profit,
                effective_user:$scope.effective_user,
                list:check_val
            }
            precalculaService.overrideAddone(postData).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    $state.go('app.CommissionStatistics.AgentSetting');
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };





});