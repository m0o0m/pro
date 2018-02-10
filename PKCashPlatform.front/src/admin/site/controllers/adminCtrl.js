/**
 * Created by apple on 17/12/19.
 */
angular.module('app.site').controller('adminCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){

        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            };
            siteService.adtaAdmin(postData).then(function (response) {
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

                name: $scope.data.name,
                account:$scope.data.account,
                domin: $scope.domin,
                remack: $scope.data.remack

            };
            siteService.adtaAdminModify(postData).then(function (data) {
                if(data.data==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));
                }else {
                    popupSvc.smallBox("fail",data.msg);
                };
            });
        };

        $scope.sub = function () {
            var postData={

                name: $scope.name,
                account:$scope.account,
                domin: $scope.domins,
                remack: $scope.remack

            };
            siteService.adtaAdminAdd(postData).then(function (data) {
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
                siteService.adtaAdminDel(item.id).then(function (response) {
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


    });