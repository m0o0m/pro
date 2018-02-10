angular.module('app.Precalcula').controller('queryCtrl',
    function($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state,$stateParams){
   //初始化值
    $scope.s = "1";
        if( $scope.s=="1"){
            $scope.bancks = true;
            $scope.bancks_1 = false;
        };
    $scope.change = function (s) {
        if(s=="2"){
            $scope.bancks = false;
            $scope.bancks_1 = true;
        }else {
            $scope.bancks = true;
            $scope.bancks_1 = false;
        }
    }


    var GetAllEmployee = function () {
        var postData = {
            site_index_id:$stateParams.site_index_id,
            rtype:$stateParams.rtype,
            typeed:$stateParams.typeed,
            v_type:$stateParams.v_type,
            start_time:$stateParams.start_time,
            end_time:$stateParams.end_time
        }
        precalculaService.statistics(postData).then(function (response) {
          console.log( response);
                $scope.paginationConf.totalItems = response.list.length;
                $scope.list = response.list;
                $scope.arr = response.arr;
                $scope.toal = response.Total;
                $scope.bi  = $scope.arr.length+5;
        });


    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
   $scope.banck= 1;

    $scope.Deposit = function () {
        var arr = [];
        $('input[name="test"]:checked').each(function(){
            arr.push($(this).val()*1);
        });
        var postData={
            id:arr,
            num:$scope.accounts
        }
        precalculaService.depositDiscouny(postData).then(function (data) {
            if(data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));

            }else {
                popupSvc.smallBox("fail",data.msg);
            }
        });
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
