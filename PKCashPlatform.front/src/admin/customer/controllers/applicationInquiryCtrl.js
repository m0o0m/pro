
angular.module('app.customer').controller('applicationInquiryCtrl',
    function ($scope, popupSvc, commonService, customerApplicationInquiryService, $rootScope, APP_CONFIG) {
        //获取站点
        commonService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });

        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
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

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                indexId: $scope.indexId,
                status: $scope.status,
                account: $scope.account,
                startTime: $scope.startTime,
                endTime: $scope.endTime,
            };
            customerApplicationInquiryService.getList(postData).then(function (response) {
                $scope.list = response.data;
                $scope.paginationConf.totalItems = response.meta.count;
            })
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search=function(){
            GetAllEmployee()
        }


        $scope.detail=function(item){
            $scope.detailData=item;
        }

    });