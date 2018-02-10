angular.module('app.ReportForm').controller('ArrearsCtrl',
    function ($scope, popupSvc, commonService, financeArrearsService, $rootScope, APP_CONFIG) {
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
                status: $scope.status
            };
            financeArrearsService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };

        $scope.add = function () {
            var siteList=[];
            var list=$(".check1");
            for(var i=0; i<list.length; i++){
                if($(list[i]).prop("checked")){
                    var id=$(list[i]).parent().attr("data-site");
                    siteList.push(id);
                }
            }
            $scope.addData=angular.extend($scope.addData, {
                "siteList": siteList
            });
            console.log($scope.addData);

            financeArrearsService.add($scope.addData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.modify=function(item){
            $scope.modifyData=angular.copy(item);
        };


        $scope.modifySubmit = function () {
            var siteList=[];
            var list=$(".check2");
            for(var i=0; i<list.length; i++){
                if($(list[i]).prop("checked")){
                    var id=$(list[i]).parent().attr("data-site");
                    siteList.push(id);
                }
            }
            $scope.modifyData=angular.extend($scope.modifyData, {
                "siteList": siteList
            });
            console.log($scope.modifyData);
            financeArrearsService.modify($scope.modifyData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

    });
