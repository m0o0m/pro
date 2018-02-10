    angular.module('app.AccessMoney').controller('DepositCtrl',
    function(httpSvc,popupSvc,$scope,AccessMoneyService,APP_CONFIG,$rootScope){
        //获取存款项目
        $scope.option_deposit = APP_CONFIG.option.option_deposit_item;
        //获取层级
        AccessMoneyService.getLevel().then(function (response) {
            console.log(response);
            $scope.levelList = response.data;
        })
        $scope.search=function(){
            var accountkey = $scope.accountkey
            AccessMoneyService.memberSearch(accountkey).then(function (response) {
                $scope.data= angular.copy(response.data[0]);
            });
        };

        $scope.commit=function(){
            var val=$("#account").val();
            $scope.account = val.split('、');
            var checkbox1 = document.getElementById("checkbox1");
            if(checkbox1.checked){
                $scope.checkbox1 =1
            }else {
                $scope.checkbox1 =0
            }
            var checkbox2 = document.getElementById("checkbox2");
            if(checkbox2.checked){
                $scope.checkbox2 =1
            }else {
                $scope.checkbox2 =0
            }
            var checkbox3 = document.getElementById("checkbox3");
            if(checkbox3.checked){
                $scope.checkbox3 =1
            }else {
                $scope.checkbox3 =0
            }
            var checkbox4 = document.getElementById("checkbox4");
            if(checkbox4.checked){
                $scope.checkbox4 =1
            }else {
                $scope.checkbox4 =0
            }

            var postData = {
                types: 1,
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
            AccessMoneyService.manualAccessBatch(postData).then(function (response) {
                if (response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };

        $scope.commit1=function(){
            var checkbox5 = document.getElementById("checkbox5");
            if(checkbox5.checked){
                $scope.checkbox5 =1
            }else {
                $scope.checkbox5 =0
            }
            var checkbox6 = document.getElementById("checkbox6");
            if(checkbox6.checked){
                $scope.checkbox6 =1
            }else {
                $scope.checkbox6 =0
            }
            var checkbox7 = document.getElementById("checkbox7");
            if(checkbox7.checked){
                $scope.checkbox7 =1
            }else {
                $scope.checkbox7 =0
            }
            var checkbox8 = document.getElementById("checkbox8");
            if(checkbox8.checked){
                $scope.checkbox8 =1
            }else {
                $scope.checkbox8 =0
            }

            var level_id = [];
            level_id.push($scope.level_id);
            var postData = {
                types: 2,
                level_id: level_id,
                money: $scope.deposit_amount1*1,
                is_deposit_discount:$scope.checkbox5*1,
                deposit_discount: $scope.deposit_preference1*1,
                is_remit_discount:$scope.checkbox6*1,
                remit_discount:$scope.remittance_discount1*1,
                is_code_count:$scope.checkbox7*1,
                code_count:$scope.audit1*1,
                is_routine_check:$scope.checkbox8*1,
                deposit_type:$scope.deposit_type1*1,
                is_write_rebate:$scope.is_write_rebate1*1,
                remark:$scope.remark1
            };
            AccessMoneyService.manualAccessBatch(postData).then(function (response) {
                if (response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };


        $scope.reset=function(){
            $scope.data={};
        };
        //添加到账号
        $scope.addAcount=function(){
            var val=$("#account").val();
            if(val.indexOf($scope.key)==-1){
                if(val==''){
                    val=$scope.key;
                }else{
                    val=val+"、"+$scope.key;
                }
            }
            $("#account").val(val);
        };
    });