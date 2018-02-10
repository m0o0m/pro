/**
 * Created by apple on 17/12/20.
 */

angular.module('app.site').controller('MaintenancesetCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        //获取数据
        siteService.maintencesttings().then(function (response) {
            $scope.data = response.data.data;
        });
        //获取站点
        $scope.detail=function (item) {
          siteService.thirdDropf(item.id).then(function (res) {
              $scope.site = res.data.data;
          });
            $scope.res = item;
        };

        //修改
        $scope.modify = function () {
            $scope.arr = [];
            $('input[class="disable"]:checked').each(function(){
                $scope.arr.push($(this).val());
            });
            console.log($scope.arr);

            var postData= {
                site_id : $scope.arr,
                main : $scope.res.main
            };
            siteService.savesettings(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        }

    });

