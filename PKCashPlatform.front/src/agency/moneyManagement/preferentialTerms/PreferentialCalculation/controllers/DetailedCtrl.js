angular.module('app.Precalcula').controller('DetailedCtrl',
    function ($scope, popupSvc, siteService, preferentialQuiryService, $rootScope, APP_CONFIG,$stateParams) {
        $scope.id = $stateParams.id;

        var GetAllEmployee = function () {
            preferentialQuiryService.getDetail({
                id: $stateParams.id
            }).then(function (response) {
                $scope.paginationConf.totalItems = response.meta.count;
                $scope.list = response.data.list;
                $scope.arr = response.data.Arr;
                $scope.total = response.data.bet_money;
                $scope.bi  = $scope.arr.length+5;
            });

        };
        GetAllEmployee();
        $scope.paginationConf = {
            currentPage: 1,
            itemsPerPage: APP_CONFIG.PAGE_SIZE_DEFAULT,
        };
        $scope.$watch('paginationConf.currentPage + paginationConf.itemsPerPage', GetAllEmployee);

        $scope.vaules = function (val) {
            if (val == "") {
                $scope.isture = true;
                $scope.isture_1 = false;
            } else {
                $scope.isture = false;
                $scope.isture_1 = true;
            }
        };
        //全选
        var sels = document.getElementsByClassName('selected');
        $scope.all = function (m) {
            for (var i = 0; i < sels.length; i++) {
                if (m === true) {
                    sels[i].checked = true;
                } else {
                    sels[i].checked = false;
                }
            }
        };
        //点击冲销

        $scope.sub = function () {
            var arr = [];
            $('input[name="test"]:checked').each(function () {
                arr.push($(this).val() * 1);
            });
            preferentialQuiryService.retreatWaterEdit(arr).then(function (response) {
                if(response==null){
                    popupSvc.smallBox("success",$rootScope.getWord('success'));

                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });
        };


    });
