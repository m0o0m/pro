/**
 * Created by apple on 17/11/21.
 */
/**
 * Created by apple on 17/9/8.
 */
/**
 * Created by apple on 17/8/15.
 */
angular.module('app.customer').controller('AddHierarchyCtrl',
    function($scope,$state,httpSvc,popupSvc,resourceSvc){
        console.log('zxjj')


        $scope.siteId=function () {
            $scope.site_index_id=$("#site_index_id")[0].value;
        }

        $scope.sumbit=function () {
            $scope.formData=angular.extend($scope.formData,{
                pay_set_id: $("#pay_set")[0].value*1,
                site_index_id: $scope.site_index_id,
            })
            httpSvc.post("/member/level",$scope.formData)
                .then(function (response) {
                    if(response===null){
                        popupSvc.smallBox("success","添加成功");
                        window.history.go(-1);
                    }else {
                        popupSvc.smallBox("fail","添加失败," + response.msg);
                    }
                },function (error) {

                })

        }
    });
