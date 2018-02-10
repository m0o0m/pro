angular.module('app.CommissionStatistics').controller('AgentSettingCtrl',
    function ($scope, popupSvc, siteService, agentSettingService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
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
                site_id: $scope.site_id,
                site_index_id: $scope.site_index_id
            }
            agentSettingService.getList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.arr;
                $scope.bet_amount = response.data.bet_money;
            });
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        //搜索
        $scope.search = function () {
            GetAllEmployee();
        };

        $scope.del = function (id) {
            var fn=function() {
                agentSettingService.del({
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
            popupSvc.smartMessageBox($rootScope.getWord("confirmStorage")+"？", fn);
        };

        $scope.setting = function () {
            $scope.is_show = true;
        }
    });

