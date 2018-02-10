/**
 * Created by apple on 17/11/21.
 */
angular.module('app.customer').controller('hierarchicalManagCtrl',
    function(httpSvc,popupSvc,$scope,APP_CONFIG,customerService){
        //获取站点
        $scope.siteId = function () {
            customerService.getSite().then(function (response) {
                $scope.siteJson  = response.data;
            });
        };
        $scope.siteId();
        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_id: $scope.site_index_id,
                site_name:$scope.site_name
            };

            customerService.getHierarchicalList(postData).then(function (response) {
                $scope.list = response.data;
                $scope.paginationConf.totalItems = response.meta.count;
            })
        };

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        //点击搜索
        $scope.search = function () {
            GetAllEmployee();
        };
        //筛选展开
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };

    });
