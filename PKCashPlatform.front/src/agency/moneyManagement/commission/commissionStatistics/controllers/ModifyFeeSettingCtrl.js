angular.module('app.CommissionStatistics').controller('ModifyFeeSettingCtrl',
    function ($scope, popupSvc, siteService, precalculaService, $rootScope, APP_CONFIG,$state,$stateParams) {
        //获取站点
        siteService.getSite().then(function (response) {
            $scope.siteJson = response.data;
        });
        //获取商品
        precalculaService.comboPlatform().then(function (response) {
            $scope.commodity = response.Children;
            console.log($scope.commodity);
        });
        //获取详情
        precalculaService.overrideDetail($stateParams.id).then(function (res) {
           $scope.data = res.data;
            $scope.params = $scope.data.list;
            var arrId =[];
            for (var i =0; i<$scope.params.length;i++){
                arrId.push($scope.params[i].id);
            };
            $scope.isSelected = function (id) {
                return $.inArray(id, arrId)!=-1;
            };
        });

        //修改后提交
        $scope.edit = function () {
            var check_val = [];
            var test  = document.getElementsByClassName('test');
            for (var j=0;j<test.length;j++){
                if(test[j].checked){
                    var parent =$(test[j]).parent()[0];
                    var obj ={
                        platform_id:test[j].value*1,
                        proportion:$(parent).find('.inputse')[0].value*1
                    };
                    check_val.push(obj);
                };
            };
            var postData = {
                site_index_id:$scope.data.site_index_id,
                self_profit:$scope.data.self_profit,
                effective_user:$scope.data.effective_user,
                list:check_val
            };
            precalculaService.overrideModify(postData).then(function (response) {
                if(response===null){
                    popupSvc.smallBox("success",$rootScope.getWord("success"));
                    $state.go('app.Precalcula.Precalcula');
                }else {
                    popupSvc.smallBox("fail",response.msg);
                }
            });


        };

        $scope.ckecked = function (event) {
            var parent = $(event.target).parent()[0];
            var input= $(parent).find(".inputsese")[0];
            if(event.target.checked){
                $(input).removeClass('ng-hide');
                $(input).show();
            }else {
                $(input).hide();
            }
        };


    });