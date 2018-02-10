/**
 * Created by apple on 17/12/19.
 */
angular.module('app.site').controller('dataCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService,$stateParams){
        $scope.pag = "视讯配置";
        $scope.packge = function () {
            siteService.videoConfifurtion( $stateParams.id).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                    $scope.pag="配置成功";
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        }


    });