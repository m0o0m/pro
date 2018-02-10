angular.module('app.AccessMoney').controller('AccessMoneyCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope){
        //获取存款项目
        $scope.option_deposit = APP_CONFIG.option.option_deposit_item;
        //会员账号搜索
        $scope.search=function(){
            var accountkey = $scope.accountkey;
            AccessMoneyService.memberSearch(accountkey).then(function (response) {
                $scope.data = response.data;
                $scope.account = response.data[0].account;
                $scope.balance = response.data[0].balance;
                $scope.realname = response.data[0].realname;
            });
        };

        $scope.commit=function(){
            var checkbox1 = document.getElementById("checkbox1");
            console.log(checkbox1);
            if(checkbox1.checked){
                $scope.checkbox1 =1;
            }else {
                $scope.checkbox1 =0;
            }
            var checkbox2 = document.getElementById("checkbox2");
            if(checkbox2.checked){
                $scope.checkbox2 =1;
            }else {
                $scope.checkbox2 =0;
            }
            var checkbox3 = document.getElementById("checkbox3");
            if(checkbox3.checked){
                $scope.checkbox3 =1;
            }else {
                $scope.checkbox3 =0;
            }
            var checkbox4 = document.getElementById("checkbox4");
            if(checkbox4.checked){
                $scope.checkbox4 =1;
            }else {
                $scope.checkbox4 =0;
            }

            var postData = {
                account: $scope.account,
                money: $scope.deposit_amount*1,
                is_deposit_discount:$scope.checkbox1*1,
                deposit_discount: $scope.deposit_preference*1,
                is_remit_discount:$scope.checkbox2*1,
                remit_discount:$scope.remittance_discount*1,
                is_code_count:$scope.checkbox3*1,
                code_count:$scope.audit*1,
                is_routine_check:$scope.checkbox4*1,
                deposit_type:$scope.deposit_type*1,
                is_write_rebate:$scope.is_write_rebate*1,
                remark:$scope.remark
            };

            AccessMoneyService.manualAccess(postData).then(function (response) {
                if (response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            })
        }
    });