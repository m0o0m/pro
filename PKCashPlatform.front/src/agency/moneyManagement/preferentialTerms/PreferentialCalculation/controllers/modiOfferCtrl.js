/**
 * Created by apple on 17/12/14.
 */
angular.module('app.Precalcula').controller('modiOfferCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$stateParams,$state){
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取层级
        precalculaService.getLevel().then(function (response) {
            $scope.drp = response.data;
        });
        //获取详情
        precalculaService.retreatWaterSetDetail($stateParams.id).then(function (res) {
            $scope.data = res.data;
            $scope.params = $scope.data.params;
            var arrId =[];
            for (var i =0; i<$scope.params.length;i++){
             arrId.push($scope.params[i].product_id);
            };
            $scope.isSelected = function (id) {
                return $.inArray(id, arrId)!=-1;
            };
        });
        //获取商品
        precalculaService.comboPlatform().then(function (response) {
            $scope.commodity = response.Children;
        });

        //修改后提交
        $scope.edit = function () {
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
                site_index_id:$scope.data.site_index_id,
                level_id:$scope.data.level_id,
                valid_money:$scope.data.valid_money,
                discount_up:$scope.data.discount_up,
                params:check_val
            };
            precalculaService.RETREAT_WATER_SET_EDIT(postData).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    $state.go('app.CommissionStatistics.AgentSetting');
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };



    });
