angular.module('app.membershipReturns').controller('settingCtrl',  function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService){
    $scope.sitId = function (site_index_id) {
        MembershipReturnsService.getSiteSelect(site_index_id).then(function (response) {
            $scope.sharedJson = response.data.data;
        });
    };
    var user = JSON.parse($LocalStorage.getItem("user"));
    $scope.isSuperAdmin = user.site_index_id === '';
    if ($scope.isSuperAdmin === false) {
        //获取全部站点
        $scope.sitId();
    } else {
        console.log($scope.siteId);
        $scope.sitId(user.site_index_id);
    }
    $scope.json = APP_CONFIG.option;
    var GetAllEmployee = function () {
        $scope.year=new Date().getFullYear();
        console.log($scope.year);
        console.log($scope.yy);
        console.log($scope.mm);
        console.log($scope.site_index_id);
        if($scope.yy==undefined){
            $scope.yy = "";
        }
        if($scope.mm==undefined){
            $scope.mm = "";
        }
        if($scope.site_index_id==undefined){
            $scope.site_index_id = "";
        }
        var postData = {
            month:$scope.mm,
            site_index_id:$scope.site_index_id,
            year:$scope.yy,
        }
        MembershipReturnsService.getRebateList(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data.data;
        })
    }
    GetAllEmployee();

    $scope.search = function () {
        GetAllEmployee();
    }

    $scope.detailed = function (id) {
        console.log(id);
        $state.go("app.MembershipReturns.Detail",{detailedID:id})
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
