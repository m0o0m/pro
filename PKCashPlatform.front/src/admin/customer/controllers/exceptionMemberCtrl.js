angular.module('app.customer').controller('exceptionMemberCtrl',
    function ($scope, popupSvc, commonService, customerExceptionMemberService, $rootScope, APP_CONFIG) {
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

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                type: $scope.type,
                key: $scope.key
            };
            customerExceptionMemberService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };
        //立即处理
        $scope.deal = function (id) {
            var deal = function () {
                customerExceptionMemberService.deal({
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
            popupSvc.smartMessageBox("确定处理？", deal);
        };



        $scope.selectAll= function ($event) {
            var target = $event.target;
            var isCheck=$(target).is(':checked');
            var list=$("tbody").find('input[type="checkbox"]');
            for(var i=0;i<list.length;i++){
                if(isCheck){
                    $(list[i]).prop("checked",true);
                }else {
                    $(list[i]).prop("checked",false);
                }
            }
        };


    });