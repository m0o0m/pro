angular.module('app.administrators').controller('PlayAccountsCtrl',
    function(BusinessService,httpSvc,popupSvc,DTOptionsBuilder,DTColumnBuilder,$http,$scope,$rootScope,$compile,APP_CONFIG){
        //TODO 
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        // 停用
        $scope.disable=function () {
            popupSvc.smartMessageBox("确定停用此账户权限?","停用成功","停用失败");
        }

        // 踢线
        $scope.kick=function () {
            popupSvc.smartMessageBox("确定踢线？","踢线成功","踢线失败");
        }


        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            }

            httpSvc.get("/A0001.json",postData).then(function (response) {
                $scope.paginationConf.totalItems = response.length;
                $scope.list = response;
            })

        }

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: 20,
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);



    });
