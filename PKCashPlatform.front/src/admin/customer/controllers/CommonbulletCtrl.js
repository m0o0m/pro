angular.module('app.customer').controller('CommonbulletCtrl',
    function ($scope, popupSvc, commonService, customerCommonbulletService, $rootScope, APP_CONFIG) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        $scope.paginationConf1 = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee1 = function () {
            var postData = {
                page: $scope.paginationConf1.currentPage,
                pageSize: $scope.paginationConf1.itemsPerPage,
                site: $scope.site,
                id: $scope.id
            };

            customerCommonbulletService.getBullet(postData).then(function (response) {
                $scope.paginationConf1.totalItems = response.meta.count;
                $scope.list1 = response.data;
            });
        };

        $scope.$watch('paginationConf1.currentPage + paginationConf1.itemsPerPage', GetAllEmployee1);
        $scope.search1=function(){
            GetAllEmployee1()
        };

        $scope.paginationConf2 = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee2 = function () {
            var postData = {
                page: $scope.paginationConf2.currentPage,
                pageSize: $scope.paginationConf2.itemsPerPage,
                site: $scope.site2,
                status: $scope.status
            };

            customerCommonbulletService.getAnimation(postData).then(function (response) {
                $scope.paginationConf2.totalItems = response.meta.count;
                $scope.list2 = response.data;
            });
        };

        $scope.$watch('paginationConf2.currentPage + paginationConf2.itemsPerPage', GetAllEmployee2);
        $scope.search2=function(){
            GetAllEmployee2()
        };

        $scope.disable=function(id,status){
            var deal = function () {
                customerCommonbulletService.disable({
                    id: id,
                    status: status===1?2:1
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                        GetAllEmployee2();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox("确定处理？", deal);
        }

    });
