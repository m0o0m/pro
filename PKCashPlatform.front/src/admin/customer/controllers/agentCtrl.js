/**
 * Created by apple on 17/11/21.
 */
/**
 * Created by apple on 17/11/16.
 */
angular.module('app.customer').controller('agentCtrl',
    function(httpSvc,popupSvc,$stateParams,$http,$scope,APP_CONFIG,$state,customerService){

    //获取站点
    $scope.siteId = function () {
        customerService.getSite().then(function (response) {
            $scope.sharedJson = response.data;
        });
    };
    $scope.siteId();

    //获取子站点
    $scope.childernSiteId = function () {
        customerService.getChildernSite().then(function (response) {
            $scope.childernJson = response.data;
        });
    };
    $scope.childernSiteId();

    $scope.stateJson = APP_CONFIG.option.option_status;
    $scope.hierarchyJson = APP_CONFIG.option.option_hierarchy;
    $scope.agent = '1';
    $scope.extensionJson = APP_CONFIG.option.option_extension;
    $scope.type = '1';

    var GetAllEmployee = function () {
        var postData = {
            site_id:$scope.site_id,
            site_index_id:$scope.site_index_id,
            status:$scope.statusd,
            type:$scope.agent,
            spread_id:$scope.type,
            name:$scope.name,
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage
        };

        customerService.getProxyList(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.meta.count;
            $scope.list = response.data;
            $scope.arr = response.arr
        })

    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    //搜索
    $scope.search = function () {
        GetAllEmployee();
    };

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
