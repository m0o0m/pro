angular.module('app.quotaStatistics').controller('SeparateApplicationCtrl', function(popupSvc,$scope,$rootScope,QuotaStatisticsService){

    //获取转入转出项目
    QuotaStatisticsService.getTypeList().then(function(response){
        console.log(response);
        $scope.listData = response.data.data;
    });

    $scope.submit = function () {
        if($scope.username == undefined || $scope.ctype == undefined || $scope.vtype == undefined || $scope.money == undefined || $scope.do_time == undefined || $scope.remark == undefined){
            popupSvc.smallBox("fail","请输入完整！")
        }
        var postData = {
            username: $scope.username,
            ctype: $scope.ctype*1,
            vtype: $scope.vtype*1,
            money: $scope.money*1,
            do_time: $scope.do_time,
            remark: $scope.remark
        };
        console.log(postData);
        QuotaStatisticsService.getSubDropList(postData).then(function(response){
            console.log(response.data);
            if(response.data.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    }

});
