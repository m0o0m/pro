angular.module('app.AccessMoney').controller('ArtificialWithdrawalCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope){
        //获取存款项目
        $scope.option_deposit = APP_CONFIG.option.option_deposit_item;
        //会员账号搜索
        $scope.search=function(){
            var accountkey = $scope.accountkey
            AccessMoneyService.memberSearch(accountkey).then(function (response) {
                $scope.data = response.data;
                $scope.account = response.data[0].account;
                $scope.balance = response.data[0].balance;
                $scope.realname = response.data[0].realname;
            });
        };
        $scope.commit=function(){
            var postData = {
                account:$scope.account,
                money:$scope.money*1,
                deposit_type:$scope.deposit_type*1,
                remark: $scope.remark
            };
            console.log(postData);
            AccessMoneyService.manualWithdrawal(postData).then(function (response) {
                if (response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };
    });