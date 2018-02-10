angular.module('app.site').controller('informationHistoryCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
        //获取站点
        siteService.thirdDropf().then(function (response) {
            $scope.siteJson = response.data.data;
        });
        //获取下拉框类型选择
        $scope.information = APP_CONFIG.option.option_information_type;
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };


        var GetAllEmployee = function () {
            var postData = {
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            };
            siteService.copywaritingHistory(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.list;
            });

        }
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);


        //删除
        $scope.delete=function (ids) {
            var del = function () {

            }
            popupSvc.smartMessageBox("确定删除？",del);
        };


    });