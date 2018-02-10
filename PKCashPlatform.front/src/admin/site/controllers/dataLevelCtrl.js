/**
 * Created by apple on 17/12/19.
 */
angular.module('app.site').controller('dataLevelCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        //获取站点
        siteService.thirdDropf().then(function (response) {
            $scope.siteJson = response.data.data;
        });
        $scope.option_onOff = APP_CONFIG.option.option_onOff

        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            };
            siteService.hierchicalData(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta[0].count;
                $scope.list =response.data.data;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };
        //获取单个详情
        $scope.detali = function (item) {
          $scope.data = item;
        };
        //修改
        $scope.sumbit = function () {
            var postData={
                id:$scope.data.id,
                name:$scope.data.name,
                type:$scope.data.type
            };
            siteService.hierchicalDataModify(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };
        //删除
        $scope.del = function (item) {
            var sure = function () {
                siteService.hierchicalDataDel(item.id).then(function (response) {
                    if(response.data===null){
                        item.status = status;
                        popupSvc.smallBox("success",$rootScope.getWord("success"));

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                });

            }
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),sure);
        };
        //添加
        $scope.add = function () {
            var postData={
                name :$scope.name,
                type:$scope.type,
                id:$scope.id
            };
           siteService.hierchicalDataAdd(postData).then(function (data) {
               if(data.data==null){
                   popupSvc.smallBox("success",$rootScope.getWord('success'));
               }else {
                   popupSvc.smallBox("fail",data.msg);
               };
           });
        };

    });