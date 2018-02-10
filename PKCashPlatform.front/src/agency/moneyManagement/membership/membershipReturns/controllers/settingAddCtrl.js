angular.module('app.membershipReturns').controller('settingAddCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService,$stateParams){
    $scope.discountID = $stateParams.discountID;
    $scope.sitId = function (site_index_id) {
        MembershipReturnsService.getSiteSelect(site_index_id).then(function (response) {
            console.log(response);
            $scope.names = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        $scope.sitId(user.site_index_id);
    }
    console.log($scope.discountID);
    var GetAllEmployee = function () {
        var postData = {
            id:$scope.discountID
        }
        MembershipReturnsService.getRebateSetAddAll(postData).then(function (response) {
            $scope.list = response.data;
            $scope.buy = response.money;
            $scope.sale = response.upper_limit;
            console.log($scope.list);
        })
    };
     GetAllEmployee();
    //点击提交
    $scope.sub = function () {
        console.log($scope.site_index_id);
        console.log($scope.money);
        var check_val = [];
        var test  = document.getElementsByClassName('test');
        for (var j=0;j<test.length;j++){
            var obj ={
                id:test[j].value*1,
                proportion:$(test[j]).parent().parent().find('.inputse')[0].value/100
            };
            check_val.push(obj);
        };
        console.log(check_val);

        var postData = {
            data:check_val,
            site_index_id:$scope.site_index_id,
            upper_limit:$scope.sale*1,
            money:$scope.buy*1
        }
        MembershipReturnsService.getRebateSetSubmit(postData).then(function (response) {
            console.log(response);
            if(response.data===null){
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        })
    };

});

