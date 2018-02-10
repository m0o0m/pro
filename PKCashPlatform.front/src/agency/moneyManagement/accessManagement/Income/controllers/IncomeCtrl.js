angular.module('app.Income').controller('IncomeCtrl', function (httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope,$interval) {
        //初始化下拉框
        $scope.option_handle = APP_CONFIG.option.option_handle;
        $scope.shebei_opition = APP_CONFIG.option.shebei;
        $scope.option_query_criteria = APP_CONFIG.option.option_query_criteria;
        $scope.refresh_time_opition=APP_CONFIG.option.refresh_time;
        $scope.option_discount = APP_CONFIG.option.option_discount;


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
            $scope.siteJson = response.data;
            console.log($scope.site);
         });



        //获取层级
        AccessMoneyService.getLevel().then(function (response) {
            console.log(response);
            $scope.levelList = response.data;
        })

        //获取代理账号
        AccessMoneyService.getAgencySelect().then(function(response){
            $scope.agencyList = response.data.data;
    })

        //获取收款账号


        var GetAllEmployee = function () {
            var postData = {
                page: $scope.paginationConf.currentPage,
                page_size: $scope.paginationConf.itemsPerPage,
                site_index_id: $scope.site_index_id,
                agency_account: $scope.agency_account,
                level: $scope.level,
                status: $scope.status*1,
                start_time: $scope.start_time,
                end_time: $scope.end_time,
                upper_limit: $scope.upper_limit,
                lower_limit: $scope.lower_limit,
                client_type: $scope.client_type*1,
                payment_account: $scope.payment_account,
                select_by: $scope.select_by,
                conditions: $scope.conditions,
                is_discount: $scope.is_discount*1
            };
            AccessMoneyService.companyIncome(postData).then(function (response) {
                    if(!response.code){
                        $scope.paginationConf.totalItems = response.data[0].total_count;
                        $scope.subtotal=response.data[0];
                        $scope.list = response.data[0].AllData;
                        console.log(response.data[0]);
                    }else{
                        $scope.paginationConf.totalItems = 0;
                        $scope.subtotal = null;
                        $scope.list = null;
                    }
            });


        };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.sure=function(id){
            var confirm = function () {
                AccessMoneyService.confirmMoney(id).then(function (response) {
                    if(response.data==null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                });

            };
            popupSvc.smartMessageBox($rootScope.getWord('confirmationOperation'),confirm);
        };

        $scope.cancle=function(id){
            $scope.cancle_id = id;
            var del = function () {
                AccessMoneyService.cancleMoney($scope.cancle_id).then(function (response) {
                    if(response==null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));
                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                })
            }
            popupSvc.smartMessageBox($rootScope.getWord('confirmationOperation'),del);
        }
        
        $scope.search = function () {
            GetAllEmployee();
        }

        $scope.refresh = $interval(function(){
            console.log(11111111111)
        },100000000)
        $interval.cancel($scope.refresh);

        $scope.onChange=function(){
            $interval.cancel($scope.refresh);
            $scope.refresh = $interval(function(){
                GetAllEmployee();
            },$scope.refresh_time)
        }

    });
