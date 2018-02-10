angular.module('app.AccessMoney').controller('QuotaCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope){
        $scope.search=function(){
            var accountkey = $scope.accountkey
            AccessMoneyService.memberSearch(accountkey).then(function (response) {
                $scope.data= angular.copy(response.data[0]);
            });
            AccessMoneyService.typeList().then(function (response) {
                console.log(response.data.data);
                $scope.listData = response.data.data;
            })
        };

        //获取转出项目金额
        $scope.zcmoney = function () {
            AccessMoneyService.memberBlance($scope.data.account,$scope.for_type*1).then(function (response) {
                 $scope.yueMoney = response.data.data[0].balance;
            });
            //获取转入转出项目
            AccessMoneyService.typeList().then(function (response) {
                console.log(response.data.data);
                $scope.listData = response.data.data;
            })
        };

        //确定提交
        $scope.commit=function(){
            var postdata = {
                account:$scope.data.account,
                money:$scope.data.transfer_amount*1,
                from_type:$scope.from_type*1,
                for_type:$scope.for_type*1,
                remark:$scope.data.remark,
                do_user_type:1
            };
            AccessMoneyService.quotqSumbit(postdata).then(function (response) {
                    if(response == null){
                        popupSvc.smallBox("success",$rootScope.getWord("success"));
                    }else {
                        popupSvc.smallBox("fail", response.msg);
                    }
            });
        };

        $scope.reset=function(){
            popupSvc.smallBox("success",$rootScope.getWord("success"))
        }
    });