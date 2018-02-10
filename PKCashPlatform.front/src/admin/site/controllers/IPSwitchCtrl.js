/**
 * Created by apple on 17/11/20.
 */
angular.module('app.site').controller('IPSwitchCtrl',
    function(httpSvc,popupSvc,$http,$scope,$rootScope,$compile,APP_CONFIG,siteService){
    //获取ip
    $scope.option_ip = APP_CONFIG.option.option_ip;

    var GetAllEmployee = function () {
        var postData = {
            start_ip:$scope.start_ip,
            end_ip: $scope.end_ip,
        };
        siteService.IPSwitching(postData).then(function (response) {
                $scope.paginationConf.totalItems = response.data.meta.count;
                $scope.list = response.data.data;
                console.log(response.data.meta.count);
        });

    }
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
    };

    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    //添加ip操作
    $scope.add = function () {
        var postData={
              ip_id:$scope.ip_id,
              content:$scope.content,
              status:$scope.status
        };
        siteService.IPSwitchAdd(postData).then(function (data) {
            if(data.data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
            }else {
                popupSvc.smallBox("fail",data.msg);
            };
        });
    };
    //获取详情
    $scope.getID = function (item) {
        $scope.data = item;
        console.log($scope.data);
        $scope.data.id.toString();
        console.log($scope.data.id);
    };
    //修改后提交
    $scope.modify = function () {
        siteService.IPSwitchModify($scope.data).then(function (data) {
            if(data.data==null){
                popupSvc.smallBox("success",$rootScope.getWord('success'));
            }else {
                popupSvc.smallBox("fail",data.msg);
            };
        });
    };



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