/**
 * Created by apple on 17/12/18.
 */
/**
 * Created by apple on 17/12/18.
 */
angular.module('app.site').controller('negativeCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService,$stateParams){

        siteService.negative($stateParams.id).then(function (response) {
            $scope.list = response.data.list;
            $scope.name = response.data.product_name;
        })
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        $scope.add = function () {

            var check_val = [];
            var test  = document.getElementsByClassName('test');
            for (var j=0;j<test.length;j++){
                    var parent =$(test[j]).parent()[0];
                    var obj ={
                        platform_id:$(test[j])[0].innerHTML,
                        proportion:$(parent).find('.values')[0].value
                    };
                check_val.push(obj);

            };
            var postData ={
                arr:check_val
            }
            siteService.negativeAdd(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });

        };


    });