angular.module('app.membershipReturns').controller('DiscountsCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService){
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
    var GetAllEmployee = function () {
        if($scope.site_index_id == undefined){
            $scope.site_index_id = "";
        }
        var postData = {
            site_index_id:$scope.site_index_id
        };

        MembershipReturnsService.getRebateSetGetAll(postData).then(function (response) {
            $scope.list = response.list;
            $scope.name = response.product_name;
        })
    };
    GetAllEmployee();

    $scope.search = function () {
        GetAllEmployee();
    };

    $scope.modify = function (id) {
        $state.go("app.MembershipReturns.settingModify",{discountID:id});
    };
    $scope.add = function () {
        $state.go("app.MembershipReturns.settingAdd");
    };

    $scope.delete = function (id) {
        var postData = {
            id:id
        };
        var del = function () {
            MembershipReturnsService.getRebateSetDel(postData).then(function (response) {
                console.log(response);
                if(response.data===null){
                    GetAllEmployee();
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        };
        popupSvc.smartMessageBox($rootScope.getWord("confirmationOperation"),del);
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

