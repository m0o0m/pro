angular.module('app.CashManagement').controller('MonitorCtrl', function($interval, AccessMoneyService,httpSvc, popupSvc, $stateParams, $scope, $rootScope, APP_CONFIG, $state, $LocalStorage){
 console.log(12);
   //获取公司入款
   AccessMoneyService.monitorDeposit().then(function (res) {
    $scope.list1 = res;
 });

   //获取线上入款
    $scope.online = function () {
        AccessMoneyService.MonitorOnline().then(function (res) {
            $scope.list2 = res;
        })
    };
    //获取出款管理
    $scope.Monitors = function () {
        AccessMoneyService.MonitorMoney().then(function (res) {
            $scope.list3 = res.list;
        })
    };







    $scope.sure=function(){
        var del = function () {

        }
        popupSvc.smartMessageBox("是否确认此操作",del);
    }

    $scope.cancle=function(){
        var del = function () {

        }
        popupSvc.smartMessageBox("是否确认此操作",del);
    }

});
