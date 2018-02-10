angular.module('app.membershipReturns').controller('MemberPromotionCtrl', function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,$state,$LocalStorage,MembershipReturnsService){
    $scope.sitId = function (site_index_id) {
        MembershipReturnsService.getSiteSelect(site_index_id).then(function (response) {
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
        if($scope.register_ip == undefined){
            $scope.register_ip = "";
        }
        if($scope.site_index_id == undefined){
            $scope.site_index_id = "";
        }
        if($scope.spread_id == undefined){
            $scope.spread_id = "";
        }
        if($scope.account == undefined){
            $scope.account = "";
        }
        var postData = {
            site_index_id: $scope.site_index_id,
            account: $scope.account,
            register_ip: $scope.register_ip,
            spread_id: $scope.spread_id
        };;
        MembershipReturnsService.getSpreadInfo(postData).then(function (response) {
            console.log(response);
            $scope.list = response.data;
            $scope.paginationConf.totalItems = response.meta.count;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: 20
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    $scope.search = function () {
        GetAllEmployee();
    };

    $scope.user = function (id) {
        $state.go("app.MembershipReturns.numDetail",{id:id})
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
