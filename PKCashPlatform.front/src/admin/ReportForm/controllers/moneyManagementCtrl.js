angular.module('app.ReportForm').controller('moneyManagementCtrl',
    function ($scope, popupSvc, financeMoneyManagementService, $rootScope, APP_CONFIG) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };

        financeMoneyManagementService.getHierarchy({

        }).then(function (response) {
            $scope.hierarchyList = response.data;
        });
        $scope.paginationConf1 = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee1 = function () {
            var postData = {
                page: $scope.paginationConf1.currentPage,
                pageSize: $scope.paginationConf1.itemsPerPage,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                account: $scope.account
            };
            financeMoneyManagementService.getThird(postData).then(function (response) {
                $scope.paginationConf1.totalItems = response.meta.count;
                $scope.list1 = response.data;
            });

        };
        $scope.$watch('paginationConf1.currentPage + paginationConf1.itemsPerPage', GetAllEmployee1);
        $scope.search = function () {
            GetAllEmployee();
        };


        $scope.paginationConf2 = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        var GetAllEmployee2= function () {
            var postData = {
                page: $scope.paginationConf2.currentPage,
                pageSize: $scope.paginationConf2.itemsPerPage,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                account: $scope.account
            };
            financeMoneyManagementService.getBank(postData).then(function (response) {
                $scope.paginationConf2.totalItems = response.meta.count;
                $scope.list2 = response.data;
            });

        };
        $scope.$watch('paginationConf2.currentPage + paginationConf2.itemsPerPage', GetAllEmployee2);


        $scope.addThird=function(){
            financeMoneyManagementService.addThird($scope.formData1).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee1();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.addBank=function(){
            financeMoneyManagementService.addBank($scope.formData2).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee2();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.modifyThird=function(item){
            $scope.modifyData1=angular.copy(item);
        };
        $scope.modifyBank=function(item){
            $scope.modifyData2=angular.copy(item);
        };

        $scope.modifyThirdSubmit=function(){
            financeMoneyManagementService.modifyThird($scope.modifyData1).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee1();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.modifyBankSubmit=function(){
            financeMoneyManagementService.modifyBank($scope.modifyData2).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee2();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

    });

