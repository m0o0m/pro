angular.module('app.membershipReturns').controller('TranslationResultCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService,$stateParams){
    $scope.site = $stateParams.site;
    $scope.startTime = $stateParams.startTime;
    $scope.endTime = $stateParams.endTime;
    $scope.condition = $stateParams.condition;

    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$scope.site,
            start_time:$scope.startTime,
            end_time:$scope.endTime,
            is_rebate:$scope.condition
        };
        MembershipReturnsService.getUserRebateSearch(postData).then(function (response) {
            $scope.list = response.list;
            $scope.name = response.product_name;
            $scope.total = response.total;
        })
    };

    GetAllEmployee();

    $scope.doallcheck = function(){
        var allche = document.getElementById("all");
        var che = document.getElementsByClassName("test");
        if(allche.checked == true){
            for(var i=0;i<che.length;i++){
                che[i].checked="checked";
            }
        }else{
            for(var i=0;i<che.length;i++){
                che[i].checked=false;
            }
        }
    };
    $scope.sub = function () {
        var check_val = [];
        var test  = document.getElementsByClassName('test');
        for (var j=0;j<test.length;j++){
            if(test[j].checked) {
                var obj ={
                    id:test[j].value*1
                };
                check_val.push(obj);
            }
        }
        var postData = {
            deposit_id:check_val,
            deposit:$scope.code
        };
        MembershipReturnsService.getUserRebateDeposit(postData).then(function (response) {
            if(response.data==null){
                GetAllEmployee();
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };
});
