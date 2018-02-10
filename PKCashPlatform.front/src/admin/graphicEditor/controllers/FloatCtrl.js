angular.module('app.GraphicEditor').controller('FloatCtrl',
    function ($scope, popupSvc, siteService, floatService, $rootScope, APP_CONFIG) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        var GetAllEmployee = function () {
            floatService.getList({
                site: $scope.site
            }).then(function (response) {
                $scope.list = response.data;
            });
        };
        GetAllEmployee();

        //搜索
        $scope.search = function () {
            GetAllEmployee();
        };

        $scope.modify = function (item) {
            $scope.formData = item;
        };
        //修改提交
        $scope.submit = function () {
            floatService.modify($scope.formData).then(function (response) {
                if (response === null) {
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    GetAllEmployee();
                } else {
                    popupSvc.smallBox("fail", response.msg);
                }
            });
        }
    });