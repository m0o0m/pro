angular.module('app.customer').controller('operationLogCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,APP_CONFIG,customerService){
    $scope.toggleAdd = function () {
        if (!$scope.newTodo) {
            $scope.newTodo = {
                state: 'Important'
            };
        } else {
            $scope.newTodo = undefined;
        }
    };

    //获取站点
    $scope.siteId = function () {
        customerService.getSite().then(function (response) {
            $scope.siteJson  = response.data;
        });
    };
    $scope.siteId();

    $scope.loginsKeyJson = APP_CONFIG.option.option_login_id;
    $scope.key = '1';
    //操作日志
    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.appsite,
            key:$scope.key,
            strat_time:$scope.startTime,
            end_time:$scope.endTime,
            value:$scope.info,
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        };
        customerService.getSiteLog(postData).then(function (response) {
            $scope.list = response.data;
            $scope.paginationConf.totalItems = response.meta.count;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.logSearch = function () {
        GetAllEmployees();
    };


    $scope.keyNameJson = APP_CONFIG.option.option_login_id;
    $scope.keyName = '1';
    $scope.deviceNameJson = APP_CONFIG.option.option_reg;
    $scope.deviceName = '1';
    //自动稽核
    var GetAllEmployeeAuto = function () {
        var postData = {
            site_id:$scope.siteName,
            key:$scope.keyName,
            value:$scope.valueName,
            device:$scope.deviceName,
            pageIndex: $scope.paginationConf.currentPage,
            pageSize: $scope.paginationConf.itemsPerPage
        };
        customerService.getAutoAudit(postData).then(function (response) {
            $scope.autoList = response.data;
            $scope.paginationConf.totalItems = response.meta.count;
        })
    };
    $scope.paginationConfAuto = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployeeAuto);

    $scope.autoSearch = function () {
        GetAllEmployeeAuto();
    }

});

