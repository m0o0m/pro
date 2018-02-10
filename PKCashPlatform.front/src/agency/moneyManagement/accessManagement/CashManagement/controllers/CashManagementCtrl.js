angular.module('app.CashManagement').controller('CashManagementCtrl',
    function ($interval, AccessMoneyService,httpSvc, popupSvc, $stateParams, $scope, $rootScope, APP_CONFIG, $state, $LocalStorage
    ) {
        $scope.user = JSON.parse($LocalStorage.getItem("user"));
        $scope.toggleAdd = function () {
            if (!$scope.newTodo) {
                $scope.newTodo = {
                    state: 'Important'
                };
            } else {
                $scope.newTodo = undefined;
            }
        };
        //获取站点下拉框
        AccessMoneyService.getDropSelect().then(function (response) {
            $scope.site = response.data;
            console.log($scope.site);
        });
       //下拉框初始化
        $scope.shebei_opition = APP_CONFIG.option.shebei;
        $scope.money_status_opition = APP_CONFIG.option.money_status;
        $scope.automatic_opition = APP_CONFIG.option.automatic;
        $scope.select_by_opition = APP_CONFIG.option.select_by;
        $scope.refresh_time_opition=APP_CONFIG.option.refresh_time;

        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id,
                level: $scope.level,
                out_status: $scope.out_status*1,
                start_time: $scope.start_time,
                end_time: $scope.end_time,
                upper_limit: $scope.upper_limit,
                lower_limit: $scope.lower_limit,
                client_type: $scope.client_type*1,
                select_by: $scope.select_by,
                conditions: $scope.conditions,
                automatic: $scope.automatic*1
            };
            AccessMoneyService.getMoney(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.total_count;
                if(!response.code){
                    $scope.paginationConf.totalItems = response.data.total_count;
                    $scope.subtotal=response.data;
                    $scope.list = response.data.OutMoneyList;
                }else{
                    $scope.paginationConf.totalItems = 0;
                    $scope.subtotal = null;
                    $scope.list = null;
                }
            })
        };
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: 10
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
        $scope.search = function () {
            GetAllEmployee();
        };
        $scope.reason=function(reason){
            var del = function () {

            };
            popupSvc.smartMessageBox(reason,del);
        };
        //预备出款
        $scope.out=function(id,agency_id){
            var del = function () {
                httpSvc.put("/prepare_out",{
                    id: id
                    //agency_id: agency_id
                }).then(function (response) {
                    if(response === null){
                        GetAllEmployee();
                        popupSvc.smallBox("success",$rootScope.getWord("success"))
                    }else {
                        popupSvc.smallBox("fail",response.msg)
                    }
                }, function (data) {
                    popupSvc.smallBox("fail",data.msg)
                })
            };
            popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
        };

        //确认出款
        $scope.sure1=function(id,agency_id,member_account,outward_money){
            $scope.sure_id=id;
            $scope.sure_agency_id=agency_id;
            $scope.memberAccount=member_account;
            $scope.outMoney=outward_money;
        };
        //拒绝出款
        $scope.refuse1=function(id,agency_id){
            $scope.refuse_id=id;
            $scope.refuse_agency_id=agency_id;

        };
        //取消出款
        $scope.cancle1=function(id,agency_id){
            $scope.cancle_id=id;
            $scope.cancle_agency_id=agency_id;

        };

        //确认出款
        $scope.sure=function(){
            AccessMoneyService.confirmMoney($scope.sure_id,$scope.sure_agency_id).then(function (response) {
                if(response.data===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };

        //拒绝出款
        $scope.refuse=function(){
                var reason = $scope.refuseReason;
            AccessMoneyService.refuseMoney($scope.refuse_id, $scope.refuse_agency_id,reason).then(function (response) {
                if(response.data===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        };

        //取消出款
        $scope.cancle=function(){
            AccessMoneyService.cancleMoney($scope.cancle_id).then(function (response) {
                if(response==null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        };
        //确认
        $scope.refresh = $interval(function(){
        },100000000);
        $interval.cancel($scope.refresh);
        $scope.onChange=function(){
            $interval.cancel($scope.refresh);
            $scope.refresh = $interval(function(){
                GetAllEmployee();
            },$scope.refresh_time)
        };
    });
