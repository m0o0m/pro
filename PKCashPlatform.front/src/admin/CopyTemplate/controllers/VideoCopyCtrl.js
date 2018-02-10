angular.module('app.CopyTemplate').controller('VideoCopyCtrl', function ($scope,$state,httpSvc,popupSvc,APP_CONFIG,copyTemplateService,$rootScope){

    var GetAllEmployee = function () {
        copyTemplateService.videoCopyTemplate().then(function (response) {
            $scope.list = response.data.data;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    console.log($scope.paginationConf.itemsPerPage);
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

    //修改获取
    $scope.modify = function (id,$index) {
        $scope.status_m = $scope.list[$index].status;
        $scope.name_m = $scope.list[$index].name;
        $scope.number_m = $scope.list[$index].num;
        $scope.id = $scope.list[$index].id;
        //修改提交
        $scope.submit_m = function (item) {
            var status=$('input:radio[name="bb"]:checked').val();
            var postData = {
                id:$scope.id*1,
                num:$scope.number_m*1,
                name:$scope.name_m,
                status:status*1
            };
            copyTemplateService.putCpyTemplateStatus(postData).then(function (response) {
                if (response) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            });
        };
    };
    $scope.submit_a = function () {
        var status_a=$('input:radio[name="aa"]:checked').val();
        var postData = {
            name:$scope.name,
            num:$scope.number*1,
            status:status_a*1
        };
        copyTemplateService.postCpyTemplateStatus(postData).then(function (response) {
            if (response) {
                popupSvc.smallBox("success", $rootScope.getWord("success"));
            } else {
                popupSvc.smallBox("fail", response.data.msg);
            }
        });
    }
});