angular.module('app.administrators').controller('AgentEditCtrl',
    function(httpSvc,popupSvc,$scope,CONFIG,$state,resourceSvc,$stateParams){
        $scope.edit = $stateParams.editid;
        console.log($scope.edit);
        httpSvc.get("/agent/third/info",{
            id: $scope.edit
        }).then(function (response) {
            console.log(response);
            $scope.detail = response.data;
        }, function (error) {

        })
        //修改后提交
        $scope.submit = function () {
            console.log($scope.detail.ch_name);
            console.log($scope.detail.us_name);
            console.log($scope.detail.card);
            console.log($scope.detail.phone);
            console.log($scope.detail.email);
            console.log($scope.detail.qq);
            console.log($scope.detail.province_id);
            console.log($scope.detail.city_id);
            console.log($scope.detail.area_id);
            console.log($scope.detail.remark);
            if($scope.detail.ch_name==''||$scope.detail.us_name==''||$scope.detail.card==''||$scope.detail.phone==''||$scope.detail.email==''||$scope.detail.qq==''||$scope.detail.province_id==''||$scope.detail.city_id==''||$scope.detail.area_id==''||$scope.detail.remark==''){
                popupSvc.smallBox("fail","请输入完整！");
            }else{
                httpSvc.put("/agent/third/info",{
                    agency_id:$scope.edit,
                    ch_name:$scope.detail.ch_name,
                    us_name:$scope.detail.us_name,
                    card:$scope.detail.card,
                    phone:$scope.detail.phone,
                    email:$scope.detail.email,
                    qq:$scope.detail.qq,
                    province_id:$scope.detail.province_id*1,
                    city_id:$scope.detail.city_id*1,
                    area_id:$scope.detail.area_id*1,
                    remark:$scope.detail.remark
                }).then(function (response) {
                    console.log(response);
                    if(response===null){
                        popupSvc.smallBox("success","操作成功");
                        $state.go('app.administrators.agent')

                    }else {
                        popupSvc.smallBox("fail",response.msg);
                    }
                }, function (error) {

                })
            }

        }
    });
