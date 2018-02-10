angular.module('app.memberMessage').controller('memberMessageCtrl',
function(httpSvc, popupSvc, memberMessageService, $stateParams, DTColumnBuilder, $http, $scope, $rootScope, APP_CONFIG, $state, $LocalStorage){
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };

    //获取代理
    $scope.sitId = function (site_index_id) {
        memberMessageService.getDropSelect(site_index_id).then(function (response) {
            $scope.sharedJson = response.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        console.log($scope.siteId);
        $scope.sitId(user.site_index_id);
    }

    //获取类型下拉
    $scope.typeSelect = APP_CONFIG.option.select_by;
    console.log($scope.typeSelect);

    var GetAllEmployee = function () {
         var postData = {
             page: $scope.paginationConf.currentPage,
             page_size: $scope.paginationConf.itemsPerPage,
             site:$scope.site,
             type:$scope.type,
             account:$scope.account,
             start_time:$scope.start_time,
             end_time:$scope.end_time
         };
        memberMessageService.setMemberNews(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.search = function () {
        GetAllEmployee();
    };
    $scope.delete = function (id) {
        var postData = {
            id:id
        };
        var del = function () {
            memberMessageService.deleteNews(postData).then(function (response) {
                if (response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    GetAllEmployee();
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
    }
});