angular.module('app.notice').controller('systemLogCtrl',
function($scope,$state,httpSvc,APP_CONFIG,popupSvc,noticeService){
    var GetAllEmployee = function () {
         var postData = {
            page: $scope.paginationConf.currentPage,
            page_size: $scope.paginationConf.itemsPerPage,
            key:$scope.key
        };
        noticeService.setSystermNotice(postData).then(function (response) {
            $scope.paginationConf.totalItems = response.data.meta.count;
            $scope.list = response.data.data;
        });
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    //搜索
    $scope.search=function(){
        GetAllEmployee();
    }
});