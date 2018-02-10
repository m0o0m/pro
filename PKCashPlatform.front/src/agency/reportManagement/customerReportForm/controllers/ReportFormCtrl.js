angular.module('app.ReportForm').controller('ReportFormCtrl',
    function(httpSvc,popupSvc,$scope,reportformService,APP_CONFIG,$rootScope,$stateParams,$state){
    //获取商品
    reportformService.getLevel().then(function (res) {
        $scope.res = res.data;
    });
    //获取站点
     reportformService.getDropSelect().then(function (response) {
         $scope.site = response.data;
     }) ;
    //类型
     $scope.option_time_zone  = APP_CONFIG.option.option_time_zone


    $scope.detail = function () {
        var arr = [];
        $('input[name="radio-inline2"]:checked').each(function(){
            arr.push($(this).val());
        });
        console.log($scope.names)
        $state.go('app.ReportForm.ReportDetails',{
            site_index_id:$scope.site_id,
            time_zone:$scope.type,
            username:$('.names').val(),
            rtype:$scope.Pattern,
            v_type:arr,
            start_time:$scope.start_time,
            end_time:$scope.end_time
        });
    };




});

