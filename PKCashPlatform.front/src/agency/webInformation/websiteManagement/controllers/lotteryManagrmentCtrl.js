angular.module('app.websiteManagement').controller('lotteryManagrmentCtrl',
    function ($scope, popupSvc, lotteryManagrmentService, $rootScope, APP_CONFIG) {
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
        lotteryManagrmentService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        var GetAllEmployee = function () {
            var postData = {
                site: $scope.site,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
            };
            lotteryManagrmentService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
            });
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.search = function () {
            GetAllEmployee();
        };

        //图片预览
        $scope.see = function (img) {
            $scope.img_src = img;
        };

        $scope.upload = function (id) {
            $scope.id = id;
            lotteryManagrmentService.getEnclosure({
                id: $scope.id
            }).then(function (response) {
                $scope.enclosure = response.data;
            });
        };

        $scope.select = function (item) {
            $scope.title = item.title;
            $scope.url = item.url;
            $scope.enclosure_id = item.id;
        };
        $scope.modifyTitle = function (item) {
            lotteryManagrmentService.modifyEnclosure({
                id: item.id,
                title: item.title
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };
        $scope.delete = function (item) {
            lotteryManagrmentService.deleteEnclosure({
                id: item.id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    lotteryManagrmentService.getEnclosure({
                        id: $scope.id
                    }).then(function (response) {
                        $scope.enclosure = response.data;
                    });
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.submit = function () {
            lotteryManagrmentService.selectEnclosure({
                id: $scope.id,
                enclosure_id: $scope.enclosure_id
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };

        $scope.open = function (item) {
            $scope.order = item.order;
            $scope.new_order = item.order;
            $scope.id = item.id;
            $scope.name = item.name;
        };
        $scope.modify = function () {
            lotteryManagrmentService.modifyOrder({
                order: $scope.order,
                id: $scope.id,
                new_order: $scope.new_order
            }).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        };


    });