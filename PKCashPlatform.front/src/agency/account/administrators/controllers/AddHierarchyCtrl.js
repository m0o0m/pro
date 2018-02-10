angular.module('app.administrators').controller('AddHierarchyCtrl',
    function(AccountService,$scope,$state,httpSvc,popupSvc,resourceSvc){
        var user=JSON.parse(resourceSvc.getSession("user"));

        $scope.isSuperAdmin=user.site_index_id==='';

        if($scope.isSuperAdmin){
            //获取站点
            AccountService.getFirstDropSelect().then(function (response) {
                $scope.siteJson=response;
            });
        }else{
            $scope.site_index_id=user.site_index_id;
        }

        $scope.siteId=function () {
            $scope.site_index_id=$("#site_index_id")[0].value;
        };

        $scope.sumbit=function () {
            $scope.formData=angular.extend($scope.formData,{
                pay_set_id: $("#pay_set")[0].value*1,
                site_index_id: $scope.site_index_id,
            });
            AccountService.putMemberLevel($scope.formData)
            .then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success","添加成功");
                    window.history.go(-1);
                }else {
                    popupSvc.smallBox("fail","添加失败," + response.msg);
                }
            });
        };
});
