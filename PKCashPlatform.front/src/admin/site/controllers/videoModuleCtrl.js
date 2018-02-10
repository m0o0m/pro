/**
 * Created by apple on 17/12/18.
 */
angular.module('app.site').controller('videoModuleCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService,$stateParams){

    siteService.moduleManagemnet($stateParams.id).then(function (response) {
        $scope.data =response.data.data;
        console.log($scope.data);
    });

    $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        //提交
        $scope.se = function () {
            $scope.arr = [];
            $('input[class="disable"]:checked').each(function(){
                $scope.arr.push($(this).val());
            });
            var postData = {
                id:$scope.arr
            }
            siteService.modeular(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };


    });