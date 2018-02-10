angular.module('app.Precalcula').controller('PreferentialStatisticsCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state){
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取层级
        precalculaService.getLevel().then(function (response) {
            $scope.drp = response.data;
        })
       $scope.option_Preferential_member = APP_CONFIG.option.option_Preferential_member

    $scope.change = function (id) {
        console.log(id);
        if(id ==1){
            $scope.istrue =true;
            $scope.istrue_1 = false;
        }else {
            $scope.istrue = false;
            $scope.istrue_1 = true;
        }
    }

    $scope.se = function () {
        var arr = [];
        $('input[name="radio-inline2"]:checked').each(function(){
            arr.push($(this).val());
        });
        $state.go('app.Precalcula.query',{
            site_index_id:$scope.site,
            rtype:$scope.rtype,
            typeed:$scope.typeed,
            v_type:arr,
            start_time:$scope.start_time,
            end_time:$scope.end_time
        });
    };


});

