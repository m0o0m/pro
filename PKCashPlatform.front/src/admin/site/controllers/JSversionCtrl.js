/**
 * Created by apple on 17/11/20.
 */
/**
 * Created by apple on 17/11/20.
 */
angular.module('app.site').controller('JSversionCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
    //获取WEP数据
   siteService.jsWep().then(function (res) {
       $scope.list = res.data.data;
   });
   //获取总表数据
   siteService.JSTable().then(function (res) {
       $scope.list2 = res.data.data;
   });
   //获取PC端数据
    siteService.JSPc().then(function (res) {
        $scope.list1 = res.data.data;
    })

    //获取详情
    $scope.detail = function (id,type_id) {
        siteService.JSDetail(id,type_id).then(function (res) {
            $scope.data = res.data.data;
        });
    };


//删除
    $scope.delete = function (id,type_id) {
        var del = function () {
            siteService.JSDel(id,type_id).then(function (response) {
                if(response.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        }
        popupSvc.smartMessageBox($rootScope.getWord('confirmationOperation'),del);
    }
    //生成web版本
    $scope.add = function () {
        var add = function () {
            siteService.genrateWep($scope.account).then(function (response) {
                if(response.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                };
            });
        };
        popupSvc.smartMessageBox($rootScope.getWord('confirmationOperation'),add);
    };
    //生成PC端
    $scope.addl = function () {
        var add = function () {
            siteService.genratePc($scope.account).then(function (response) {
                if(response.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                };
            });
        }
        popupSvc.smartMessageBox($rootScope.getWord('confirmationOperation'),add);
    };
    //编辑总表
     $scope.geteQuipment = function (Quipment) {
         $scope.Quipment =Quipment;
     };
    //修改后提交
    $scope.submiteds = function () {
        siteService.tableModify($scope.Quipment).then(function (response) {
            if(response.data==null){
                popupSvc.smallBox("success",$rootScope.getWord("success"));
            }else {
                popupSvc.smallBox("fail",response.msg);
            };
        });
    };

    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {

            $scope.newTodo = undefined;
        }
    };
});