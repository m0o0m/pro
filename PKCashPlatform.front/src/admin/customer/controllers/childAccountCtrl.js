angular.module('app.customer').controller('childAccountCtrl',
    function(httpSvc,popupSvc,$scope,APP_CONFIG,customerService,$rootScope){

        $scope.option_stateJson = APP_CONFIG.option.option_state;
        $scope.lineStat = '-1';
        $scope.option_typrJson = APP_CONFIG.option.option_name_type;
        $scope.type = '1';
        $scope.option_accountJson = APP_CONFIG.option.option_accounts_type;
        $scope.accountType = '1';
        //获取站点
        $scope.siteId = function () {
            customerService.getSite().then(function (response) {
                $scope.siteJson  = response.data;
            });
        };
        $scope.siteId();

        var GetAllEmployee = function () {
            var postData = {
                site_id: $scope.site_index_id,
                start_time:$scope.startTime,
                end_time:$scope.endTime,
                is_login: $scope.lineState,
                type:$scope.type,
                name: $scope.searchKey,
                account:$scope.accountType,
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage
            };
            customerService.getChilderList(postData).then(function (response) {
                $scope.list = response.data;
                $scope.paginationConf.totalItems = response.meta.count;
            })
        };

        //分页初始化
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
        };

        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        //点击修改时获取ID;
        $scope.getID = function (item) {
            console.log(item);
            $scope.account = item.account;
            $scope.id = item.id;
            $scope.username = item.username;
            $scope.sitename = item.site_index_id;
        };
        //修改后提交
        $scope.submit = function () {
            var postData = {
                site_index_id: $scope.sitename,
                username:$scope.username,
                id:$scope.id,
                account: $scope.account,
                password:$scope.password,
                confirmPassword: $scope.confirm_password
            };
            customerService.getChilderPut(postData).then(function (response) {
                if (response.data.data === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        //点击搜索
        $scope.search = function () {
            GetAllEmployee()
        };
        // 停用
        $scope.disable=function (item) {
            var able = function () {
                var postData = {
                    id:item.id,
                    site_index_id:item.site_index_id,
                    status:item.status
                };
                customerService.getChilderStatus(postData).then(function (response) {
                    if (response.data === null) {
                        if(item.status==1){
                            item.status = 2;
                        }else{
                            item.status = 1;
                        }
                        popupSvc.smallBox("success", $rootScope.getWord("success"));
                    } else {
                        popupSvc.smallBox("fail", response.data.msg);
                    }
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"), able);
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
