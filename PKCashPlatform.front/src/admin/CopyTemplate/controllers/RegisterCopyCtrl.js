angular.module('app.CopyTemplate').controller('registerCopyCtrl',
    function($scope,$state,httpSvc,$stateParams,APP_CONFIG,popupSvc,copyTemplateService,$rootScope){
    var GetAllEmployee = function () {
        copyTemplateService.copyTemplate().then(function (response) {
            $scope.list = response.data.data.list;
        })
    };
    $scope.paginationConf = {
        currentPage: 1,
        itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT
    };
    console.log($scope.paginationConf.itemsPerPage);
    $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);
    $scope.add = function () {
        $state.go('app.CopyTemplate.AddRegisterCopy');
    };
    $scope.modify = function (id,$index) {
        console.log($scope.list[$index]);
        $state.go('app.CopyTemplate.ModifyRegisterCopy',{
            id:id,
            data:$scope.list[$index]
        });
    };


    $scope.disable = function (item) {
        var status = 2;
        if (item.status === 0 || item.status === 1) {
            status = 2;
        } else {
            status = 1;
        }
        //1正常2禁用
        var sure = function () {
            copyTemplateService.copyTemplateStatus(item.id, status).then(function (response) {
                if (response) {
                    item.status = status;
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            });
        };
        popupSvc.smartMessageBox($rootScope.getWord("确定更改状态") + "?", sure);
    };

});