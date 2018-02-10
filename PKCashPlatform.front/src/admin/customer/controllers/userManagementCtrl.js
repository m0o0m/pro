angular.module('app.customer').controller('userManagementCtrl',
    function ($scope, $state, httpSvc, APP_CONFIG, popupSvc,customerService,$rootScope) {
        //获取站点
        $scope.siteId = function () {
            customerService.getSite().then(function (response) {
                $scope.sharedJson = response.data;
            });
        };
        $scope.siteId();

        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };

        $scope.statusJson = APP_CONFIG.option.option_status;
        $scope.accountTYpeJson = APP_CONFIG.option.option_account_type;
        $scope.account_type = '1';
        $scope.equipmentJson = APP_CONFIG.option.option_reg;

        var GetAllEmployee = function () {
            var postData = {
                site:$scope.site,
                arr:$scope.moreSite,
                account:$scope.account,
                login_ip:$scope.loginIp,
                type:$scope.type,
                start_time:$scope.startTime,
                end_time:$scope.endTime,
                account_type:$scope.account_type,
                equipment:$scope.equipment,
                pageIndex: $scope.paginationConf.currentPage,
                pageSize: $scope.paginationConf.itemsPerPage
            };
            customerService.getUserList(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data;
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
        }

        // 停用
        $scope.disable = function (item) {
            var status = 2;
            if (item.status === 0 || item.status === 1) {
                status = 2;
            } else {
                status = 1;
            }
            //1正常2禁用
            var sure = function () {
                customerService.setMemberStatus(item.id, status).then(function (response) {
                    if (response) {
                        item.status = status;
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                });
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation") + "?", sure);
        };

        //会员资料
        $scope.Info = function (id) {
            customerService.setuserInforment(id).then(function (response) {
                $scope.account_name = response.data.data[0].account;
                $scope.username = response.data.data[0].username;
                $scope.userdata = response.data.data[0].userdata;
                $scope.country = response.data.data[0].country;
                $scope.bank = response.data.data[0].bank;
                $scope.bank_account = response.data.data[0].bank_account;
                $scope.password = response.data.data[0].password;
                $scope.address = response.data.data[0].address;
                $scope.remarks = response.data.data[0].remarks;
                $scope.registered_ip = response.data.data[0].registered_ip;
                $scope.registered_time = response.data.data[0].registered_time;
                $scope.last_time = response.data.data[0].last_time;
                $scope.last_ip = response.data.data[0].last_ip;
            });
        }
    });