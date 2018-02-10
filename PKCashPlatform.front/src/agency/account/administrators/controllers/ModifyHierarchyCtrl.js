angular.module('app.administrators').controller('ModifyHierarchyCtrl',
    function($scope,$stateParams,httpSvc,popupSvc){

        httpSvc.get("/member/level/info",{
            level_id: $stateParams.level_id,
            site_index_id: $stateParams.site_index_id,
        }).then(function (response) {
            $scope.formData=angular.copy(response.data);
        })
        $scope.submit=function () {
            $scope.formData=angular.extend($scope.formData,{
                start_time: $("#StartDateline")[0].value,
                end_time: $("#EndDateline")[0].value,
                old_level_id: $stateParams.level_id,
                new_level_id: $scope.formData.level_id,
                site_index_id: $stateParams.site_index_id
            });
            delete $scope.formData.level_id;
            httpSvc.put("/member/level",$scope.formData)
            .then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success","修改成功");
                    window.history.go(-1);
                }else {
                    popupSvc.smallBox("fail","修改失败");
                }
            },function (error) {
                
            })
        }

    });
