angular.module('app.GraphicEditor').controller('AdvertisementCtrl',
    function ($scope, popupSvc, siteService, advertisementService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.thirdDropf().then(function (res) {
            $scope.site = res.data.data;
        });
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
                page_size: $scope.paginationConf.itemsPerPage,
                site: $scope.site,
                name: $scope.name
            };
            advertisementService.getList(postData).then(function (response) {
                //$scope.paginationConf.totalItems = response.data.count;
                $scope.logo = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search=function(){
            GetAllEmployee();
        };
        $scope.sort = function (type,id,order) {
            advertisementService.sort({
                id: id,
                type: type,
                order: order
            }).then(function (response) {

            });
        };

        $scope.disable = function (id, status) {
            advertisementService.disable({
                id: id,
                status: status===1?2:1
            }).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        }
    });