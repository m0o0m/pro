angular.module('app.Precalcula').controller('NewOfferCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state){
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取层级
        precalculaService.getLevel().then(function (response) {
            $scope.drp = response.data;
        });
        //获取商品
        precalculaService.comboPlatform().then(function (response) {
            $scope.commodity = response.Children;
            console.log($scope.commodity);
        })


   //确认提交
    $scope.submit = function () {
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
            site_index_id:$scope.site,
            level_id:$scope.drp_id,
            valid_money:$scope.valid_money,
            discount_up:$scope.discount_up,
            params:check_val
        }
        precalculaService.retreatWaterSetAdd(postData).then(function (response) {
            if(response===null){
                popupSvc.smallBox("success",$rootScope.getWord("success"));
                $state.go('app.Precalcula.Precalcula');
            }else {
                popupSvc.smallBox("fail",response.msg);
            }
        });
    };


});
