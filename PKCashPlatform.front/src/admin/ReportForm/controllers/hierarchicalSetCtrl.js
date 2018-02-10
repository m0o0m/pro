angular.module('app.ReportForm').controller('hierarchicalSetCtrl',
    function ($scope, popupSvc, commonService, financeHierarchicalService, $rootScope, APP_CONFIG) {
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
            };
            financeHierarchicalService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.platforms;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.del = function (id) {
            var fn=function() {
                financeHierarchicalService.del({
                    id: id
                }).then(function (response) {
                    if (response === null) {
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                        GetAllEmployee();
                    } else {
                        popupSvc.smallBox("fail", response.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation")+"？", fn);
        };
        $scope.add=function(){
            var list=$(".platforms");
            var proportion=[];
            for(var i=0; i<list.length; i++){
                proportion.push(list[i].value)
            }
            financeHierarchicalService.add({
                lid: $scope.addData.lid,
                level_name: $scope.addData.level_name,
                talk: $scope.addData.talk,
                proportion: proportion
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.modify=function(item){
            $scope.modifyData=item;
            var proportion=item.proportion;
            var list=$(".platformss");
            for(var i=0; i<list.length; i++){
                list[i].value=proportion[i]
            }
        };
        $scope.modifysubmit=function(){
            var list=$(".platformss");
            var proportion=[];
            for(var i=0; i<list.length; i++){
                proportion.push(list[i].value)
            }
            financeHierarchicalService.modify({
                id: $scope.modifyData.id,
                lid: $scope.modifyData.lid,
                level_name: $scope.modifyData.level_name,
                talk: $scope.modifyData.talk,
                proportion: proportion
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }


    });
