
angular.module('app.Platform').controller('PermissionConfigCtrl',
    function($scope, popupSvc,$LocalStorage, $rootScope, APP_CONFIG,PlatformService,$stateParams){
        var postData = {
            id: $stateParams.id,
            typed:$stateParams.role_mark
        };
        PlatformService.getRolePermissionGet(postData).then(function (response) {
            console.log(response);
            $scope.role_name = response.data.data.role_name;
            $scope.status = response.data.data.status;
            $scope.is_operate = response.data.data.is_operate;
            $scope.list = response.data.data.ps;
            var role_name = document.getElementById("role_name");
            if($scope.is_operate){
                role_name.setAttribute("disabled","disabled");
            }
        });

        $scope.selectAll= function ($event) {
            var target = $event.target;
            console.log(target);
            var isCheck=$(target).is(':checked');
            var list=$(target).parent().parent().parent().find('.check');
            for(var i=0;i<list.length;i++){
                if(isCheck){
                    $(list[i]).prop("checked",true);
                }else {
                    $(list[i]).prop("checked",false);
                }
            }
        };

        $scope.submit=function(){
            var status=$('input:radio[name="status"]:checked').val();
            var arr=[];
            var list=$('.check');
            for(var i=0;i<list.length;i++){
                var id=$(list[i]).val();
                if($(list[i]).is(':checked')){
                    arr.push(id*1);
                }
            }
            var postData = {
                role_id: $stateParams.id*1,
                role_name: $scope.role_name,
                status: status*1,
                permission_id: arr
            };
            PlatformService.getRolePermissionPost(postData).then(function (response) {
                console.log(response);
                if(response.data.data===null){
                    popupSvc.smallBox("success", $rootScope.getWord("success"));
                    history.go(-1);
                } else {
                    popupSvc.smallBox("fail", response.data.msg);
                }
            })
        }
    });
